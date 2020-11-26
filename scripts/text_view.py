#!env python3

from generate import Component, Property, init_method, Param


if __name__ == "__main__":
    w = Component(
        Type="appkit.TextView",
        super_type='appkit.Text',
        description="A view that draws text and handles user interactions with that text.",
        properties=[
            Property(name="textContainer", Type='appkit.TextContainer', description='the receiver’s text container'),
        ],
        methods=[
            init_method(name="initWithFrame", Type="appkit.TextView", params=[Param(name='frame', Type='foundation.Rect')]),
        ],
    )
    w.generate_code()
