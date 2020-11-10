#!env python3

from generate_widget import *


if __name__ == "__main__":
    generator = Generator(
        Type="TextContainer",
        super_type='foundation.Object',
        description="a region where text layout occurs.",
        init_method=InitMethod(param_name='size', param_type='Size'),
        properties=[
            Property(name="size", Type='Size', description='the size of the text container’s bounding rectangle'),
            Property(name="widthTracksTextView", Type='bool',
                     description='whether the text container adjusts the width of its bounding rectangle when its text view resizes'),
            Property(name="heightTracksTextView", Type='bool',
                     description='whether the text container adjusts the height of its bounding rectangle when its text view resizes'),
        ],
    )
    generator.generate_widgets()
