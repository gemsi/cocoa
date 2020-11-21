#!env python3

from .generate import Component, InitMethod, Param

if __name__ == "__main__":
    w = Component(
        Type="appkit.SecureTextField",
        super_type='appkit.TextField',
        delegate_type='TextFieldDelegate',
        description="a text field that hides the typed text",
        init_method=InitMethod(name="initWithFrame", params=[Param(name='frame', Type='foundation.Rect')]),
        properties=[],
        extend_delegate=True,
    )
    w.generate_code()
