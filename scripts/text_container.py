#!env python3

from generate import *


if __name__ == "__main__":
    w = Component(
        Type="appkit.TextContainer",
        super_type='foundation.Object',
        description="a region where text layout occurs.",
        properties=[
            Property(name="size", Type='foundation.Size', description='the size of the text container’s bounding rectangle'),
            Property(name="widthTracksTextView", Type='bool', getter_prefix_is=False,
                     description='whether the text container adjusts the width of its bounding rectangle when its text view resizes'),
            Property(name="heightTracksTextView", Type='bool', getter_prefix_is=False,
                     description='whether the text container adjusts the height of its bounding rectangle when its text view resizes'),
        ],
        methods=[
            init_method(name="initWithSize", Type="appkit.TextContainer", params=[Param(name='size', Type='foundation.Size')]),
        ]
    )
    w.generate_code()
