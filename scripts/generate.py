import os
import re
import textwrap
from dataclasses import dataclass, field
from typing import List, Optional, Tuple, Set

cgo_type_dict = {
    'bool': 'bool',
    'byte': 'char',
    'int8': 'schar',
    'uint8': 'uchar',
    'int16': 'short',
    'uint16': 'ushort',
    'int': 'long',
    'uint': 'ulong',
    'int32': 'int',
    'uint32': 'uint',
    'int64': 'long',
    'uint64': 'ulong',
    'float32': 'float',
    'float64': 'double',
}

c_type_dict = {
    'bool': 'bool',
    'byte': 'char',
    'int8': 'signed char',
    'uint8': 'unsigned char',
    'int16': 'short',
    'uint16': 'unsigned short',
    'int': 'long',
    'uint': 'unsigned long',
    'int32': 'int',
    'uint32': 'uint',
    'int64': 'long',
    'uint64': 'unsigned long',
    'float32': 'float',
    'float64': 'double',
}

geo_struct_types = {'foundation.Rect', 'foundation.Point', 'foundation.Size', 'foundation.EdgeInsets'}


def type_part(Type: str) -> str:
    idx = Type.find('.')
    if idx >= 0:
        return Type[idx + 1:]
    return Type


def package_part(Type: str) -> str:
    idx = Type.find('.')
    if idx >= 0:
        return Type[:idx]
    return ''


def split_type(Type: str) -> Tuple[str, str]:
    idx = Type.find('.')
    if idx >= 0:
        return Type[:idx], Type[idx + 1:]
    return '', Type


def c_type(Type: str) -> str:
    if Type == '':
        return 'void'
    elif Type in c_type_dict:
        return c_type_dict[Type]
    elif Type in geo_struct_types:
        return 'NS' + type_part(Type)
    elif Type == 'string':
        return 'const char*'
    else:
        # return 'NS' + type_part(Type)
        return 'void*'


def objc_type(Type: str) -> str:
    if Type == '':
        return 'void'
    elif Type in c_type_dict:
        if Type == 'bool':
            return 'BOOL'
        return c_type_dict[Type]
    elif Type in geo_struct_types:
        return 'NS' + type_part(Type)
    elif Type == 'string':
        return 'NSString*'
    else:
        return 'NS' + type_part(Type) + '*'


def cgo_export_type(Type: str) -> str:
    if Type == '':
        return ''
    elif Type in c_type_dict:
        return Type
    elif Type in geo_struct_types:
        return 'C.' + type_part(Type)
    elif Type == 'string':
        return '*C.char'
    else:
        return 'unsafe.Pointer'


@dataclass
class Param:
    name: str
    Type: str
    go_alias: str = ''  # go alias type, enum, etc.
    objc_param_name: str = ''  # objc param def name
    array: bool = False

    def go_def_code(self, current_pkg: str) -> str:
        pkg, type_name = split_type(self.go_alias if self.go_alias else self.Type)
        def_type = type_name if pkg == '' or current_pkg == pkg else pkg + '.' + type_name
        if self.array:
            def_type = '[]' + def_type
        return self.name + ' ' + def_type

    def c_def_code(self) -> str:
        if self.array:
            # Note: only support NSObject array
            return f'Array {self.name}'
        return c_type(self.Type) + ' ' + self.name

    def objc_def_code(self) -> str:
        code = f'({objc_type(self.Type)}){self.name}'
        if self.objc_param_name:
            code = self.objc_param_name + ':' + code
        return code

    def cgo_export_def_code(self) -> str:
        return f'{self.name} {cgo_export_type(self.Type)}'

    def go_to_c_code(self) -> Tuple[List[str], str]:
        """convert go to c types in go code"""
        if self.array:
            # Note: only support NSObject array
            c_name = 'c' + cap(self.name)
            codes = [
                f'{c_name}Data := make([]unsafe.Pointer, len({self.name}))',
                f'for idx, v := range {self.name} {{',
                f'\t{c_name}Data[idx] = v.Ptr()',
                '}',
                f'{c_name} := C.Array{{data:unsafe.Pointer(&{c_name}Data[0]), len:C.int(len({self.name}))}}',
            ]
            return codes, c_name
        elif self.Type in cgo_type_dict:
            return [], f'C.{cgo_type_dict[self.Type]}({self.name})'
        elif self.Type == 'string':
            cstr_name = 'c' + cap(self.name)
            codes = [
                f'{cstr_name} := C.CString({self.name})',
                f'defer C.free(unsafe.Pointer({cstr_name}))',
            ]
            return codes, cstr_name
        elif self.Type in geo_struct_types:
            return [], f'toNS{type_part(self.Type)}({self.name})'
        else:
            return [], f'toPointer({self.name})'

    def c_to_objc_code(self) -> Tuple[List[str], str]:
        """convert c types to objc types"""
        pkg, type_name = split_type(self.go_alias if self.go_alias else self.Type)
        if self.array:
            # Note: only support NSObject array
            objc_name = f'objc{cap(self.name)}'
            codes = [
                f'NSMutableArray* {objc_name} = [[NSMutableArray alloc] init];',
                f'NS{type_name}** {self.name}Data = (NS{type_name}**){self.name}.data;',
                f'for (int i = 0; i < {self.name}.len; i++) {{',
                f'    [{objc_name} addObject:{self.name}Data[i]];',
                '}',
            ]
            return codes, objc_name
        elif self.Type == 'string':
            return [], f'[NSString stringWithUTF8String:{self.name}]'
        else:
            return [], self.name

    def objc_to_c_code(self) -> str:
        """convert objc type to c type"""
        if self.Type == 'string':
            return f'[{self.name} UTF8String]'
        return self.name

    def cgo_export_to_go_code(self, current_pkg: str) -> str:
        """convert cgo export type to go types"""
        if self.Type in c_type_dict:
            return self.name
        elif self.Type in geo_struct_types:
            return 'to' + type_part(self.Type)
        elif self.Type == 'string':
            return f'C.GoString({self.name})'
        else:
            p, t = split_type(self.Type)
            if p != current_pkg:
                return f'{p}.Make{t}({self.name})'
            else:
                return f'Make{t}({self.name})'


@dataclass
class Return:
    Type: str
    go_alias: str = ''  # go alias type, enum, etc.
    array: bool = False

    def go_def_code(self, current_pkg: str) -> str:
        pkg, type_name = split_type(self.go_alias if self.go_alias else self.Type)
        def_type = type_name if pkg == '' or current_pkg == pkg else pkg + '.' + type_name
        if self.array:
            def_type = '[]' + def_type
        return def_type

    def c_def_code(self) -> str:
        t = c_type(self.Type)
        if self.array:
            t = 'Array'
        return t

    def objc_def_code(self) -> str:
        return objc_type(self.Type)

    def c_to_go_code(self, return_str: str, current_pkg: str) -> Tuple[List[str], str]:
        """convert c types to go types, in go code"""
        if self.Type == '':
            return [], return_str
        elif self.array:
            p, t = split_type(self.Type)
            if p and p != current_pkg:
                make_func = f'{p}.Make{t}'
            else:
                make_func = f'Make{t}'
            pkg, type_name = split_type(self.go_alias if self.go_alias else self.Type)
            def_type = type_name if pkg == '' or current_pkg == pkg else pkg + '.' + type_name
            codes = [
                f'var cArray C.Array =  {return_str}',
                f'defer C.free(cArray.data)',
                f'result := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(cArray.data))[:cArray.len:cArray.len]',
                f'var goResult = make([]{def_type}, len(result))',
                f'for idx, r := range result {{',
                f'\tgoResult[idx] = {make_func}(r)',
                '}',
            ]
            return codes, 'goResult'
        elif self.Type in cgo_type_dict:
            if self.go_alias:
                return [], f'{self.go_alias}({return_str})'
            return [], f'{self.Type}({return_str})'
        elif self.Type == 'string':
            return [], f'C.GoString({return_str})'
        elif self.Type in geo_struct_types:
            t = type_part(self.Type)
            return [], f'to{t}({return_str})'
        else:
            p, t = split_type(self.Type)
            if p and p != current_pkg:
                return [], f'{p}.Make{t}({return_str})'
            else:
                return [], f'Make{t}({return_str})'

    def objc_to_c_code(self, return_str: str) -> Tuple[List[str], str]:
        """convert objc type to c type"""
        if self.array:
            codes = [
                f'NSArray* array = {return_str};',
                'int count = [array count];',
                'void** data = malloc(count * sizeof(void*));',
                'for (int i = 0; i < count; i++) {',
                '\t data[i] = [array objectAtIndex:i];',
                '}',
                'Array result;',
                'result.data = data;',
                'result.len = count;',
            ]
            return codes, 'result'
        elif self.Type == 'string':
            return [], f'[{return_str} UTF8String]'
        return [], return_str


@dataclass
class Method:
    name: str
    description: str
    params: List[Param] = field(default_factory=list)
    return_value: Return = Return('')
    do_alloc: bool = False  # for init method
    static: bool = False  # for static method
    register_delegate: bool = False  # if register delegate and action target
    go_func_name: str = ''

    def init_env(self, has_delegate: bool, has_target: bool):
        self.has_delegate = has_delegate
        self.has_target = has_target

    def go_interface_code(self, current_pkg: str) -> List[str]:
        if self.do_alloc or self.static:
            return []
        if self.go_func_name:
            name = self.go_func_name
        else:
            name = cap(self.name)
        params_str = ', '.join([p.go_def_code(current_pkg) for p in self.params])
        go_def_return = self.return_value.go_def_code(current_pkg)
        if go_def_return:
            go_def_return = ' ' + go_def_return
        return [
            f'// {name} {self.description}',
            name + '(' + params_str + ')' + go_def_return,
        ]

    def go_impl_code(self, current_pkg: str, receiver_type: str) -> List[str]:
        if self.go_func_name:
            go_func_name = self.go_func_name
        else:
            go_func_name = cap(self.name)
        params_str = ', '.join([p.go_def_code(current_pkg) for p in self.params])

        go_def_return = self.return_value.go_def_code(current_pkg)
        if go_def_return:
            go_def_return = ' ' + go_def_return
        if not self.do_alloc and not self.static:
            receiver = receiver_type[0].lower()
            receiver_str = receiver + ' *NS' + receiver_type
            codes = ['func (' + receiver_str + ') ' + go_func_name + '(' + params_str + ')' + go_def_return + ' {', ]
            args = [f'{receiver}.Ptr()']
        else:
            codes = ['func ' + go_func_name + '(' + params_str + ')' + go_def_return + ' {', ]
            args = []

        for param in self.params:
            convert_codes, arg = param.go_to_c_code()
            for convert_code in convert_codes:
                codes.append('\t' + convert_code)
            args.append(arg)
        c_args_str = ', '.join(args)
        convert_codes, return_str = self.return_value.c_to_go_code(
            f'C.{receiver_type}_{cap(self.name)}({c_args_str})', current_pkg)
        for convert_code in convert_codes:
            codes.append('\t' + convert_code)
        if self.register_delegate and (self.has_delegate or self.has_target) and self.return_value.Type != '':
            codes.append(f'\tinstance := {return_str}')
            codes.append(f'\tinstance.setDelegate()')
            return_str = 'instance'
        if self.return_value.Type != '':
            return_str = 'return ' + return_str
        codes.append('\t' + return_str)
        codes.append('}')
        return codes

    def c_h_code(self, receiver_type: str) -> List[str]:
        params: List[str] = []
        if not self.do_alloc and not self.static:
            params.append('void* ptr')
        params.extend([p.c_def_code() for p in self.params])
        c_params_str = ', '.join(params)
        return [f'{self.return_value.c_def_code()} {receiver_type}_{cap(self.name)}({c_params_str});']

    def objc_m_code(self, receiver_type: str) -> List[str]:
        params: List[str] = []
        if not self.do_alloc and not self.static:
            params.append('void* ptr')
        params.extend([p.c_def_code() for p in self.params])
        c_params_str = ', '.join(params)
        ns_var_name = de_cap(receiver_type)
        codes = [
            f'{self.return_value.c_def_code()} {receiver_type}_{cap(self.name)}({c_params_str}) {{',
        ]
        if self.do_alloc:
            codes.append(f'\tNS{receiver_type}* {ns_var_name} = [NS{receiver_type} alloc];')
        elif not self.static:
            codes.append(f'\tNS{receiver_type}* {ns_var_name} = (NS{receiver_type}*)ptr;')
        else:
            ns_var_name = f'NS{receiver_type}'
        call_code = '[' + ns_var_name + ' ' + de_cap(self.name)
        if len(self.params) > 0:
            convert_codes, arg = self.params[0].c_to_objc_code()
            for c in convert_codes:
                codes.append('    ' + c)
            call_code += ':' + arg
            for param in self.params[1:]:
                convert_codes, arg = param.c_to_objc_code()
                for c in convert_codes:
                    codes.append('    ' + c)
                call_code += ' ' + (param.objc_param_name if param.objc_param_name else param.name) + ':' + arg
        call_code += ']'
        convert_codes, call_code = self.return_value.objc_to_c_code(call_code)
        for c in convert_codes:
            codes.append('\t' + c)
        if self.do_alloc:
            call_code = f'[{call_code} autorelease]'
        if self.return_value.Type != '':
            call_code = 'return ' + call_code
        codes.append('\t' + call_code + ';')
        codes.append('}')
        return codes


def init_method(name: str, Type: str, params: List[Param] = field(default_factory=list), go_func_name: str = '', description: str = '') -> Method:
    Type = type_part(Type)
    if not go_func_name:
        go_func_name = f'New{Type}'
    if not description:
        description = f'return a new {Type} instance'

    return Method(name=name, params=params, return_value=Return(Type=Type), go_func_name=go_func_name, description=description, do_alloc=True, register_delegate=True)


@dataclass
class DelegateMethod:
    name: str
    params: List[Param]
    description: str
    return_value: Return = Return('')
    go_func_name: str = ''
    default_return_value: str = ''

    def __post_init__(self):
        pass

    def init_env(self, current_pkg: str, receiver_type: str):
        self.current_pkg = current_pkg
        self.receiver_type = receiver_type
        self.receiver_name = receiver_type[0].lower()
        self.go_params_str = ', '.join([p.go_def_code(current_pkg) for p in self.params])
        self.objc_func_name = de_cap(self.name)
        self.go_def_return = self.return_value.go_def_code(current_pkg)
        if self.go_def_return:
            self.go_def_return = ' ' + self.go_def_return
        if not self.go_func_name:
            self.go_func_name = cap(self.name)
        self.cgo_export_func_name = f'{self.receiver_type}_Delegate_{self.go_func_name}'
        self.callback_field_name = de_cap(self.go_func_name)
        self.callback_get_func = '_' + self.callback_field_name
        self.callback_type = 'func(' + self.go_params_str + ')' + self.go_def_return

    def go_interface_code(self) -> List[str]:
        return [
            f'// {self.go_func_name} {self.description}',
            f'{self.go_func_name}(callback {self.callback_type})',
            f'{self.callback_get_func}() {self.callback_type}',
        ]

    def go_impl_code(self) -> List[str]:
        receiver_str = self.receiver_name + ' *NS' + self.receiver_type
        codes = [
            f'func ({receiver_str}) {self.go_func_name}(callback {self.callback_type}) {{',
            f'\t{self.receiver_name}.{self.callback_field_name} = callback',
            '}',
            '',
            f'func ({receiver_str}) {self.callback_get_func}() {self.callback_type} {{',
            f'\treturn {self.receiver_name}.{self.callback_field_name}',
            '}',
        ]
        return codes

    def go_callback_field_code(self) -> List[str]:
        return [f'{self.callback_field_name} func({self.go_params_str}){self.go_def_return}']

    def cgo_export_code(self) -> List[str]:
        param_name = de_cap(self.receiver_type)
        params_def_str = ', '.join([param.cgo_export_def_code() for param in self.params])
        args_str = ', '.join([param.cgo_export_to_go_code(self.current_pkg) for param in self.params])
        return_str = self.return_value.go_def_code(self.current_pkg)
        if return_str:
            return_str = ' ' + return_str
        codes = [
            f'//export {self.cgo_export_func_name}',
            f'func {self.cgo_export_func_name}(id int64, {params_def_str}){return_str} {{',
            f'\t{param_name} := resources.Get(id).({self.receiver_type})',
            f'\tcallback := {param_name}.{self.callback_get_func}()',
            f'\tif callback != nil {{',
        ]
        return_state = f'\t\tcallback({args_str})'
        if return_str:
            return_state = 'return ' + return_state
        codes.append(return_state)
        codes.append('\t}')
        if return_str:
            codes.append(f'\treturn {self.default_return_value}')
        codes.append('}')
        return codes

    def objc_m_code(self) -> List[str]:
        params_def_str = ' '.join([param.objc_def_code() for param in self.params])
        args_str = ', '.join([param.objc_to_c_code() for param in self.params])
        # TODO: convert ns types to c types?
        return [
            f'- ({self.return_value.objc_def_code()}){self.objc_func_name}:{params_def_str} {{',
            f'\treturn {self.cgo_export_func_name}([self goID], {args_str});',
            '}',
        ]


@dataclass
class ActionMethod:
    name: str
    params: List[Param] = field(default_factory=list)
    description: str = ''
    return_value: Return = Return('')

    def __post_init__(self):
        pass

    def init_env(self, current_pkg: str, receiver_type: str):
        if len(self.params) == 0:
            self.params = [Param(name='sender', Type='foundation.Object')]

        self.current_pkg = current_pkg
        self.receiver_type = receiver_type
        self.receiver_name = receiver_type[0].lower()
        self.go_func_name = 'Set' + cap(self.name)
        self.go_params_str = ', '.join([p.go_def_code(current_pkg) for p in self.params])
        self.cgo_export_func_name = f'{self.receiver_type}_Target_{cap(self.name)}'
        self.objc_func_name = 'on' + cap(self.name)
        self.callback_field_name = de_cap(self.name)
        self.go_def_return = self.return_value.go_def_code(current_pkg)
        if self.go_def_return:
            self.go_def_return = ' ' + self.go_def_return

    def go_interface_code(self) -> List[str]:
        return [
            f'// {self.go_func_name} {self.description}',
            self.go_func_name + '(callback func(' + self.go_params_str + ')' + self.go_def_return + ')',
        ]

    def go_impl_code(self) -> List[str]:
        receiver_str = self.receiver_name + ' *NS' + self.receiver_type
        codes = [
            'func (' + receiver_str + ') ' + self.go_func_name +
            '(callback func(' + self.go_params_str + ')' + self.go_def_return + ')' + ' {',
            f'\t{self.receiver_name}.{self.callback_field_name} = callback',
            '}'
        ]
        return codes

    def go_callback_field_code(self) -> List[str]:
        return [f'{self.callback_field_name} func({self.go_params_str})']

    def cgo_export_code(self) -> List[str]:
        param_name = de_cap(self.receiver_type)
        params_def_str = ' '.join([param.cgo_export_def_code() for param in self.params])
        args_str = ' '.join([param.cgo_export_to_go_code(self.current_pkg) for param in self.params])
        return [
            f'//export {self.cgo_export_func_name}',
            f'func {self.cgo_export_func_name}(id int64, {params_def_str}) {{',
            f'\t{param_name} := resources.Get(id).(*NS{self.receiver_type})',
            f'\tif {param_name}.{self.callback_field_name} != nil {{',
            f'\t\t{param_name}.{self.callback_field_name}({args_str})',
            '\t}',
            '}'
        ]

    def objc_m_code(self) -> List[str]:
        params_def_str = ' '.join([param.objc_def_code() for param in self.params])
        args_str = ' '.join([param.objc_to_c_code() for param in self.params])
        # TODO: convert ns types to c types?
        return [
            f'- ({self.return_value.objc_def_code()}){self.objc_func_name}:{params_def_str} {{',
            f'\treturn {self.cgo_export_func_name}([self goID], {args_str});',
            '}',
        ]


@dataclass
class Property:
    name: str  # the property name
    Type: str  # the property type
    description: str
    readonly: bool = False
    go_alias_type: str = ''  # the go alias type, for enum etc..
    getter_prefix_is: bool = True
    array: bool = False

    def getter(self) -> Method:
        return Method(
            name='is' + cap(self.name) if self.Type == 'bool' and self.getter_prefix_is else self.name,
            params=[],
            return_value=Return(Type=self.Type, go_alias=self.go_alias_type, array=self.array),
            description='return ' + self.description,
        )

    def setter(self) -> Optional[Method]:
        if self.readonly:
            return None
        return Method(
            name='set' + cap(self.name),
            params=[Param(name=self.name, Type=self.Type, go_alias=self.go_alias_type, array=self.array)],
            return_value=Return(Type=''),
            description='set ' + self.description,
        )


@dataclass
class Component:
    Type: str
    super_type: str
    description: str
    properties: List[Property]
    methods: List[Method] = field(default_factory=list)
    delegate_type: str = ''
    delegate_methods: List[DelegateMethod] = field(default_factory=list)
    action_methods: List[ActionMethod] = field(default_factory=list)
    extend_delegate: bool = False
    additional_objc_imports: List[str] = field(default_factory=list)

    def __post_init__(self):
        self.pkg, self.type_name = split_type(self.Type)
        self.file_name = camel_to_underscore(self.type_name)
        self.ns_type = 'NS' + self.type_name
        self.super_package, self.super_type_name = split_type(self.super_type)
        if self.delegate_type == '':
            self.delegate_type = type_part(self.Type) + 'Delegate'
        for method in self.methods:
            method.init_env(self.delegate_methods or self.extend_delegate, self.action_methods)
        for method in self.delegate_methods:
            pkg, type_name = split_type(self.Type)
            method.init_env(pkg, type_name)
        for method in self.action_methods:
            pkg, type_name = split_type(self.Type)
            method.init_env(pkg, type_name)

    def generate_go_file(self):
        pkg, type_name = split_type(self.Type)
        ns_name = 'NS' + type_name
        super_package, super_type_name = split_type(self.super_type)
        if super_package == pkg:
            super_name = super_type_name
            super_ns_name = 'NS' + super_type_name
            super_make_name = 'Make' + super_type_name
        else:
            super_name = self.super_type
            super_ns_name = super_package + '.NS' + super_type_name
            super_make_name = super_package + '.Make' + super_type_name

        go_file_path = f'../{self.pkg}/{self.file_name}.go'

        cgo_imports_str = textwrap.dedent(f'''
        // #cgo CFLAGS: -x objective-c
        // #cgo LDFLAGS: -framework Cocoa
        // #include "{self.file_name}.h"
        import "C"
        ''')

        interface_str = textwrap.dedent(f'''
        // {type_name} is {self.description}
        type {type_name} interface {{
        \t{super_name}
        ''')

        make_method_str = textwrap.dedent(f'''
        // Make{type_name} create a {type_name} from native pointer
        func Make{type_name}(ptr unsafe.Pointer) *{ns_name} {{
        \tif ptr == nil {{
        \t\treturn nil
        \t}}
        \treturn &{ns_name}{{
        \t\tNS{super_type_name}: *{super_make_name}(ptr),
        \t}}
        }}''')

        with open(go_file_path, 'w+') as out:
            print(f'package {pkg}', file=out)
            print(cgo_imports_str, file=out)
            print('import (', file=out)
            for s in self.get_go_imports():
                print(f'\t"{s}"', file=out)
            print(')', file=out)

            # interface
            print(interface_str, file=out)
            # interface properties
            for p in self.properties:
                getter = p.getter()
                for line in getter.go_interface_code(pkg):
                    print('\t' + line, file=out)
                setter = p.setter()
                if setter is not None:
                    for line in setter.go_interface_code(pkg):
                        print('\t' + line, file=out)
            # interface methods
            for method in self.methods:
                for line in method.go_interface_code(pkg):
                    print('\t' + line, file=out)
            # delegate call backs
            for method in self.delegate_methods:
                for line in method.go_interface_code():
                    print('\t' + line, file=out)
            # action call backs
            for method in self.action_methods:
                for line in method.go_interface_code():
                    print('\t' + line, file=out)
            print('}', file=out)

            # struct impl
            print(file=out)
            print(f'var _ {type_name} = (*{ns_name})(nil)', file=out)
            print(file=out)
            print(f'type {ns_name} struct {{', file=out)
            print(f'\t{super_ns_name}', file=out)
            for method in self.delegate_methods:
                for line in method.go_callback_field_code():
                    print(f'\t{line}', file=out)
            for method in self.action_methods:
                for line in method.go_callback_field_code():
                    print(f'\t{line}', file=out)
            print('}', file=out)
            print(make_method_str, file=out)

            # delegate and target
            if self.delegate_methods or self.extend_delegate or self.action_methods:
                receiver = self.type_name[0].lower()
                receiver_str = receiver + ' *' + self.ns_type
                print(f'func ({receiver_str}) setDelegate() {{', file=out)
                print('\tid := resources.NextId()', file=out)
                print(f'\tC.{self.type_name}_RegisterDelegate({receiver}.Ptr(), C.long(id))', file=out)
                print(f'\tresources.Store(id, {receiver})', file=out)
                print(f'\tfoundation.AddDeallocHook({receiver}, func() {{', file=out)
                print('\t\tresources.Delete(id)', file=out)
                print('\t})', file=out)
                print('}', file=out)
            # properties impl
            for p in self.properties:
                getter = p.getter()
                print(file=out)
                for line in getter.go_impl_code(pkg, type_name):
                    print(line, file=out)
                setter = p.setter()
                if setter is not None:
                    print(file=out)
                    for line in setter.go_impl_code(pkg, type_name):
                        print(line, file=out)
            # methods
            for method in self.methods:
                print(file=out)
                for line in method.go_impl_code(pkg, type_name):
                    print(line, file=out)
            # call backs
            for method in self.delegate_methods:
                print(file=out)
                for line in method.go_impl_code():
                    print(line, file=out)
            for method in self.action_methods:
                print(file=out)
                for line in method.go_impl_code():
                    print(line, file=out)

            # delegate/action cgo exports funcs
            for method in self.delegate_methods:
                print(file=out)
                for line in method.cgo_export_code():
                    print(line, file=out)
            for method in self.action_methods:
                print(file=out)
                for line in method.cgo_export_code():
                    print(line, file=out)

    def get_go_imports(self) -> List[str]:
        imports: Set[str] = {
            "unsafe",
        }
        pkg = package_part(self.Type)
        types: List[str] = [self.super_type]
        types.extend([p.Type for p in self.properties])
        for method in self.methods:
            types.extend([param.Type for param in method.params])
            types.append(method.return_value.Type)
        for t in types:
            if not t:
                continue
            t_pkg = package_part(t)
            if not t_pkg:
                continue
            if pkg != t_pkg:
                imports.add(f'github.com/hsiafan/cocoa/{t_pkg}')
        return sorted(list(imports))

    def generate_c_header_file(self):
        pkg, type_name = split_type(self.Type)
        file_name = camel_to_underscore(type_name)

        # header files
        header_file_path = f'../{pkg}/{file_name}.h'
        with open(header_file_path, 'w+') as out:
            for s in self.get_c_imports():
                print(f'#import <{s}>', file=out)

            print(file=out)
            if self.delegate_methods or self.extend_delegate or self.action_methods:
                print(f'void {self.type_name}_RegisterDelegate(void *ptr, long goID);', file=out)
            # properties header definition
            for p in self.properties:
                getter = p.getter()
                for line in getter.c_h_code(type_name):
                    print(line, file=out)
                setter = p.setter()
                if setter is not None:
                    for line in setter.c_h_code(type_name):
                        print(line, file=out)
            print(file=out)
            # methods
            for method in self.methods:
                for line in method.c_h_code(type_name):
                    print(line, file=out)

    def generate_delegate_header_file(self):
        if not self.delegate_methods and not self.action_methods and not self.extend_delegate:
            return
        delegate_header_file_path = f'../{self.pkg}/{self.file_name}_delegate.h'
        with open(delegate_header_file_path, 'w+') as out:
            print(f'#import <Appkit/{self.ns_type}.h>', file=out)
            print(f'#import "_cgo_export.h"', file=out)
            base_delegate_class = 'NSObject'
            if self.extend_delegate:
                print(f'#import "{camel_to_underscore(self.super_type_name)}_delegate.h"', file=out)
                base_delegate_class = f'GoNS{self.super_type_name}Delegate'
            # delegate
            print(file=out)
            if self.delegate_methods or self.extend_delegate:
                print(f'@interface Go{self.ns_type}Delegate : {base_delegate_class} <NS{self.delegate_type}>', file=out)
                if not self.extend_delegate:
                    print('@property (assign) long goID;', file=out)
                print(f'@end', file=out)
            # action target
            if self.action_methods:
                print(f'@interface {self.ns_type}Handler : NSObject', file=out)
                print('@property (assign) long goID;', file=out)
                print(f'@end', file=out)

    def generate_objc_m_file(self):
        m_file_path = f'../{self.pkg}/{self.file_name}.m'
        with open(m_file_path, 'w+') as out:
            print(f'#import <Appkit/{self.ns_type}.h>', file=out)
            print(f'#import "{self.file_name}.h"', file=out)
            for import_ in self.additional_objc_imports:
                print(f'#import <{import_}.h>', file=out)
            if self.delegate_methods or self.extend_delegate or self.action_methods:
                print(f'#import "{self.file_name}_delegate.h"', file=out)

            # delegate and target impl
            if self.delegate_methods or self.extend_delegate:
                print(file=out)
                print(f'@implementation Go{self.ns_type}Delegate', file=out)
                for method in self.delegate_methods:
                    print(file=out)
                    for line in method.objc_m_code():
                        print(line, file=out)
                print(file=out)
                print(f'@end', file=out)
            # action target
            if self.action_methods:
                print(file=out)
                print(f'@implementation {self.ns_type}Handler', file=out)
                for method in self.action_methods:
                    print(file=out)
                    for line in method.objc_m_code():
                        print(line, file=out)
                print(file=out)
                print(f'@end', file=out)

            # set delegate
            if self.delegate_methods or self.extend_delegate or self.action_methods:
                print(f'void {self.type_name}_RegisterDelegate(void *ptr, long goID) {{', file=out)
                var_name = de_cap(self.type_name)
                print(f'\t{self.ns_type}* {var_name} = ({self.ns_type}*)ptr;', file=out)
                if self.delegate_methods or self.extend_delegate:
                    print(
                        f'\tGoNS{self.type_name}Delegate* delegate = [[[GoNS{self.type_name}Delegate alloc] init] autorelease];', file=out)
                    print('\t[delegate setGoID:goID];', file=out)
                    print(f'\t[{var_name} setDelegate:delegate];', file=out)
                if self.action_methods:
                    print(
                        f'\tNS{self.type_name}Handler* handler = [[[NS{self.type_name}Handler alloc] init] autorelease];', file=out)
                    print('\t[handler setGoID:goID];', file=out)
                    print(f'\t[{var_name} setTarget:handler];', file=out)
                    for method in self.action_methods:
                        print(f'\t[{var_name} set{cap(method.name)}:@selector(on{cap(method.name)}:)];', file=out)
                print('}', file=out)

            # properties header definition
            for p in self.properties:
                getter = p.getter()
                print(file=out)
                for line in getter.objc_m_code(self.type_name):
                    print(line, file=out)
                setter = p.setter()
                if setter is not None:
                    print(file=out)
                    for line in setter.objc_m_code(self.type_name):
                        print(line, file=out)
            # methods
            for method in self.methods:
                print(file=out)
                for line in method.objc_m_code(self.type_name):
                    print(line, file=out)

    def get_c_imports(self) -> List[str]:
        imports: Set[str] = {
            'stdlib.h',
            'utils.h'
        }
        types: List[str] = []
        types.extend([p.Type for p in self.properties])
        for method in self.methods:
            types.extend([param.Type for param in method.params])
            types.append(method.return_value.Type)
        for t in types:
            if not t:
                continue
            if t in geo_struct_types:
                imports.add('Foundation/NSGeometry.h')
            elif t == 'bool':
                imports.add('stdbool.h')
        return sorted(list(imports))

    def generate_code(self):
        print('generate code for', self.Type)
        package_path = '../' + self.pkg
        try:
            os.makedirs(package_path)
        except FileExistsError:
            pass
        self.generate_go_file()
        self.generate_c_header_file()
        self.generate_delegate_header_file()
        self.generate_objc_m_file()


def camel_to_underscore(s: str) -> str:
    return re.sub(r'(?<!^)(?=[A-Z])', '_', s).lower()


def cap(s: str) -> str:
    return s[:1].upper() + s[1:]


def de_cap(s: str) -> str:
    return s[:1].lower() + s[1:]


def to_receiver_name(s: str) -> str:
    return s[0].lower()
