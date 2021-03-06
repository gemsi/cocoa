#!env python3

from generate import Component, Property, Param, Method, Return, DelegateMethod, init_method


if __name__ == "__main__":
    w = Component(
        Type="appkit.Window",
        super_type='appkit.Responder',
        description="a window that an app displays on the screen.",
        properties=[
            Property(name='title', Type='string',
                     description='the string that appears in the title bar of the window or the path to the represented file'),
            Property(name='contentView', Type='appkit.View',
                     description='the window’s content view, the highest accessible NSView object in the window’s view hierarchy'),
            Property(name='styleMask', Type='uint', go_alias_type='WindowStyleMask', description='the style mask'),
        ],
        methods=[
            init_method(
                name='initWithContentRect',
                Type="appkit.Window",
                params=[
                    Param(name='rect', Type='foundation.Rect'),
                    Param(name='styleMask', Type='uint', go_alias='WindowStyleMask'),
                    Param(name='backing', Type='uint', go_alias='BackingStoreType'),
                    Param(name='Defer', Type='bool', objc_param_name='defer'),
                ],
            ),
            Method(name='center', params=[],
                   description='sets the window’s location to the center of the screen'),
            Method(name='makeKeyAndOrderFront', params=[Param(name='sender', Type='foundation.Object')],
                   description='moves the window to the front of the screen list, within its level, and makes it the key window; that is, it shows the window.'),
            Method(name='makeFirstResponder', params=[Param(name='responder', Type='appkit.Responder')], return_value=Return('bool'),
                   description='attempts to make a given responder the first responder for the window.nil makes the window its first responder.'),
        ],
        delegate_methods=[
            DelegateMethod(
                name='windowDidResize',
                params=[Param(name='notification', Type='foundation.Notification')],
                return_value=Return(''),
                description='the window has been resized'
            ),
            DelegateMethod(
                name='windowDidMove',
                params=[Param(name='notification', Type='foundation.Notification')],
                return_value=Return(''),
                description='the window has moved'
            ),
            DelegateMethod(
                name='windowDidMiniaturize',
                params=[Param(name='notification', Type='foundation.Notification')],
                return_value=Return(''),
                description='the window has been minimized'
            ),
            DelegateMethod(
                name='windowDidDeminiaturize',
                params=[Param(name='notification', Type='foundation.Notification')],
                return_value=Return(''),
                description='the window has been deminimized'
            )
        ]
    )
    w.generate_code()
