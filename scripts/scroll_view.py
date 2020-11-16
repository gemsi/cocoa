#!env python3

from generate_widget import Generator, Property, InitMethod


if __name__ == "__main__":
    generator = Generator(
        Type = "ScrollView", 
        super_type = 'View',
        description = "A view that displays a portion of a document view and provides scroll bars that allow the user to move the document view within the scroll view.",
        init_method=InitMethod(param_name='frame', param_type='foundation.Rect'),
        properties = [
            Property(name='hasVerticalScroller', Type = 'bool', prefixIs=False, description='whether the scroll view has a vertical scroller'),
            Property(name='hasHorizontalScroller', Type = 'bool', prefixIs=False, description='whether the scroll view has a horizontal scroller'),
            Property(name='documentView', Type = 'View', description='the view the scroll view scrolls within its content view'),
            Property(name='borderType', Type = 'uint64', go_alias_type='BorderType', description='the appearance of the scroll view’s border'),
            Property(name='contentSize', Type = 'foundation.Size', readonly=True, description='the size of the scroll view’s content view'),
        ],
    )
    generator.generate_widgets()