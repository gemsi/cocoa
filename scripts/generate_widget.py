import sys
import re
import os
from dataclasses import dataclass
from typing import List, Optional, Tuple
import textwrap


cgo_type_dict = {
    'bool': 'bool',
    'byte': 'char',
    'int8': 'schar',
    'uint8': 'uchar',
    'int16': 'short',
    'uint16': 'ushort',
    'int': 'int',
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
    'int': 'int',
    'int32': 'int',
    'uint32': 'uint',
    'int64': 'long',
    'uint64': 'unsigned long',
    'float32': 'float',
    'float64': 'double',
}

struct_types = {'foundation.Rect', 'foundation.Point', 'foundation.Size'}


def type_part(Type: str) -> str:
    idx = Type.find('.')
    if idx >= 0:
        return Type[idx+1:]
    return Type


def package_part(Type: str) -> str:
    idx = Type.find('.')
    if idx >= 0:
        return Type[:idx]
    return ''


def objc_type(Type: str) -> str:
    if Type == '':
        return 'void'
    elif Type in c_type_dict:
        return c_type_dict[Type]
    elif Type in struct_types:
        return 'NS' + type_part(Type)
    elif Type == 'string':
        return 'const char*'
    else:
        # return 'NS' + type_part(Type)
        return 'void*'


@dataclass
class Param:
    name: str
    Type: str
    go_alias: str = ''  # go alias type, enum, etc.
    objc_param_name: str = ''  # objc param def name

    def go_def_code(self) -> str:
        if self.go_alias:
            return self.name + ' ' + self.go_alias
        return self.name + ' ' + self.Type

    def objc_def_code(self) -> str:
        return objc_type(self.Type) + ' ' + self.name

    def go_convert_code(self) -> str:
        # TODO: cgo convert code
        if self.Type in cgo_type_dict:
            return f'C.{cgo_type_dict[self.Type]}({self.name})'
        elif self.Type == 'string':
            return f'C.CString({self.name})'
        elif self.Type in struct_types:
            return f'toNS{type_part(self.Type)}({self.name})'
        else:
            return f'{self.name}.Ptr()'

    def objc_convert_code(self) -> str:
        if self.Type == 'string':
            return f'[NSString stringWithUTF8String:{self.name}]'
        else:
            return self.name


@dataclass
class ReturnValue:
    Type: str
    go_alias: str = ''  # go alias type, enum, etc.

    def go_def_code(self) -> str:
        if self.go_alias:
            return self.go_alias
        return self.Type

    def objc_def_code(self) -> str:
        return objc_type(self.Type)

    def go_convert_code(self, return_str: str) -> str:
        if self.Type == '':
            return return_str
        elif self.Type in cgo_type_dict:
            if self.go_alias:
                return f'{self.go_alias}({return_str})'
            return f'{self.Type}({return_str})'
        elif self.Type == 'string':
            return f'C.GoString({return_str})'
        elif self.Type in struct_types:
            t = type_part(self.Type)
            return f'to{t}({return_str})'
        else:
            t = type_part(self.Type)
            p = package_part(self.Type)
            if p:
                p += '.'
            return f'{p}Make{t}({return_str})'

    def objc_convert_code(self, return_str: str) -> str:
        if self.Type == 'string':
            return f'[{return_str} UTF8String]'
        return return_str


@dataclass
class Method:
    name: str
    params: List[Param]
    return_value: ReturnValue
    description: str

    def go_interface_code(self) -> List[str]:
        name = cap(self.name)
        params_str = ' '.join([p.go_def_code() for p in self.params])
        go_def_return = self.return_value.go_def_code()
        if go_def_return:
            go_def_return = ' ' + go_def_return
        return [
            f'// {name} {self.description}',
            name + '(' + params_str + ')' + go_def_return,
        ]

    def go_impl_code(self, receiver_type: str) -> List[str]:
        receiver = receiver_type[0].lower()
        name = cap(self.name)
        params_str = ' '.join([p.go_def_code() for p in self.params])
        receiver_str = receiver + ' *NS' + receiver_type
        go_def_return = self.return_value.go_def_code()
        if go_def_return:
            go_def_return = ' ' + go_def_return
        codes = ['func (' + receiver_str + ') ' + name + '(' + params_str + ')' + go_def_return + ' {', ]
        args = [f'{receiver}.Ptr()']
        for param in self.params:
            convert_code = param.go_convert_code()
            if param.Type == 'string':
                cstr_name = 'c_' + param.name
                codes.append(f'\t{cstr_name} := {convert_code}')
                codes.append(f'\tdefer C.free(unsafe.Pointer({cstr_name}))')
                args.append(cstr_name)
            else:
                args.append(convert_code)
        c_args_str = ', '.join(args)
        return_str = self.return_value.go_convert_code(f'C.{receiver_type}_{name}({c_args_str})')
        if self.return_value.Type != '':
            return_str = 'return ' + return_str
        codes.append('\t' + return_str)
        codes.append('}')
        return codes

    def objc_h_code(self, receiver_type: str) -> List[str]:
        c_params_str = ' '.join([p.objc_def_code() for p in self.params])
        if c_params_str:
            c_params_str = f'void* ptr, {c_params_str}'
        else:
            c_params_str = f'void* ptr'
        return [f'{self.return_value.objc_def_code()} {receiver_type}_{cap(self.name)}({c_params_str});']

    def objc_m_code(self, receiver_type: str) -> List[str]:
        c_params_str = ' '.join([p.objc_def_code() for p in self.params])
        if c_params_str:
            c_params_str = f'void* ptr, {c_params_str}'
        else:
            c_params_str = f'void* ptr'
        ns_var_name = de_cap(receiver_type)
        codes = [
            f'{self.return_value.objc_def_code()} {receiver_type}_{cap(self.name)}({c_params_str}) {{',
            f'\tNS{receiver_type}* {ns_var_name} = (NS{receiver_type}*)ptr;',
        ]
        call_code = '[' + ns_var_name + ' ' + de_cap(self.name)
        if len(self.params) > 0:
            call_code += ':' + self.params[0].objc_convert_code()
            for param in self.params[1:]:
                call_code += ' ' + (param.objc_param_name if param.objc_param_name else param.name) + ':' + param.name
        call_code += ']'
        call_code = self.return_value.objc_convert_code(call_code)
        if self.return_value.Type != '':
            call_code = 'return ' + call_code
        codes.append('\t' + call_code + ';')
        codes.append('}')
        return codes


@dataclass
class InitMethod:
    param_name: str
    param_type: str


@dataclass
class Property:
    name: str  # the property name
    Type: str  # the property type
    description: str
    readonly: bool = False
    go_alias_type: str = ''  # the go alias type, for enum etc..
    prefixIs: bool = True

    def getter(self) -> Method:
        return Method(
            name='is' + cap(self.name) if self.Type == 'bool' and self.prefixIs else self.name,
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
class Generator:
    Type: str
    super_type: str
    description: str
    init_method: InitMethod
    properties: List[Property]

    def generate_go_file(self):
        name = self.Type
        super_type = self.super_type
        ns_name = 'NS' + name
        dot_idx = super_type.find('.')
        if dot_idx > 0:
            super_ns_name = super_type[:dot_idx+1] + \
                'NS' + super_type[dot_idx+1:]
            super_field_name = 'NS' + super_type[dot_idx+1:]
            super_make_name = super_type[:dot_idx +
                                         1] + 'Make' + super_type[dot_idx+1:]
        else:
            super_ns_name = 'NS' + super_type
            super_field_name = 'NS' + super_type
            super_make_name = 'Make' + super_type
        file_name = camel_to_underscore(name)
        package_path = '../appkit/widget'
        try:
            os.makedirs(package_path)
        except FileExistsError:
            pass

        go_file_path = f'{package_path}/{file_name}.go'

        cgo_imports_str = textwrap.dedent(f'''
        // #cgo CFLAGS: -x objective-c
        // #cgo LDFLAGS: -framework Cocoa
        // #include "{file_name}.h"
        import "C"
        ''')

        self.imports = [
            "github.com/hsiafan/cocoa/foundation",
            "unsafe",
        ]

        interface_str = textwrap.dedent(f'''
        // {self.description}
        type {name} interface {{
        \t{super_type}
        ''')

        struct_str = textwrap.dedent(f'''
        var _ {name} = (*{ns_name})(nil)

        type {ns_name} struct {{
        \t{super_ns_name}
        }}''')

        make_method_str = textwrap.dedent(f'''
        // Make create a {name} from native pointer
        func Make{name}(ptr unsafe.Pointer) *{ns_name} {{
        \treturn &{ns_name}{{*{super_make_name}(ptr)}}
        }}''')

        init_param_name = self.init_method.param_name
        init_param_type = self.init_method.param_type
        init_method_str = textwrap.dedent(f'''
        // New create new {name}
        func New{name}({init_param_name} {init_param_type}) {name} {{
        \tid := resources.NextId()
        \tptr := C.{name}_New(C.long(id), toNS{type_part(init_param_type)}({init_param_name}))

        \tv := &{ns_name}{{
        \t\t{super_field_name}: *{super_make_name}(ptr),
        \t}}

        \tresources.Store(id, v)

        \tfoundation.AddDeallocHook(v, func() {{
        \t\tresources.Delete(id)
        \t}})

        \treturn v
        }}''')

        with open(go_file_path, 'w+') as out:
            print(f'package widget', file=out)
            print(cgo_imports_str, file=out)
            self.imports.sort()
            print('import (', file=out)
            for s in self.imports:
                print(f'\t"{s}"', file=out)
            print(')', file=out)

            # interface
            print(interface_str, file=out)
            # interface properties
            for property in self.properties:
                getter = property.getter()
                for line in getter.go_interface_code():
                    print('\t' + line, file=out)
                setter = property.setter()
                if setter is not None:
                    for line in setter.go_interface_code():
                        print('\t' + line, file=out)
            print('}', file=out)

            # struct impl
            print(struct_str, file=out)
            print(make_method_str, file=out)
            print(init_method_str, file=out)
            # properties impl
            for property in self.properties:
                getter = property.getter()
                print(file=out)
                for line in getter.go_impl_code(self.Type):
                    print(line, file=out)
                setter = property.setter()
                if setter is not None:
                    print(file=out)
                    for line in setter.go_impl_code(self.Type):
                        print(line, file=out)

    def generate_objc_file(self):
        name = self.Type
        ns_name = 'NS' + name
        file_name = camel_to_underscore(name)
        var_name = file_name
        package_path = f'../appkit/widget'
        try:
            os.makedirs(package_path)
        except FileExistsError:
            pass

        imports = [
            'stdlib.h',
            'Foundation/NSGeometry.h',
        ]

        init_param_name = self.init_method.param_name
        init_param_type = 'NS' + type_part(self.init_method.param_type)
        init_param_mname = 'initWith' + cap(self.init_method.param_name)
        new_method_str = f'\nvoid* {name}_New(long id, {init_param_type} {init_param_name});'
        new_method_impl = textwrap.dedent(f'''
        void* {name}_New(long id, {init_param_type} {init_param_name}) {{
            {ns_name}* {var_name} = [[[{ns_name} alloc] {init_param_mname}:{init_param_name}] autorelease];
            return (void*){var_name};
        }}''')
        m_import_str = textwrap.dedent(f'''#import <Appkit/{ns_name}.h>
        #import "{file_name}.h"
        #include "_cgo_export.h"
        ''')

        # header files
        header_file_path = f'{package_path}/{file_name}.h'
        with open(header_file_path, 'w+') as out:
            for s in imports:
                print(f'#import <{s}>', file=out)
            print(new_method_str, file=out)
            # properties header defination
            for property in self.properties:
                getter = property.getter()
                for line in getter.objc_h_code(self.Type):
                    print(line, file=out)
                setter = property.setter()
                if setter is not None:
                    for line in setter.objc_h_code(self.Type):
                        print(line, file=out)

        # m files
        m_file_path = f'{package_path}/{file_name}.m'
        with open(m_file_path, 'w+') as out:
            print(m_import_str, file=out)
            print(new_method_impl, file=out)
            # properties header defination
            for property in self.properties:
                getter = property.getter()
                for line in getter.objc_m_code(self.Type):
                    print(line, file=out)
                setter = property.setter()
                if setter is not None:
                    for line in setter.objc_m_code(self.Type):
                        print(line, file=out)

    def generate_widgets(self):
        self.generate_go_file()
        self.generate_objc_file()


def camel_to_underscore(s: str) -> str:
    return re.sub(r'(?<!^)(?=[A-Z])', '_', s).lower()


def cap(s: str) -> str:
    return s[:1].upper() + s[1:]


def de_cap(s: str) -> str:
    return s[:1].lower() + s[1:]


def to_receiver_name(s: str) -> str:
    return s[0].lower()
