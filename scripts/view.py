#!env python3

from generate_widget import Generator, Property, InitMethod


if __name__ == "__main__":
    generator = Generator(
        Type = "View", 
        super_type = 'Responder',
        description = "The infrastructure for drawing, printing, and handling events in an app.",
        init_method=InitMethod(param_name='frame', param_type='foundation.Rect'),
        properties = [
            Property(name='frame', Type = 'foundation.Rect', readonly=False, description='the view’s frame rectangle, which defines its position and size in its superview’s coordinate system'),
            Property(name='autoresizingMask', Type = 'uint64', go_alias_type='AutoresizingMaskOptions', readonly=False, description='the view’s frame rectangle, which defines its position and size in its superview’s coordinate system'),
        ],
    )
    generator.generate_widgets()