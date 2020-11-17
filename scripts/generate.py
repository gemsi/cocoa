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

geo_struct_types = {'foundation.Rect', 'foundation.Point', 'foundation.Size'}


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

    def go_def_code(self, current_pkg: str) -> str:
        if self.go_alias:
            return self.name + ' ' + self.go_alias
        pkg, type_name = split_type(self.Type)
        if current_pkg == pkg:
            return self.name + ' ' + type_name
        else:
            return self.name + ' ' + self.Type

    def c_def_code(self) -> str:
        return c_type(self.Type) + ' ' + self.name

    def objc_def_code(self) -> str:
        return f'({objc_type(self.Type)}){self.name}'

    def cgo_export_def_code(self) -> str:
        return f'{self.name} {cgo_export_type(self.Type)}'

    def go_to_c_code(self) -> str:
        """convert go to c types in go code"""
        if self.Type in cgo_type_dict:
            return f'C.{cgo_type_dict[self.Type]}({self.name})'
        elif self.Type == 'string':
            return f'C.CString({self.name})'
        elif self.Type in geo_struct_types:
            return f'toNS{type_part(self.Type)}({self.name})'
        else:
            return f'{self.name}.Ptr()'

    def c_to_objc_code(self) -> str:
        """convert c types to objc types"""
        if self.Type == 'string':
            return f'[NSString stringWithUTF8String:{self.name}]'
        else:
            return self.name

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
class ReturnValue:
    Type: str
    go_alias: str = ''  # go alias type, enum, etc.

    def go_def_code(self, current_pkg: str) -> str:
        if self.go_alias:
            return self.go_alias
        pkg, type_name = split_type(self.Type)
        if current_pkg == pkg:
            return type_name
        else:
            return self.Type

    def c_def_code(self) -> str:
        return c_type(self.Type)

    def objc_def_code(self) -> str:
        return objc_type(self.Type)

    def c_to_go_code(self, return_str: str, current_pkg: str) -> str:
        """convert c types to go types, in go code"""
        if self.Type == '':
            return return_str
        elif self.Type in cgo_type_dict:
            if self.go_alias:
                return f'{self.go_alias}({return_str})'
            return f'{self.Type}({return_str})'
        elif self.Type == 'string':
            return f'C.GoString({return_str})'
        elif self.Type in geo_struct_types:
            t = type_part(self.Type)
            return f'to{t}({return_str})'
        else:
            p, t = split_type(self.Type)
            if p and p != current_pkg:
                return f'{p}.Make{t}({return_str})'
            else:
                return f'Make{t}({return_str})'

    def objc_to_c_code(self, return_str: str) -> str:
        """convert objc type to c type"""
        if self.Type == 'string':
            return f'[{return_str} UTF8String]'
        return return_str


@dataclass
class Method:
    name: str
    description: str
    params: List[Param] = field(default_factory=list)
    return_value: ReturnValue = ReturnValue('')

    def go_interface_code(self, current_pkg: str) -> List[str]:
        name = cap(self.name)
        params_str = ' '.join([p.go_def_code(current_pkg) for p in self.params])
        go_def_return = self.return_value.go_def_code(current_pkg)
        if go_def_return:
            go_def_return = ' ' + go_def_return
        return [
            f'// {name} {self.description}',
            name + '(' + params_str + ')' + go_def_return,
        ]

    def go_impl_code(self, current_pkg: str, receiver_type: str) -> List[str]:
        receiver = receiver_type[0].lower()
        name = cap(self.name)
        params_str = ' '.join([p.go_def_code(current_pkg) for p in self.params])
        receiver_str = receiver + ' *NS' + receiver_type
        go_def_return = self.return_value.go_def_code(current_pkg)
        if go_def_return:
            go_def_return = ' ' + go_def_return
        codes = ['func (' + receiver_str + ') ' + name + '(' + params_str + ')' + go_def_return + ' {', ]
        args = [f'{receiver}.Ptr()']
        for param in self.params:
            convert_code = param.go_to_c_code()
            if param.Type == 'string':
                cstr_name = 'c_' + param.name
                codes.append(f'\t{cstr_name} := {convert_code}')
                codes.append(f'\tdefer C.free(unsafe.Pointer({cstr_name}))')
                args.append(cstr_name)
            else:
                args.append(convert_code)
        c_args_str = ', '.join(args)
        return_str = self.return_value.c_to_go_code(f'C.{receiver_type}_{name}({c_args_str})', current_pkg)
        if self.return_value.Type != '':
            return_str = 'return ' + return_str
        codes.append('\t' + return_str)
        codes.append('}')
        return codes

    def c_h_code(self, receiver_type: str) -> List[str]:
        c_params_str = ' '.join([p.c_def_code() for p in self.params])
        if c_params_str:
            c_params_str = f'void* ptr, {c_params_str}'
        else:
            c_params_str = f'void* ptr'
        return [f'{self.return_value.c_def_code()} {receiver_type}_{cap(self.name)}({c_params_str});']

    def objc_m_code(self, receiver_type: str) -> List[str]:
        c_params_str = ' '.join([p.c_def_code() for p in self.params])
        if c_params_str:
            c_params_str = f'void* ptr, {c_params_str}'
        else:
            c_params_str = f'void* ptr'
        ns_var_name = de_cap(receiver_type)
        codes = [
            f'{self.return_value.c_def_code()} {receiver_type}_{cap(self.name)}({c_params_str}) {{',
            f'\tNS{receiver_type}* {ns_var_name} = (NS{receiver_type}*)ptr;',
        ]
        call_code = '[' + ns_var_name + ' ' + de_cap(self.name)
        if len(self.params) > 0:
            call_code += ':' + self.params[0].c_to_objc_code()
            for param in self.params[1:]:
                call_code += ' ' + (param.objc_param_name if param.objc_param_name else param.name) + ':' + param.name
        call_code += ']'
        call_code = self.return_value.objc_to_c_code(call_code)
        if self.return_value.Type != '':
            call_code = 'return ' + call_code
        codes.append('\t' + call_code + ';')
        codes.append('}')
        return codes


@dataclass
class DelegateMethod:
    name: str
    params: List[Param]
    description: str
    return_value: ReturnValue = ReturnValue('')

    def __post_init__(self):
        pass

    def init_env(self, current_pkg: str, receiver_type: str):
        self.current_pkg = current_pkg
        self.receiver_type = receiver_type
        self.receiver_name = receiver_type[0].lower()
        self.go_func_name = cap(self.name)
        self.go_params_str = ' '.join([p.go_def_code(current_pkg) for p in self.params])
        self.cgo_export_func_name = f'{self.receiver_type}_Delegate_{cap(self.name)}'
        self.objc_func_name = de_cap(self.name)
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
class ActionMethod:
    name: str
    params: List[Param] = field(default_factory=list)
    description: str = ''
    return_value: ReturnValue = ReturnValue('')

    def __post_init__(self):
        pass

    def init_env(self, current_pkg: str, receiver_type: str):
        if len(self.params) == 0:
            self.params = [Param(name='sender', Type='foundation.Object')]

        self.current_pkg = current_pkg
        self.receiver_type = receiver_type
        self.receiver_name = receiver_type[0].lower()
        self.go_func_name = 'Set' + cap(self.name)
        self.go_params_str = ' '.join([p.go_def_code(current_pkg) for p in self.params])
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
class InitMethod:
    name: str
    params: List[Param] = field(default_factory=list)

    def init_env(self, has_delegate: bool, action_methods: List[ActionMethod]):
        self.has_delegate = has_delegate
        self.action_methods = action_methods

    def go_code(self, Type: str, super_type: str) -> List[str]:
        current_pkg, type_name = split_type(Type)
        params_str = ', '.join([p.go_def_code(current_pkg) for p in self.params])
        c_args_str = ', '.join([param.go_to_c_code() for param in self.params])
        pkg, type_name = split_type(Type)
        super_package, super_type_name = split_type(super_type)
        if super_package == pkg:
            super_make_name = 'Make' + super_type_name
        else:
            super_make_name = super_package + '.Make' + super_type_name
        return [
            f"// New{type_name} create new {type_name}",
            f'func New{type_name}({params_str}) {type_name} {{',
            '\tid := resources.NextId()',
            f'\tptr := C.{type_name}_{self.name}(C.long(id), {c_args_str})',
            f'\tv := &NS{type_name}{{',
            f'\t\tNS{super_type_name}: *{super_make_name}(ptr),',
            '\t}',
            '\tresources.Store(id, v)',
            '\tfoundation.AddDeallocHook(v, func() {',
            '\t\tresources.Delete(id)',
            '\t})',
            '\treturn v',
            '}',
        ]

    def c_h_code(self, Type: str) -> List[str]:
        c_params_str = ', '.join([p.c_def_code() for p in self.params])
        if c_params_str:
            c_params_str = f'long goID, {c_params_str}'
        else:
            c_params_str = f'long goID'
        return [f'void* {Type}_{self.name}({c_params_str});']

    def objc_m_code(self, Type: str) -> List[str]:
        c_params_str = ', '.join([p.c_def_code() for p in self.params])
        if c_params_str:
            c_params_str = f'long goID, {c_params_str}'
        else:
            c_params_str = f'long goID'
        var_name = camel_to_underscore(Type)
        params_str = self.params[0].c_to_objc_code()
        for param in self.params[1:]:
            params_str += ' ' + (param.objc_param_name if param.objc_param_name else param.name) + ':' + param.name

        codes = [
            f'void* {Type}_{self.name}({c_params_str}) {{',
            f'\tNS{Type}* {var_name} = [[[NS{Type} alloc] {self.name}:{params_str}] autorelease];',
        ]

        if self.has_delegate:
            codes.extend([
                f'\tGoNS{Type}Delegate* delegate = [[GoNS{Type}Delegate alloc] init];',
                '\t[delegate setGoID:goID];',
                f'\t[{var_name} setDelegate:delegate];'
            ])
        if self.action_methods:
            codes.extend([
                f'\tNS{Type}Handler* handler = [[[NS{Type}Handler alloc] init] autorelease];',
                '\t[handler setGoID:goID];',
                f'\t[{var_name} setTarget:handler];',
            ])
            for method in self.action_methods:
                codes.append(f'\t[{var_name} set{cap(method.name)}:@selector(on{cap(method.name)}:)];')
        codes.extend([
            f'\treturn (void*){var_name};',
            '}',
        ])
        return codes


@dataclass
class Property:
    name: str  # the property name
    Type: str  # the property type
    description: str
    readonly: bool = False
    go_alias_type: str = ''  # the go alias type, for enum etc..
    getter_prefix_is: bool = True

    def getter(self) -> Method:
        return Method(
            name='is' + cap(self.name) if self.Type == 'bool' and self.getter_prefix_is else self.name,
            params=[],
            return_value=ReturnValue(Type=self.Type, go_alias=self.go_alias_type),
            description='return ' + self.description,
        )

    def setter(self) -> Optional[Method]:
        if self.readonly:
            return None
        return Method(
            name='set' + cap(self.name),
            params=[Param(name=self.name, Type=self.Type, go_alias=self.go_alias_type)],
            return_value=ReturnValue(Type=''),
            description='set ' + self.description,
        )


@dataclass
class Component:
    Type: str
    super_type: str
    description: str
    properties: List[Property]
    init_method: Optional[InitMethod] = None
    methods: List[Method] = field(default_factory=list)
    delegate_type: str = ''
    delegate_methods: List[DelegateMethod] = field(default_factory=list)
    action_methods: List[ActionMethod] = field(default_factory=list)
    extend_delegate: bool = False

    def __post_init__(self):
        self.pkg, self.type_name = split_type(self.Type)
        self.file_name = camel_to_underscore(self.type_name)
        self.ns_type = 'NS' + self.type_name
        self.super_package, self.super_type_name = split_type(self.super_type)
        if self.delegate_type == '':
            self.delegate_type = type_part(self.Type) + 'Delegate'
        if self.init_method is not None:
            self.init_method.init_env(len(self.delegate_methods) > 0, self.action_methods)
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

        go_file_path = f'./{self.pkg}/{self.file_name}.go'

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
            if self.init_method is not None:
                print(file=out)
                for line in self.init_method.go_code(self.Type, self.super_type):
                    print(line, file=out)
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
        if self.init_method is not None:
            types.extend([param.Type for param in self.init_method.params])
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
        ns_name = 'NS' + type_name
        file_name = camel_to_underscore(type_name)
        package_path = './' + pkg
        try:
            os.makedirs(package_path)
        except FileExistsError:
            pass

        # header files
        header_file_path = f'{package_path}/{file_name}.h'
        with open(header_file_path, 'w+') as out:
            for s in self.get_c_imports():
                print(f'#import <{s}>', file=out)
            if self.init_method is not None:
                print(file=out)
                for line in self.init_method.c_h_code(type_name):
                    print(line, file=out)

            print(file=out)
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
        delegate_header_file_path = f'./{self.pkg}/{self.file_name}_delegate.h'
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
                print('@property (assign) long goID;', file=out)
                print(f'@end', file=out)
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
                print(f'@interface {self.ns_type}Handler : NSObject', file=out)
                print('@property (assign) long goID;', file=out)
                print(f'@end', file=out)
                print(file=out)
                print(f'@implementation {self.ns_type}Handler', file=out)
                for method in self.action_methods:
                    print(file=out)
                    for line in method.objc_m_code():
                        print(line, file=out)
                print(file=out)
                print(f'@end', file=out)

    def generate_objc_m_file(self):
        m_file_path = f'./{self.pkg}/{self.file_name}.m'
        with open(m_file_path, 'w+') as out:
            print(f'#import <Appkit/{self.ns_type}.h>', file=out)
            print(f'#import "{self.file_name}.h"', file=out)
            if self.delegate_methods or self.action_methods:
                print(f'#import "{self.file_name}_delegate.h"', file=out)
            # init method
            if self.init_method is not None:
                print(file=out)
                for line in self.init_method.objc_m_code(self.type_name):
                    print(line, file=out)
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
        }
        types: List[str] = []
        if self.init_method is not None:
            types.extend([param.Type for param in self.init_method.params])
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
        package_path = './' + self.pkg
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
