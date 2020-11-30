#!env python3

from generate import Component, Property, Method, Return, Param, init_method


if __name__ == "__main__":
    w = Component(
        Type="appkit.LayoutXAxisAnchor",
        super_type='appkit.LayoutAnchor',
        objc_imports=['Appkit/NSLayoutAnchor'],
        description="a factory class for creating horizontal layout constraint objects using a fluent API",
        properties=[],
        methods=[
            Method(
                name='constraintEqualToSystemSpacingAfterAnchor',
                params=[
                    Param(name='anchor', Type='appkit.LayoutXAxisAnchor'),
                    Param(name='multiplier', Type='float64', objc_param_name='multiplier'),
                ],
                return_value=Return(Type='appkit.LayoutConstraint'),
                description='returns a constraint that defines by how much the current anchor trails the specified anchor',
            ),
            Method(
                name='constraintGreaterThanOrEqualToSystemSpacingAfterAnchor',
                params=[
                    Param(name='anchor', Type='appkit.LayoutXAxisAnchor'),
                    Param(name='multiplier', Type='float64', objc_param_name='multiplier'),
                ],
                return_value=Return(Type='appkit.LayoutConstraint'),
                description='returns a constraint that defines the minimum amount by which the current anchor trails the specified anchor',
            ),
            Method(
                name='constraintLessThanOrEqualToSystemSpacingAfterAnchor',
                params=[
                    Param(name='anchor', Type='appkit.LayoutXAxisAnchor'),
                    Param(name='multiplier', Type='float64', objc_param_name='multiplier'),
                ],
                return_value=Return(Type='appkit.LayoutConstraint'),
                description='returns a constraint that defines the maximum amount by which the current anchor trails the specified anchor',
            ),
            Method(
                name='anchorWithOffsetToAnchor',
                params=[
                    Param(name='otherAnchor', Type='appkit.LayoutXAxisAnchor'),
                ],
                return_value=Return(Type='appkit.LayoutDimension'),
                description='creates a layout dimension object from two anchors',
            ),
            
        ]
    )
    w.generate_code()
