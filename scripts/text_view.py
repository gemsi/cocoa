#!env python3

from generate import Component, Property, init_method, Param, Method, Return


if __name__ == "__main__":
    w = Component(
        Type="appkit.TextView",
        super_type='appkit.Text',
        description="A view that draws text and handles user interactions with that text.",
        properties=[
            Property(name="allowsUndo", Type='bool', getter_prefix_is=False),
            Property(name="textContainer", Type='appkit.TextContainer', description='the receiver’s text container'),
        ],
        methods=[
            init_method(name="initWithFrame", Type="appkit.TextView", params=[Param(name='frame', Type='foundation.Rect')]),
            Method(name='scrollableTextView', return_value=Return('appkit.ScrollView'), static=True),
        ],
        override_methods=[
'''
- (BOOL)performKeyEquivalent:(NSEvent *)event {
    if ([event type] == NSEventTypeKeyDown) {
        if (([event modifierFlags] & NSEventModifierFlagDeviceIndependentFlagsMask) == NSEventModifierFlagCommand) {
            if ([[event charactersIgnoringModifiers] isEqualToString:@"x"]) {
                if ([NSApp sendAction:@selector(cut:) to:nil from:self]) {
                    return YES;
                }
            } else if ([[event charactersIgnoringModifiers] isEqualToString:@"c"]) {
                if ([NSApp sendAction:@selector(copy:) to:nil from:self]) {
                    return YES;
                }
            } else if ([[event charactersIgnoringModifiers] isEqualToString:@"v"]) {
                if ([NSApp sendAction:@selector(paste:) to:nil from:self]) {
                    return YES;
                }
            } else if ([[event charactersIgnoringModifiers] isEqualToString:@"z"]) {
                if ([NSApp sendAction:@selector(undo:) to:nil from:self]) {
                    return YES;
                }
            } else if ([[event charactersIgnoringModifiers] isEqualToString:@"a"]) {
                if ([NSApp sendAction:@selector(selectAll:) to:nil from:self]) {
                    return YES;
                }
            }
        } else if (([event modifierFlags] & NSEventModifierFlagDeviceIndependentFlagsMask) == (NSEventModifierFlagCommand | NSEventModifierFlagShift)) {
            if ([[event charactersIgnoringModifiers] isEqualToString:@"Z"]) {
                if ([NSApp sendAction:@selector(redo:) to:nil from:self]) {
                    return YES;
                }
            }
        }
    }
    return [super performKeyEquivalent:event];
}'''
        ]
    )
    w.generate_code()
