#!env python3

from generate import Component, Property, Method, Return, Param, init_method, ActionMethod


if __name__ == "__main__":
    w = Component(
        Type="appkit.MenuItem",
        super_type='foundation.Object',
        description="a command item in an app menu",
        properties=[
            # Property(name='enabled', Type='bool', description='whether the menu item is enabled'),
            # Property(name='hidden', Type='bool', description='whether the menu item is hidden'),
            # Property(name='title', Type='string', description='the menu item\'s title'),
            # Property(name='submenu', Type='appkit.Menu', description='the submenu of the menu item'),
            # Property(name='hasSubmenu', Type='bool', getter_prefix_is=False, readonly=True,
            #          description='whether the menu item has a submenu'),
            # Property(name='separatorItem', Type='bool', readonly=True,
            #          description='a menu item that is used to separate logical groups of menu commands'),
            # Property(name='menu', Type='appkit.Menu', description='the menu item’s menu'),
            # Property(name='toolTip', Type='string', description='a help tag for the menu item'),
            # Property(name='highlighted', Type='bool', readonly=True,
            #          description='whether the menu item should be drawn highlighted'),
            # Property(name='keyEquivalent', Type='string', description='the menu item’s unmodified key equivalent'),
            # Property(name='keyEquivalentModifierMask', Type='uint', go_alias_type='EventModifierFlags',description='the menu item’s unmodified key equivalent'),
            # Property(name='userKeyEquivalent', Type='string', readonly=True, description='Tte user-assigned key equivalent for the menu item'),
            # Property(name='alternate', Type='bool', description='a Boolean value that marks the menu item as an alternate to the previous menu item'),
            # Property(name='indentationLevel', Type='int', description='the menu item indentation level for the menu item'),
            # Property(name='view', Type='appkit.View', description='the content view for the menu item'),
            # Property(name='allowsKeyEquivalentWhenHidden', Type='bool', getter_prefix_is=False, description=''),
            # Property(name='usesUserKeyEquivalents', Type='bool', getter_prefix_is=False, static=True, description='whether menu items conform to user preferences for key equivalents'),
            Property(name='state', Type='int', go_alias_type='ControlStateValue',
                     description='the state of the menu item'),
            Property(name='tag', Type='int', description='the tag of the menu item'),
            Property(name='separatorItem', Type='bool', readonly=True,
                     description='if is a menu item that is used to separate logical groups of menu commands'),
        ],
        methods=[
            # init_method(
            #     name='initWithTitle',
            #     Type='appkit.MenuItem',
            #     params=[
            #         Param(name='title', Type='string'),
            #         Param(name='charCode', Type='string', objc_param_name='keyEquivalent'),
            #     ],
            # ),
            Method(
                name='separatorItem',
                go_func_name='NewSeparatorItem',
                static=True,
                return_value=Return('appkit.MenuItem'),
                description='returns a menu item that is used to separate logical groups of menu commands',
            )
        ],
        action_methods=[
            # ActionMethod('action', ''),
        ]
    )
    w.generate_code()
