#!env python3

from generate import Component, Property, init_method, Param, Method, Return

if __name__ == "__main__":
    w = Component(
        Type="appkit.StackView",
        super_type='appkit.View',
        description="a view that arranges an array of views horizontally or vertically and updates their placement and sizing when the window size changes",
        properties=[
            Property(name='views', Type='appkit.View', readonly=True, array=True,
                     description='the array of views owned by the stack view'),
            Property(name='detachedViews', Type='appkit.View', readonly=True, array=True,
                     description='the detached views from all the stack view’s gravity areas'),
            Property(name='orientation', Type='int', go_alias_type='UserInterfaceLayoutOrientation',
                     description='the horizontal or vertical layout direction of the stack view'),
            Property(name='alignment', Type='int', go_alias_type='LayoutAttribute',
                     description='the view alignment within the stack view'),
            Property(name='spacing', Type='float64',
                     description='the minimum spacing, in points, between adjacent views in the stack view'),
            Property(name='edgeInsets', Type='foundation.EdgeInsets',
                     description='the geometric padding, in points, inside the stack view, surrounding its views'),
            Property(name='distribution', Type='int', go_alias_type='StackViewDistribution',
                     description='distribution'),
            Property(name='arrangedSubviews', Type='appkit.View', readonly=True, array=True,
                     description='the array of views arranged by the stack view'),
            Property(name='detachesHiddenViews', Type='bool', getter_prefix_is=False,
                     description='whether the stack view removes hidden views from its view hierarchy'),

        ],
        methods=[
            Method(
                name='stackViewWithViews',
                register_delegate=True,
                static=True,
                params=[Param(name='views', Type='appkit.View', array=True)],
                return_value=Return(Type='appkit.StackView'),
                description='return a stack view with views',
            ),
            Method(
                name='viewsInGravity',
                params=[Param(name='gravity', Type='int', go_alias='StackViewGravity')],
                return_value=Return(Type='appkit.View', array=True),
                description='returns the array of views in the specified gravity area in the stack view',
            ),
            Method(
                name='addView',
                params=[
                    Param(name='view', Type='appkit.View'),
                    Param(name='gravity', Type='int', go_alias='StackViewGravity', objc_param_name='inGravity'),
                ],
                description='adds a view to the end of the stack view gravity area',
            ),
            Method(
                name='insertView',
                params=[
                    Param(name='view', Type='appkit.View'),
                    Param(name='index', Type='uint', objc_param_name='atIndex'),
                    Param(name='gravity', Type='int', go_alias='StackViewGravity', objc_param_name='inGravity'),
                ],
                description='adds a view to a stack view gravity area at a specified index position',
            ),
            Method(
                name='setViews',
                params=[
                    Param(name='views', Type='appkit.View', array=True),
                    Param(name='gravity', Type='int', go_alias='StackViewGravity', objc_param_name='inGravity'),
                ],
                description='specifies an array of views for a specified gravity area in the stack view, replacing any previous views in that area',
            ),
            Method(
                name='removeView',
                params=[
                    Param(name='view', Type='appkit.View'),
                ],
                description='removes a specified view from the stack view',
            ),
            Method(
                name='addArrangedSubview',
                params=[
                    Param(name='view', Type='appkit.View'),
                ],
                description='adds the specified view to the end of the arranged subviews list',
            ),
            Method(
                name='insertArrangedSubview',
                params=[
                    Param(name='view', Type='appkit.View'),
                    Param(name='index', Type='uint', objc_param_name='atIndex'),
                ],
                description='adds the provided view to the array of arranged subviews at the specified index',
            ),
            Method(
                name='removeArrangedSubview',
                params=[
                    Param(name='view', Type='appkit.View'),
                ],
                description='removes the provided view from the stack’s array of arranged subviews',
            ),
            Method(
                name='clippingResistancePriorityForOrientation',
                params=[
                    Param(name='orientation', Type='int', go_alias='LayoutConstraintOrientation'),
                ],
                return_value=Return(Type='float32', go_alias='LayoutPriority'),
                description='returns the Auto Layout priority for resisting clipping of views in the stack view when Auto Layout attempts to reduce the stack view’s size',
            ),
            Method(
                name='setClippingResistancePriority',
                params=[
                    Param(name='clippingResistancePriority', Type='float32', go_alias='LayoutPriority'),
                    Param(name='orientation', Type='int', go_alias='LayoutConstraintOrientation',
                          objc_param_name='forOrientation'),
                ],
                description='sets the Auto Layout priority for resisting clipping of views in the stack view when Auto Layout attempts to reduce the stack view’s size',
            ),
            Method(
                name='huggingPriorityForOrientation',
                params=[
                    Param(name='orientation', Type='int', go_alias='LayoutConstraintOrientation'),
                ],
                return_value=Return(Type='float32', go_alias='LayoutPriority'),
                description='returns the Auto Layout priority for the stack view to minimize its size to fit its contained views as closely as possible, for a specified user interface axis',
            ),
            Method(
                name='setHuggingPriority',
                params=[
                    Param(name='huggingPriority', Type='float32', go_alias='LayoutPriority'),
                    Param(name='orientation', Type='int', go_alias='LayoutConstraintOrientation',
                          objc_param_name='forOrientation'),
                ],
                description='sets the Auto Layout priority for the stack view to minimize its size, for a specified user interface axis',
            ),
            Method(
                name='customSpacingAfterView',
                params=[
                    Param(name='view', Type='appkit.View'),
                ],
                return_value=Return(Type='float64'),
                description='returns the custom spacing, in points, between a specified view in the stack view and the view that follows it',
            ),
            Method(
                name='setCustomSpacing',
                params=[
                    Param(name='spacing', Type='float64'),
                    Param(name='view', Type='appkit.View', objc_param_name='afterView'),
                ],
                description='specifies the custom spacing, in points, between a specified view and the view that follows it in the stack view',
            ),
            Method(
                name='visibilityPriorityForView',
                params=[
                    Param(name='view', Type='appkit.View'),
                ],
                return_value=Return(Type='float32', go_alias='StackViewVisibilityPriority'),
                description='returns the visibility priority for a specified view in the stack view',
            ),
            Method(
                name='setVisibilityPriority',
                params=[
                    Param(name='priority', Type='float32', go_alias='StackViewVisibilityPriority'),
                    Param(name='view', Type='appkit.View', objc_param_name='forView'),
                ],
                description='sets the Auto Layout priority for a view to remain attached to the stack view when Auto Layout reduces the stack view’s size',
            ),

        ]
    )
    w.generate_code()
