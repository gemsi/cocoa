#!env python3

from .generate import Component, Property, InitMethod


if __name__ == "__main__":
    w = Component(
        Type = "appkit.Cell", 
        super_type = 'foundation.Object',
        description = "a mechanism for displaying text or images in a view object without the overhead of a full NSView subclass",
        properties = [
            Property(name='wraps', Type = 'bool', getter_prefix_is=False, description='whether the cell wraps text whose length that exceeds the cell’s frame'),
            Property(name='scrollable', Type = 'bool', description='whether excess text scrolls past the cell’s bounds'),
            Property(name='editable', Type = 'bool', description='whether the cell is editable'),
            Property(name='selectable', Type = 'bool', description='whether the cell’s text can be selected'),
        ],
    )
    w.generate_code()
