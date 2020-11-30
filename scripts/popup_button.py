#!env python3

from generate import Component, Property, init_method, Param, Method, Return, ActionMethod


if __name__ == "__main__":
    w = Component(
        Type="appkit.PopUpButton",
        super_type='appkit.Button',
        description="a control for selecting an item from a list",
        properties=[
            Property(name='pullsDown', Type='bool', getter_prefix_is=False,
                     description='whether the button displays a pull-down or pop-up menu'),
            Property(name='autoenablesItems', Type='bool', getter_prefix_is=False,
                     description='whether the button enables and disables its items every time a user event occurs'),
            Property(name='selectedItem', Type='appkit.MenuItem', readonly=True,
                     description='the menu item that was last selected by the user'),
            Property(name='titleOfSelectedItem', Type='string', readonly=True,
                     description='the title of the item that was last selected by the user'),
            Property(name='selectedTag', Type='int', readonly=True,
                     description='the tag of the menu item that was last selected by the user'),
            Property(name='menu', Type='appkit.Menu',
                     description='the menu associated with the pop-up button'),
            Property(name='numberOfItems', Type='int', readonly=True, description='the number of items in the menu'),
            Property(name='itemArray', Type='appkit.MenuItem', array=True, readonly=True,
                     description='the array of menu item objects associated with the button'),
            Property(name='itemTitles', Type='string', array=True, readonly=True,
                     description='an array of strings corresponding to the titles of the items in the menu'),
            Property(name='lastItem', Type='appkit.MenuItem', readonly=True, description='the last item in the menu'),
            Property(name='preferredEdge', Type='int', go_alias_type='foundation.RectEdge',
                     description='the edge of the button on which to display the menu when screen space is constrained')

        ],
        methods=[
            init_method(name="initWithFrame", Type='appkit.PopUpButton', params=[
                Param(name='buttonFrame', Type='foundation.Rect'),
                Param(name='flag', Type='bool', objc_param_name='pullsDown'),
            ]),
            Method(
                name='addItemWithTitle',
                params=[Param(name='title', Type='string')],
                description='adds an item with the specified title to the end of the menu',
            ),
            Method(
                name='addItemsWithTitles',
                params=[Param(name='itemTitles', Type='string', array=True)],
                description='adds multiple items to the end of the menu',
            ),
            Method(
                name='insertItemWithTitle',
                params=[Param(name='title', Type='string'), Param(name='index', Type='int', objc_param_name='atIndex')],
                description='inserts an item at the specified position in the menu',
            ),
            Method(
                name='removeAllItems',
                description='removes all items in the receiver’s item menu',
            ),
            Method(
                name='removeItemWithTitle',
                params=[Param(name='title', Type='string')],
                description='removes the item with the specified title from the menu',
            ),
            Method(
                name='removeItemAtIndex',
                params=[Param(name='index', Type='int')],
                description='removes the item at the specified index',
            ),
            Method(
                name='selectItem',
                params=[Param(name='item', Type='appkit.MenuItem')],
                description='selects the specified menu item',
            ),
            Method(
                name='selectItemAtIndex',
                params=[Param(name='index', Type='int')],
                description='selects the item in the menu at the specified index',
            ),
            Method(
                name='selectItemWithTag',
                params=[Param(name='tag', Type='int')],
                return_value=Return(Type='bool'),
                description='selects the menu item with the specified tag',
            ),
            Method(
                name='selectItemWithTitle',
                params=[Param(name='title', Type='string')],
                description='selects the item with the specified title',
            ),
            Method(
                name='itemAtIndex',
                params=[Param(name='index', Type='int')],
                return_value=Return('appkit.MenuItem'),
                description='returns the menu item at the specified index',
            ),
            Method(
                name='itemTitleAtIndex',
                params=[Param(name='index', Type='int')],
                return_value=Return('string'),
                description='returns the title of the item at the specified index',
            ),
            Method(
                name='itemWithTitle',
                params=[Param(name='title', Type='string')],
                return_value=Return('appkit.MenuItem'),
                description='returns the menu item with the specified title',
            ),
            Method(
                name='indexOfItem',
                params=[Param(name='item', Type='appkit.MenuItem')],
                return_value=Return('int'),
                description='returns the index of the specified menu item',
            ),
            Method(
                name='indexOfItemWithTag',
                params=[Param(name='tag', Type='int')],
                return_value=Return('int'),
                description='returns the index of the menu item with the specified tag',
            ),
            Method(
                name='indexOfItemWithTitle',
                params=[Param(name='title', Type='string')],
                return_value=Return('int'),
                description='returns the index of the item with the specified title',
            ),
            # Method(
            #     name='setTitle',
            #     params=[Param(name='title', Type='string')],
            #     description='sets the string displayed in the receiver when the user isn’t pressing the mouse button',
            # ),
            Method(
                name='synchronizeTitleAndSelectedItem',
                description='ensures that the item being displayed by the receiver agrees with the selected item',
            ),
        ],
    )
    w.generate_code()
