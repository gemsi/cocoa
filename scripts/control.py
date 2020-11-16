#!env python3

from generate_widget import Generator, Property, InitMethod


if __name__ == "__main__":
    generator = Generator(
        Type = "Control", 
        super_type = 'View',
        description = "A definition of the fundamental behavior for controls, which are specialized views that notify your app of relevant events by using the target-action design pattern.",
        init_method=InitMethod(param_name='frame', param_type='foundation.Rect'),
        properties = [
            Property(name='enabled', Type = 'bool', readonly=False, description='whether the receiver reacts to mouse events'),
            Property(name='doubleValue', Type = 'float64', readonly=False, description='the value of the receiver’s cell as a double-precision floating-point number'),
            Property(name='floatValue', Type = 'float32', readonly=False, description='the value of the receiver’s cell as a single-precision floating-point number'),
            Property(name='intValue', Type = 'int', readonly=False, description='the value of the receiver’s cell as an integer'),
            Property(name='integerValue', Type = 'int64', readonly=False, description='the value of the receiver’s cell as an NSInteger value.'),
            Property(name='stringValue', Type = 'string', readonly=False, description='the value of the receiver’s cell as an NSString object'),
        ],
    )
    generator.generate_widgets()
