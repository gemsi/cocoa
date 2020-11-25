#!env python3

from generate import Component, Property, InitMethod, Param


if __name__ == "__main__":
    w = Component(
        Type="appkit.TextView",
        super_type='appkit.Text',
        description="A view that draws text and handles user interactions with that text.",
        init_method=InitMethod(name="initWithFrame", params=[Param(name='frame', Type='foundation.Rect')]),
        properties=[
            Property(name="textContainer", Type='appkit.TextContainer', description='the receiver’s text container'),
        ],
    )
    w.generate_code()
