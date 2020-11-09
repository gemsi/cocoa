#!env python3

import sys
import re
import os
from dataclasses import dataclass
from typing import List, Optional
import textwrap


helper_methods_str = '''
func toNSRect(rect geometry.Rect) C.NSRect {
	return *(*C.NSRect)(unsafe.Pointer(&rect))
}
func toRect(nsRect C.NSRect) geometry.Rect {
	return *(*geometry.Rect)(unsafe.Pointer(&nsRect))
}'''


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
class Property:
    name: str # the property name
    Type: str # the property type 
    readonly: bool
    description: str
    go_alias_type: Optional[str] = None # the go alias type, for enum etc..

@dataclass
class Generator:
    Type: str
    super_type: str
    description:str
    properties: List[Property]

    def generate_go_file(self):
        name = self.Type
        super_type = self.super_type
        ns_name = 'NS' + name
        super_ns_name = 'NS' + super_type
        package_name = name.lower()
        super_package_name = super_type.lower()
        file_name = camel_to_underscore(name)
        package_path = f'../appkit/{package_name}'
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
            "github.com/hsiafan/cocoa/foundation/geometry",
            "github.com/hsiafan/cocoa/foundation/object",
            "github.com/hsiafan/cocoa/internal",
            "unsafe",
        ]
        self.imports.append(full_pacakge_path(super_package_name))
        
        interface_str = textwrap.dedent(f'''
        // {self.description}
        type {name} interface {{
        \t{super_package_name}.{super_type}
        ''')

        struct_str = textwrap.dedent(f'''
        var _ {name} = (*{ns_name})(nil)

        type {ns_name} struct {{
        \t{super_package_name}.{super_ns_name}
        }}''')

        init_method_str = textwrap.dedent(f'''
        var resources = internal.NewResourceRegistry()

        // Make create a View from native pointer
        func Make(ptr unsafe.Pointer) *{ns_name} {{
        \treturn &{ns_name}{{*{super_package_name}.Make(ptr)}}
        }}

        // New create new {name}
        func New(frame geometry.Rect) {name} {{
        \tid := resources.NextId()
        \tptr := C.{name}_New(C.long(id), toNSRect(frame))

        \tv := &{ns_name}{{
        \t\t{super_ns_name}: *{super_package_name}.Make(ptr),
        \t}}

        \tresources.Store(id, v)

        \tobject.AddDeallocHook(v, func() {{
        \t\tresources.Delete(id)
        \t}})

        \treturn v
        }}''')

        interface_methods: List[str] = []
        implimetion_methods: List[str] = []

        for property in self.properties:
            self.generate_property_go(property, interface_methods, implimetion_methods)
        

        with open(go_file_path, 'w+') as out:
            print(f'package {package_name}', file=out)
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
            for s in implimetion_methods:
                print(s, file=out)
            print(init_method_str, file=out)
            print(helper_methods_str, file=out)


    def generate_objc_file(self):
        name = self.Type
        ns_name = 'NS' + name
        package_name = name.lower()
        file_name = camel_to_underscore(name)
        var_name = file_name
        package_path = f'../appkit/{package_name}'
        try:
            os.makedirs(package_path)
        except FileExistsError:
            pass
        
        imports = [
            'stdlib.h',
            'Foundation/NSGeometry.h',
        ]
        new_method_str = f'\nvoid* {name}_New(long id, NSRect frame);'
        new_method_impl = textwrap.dedent(f'''
        void* {name}_New(long id, NSRect frame) {{
            {ns_name}* {var_name} = [[[{ns_name} alloc] initWithFrame:frame] autorelease];
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

    def generate_property_go(self, property:Property, interface_methods:List[str], implimetion_methods:List[str]):
        name = self.Type
        receiver_name = to_receiver_name(self.Type)
        method_name = cap(property.name)
        if property.Type in cgo_type_dict:
            go_type = property.go_alias_type if property.go_alias_type else property.Type
            idx = go_type.rfind('.')
            if idx >0:
                go_package = full_pacakge_path(go_type[:idx])
                if go_package not in self.imports:
                    self.imports.append(go_package)
            cgo_type = cgo_type_dict[property.Type]
            interface_methods.append((f'\t// {method_name} return {property.description}\n\t{method_name}() {go_type}'))

            implimetion_methods.append(textwrap.dedent(f'''
            func ({receiver_name} *NS{name}) {method_name}() {go_type} {{
            \treturn {go_type}(C.{name}_{method_name}({receiver_name}.Ptr()))
            }}'''))
            
            if not property.readonly:
            
                interface_methods.append(f'''\t// Set{method_name} set {property.description}\n\tSet{method_name}(value {go_type})''')
            
                implimetion_methods.append(textwrap.dedent(f'''
                func ({receiver_name} *NS{name}) Set{method_name}(value {go_type}) {{
                \tC.{name}_Set{method_name}({receiver_name}.Ptr(), C.{cgo_type}(value))
                }}'''))
        elif property.Type in ('Rect', 'Point', 'Size'):
            go_type = f'geometry.{property.Type}'
            interface_methods.append((f'\t// {method_name} return {property.description}\n\t{method_name}() {go_type}'))

            implimetion_methods.append(textwrap.dedent(f'''
            func ({receiver_name} *NS{name}) {method_name}() {go_type} {{
            \treturn to{property.Type}(C.{name}_{method_name}({receiver_name}.Ptr()))
            }}'''))
            
            if not property.readonly:
            
                interface_methods.append(f'''\t// Set{method_name} set {property.description}\n\tSet{method_name}(value {go_type})''')
            
                implimetion_methods.append(textwrap.dedent(f'''
                func ({receiver_name} *NS{name}) Set{method_name}(value {go_type}) {{
                \tC.{name}_Set{method_name}({receiver_name}.Ptr(), toNS{property.Type}(value))
                }}'''))
        else:
            package_name = property.Type.lower()
            go_type = f'{package_name}.{property.Type}' if property.Type != self.Type else property.Type
            make_method = f'{package_name}.Make' if property.Type != self.Type else 'Make'
            go_package = full_pacakge_path(package_name)
            if go_package not in self.imports and property.Type != self.Type:
                self.imports.append(go_package)
            interface_methods.append((f'\t// {method_name} return {property.description}\n\t{method_name}() {go_type}'))

            implimetion_methods.append(textwrap.dedent(f'''
            func ({receiver_name} *NS{name}) {method_name}() {go_type} {{
            \treturn {make_method}(C.{name}_{method_name}({receiver_name}.Ptr()))
            }}'''))
            
            if not property.readonly:
            
                interface_methods.append(f'''\t// Set{method_name} set {property.description}\n\tSet{method_name}(value {go_type})''')
            
                implimetion_methods.append(textwrap.dedent(f'''
                func ({receiver_name} *NS{name}) Set{method_name}(value {go_type}) {{
                \tC.{name}_Set{method_name}({receiver_name}.Ptr(), value.Ptr())
                }}'''))

    def generate_property_objc(self, property:Property, header_funcs: List[str], m_funcs: List[str]):
        pname = cap(property.name)
        var_name = camel_to_underscore(self.Type)
        type_name = property.Type
        widget_type = self.Type
        if property.Type in c_type_dict:
            type_name = c_type_dict[property.Type]
            header_funcs.append(f'''\n{type_name} {widget_type}_{pname}(void* ptr); ''')
            m_funcs.append(textwrap.dedent(f'''
            {type_name} {widget_type}_{pname}(void* ptr) {{
                NS{widget_type}* {var_name} = (NS{widget_type}*)ptr;
                return {var_name}.{property.name};
            }}'''))
            if not property.readonly:
                header_funcs.append(f'''\nvoid {widget_type}_Set{pname}(void* ptr, {type_name} value); ''')
                m_funcs.append(textwrap.dedent(f'''
                void {widget_type}_Set{pname}(void* ptr, {type_name} value) {{
                    NS{widget_type}* {var_name} = (NS{widget_type}*)ptr;
                    return [{var_name} set{pname}:value];
                }}'''))
        elif property.Type in ('Rect', 'Point', 'Size'):
            type_name = f'NS{type_name}'
            header_funcs.append(f'''\n{type_name} {widget_type}_{pname}(void* ptr); ''')
            m_funcs.append(textwrap.dedent(f'''
            {type_name} {widget_type}_{pname}(void* ptr) {{
                NS{widget_type}* {var_name} = (NS{widget_type}*)ptr;
                return {var_name}.{property.name};
            }}'''))
            if not property.readonly:
                header_funcs.append(f'''\nvoid {widget_type}_Set{pname}(void* ptr, {type_name} value); ''')
                m_funcs.append(textwrap.dedent(f'''
                void {widget_type}_Set{pname}(void* ptr, {type_name} value) {{
                    NS{widget_type}* {var_name} = (NS{widget_type}*)ptr;
                    return [{var_name} set{pname}:value];
                }}'''))
        else:
            header_funcs.append(f'''\nvoid* {widget_type}_{pname}(void* ptr); ''')
            m_funcs.append(textwrap.dedent(f'''
            void* {widget_type}_{pname}(void* ptr) {{
                NS{widget_type}* {var_name} = (NS{widget_type}*)ptr;
                return {var_name}.{property.name};
            }}'''))
            if not property.readonly:
                header_funcs.append(f'''\nvoid {widget_type}_Set{pname}(void* ptr, void* valuePtr); ''')
                m_funcs.append(textwrap.dedent(f'''
                void {widget_type}_Set{pname}(void* ptr, void* valuePtr) {{
                    NS{widget_type}* {var_name} = (NS{widget_type}*)ptr;
                    NS{property.Type}* value = (NS{property.Type}*)valuePtr;
                    return [{var_name} set{pname}:value];
                }}'''))

def camel_to_underscore(s:str) -> str:
    return re.sub(r'(?<!^)(?=[A-Z])', '_', s).lower()

def cap(s:str) -> str:
    return s[:1].upper() + s[1:]

def to_receiver_name(s: str) -> str:
    return s[0].lower()

def full_pacakge_path(package_name:str) -> str:
    return f'github.com/hsiafan/cocoa/appkit/{package_name}'
