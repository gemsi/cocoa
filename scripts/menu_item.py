#!env python3

from generate import Component, Property, InitMethod, Method, Return, Param


if __name__ == "__main__":
    w = Component(
        Type="appkit.MenuItem",
        super_type='foundation.Object',
        description="a command item in an app menu",
        # init_method=InitMethod(name="initWithTitle", params=[Param(name='title', Type='string')]),
        properties=[
            Property(name='enabled', Type='bool', description='whether the menu item is enabled'),
            Property(name='hidden', Type='bool', description='whether the menu item is hidden'),
            Property(name='title', Type='string', description='the menu item\'s title'),
            Property(name='submenu', Type='appkit.Mennu', description='the submenu of the menu item'),
            Property(name='hasSubmenu', Type='bool', getter_prefix_is=False, readonly=True,
                     description='whether the menu item has a submenu'),
            Property(name='separatorItem', Type='bool', readonly=True,
                     description='a menu item that is used to separate logical groups of menu commands'),
            Property(name='menu', Type='appkit.Mennu', description='the menu item’s menu'),
            Property(name='toolTip', Type='string', description='a help tag for the menu item'),
            Property(name='highlighted', Type='bool', readonly=True, description='whether the menu item should be drawn highlighted'),
            

        ],
    )
    w.generate_code()
