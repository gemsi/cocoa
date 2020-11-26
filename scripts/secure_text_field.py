#!env python3

from generate import Component, Param, init_method

if __name__ == "__main__":
    w = Component(
        Type="appkit.SecureTextField",
        super_type='appkit.TextField',
        delegate_type='TextFieldDelegate',
        description="a text field that hides the typed text",
        properties=[],
        extend_delegate=True,
        methods=[
            init_method(name="initWithFrame", Type="appkit.SecureTextField",
                        params=[Param(name='frame', Type='foundation.Rect')]),
        ]
    )
    w.generate_code()
