#!env python3

from generate import Component, Property, InitMethod, Method

if __name__ == "__main__":
    w = Component(
        Type="appkit.Control",
        super_type='appkit.View',
        description="A definition of the fundamental behavior for controls, which are specialized views that notify your app of relevant events by using the target-action design pattern.",
        properties=[
            Property(name='enabled', Type='bool', description='whether the receiver reacts to mouse events'),
            Property(name='doubleValue', Type='float64',
                     description='the value of the receiver’s cell as a double-precision floating-point number'),
            Property(name='floatValue', Type='float32',
                     description='the value of the receiver’s cell as a single-precision floating-point number'),
            Property(name='intValue', Type='int', description='the value of the receiver’s cell as an integer'),
            Property(name='integerValue', Type='int',
                     description='the value of the receiver’s cell as an NSInteger value.'),
            Property(name='stringValue', Type='string',
                     description='the value of the receiver’s cell as an NSString object'),
            Property(name='cell', Type='appkit.Cell', description='the receiver’s cell object'),
        ],
        methods=[
            Method(name='sizeToFit',
                   description='resizes the receiver’s frame so that it is the minimum size needed to contain its cell')
        ],
    )
    w.generate_code()
