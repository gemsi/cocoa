#!env python3

from generate import Component, Property, Method, Return, Param, init_method


if __name__ == "__main__":
    w = Component(
        Type="appkit.LayoutConstraint",
        super_type='foundation.Object',
        description="the relationship between two user interface objects that must be satisfied by the constraint-based layout system",
        properties=[
            Property(name='active', Type='bool', 
                     description='the active state of the constraint'),
            Property(name='firstItem', Type='foundation.Object', readonly=True,
                     description='the first object participating in the constraint'),
            Property(name='firstAttribute', Type='int', readonly=True, go_alias_type='LayoutAttribute',
                     description='the attribute of the first object participating in the constraint'),
            Property(name='relation', Type='int', readonly=True, go_alias_type='LayoutRelation',
                     description='the relation between the two attributes in the constraint'),
            Property(name='secondItem', Type='foundation.Object', readonly=True,
                     description='the second object participating in the constraint'),
            Property(name='secondAttribute', Type='int', readonly=True, go_alias_type='LayoutAttribute',
                     description='the attribute of the second object participating in the constraint'),
            Property(name='multiplier', Type='float64', readonly=True,
                     description='the multiplier applied to the second attribute participating in the constraint'),
            Property(name='constant', Type='float64', readonly=True,
                     description='the constant added to the multiplied second attribute participating in the constraint'),
            Property(name='firstAnchor', Type='appkit.LayoutAnchor', readonly=True,
                     description='the first anchor that defines the constraint'),
            Property(name='secondAnchor', Type='appkit.LayoutAnchor', readonly=True,
                     description='the second anchor that defines the constraint'),
            Property(name='priority', Type='float32', go_alias_type='LayoutPriority',
                     description='the priority of the constraint'),
            Property(name='identifier', Type='string', description='the name that identifies the constraint'),
            Property(name='shouldBeArchived', Type='bool', getter_prefix_is=False,
                     description='whether the constraint should be archived by its owning view'),
        ],
        methods=[
            Method(
                name='constraintWithItem',
                static=True,
                params=[
                    Param(name='view1', Type='foundation.Object'),
                    Param(name='attr1', Type='int', go_alias='LayoutAttribute', objc_param_name='attribute'),
                    Param(name='relation', Type='int', go_alias='LayoutRelation', objc_param_name='relatedBy'),
                    Param(name='view2', Type='foundation.Object', objc_param_name='toItem'),
                    Param(name='attr2', Type='int', go_alias='LayoutAttribute', objc_param_name='attribute'),
                    Param(name='multiplier', Type='float64', objc_param_name='multiplier'),
                    Param(name='c', Type='float64', objc_param_name='constant'),
                ],
                description='creates a constraint that defines the relationship between the specified attributes of the given views',
            ),
            Method(
                name='activateConstraints',
                static=True,
                params=[
                    Param(name='constraints', Type="appkit.LayoutConstraint", array=True),
                ],
                description='activates each constraint in the specified array',
            ),
            Method(
                name='deactivateConstraints',
                static=True,
                params=[
                    Param(name='constraints', Type="appkit.LayoutConstraint", array=True),
                ],
                description='deactivates each constraint in the specified array',
            ),
        ]
    )
    w.generate_code()
