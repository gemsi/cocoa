#!env python3

from generate import Component, Property, Param, Method, Return, DelegateMethod, init_method


if __name__ == "__main__":
    w = Component(
        Type="appkit.Popover",
        super_type='appkit.Responder',
        description="a means to display additional content related to existing content on the screen",
        properties=[
            Property(name='behavior', Type='int', go_alias_type='PopoverBehavior',
                     description='the behavior of the popover'),
            Property(name='positioningRect', Type='foundation.Rect',
                     description='the rectangle within the positioning view relative to which the popover should be positioned'),
            Property(name='animates', Type='bool', getter_prefix_is=False,
                     description='if the popover is to be animated'),
            Property(name='contentSize', Type='foundation.Size', description='the content size of the popover'),
            Property(name='shown', Type='bool', readonly=True, description='the display state of the popover'),
            Property(name='detached', Type='bool', readonly=True,
                     description='whether the window created by a popover\'s detachment is automatically created'),
            Property(name='appearance', Type='appkit.Appearance', description='the appearance of the popover'),
            Property(name='effectiveAppearance', Type='appkit.Appearance', readonly=True,
                     description='the appearance that will be used when the popover is displayed onscreen'),
        ],
        methods=[
            init_method(
                name='init',
                Type="appkit.Popover",
            ),
            Method(
                name='performClose',
                params=[Param(name='sender', Type='foundation.Object')],
                description='attempts to close the popover',
            ),
            Method(
                name='close',
                description='forces the popover to close without consulting its delegate',
            ),
            Method(
                name='showRelativeToRect',
                go_func_name='ShowRelativeTo',
                params=[
                    Param(name='positioningRect', Type='foundation.Rect'),
                    Param(name='positioningView', Type='appkit.View', objc_param_name='ofView'),
                    Param(name='preferredEdge', Type='int', go_alias='foundation.RectEdge', objc_param_name='preferredEdge'),
                ],
                description='shows the popover anchored to the specified view',
            )
        ],

    )
    w.generate_code()
