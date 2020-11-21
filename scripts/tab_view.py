#!env python3

from .generate import Component, Property, InitMethod, Param, Method, Return, DelegateMethod

if __name__ == "__main__":
    w = Component(
        Type="appkit.TabView",
        super_type='appkit.View',
        description="a multipage interface that displays one page at a time",
        init_method=InitMethod(name="initWithFrame", params=[Param(name='frame', Type='foundation.Rect')]),
        properties=[
            Property(name='numberOfTabViewItems', Type='int', readonly=True,
                     description='the number of items in the tab view’s array of tab view items'),
            Property(name='tabViewType', Type='uint', go_alias_type='TabViewType', description=''),
            Property(name='tabPosition', Type='uint', go_alias_type='TabPosition', description=''),
            Property(name='tabViewBorderType', Type='uint', go_alias_type='TabViewBorderType', description=''),
            Property(name='allowsTruncatedLabels', Type='bool', getter_prefix_is=False,
                     description='if the tab view allows truncating for labels that don’t fit on a tab'),
            Property(name='selectedTabViewItem', Type='appkit.TabViewItem', readonly=True,
                     description='the tab view item for the currently selected tab'),
            Property(name='font', Type='appkit.Font',
                     description='the font used for the tab view’s label text'),
            Property(name='minimumSize', Type='foundation.Size', readonly=True,
                     description='the minimum size necessary for the tab view to display tabs in a useful way'),
            Property(name='controlSize', Type='uint', go_alias_type='ControlSize',
                     description='the size of the tab view'),
        ],
        methods=[
            Method(
                name='addTabViewItem',
                params=[Param(name='tabViewItem', Type='appkit.TabViewItem')],
                description='adds the specified tab item',
            ),
            Method(
                name='insertTabViewItem',
                params=[
                    Param(name='tabViewItem', Type='appkit.TabViewItem'),
                    Param(name='index', Type='int', objc_param_name='atIndex'),
                ],
                description='inserts the specified item into the tab view’s array of tab view items at the specified index',
            ),
            Method(
                name='removeTabViewItem',
                params=[Param(name='tabViewItem', Type='appkit.TabViewItem')],
                description='removes the specified item from the tab view’s array of tab view items',
            ),
            Method(
                name='indexOfTabViewItem',
                params=[Param(name='tabViewItem', Type='appkit.TabViewItem')],
                return_value=Return(Type='int'),
                description='returns the index of the specified item in the tab view',
            ),
            Method(
                name='tabViewItemAtIndex',
                params=[Param(name='index', Type='int')],
                return_value=Return(Type='appkit.TabViewItem'),
                description='returns the tab view item at index in the tab view’s array of items',
            ),
            Method(
                name='selectFirstTabViewItem',
                params=[Param(name='sender', Type='foundation.Object')],
                description='selects the first tab view item',
            ),
            Method(
                name='selectLastTabViewItem',
                params=[Param(name='sender', Type='foundation.Object')],
                description='selects the last tab view item',
            ),
            Method(
                name='selectNextTabViewItem',
                params=[Param(name='sender', Type='foundation.Object')],
                description='selects the next tab view item',
            ),
            Method(
                name='selectPreviousTabViewItem',
                params=[Param(name='sender', Type='foundation.Object')],
                description='selects the previous tab view item',
            ),
            Method(
                name='selectTabViewItem',
                params=[Param(name='tabViewItem', Type='appkit.TabViewItem')],
                description='selects the specified tab view item',
            ),
            Method(
                name='selectTabViewItemAtIndex',
                params=[Param(name='index', Type='int')],
                description='selects the tab view item specified by index',
            ),
        ],
        delegate_methods=[
            DelegateMethod(
                name='tabViewDidChangeNumberOfTabViewItems',
                params=[Param(name='view', Type='appkit.TabView')],
                description='informs the delegate that the number of tab view items in tabView has changed',
            ),
            DelegateMethod(
                name='tabView',
                go_func_name='ShouldSelectTabViewItem',
                params=[
                    Param(name='view', Type='appkit.TabView'),
                    Param(name='tabViewItem', Type='appkit.TabViewItem', objc_param_name='shouldSelectTabViewItem'),
                ],
                return_value=Return(Type='bool'),
                default_return_value='true',
                description='invoked just before tabViewItem in tabView is selected. Return true if the tab view item should be selected, otherwise no.',
            ),
            DelegateMethod(
                name='tabView',
                go_func_name='WillSelectTabViewItem',
                params=[
                    Param(name='view', Type='appkit.TabView'),
                    Param(name='tabViewItem', Type='appkit.TabViewItem', objc_param_name='willSelectTabViewItem'),
                ],
                description='informs the delegate that tabView is about to select tabViewItem',
            ),
            DelegateMethod(
                name='tabView',
                go_func_name='DidSelectTabViewItem',
                params=[
                    Param(name='view', Type='appkit.TabView'),
                    Param(name='tabViewItem', Type='appkit.TabViewItem', objc_param_name='didSelectTabViewItem'),
                ],
                description='informs the delegate that tabView has selected tabViewItem',
            ),
        ],
    )
    w.generate_code()
