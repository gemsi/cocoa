#!env python3

from generate_widget import Generator, Property, InitMethod


if __name__ == "__main__":
    generator = Generator(
        Type = "Text", 
        super_type = 'View',
        description = "The most general programmatic interface for objects that manage text.",
        init_method=InitMethod(param_name='frame', param_type='foundation.Rect'),
        properties = [
            Property(name="string", Type='string', description='the characters of the receiver’s text'),
            Property(name="editable", Type='bool', description='whether the receiver allows the user to edit its text'),
            Property(name="selectable", Type='bool', description='whether the receiver allows the user to select its text'),
            Property(name="fieldEditor", Type='bool', description='whether the receiver interprets Tab, Shift-Tab, and Return (Enter) as cues to end editing and possibly to change the first responder'),
            Property(name="richText", Type='bool', description='whether the receiver allows the user to apply attributes to specific ranges of the text'),
            Property(name="minSize", Type='foundation.Size', description='the receiver’s minimum size'),
            Property(name="maxSize", Type='foundation.Size', description='the receiver’s maximum size'),
            Property(name="verticallyResizable", Type='bool', description='whether the receiver changes its height to fit the height of its text.'),
            Property(name="horizontallyResizable", Type='bool', description='whether the receiver changes its width to fit the width of its text'),
        ],
    )
    generator.generate_widgets()