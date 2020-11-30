#!env python3

from generate import Component, Property, Method, Return, Param, init_method


if __name__ == "__main__":
    w = Component(
        Type="appkit.LayoutDimension",
        super_type='appkit.LayoutAnchor',
        objc_imports=['Appkit/NSLayoutAnchor'],
        description="a factory class for creating size-based layout constraint objects using a fluent API",
        properties=[],
        methods=[
            Method(
                name='constraintEqualToConstant',
                params=[
                    Param(name='constant', Type='float64', objc_param_name='constant'),
                ],
                return_value=Return(Type='appkit.LayoutConstraint'),
                description='returns a constraint that defines a constant size for the anchor’s size attribute.',
            ),
            Method(
                name='constraintEqualToAnchor',
                go_func_name='ConstraintEqualToAnchor3',
                params=[
                    Param(name='anchor', Type='appkit.LayoutDimension'),
                    Param(name='multiplier', Type='float64', objc_param_name='multiplier'),
                    Param(name='constant', Type='float64', objc_param_name='constant'),
                ],
                return_value=Return(Type='appkit.LayoutConstraint'),
                description='returns a constraint that defines the anchor’s size attribute as equal to the specified size attribute multiplied by a constant plus an offset',
            ),
            Method(
                name='constraintGreaterThanOrEqualToConstant',
                params=[
                    Param(name='constant', Type='float64', objc_param_name='constant'),
                ],
                return_value=Return(Type='appkit.LayoutConstraint'),
                description='returns a constraint that defines the minimum size for the anchor’s size attribute',
            ),
            Method(
                name='constraintGreaterThanOrEqualToAnchor',
                go_func_name='ConstraintGreaterThanOrEqualToAnchor3',
                params=[
                    Param(name='anchor', Type='appkit.LayoutDimension'),
                    Param(name='multiplier', Type='float64', objc_param_name='multiplier'),
                    Param(name='constant', Type='float64', objc_param_name='constant'),
                ],
                return_value=Return(Type='appkit.LayoutConstraint'),
                description='returns a constraint that defines the anchor’s size attribute as greater than or equal to the specified anchor multiplied by the constant',
            ),
            Method(
                name='constraintLessThanOrEqualToConstant',
                params=[
                    Param(name='constant', Type='float64', objc_param_name='constant'),
                ],
                return_value=Return(Type='appkit.LayoutConstraint'),
                description='returns a constraint that defines the maximum size for the anchor’s size attribute',
            ),
            Method(
                name='constraintLessThanOrEqualToAnchor',
                go_func_name='ConstraintLessThanOrEqualToAnchor3',
                params=[
                    Param(name='anchor', Type='appkit.LayoutDimension'),
                    Param(name='multiplier', Type='float64', objc_param_name='multiplier'),
                    Param(name='constant', Type='float64', objc_param_name='constant'),
                ],
                return_value=Return(Type='appkit.LayoutConstraint'),
                description='returns a constraint that defines the anchor’s size attribute as greater than or equal to the specified anchor multiplied by the constant plus an offset',
            ),
        ]
    )
    w.generate_code()
