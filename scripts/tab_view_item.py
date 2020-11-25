#!env python3

from generate import Component, Property, InitMethod, Param

if __name__ == "__main__":
    w = Component(
        Type="appkit.TabViewItem",
        super_type='foundation.Object',
        description="an item in a tab view",
        init_method=InitMethod(name="initWithIdentifier", params=[Param(name='identifier', Type='foundation.Object')]),
        properties=[
            Property(name='label', Type='string',
                     description='the label text for the receiver to label'),
            Property(name='toolTip', Type='string',
                     description='the tooltip displayed for the tab view item'),
            Property(name='tabState', Type='uint', go_alias_type='TabState', readonly=True,
                     description='the current display state of the tab associated with the receiver'),
            Property(name='identifier', Type='foundation.Object',
                     description='the receiver’s optional identifier object to identifier'),
            Property(name='color', Type='appkit.Color',
                     description='the background color for content in the view'),
            Property(name='view', Type='appkit.View',
                     description='the view associated with the receiver to view'),
            Property(name='initialFirstResponder', Type='appkit.View',
                     description='the initial first responder for the view associated with the receiver (the view that is displayed when a user clicks on the tab) to view'),
            Property(name='tabView', Type='appkit.TabView', readonly=True,
                     description='the parent tab view for the receiver'),

        ],
    )
    w.generate_code()
