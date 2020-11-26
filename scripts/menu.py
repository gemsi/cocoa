#!env python3

from generate import Component, Property, init_method, Method, Return, Param


if __name__ == "__main__":
    w = Component(
        Type="appkit.Menu",
        super_type='foundation.Object',
        description="an object that manages an app’s menus",
        properties=[
            Property(name='menuBarHeight', Type='float64', readonly=True,
                     description='the menu bar height for the main menu in pixels'),
            Property(name='font', Type='appkit.Font', description='the font of the menu and its submenus'),
            Property(name='autoenablesItems', Type='bool', getter_prefix_is=False,
                     description='whether the menu automatically enables and disables its menu items'),
            Property(name='title', Type='string', description='the title of the menu'),
            Property(name='minimumWidth', Type='float64',
                     description='the minimum width of the menu in screen coordinates'),
            Property(name='size', Type='foundation.Size', readonly=True,
                     description='the size of the menu in screen coordinates'),
            Property(name='propertiesToUpdate', Type='uint', go_alias_type='MenuProperties', readonly=True,
                     description='the size of the menu in screen coordinates'),
            Property(name='allowsContextMenuPlugIns', Type='bool', getter_prefix_is=False,
                     description='whether the pop-up menu allows appending of contextual menu plug-in items'),
            Property(name='showsStateColumn', Type='bool', getter_prefix_is=False,
                     description='whether the menu displays the state column'),
            Property(name='highlightedItem', Type='appkit.MenuItem', readonly=True,
                     description='indicates the currently highlighted item in the menu'),
            Property(name='userInterfaceLayoutDirection', Type='int', go_alias_type='UserInterfaceLayoutDirection',
                     description='the layout direction of menu items in the menu'),
        ],
        methods= [
        init_method(name="initWithTitle", Type='appkit.Menu', params=[Param(name='title', Type='string')]),
        ],
    )
    w.generate_code()
