#!env python3

from generate_widget import Generator, Property, InitMethod


if __name__ == "__main__":
    generator = Generator(
        Type="TextView",
        super_type='Text',
        description="A view that draws text and handles user interactions with that text.",
        init_method=InitMethod(param_name='frame', param_type='foundation.Rect'),
        properties=[
            Property(name="textContainer", Type='TextContainer', description='the receiver’s text container'),
        ],
    )
    generator.generate_widgets()
