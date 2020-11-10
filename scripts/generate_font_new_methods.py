#!env python3

import textwrap

def generate_go(name:str, description:str) -> str:
    name = name[0].upper() + name[1:]
    return textwrap.dedent(f'''
    // New{name} {description}
    func New{name}(size float64) {{
        C.Font_{name}(C.double(size))
    }}''')

def generate_h(name:str) -> str:
    name = name[0].upper() + name[1:]
    return f'void* Font_{name}(double size);'

def generate_m(name:str) -> str:
    cname = name[0].upper() + name[1:]
    return textwrap.dedent(f'''
    void* Font_{cname}(double size) {{
        return [NSFont {name}:size];
    }}''')


def generate_all_size_methods():
    methods = [
        ('boldSystemFontOfSize', 'returns the standard system font in boldface type with the specified size'),
        ('labelFontOfSize', 'returns the font used for standard interface labels in the specified size'),
        ('messageFontOfSize', 'returns the font used for standard interface items, such as button labels, menu items, and so on, in the specified size.'),
        ('menuBarFontOfSize', 'returns the font used for menu bar items, in the specified size'),
        ('menuFontOfSize', 'returns the font used for menu items, in the specified size'),
        ('controlContentFontOfSize', 'returns the font used for the content of controls in the specified size'),
        ('titleBarFontOfSize', 'returns the font used for window title bars, in the specified size'),
        ('paletteFontOfSize', 'returns the font used for palette window title bars, in the specified size'),
        ('toolTipsFontOfSize', 'returns the font used for tool tips labels, in the specified size'),
        ('systemFontSizeForControlSize', 'Returns the font size used for the specified control size.'),
    ]
    for method in methods:
        print(generate_go(method[0], method[1]))

    for method in methods:
        print(generate_h(method[0]))

    for method in methods:
        print(generate_m(method[0]))


if __name__ == '__main__':
    generate_all_size_methods()


