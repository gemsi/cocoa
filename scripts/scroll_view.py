#!env python3

from generate import Component, Property, InitMethod, Param

if __name__ == "__main__":
    w = Component(
        Type="appkit.ScrollView",
        super_type='appkit.View',
        description="a view that displays a portion of a document view and provides scroll bars that allow the user to move the document view within the scroll view",
        init_method=InitMethod(name="initWithFrame", params=[Param(name='frame', Type='foundation.Rect')]),
        properties=[
            Property(name='hasVerticalScroller', Type='bool', getter_prefix_is=False,
                     description='whether the scroll view has a vertical scroller'),
            Property(name='hasHorizontalScroller', Type='bool', getter_prefix_is=False,
                     description='whether the scroll view has a horizontal scroller'),
            Property(name='documentView', Type='appkit.View',
                     description='the view the scroll view scrolls within its content view'),
            Property(name='borderType', Type='uint', go_alias_type='BorderType',
                     description='the appearance of the scroll view’s border'),
            Property(name='contentSize', Type='foundation.Size', readonly=True,
                     description='the size of the scroll view’s content view'),
        ],
    )
    w.generate_code()
