#!env python3

from generate import Component, Property, Method, Return, Param, init_method


if __name__ == "__main__":
    w = Component(
        Type="appkit.LayoutAnchor",
        super_type='foundation.Object',
        objc_imports=['Appkit/NSLayoutAnchor', 'Appkit/NSLayoutConstraint'],
        description="a factory class for creating layout constraint objects using a fluent API",
        properties=[
            Property(name='constraintsAffectingLayout', Type='appkit.LayoutConstraint', array=True, readonly=True,
                     description='the constraints that impact the layout of the anchor'),
            Property(name='hasAmbiguousLayout', Type='bool', readonly=True, getter_prefix_is=False,
                     description='whether the constraints impacting the anchor specify its location ambiguously'),
            Property(name='name', Type='string', readonly=True,
                     description='the name assigned to the anchor for debugging purposes'),
            Property(name='item', Type='foundation.Object', readonly=True,
                     description='the layout item used to calculate the anchor\'s position'),
                     

        ],
        methods=[
            Method(
                name='constraintEqualToAnchor',
                params=[
                    Param(name='anchor', Type='appkit.LayoutAnchor'),
                ],
                return_value=Return(Type='LayoutConstraint'),
                description='returns a constraint that defines one item’s attribute as equal to another',
            ),
            Method(
                name='constraintEqualToAnchor',
                go_func_name='ConstraintEqualToAnchor2',
                params=[
                    Param(name='anchor', Type='appkit.LayoutAnchor'),
                    Param(name='constant', Type='float64', objc_param_name='constant'),
                ],
                return_value=Return(Type='LayoutConstraint'),
                description='returns a constraint that defines one item’s attribute as equal to another item’s attribute plus a constant offset',
            ),
            Method(
                name='constraintGreaterThanOrEqualToAnchor',
                params=[
                    Param(name='anchor', Type='appkit.LayoutAnchor'),
                ],
                return_value=Return(Type='LayoutConstraint'),
                description='returns a constraint that defines one item’s attribute as greater than or equal to another',
            ),
            Method(
                name='constraintGreaterThanOrEqualToAnchor',
                go_func_name='ConstraintGreaterThanOrEqualToAnchor2',
                params=[
                    Param(name='anchor', Type='appkit.LayoutAnchor'),
                    Param(name='constant', Type='float64', objc_param_name='constant'),
                ],
                return_value=Return(Type='LayoutConstraint'),
                description='returns a constraint that defines one item’s attribute as greater than or equal to another item’s attribute plus a constant offset',
            ),
            Method(
                name='constraintLessThanOrEqualToAnchor',
                params=[
                    Param(name='anchor', Type='appkit.LayoutAnchor'),
                ],
                return_value=Return(Type='LayoutConstraint'),
                description='returns a constraint that defines one item’s attribute as less than or equal to another',
            ),
            Method(
                name='constraintLessThanOrEqualToAnchor',
                go_func_name='ConstraintLessThanOrEqualToAnchor2',
                params=[
                    Param(name='anchor', Type='appkit.LayoutAnchor'),
                    Param(name='constant', Type='float64', objc_param_name='constant'),
                ],
                return_value=Return(Type='LayoutConstraint'),
                description='returns a constraint that defines one item’s attribute as less than or equal to another item’s attribute plus a constant offset',
            ),
        ]
    )
    w.generate_code()
