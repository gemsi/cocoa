import sys
import re
import os
from dataclasses import dataclass
from typing import List, Optional
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
    go_alias_type: Optional[str] = None  # the go alias type, for enum etc..


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
        func New{name}({init_param_name} foundation.{init_param_type}) {name} {{
        \tid := resources.NextId()
        \tptr := C.{name}_New(C.long(id), toNS{init_param_type}({init_param_name}))

        \tv := &{ns_name}{{
        \t\t{super_field_name}: *{super_make_name}(ptr),
        \t}}

        \tresources.Store(id, v)

        \tfoundation.AddDeallocHook(v, func() {{
        \t\tresources.Delete(id)
        \t}})

        \treturn v
        }}''')

        interface_methods: List[str] = []
        implimetion_methods: List[str] = []

        for property in self.properties:
            self.generate_property_go(
                property, interface_methods, implimetion_methods)

        with open(go_file_path, 'w+') as out:
            print(f'package widget', file=out)
            print(cgo_imports_str, file=out)
            self.imports.sort()
            print('import (', file=out)
            for s in self.imports:
                print(f'\t"{s}"', file=out)
            print(')', file=out)

            print(interface_str, file=out)
            for s in interface_methods:
                print(s, file=out)
            print('}', file=out)

            print(struct_str, file=out)

            print(make_method_str, file=out)
            print(init_method_str, file=out)
            for s in implimetion_methods:
                print(s, file=out)

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
        init_param_type = 'NS' + self.init_method.param_type
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

        header_funcs: List[str] = []
        m_funcs: List[str] = []
        for property in self.properties:
            self.generate_property_objc(property, header_funcs, m_funcs)

        header_file_path = f'{package_path}/{file_name}.h'
        with open(header_file_path, 'w+') as out:
            for s in imports:
                print(f'#import <{s}>', file=out)
            print(new_method_str, file=out)
            for s in header_funcs:
                print(s, file=out)

        m_file_path = f'{package_path}/{file_name}.m'
        with open(m_file_path, 'w+') as out:
            print(m_import_str, file=out)
            print(new_method_impl, file=out)
            for s in m_funcs:
                print(s, file=out)

    def generate_widgets(self):
        self.generate_go_file()
        self.generate_objc_file()

    def generate_property_go(self, property: Property, interface_methods: List[str], implimetion_methods: List[str]):
        name = self.Type
        receiver_name = to_receiver_name(self.Type)
        method_name = cap(property.name)
        if property.Type in cgo_type_dict:
            go_type = property.go_alias_type if property.go_alias_type else property.Type
            cgo_type = cgo_type_dict[property.Type]
            interface_methods.append(
                (f'\t// {method_name} return {property.description}\n\t{method_name}() {go_type}'))

            implimetion_methods.append(textwrap.dedent(f'''
            func ({receiver_name} *NS{name}) {method_name}() {go_type} {{
            \treturn {go_type}(C.{name}_{method_name}({receiver_name}.Ptr()))
            }}'''))

            if not property.readonly:

                interface_methods.append(
                    f'''\t// Set{method_name} set {property.description}\n\tSet{method_name}(value {go_type})''')

                implimetion_methods.append(textwrap.dedent(f'''
                func ({receiver_name} *NS{name}) Set{method_name}(value {go_type}) {{
                \tC.{name}_Set{method_name}({receiver_name}.Ptr(), C.{cgo_type}(value))
                }}'''))
        elif property.Type == 'string':
            go_type = property.Type
            interface_methods.append(
                (f'\t// {method_name} return {property.description}\n\t{method_name}() {go_type}'))

            implimetion_methods.append(textwrap.dedent(f'''
            func ({receiver_name} *NS{name}) {method_name}() {go_type} {{
            \treturn C.GoString(C.{name}_{method_name}({receiver_name}.Ptr()))
            }}'''))

            if not property.readonly:

                interface_methods.append(
                    f'''\t// Set{method_name} set {property.description}\n\tSet{method_name}(value {go_type})''')

                implimetion_methods.append(textwrap.dedent(f'''
                func ({receiver_name} *NS{name}) Set{method_name}(value {go_type}) {{
                \tcstr := C.CString(value)
                \tdefer C.free(unsafe.Pointer(cstr))
                \tC.{name}_Set{method_name}({receiver_name}.Ptr(), cstr)
                }}'''))
        elif property.Type in ('Rect', 'Point', 'Size'):
            go_type = f'foundation.{property.Type}'
            full_package = f'github.com/hsiafan/cocoa/foundation'
            if full_package not in self.imports:
                self.imports.append(full_package)
            interface_methods.append(
                (f'\t// {method_name} return {property.description}\n\t{method_name}() {go_type}'))

            implimetion_methods.append(textwrap.dedent(f'''
            func ({receiver_name} *NS{name}) {method_name}() {go_type} {{
            \treturn to{property.Type}(C.{name}_{method_name}({receiver_name}.Ptr()))
            }}'''))

            if not property.readonly:

                interface_methods.append(
                    f'''\t// Set{method_name} set {property.description}\n\tSet{method_name}(value {go_type})''')

                implimetion_methods.append(textwrap.dedent(f'''
                func ({receiver_name} *NS{name}) Set{method_name}(value {go_type}) {{
                \tC.{name}_Set{method_name}({receiver_name}.Ptr(), toNS{property.Type}(value))
                }}'''))
        else:
            go_type = f'{property.Type}'
            make_method = f'Make{property.Type}'
            interface_methods.append(
                (f'\t// {method_name} return {property.description}\n\t{method_name}() {go_type}'))

            implimetion_methods.append(textwrap.dedent(f'''
            func ({receiver_name} *NS{name}) {method_name}() {go_type} {{
            \treturn {make_method}(C.{name}_{method_name}({receiver_name}.Ptr()))
            }}'''))

            if not property.readonly:

                interface_methods.append(
                    f'''\t// Set{method_name} set {property.description}\n\tSet{method_name}(value {go_type})''')

                implimetion_methods.append(textwrap.dedent(f'''
                func ({receiver_name} *NS{name}) Set{method_name}(value {go_type}) {{
                \tC.{name}_Set{method_name}({receiver_name}.Ptr(), value.Ptr())
                }}'''))

    def generate_property_objc(self, property: Property, header_funcs: List[str], m_funcs: List[str]):
        pname = cap(property.name)
        var_name = camel_to_underscore(self.Type)
        type_name = property.Type
        widget_type = self.Type
        if property.Type in c_type_dict:
            type_name = c_type_dict[property.Type]
            header_funcs.append(
                f'''\n{type_name} {widget_type}_{pname}(void* ptr); ''')
            m_funcs.append(textwrap.dedent(f'''
            {type_name} {widget_type}_{pname}(void* ptr) {{
                NS{widget_type}* {var_name} = (NS{widget_type}*)ptr;
                return {var_name}.{property.name};
            }}'''))
            if not property.readonly:
                header_funcs.append(
                    f'''\nvoid {widget_type}_Set{pname}(void* ptr, {type_name} value); ''')
                m_funcs.append(textwrap.dedent(f'''
                void {widget_type}_Set{pname}(void* ptr, {type_name} value) {{
                    NS{widget_type}* {var_name} = (NS{widget_type}*)ptr;
                    return [{var_name} set{pname}:value];
                }}'''))
        elif property.Type in ('Rect', 'Point', 'Size'):
            type_name = f'NS{type_name}'
            header_funcs.append(
                f'''\n{type_name} {widget_type}_{pname}(void* ptr); ''')
            m_funcs.append(textwrap.dedent(f'''
            {type_name} {widget_type}_{pname}(void* ptr) {{
                NS{widget_type}* {var_name} = (NS{widget_type}*)ptr;
                return {var_name}.{property.name};
            }}'''))
            if not property.readonly:
                header_funcs.append(
                    f'''\nvoid {widget_type}_Set{pname}(void* ptr, {type_name} value); ''')
                m_funcs.append(textwrap.dedent(f'''
                void {widget_type}_Set{pname}(void* ptr, {type_name} value) {{
                    NS{widget_type}* {var_name} = (NS{widget_type}*)ptr;
                    [{var_name} set{pname}:value];
                }}'''))
        elif property.Type == 'string':
            type_name = 'const char*'
            header_funcs.append(
                f'''\n{type_name} {widget_type}_{pname}(void* ptr); ''')
            m_funcs.append(textwrap.dedent(f'''
            {type_name} {widget_type}_{pname}(void* ptr) {{
                NS{widget_type}* {var_name} = (NS{widget_type}*)ptr;
                return [{var_name}.{property.name} UTF8String];
            }}'''))
            if not property.readonly:
                header_funcs.append(
                    f'''\nvoid {widget_type}_Set{pname}(void* ptr, {type_name} value); ''')
                m_funcs.append(textwrap.dedent(f'''
                void {widget_type}_Set{pname}(void* ptr, {type_name} value) {{
                    NS{widget_type}* {var_name} = (NS{widget_type}*)ptr;
                    [{var_name} set{pname}:[NSString stringWithUTF8String:value]];
                }}'''))
        else:
            header_funcs.append(
                f'''\nvoid* {widget_type}_{pname}(void* ptr); ''')
            m_funcs.append(textwrap.dedent(f'''
            void* {widget_type}_{pname}(void* ptr) {{
                NS{widget_type}* {var_name} = (NS{widget_type}*)ptr;
                return {var_name}.{property.name};
            }}'''))
            if not property.readonly:
                header_funcs.append(
                    f'''\nvoid {widget_type}_Set{pname}(void* ptr, void* valuePtr); ''')
                m_funcs.append(textwrap.dedent(f'''
                void {widget_type}_Set{pname}(void* ptr, void* valuePtr) {{
                    NS{widget_type}* {var_name} = (NS{widget_type}*)ptr;
                    NS{property.Type}* value = (NS{property.Type}*)valuePtr;
                    [{var_name} set{pname}:value];
                }}'''))


def camel_to_underscore(s: str) -> str:
    return re.sub(r'(?<!^)(?=[A-Z])', '_', s).lower()


def cap(s: str) -> str:
    return s[:1].upper() + s[1:]


def to_receiver_name(s: str) -> str:
    return s[0].lower()
