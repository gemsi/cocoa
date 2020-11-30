#!env python3

from generate import Component, Property, Method, Return, Param, init_method


if __name__ == "__main__":
    w = Component(
        Type="appkit.LayoutYAxisAnchor",
        super_type='appkit.LayoutAnchor',
        objc_imports=['Appkit/NSLayoutAnchor'],
        description="a factory class for creating vertical layout constraint objects using a fluent API",
        properties=[],
        methods=[
            Method(
                name='constraintEqualToSystemSpacingBelowAnchor',
                params=[
                    Param(name='anchor', Type='appkit.LayoutYAxisAnchor'),
                    Param(name='multiplier', Type='float64', objc_param_name='multiplier'),
                ],
                return_value=Return(Type='appkit.LayoutConstraint'),
                description='returns a constraint that defines by how much the current anchor below the specified anchor',
            ),
            Method(
                name='constraintGreaterThanOrEqualToSystemSpacingBelowAnchor',
                params=[
                    Param(name='anchor', Type='appkit.LayoutYAxisAnchor'),
                    Param(name='multiplier', Type='float64', objc_param_name='multiplier'),
                ],
                return_value=Return(Type='appkit.LayoutConstraint'),
                description='returns a constraint that defines the minimum amount by which the current anchor below the specified anchor',
            ),
            Method(
                name='constraintLessThanOrEqualToSystemSpacingBelowAnchor',
                params=[
                    Param(name='anchor', Type='appkit.LayoutYAxisAnchor'),
                    Param(name='multiplier', Type='float64', objc_param_name='multiplier'),
                ],
                return_value=Return(Type='appkit.LayoutConstraint'),
                description='returns a constraint that defines the maximum amount by which the current anchor below the specified anchor',
            ),
            Method(
                name='anchorWithOffsetToAnchor',
                params=[
                    Param(name='otherAnchor', Type='appkit.LayoutYAxisAnchor'),
                ],
                return_value=Return(Type='appkit.LayoutYAxisAnchor'),
                description='creates a layout dimension object from two anchors',
            ),
            
        ]
    )
    w.generate_code()
