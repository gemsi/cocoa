#!env python3

from generate import Component, Property, InitMethod, Param, Method, Return

if __name__ == "__main__":
    w = Component(
        Type="appkit.SplitView",
        super_type='appkit.View',
        description="a view that arranges two or more views in a linear stack running horizontally or vertically",
        init_method=InitMethod(name="initWithFrame", params=[Param(name='frame', Type='foundation.Rect')]),
        properties=[
            Property(name='arrangesAllSubviews', Type='bool', getter_prefix_is=False,
                     description='whether the split view arranges all of its subviews as split panes'),
            Property(name='vertical', Type='bool',
                     description='the geometric orientation of the split view\'s dividers'),
            Property(name='dividerColor', Type='Color', readonly=True,
                     description='the color of the dividers that the split view draws between subviews'),
            Property(name='dividerThickness', Type='float64', readonly=True,
                     description='the thickness of the dividers for the split view'),
            Property(name='dividerStyle', Type='int', go_alias_type='SplitViewDividerStyle',
                     description='the style of divider between views'),
        ],
        methods=[
            Method(
                name='minPossiblePositionOfDividerAtIndex',
                params=[Param(name='dividerIndex', Type='int')],
                return_value=Return('float64'),
                description='returns the minimum possible position of the divider at the specified index'
            ),
            Method(
                name='maxPossiblePositionOfDividerAtIndex',
                params=[Param(name='dividerIndex', Type='int')],
                return_value=Return('float64'),
                description='returns the maximum possible position of the divider at the specified index'
            ),
            Method(
                name='setPosition',
                params=[
                    Param(name='position', Type='float64'),
                    Param(name='dividerIndex', Type='int', objc_param_name='ofDividerAtIndex'),
                ],
                description='sets the position of the divider at the specified index'
            ),
        ]
    )
    w.generate_code()
