#!env python3

from generate import Component, Property, init_method, Param, Method, Return, DelegateMethod


if __name__ == "__main__":
    w = Component(
        Type="appkit.TextField",
        super_type='appkit.Control',
        description="text that the user can select or edit and that sends its action message to its target when the user presses the Return key",
        properties=[
            Property(name='bezeled', Type='bool',
                     description='whether the receiver draws a bezeled border around its contents'),
            Property(name='drawsBackground', Type='bool', getter_prefix_is=False, description='whether the receiver’s cell draws its background color behind its text'),
            Property(name='editable', Type='bool', description='whether the user can edit the value in the text field'),
            Property(name='selectable', Type='bool', description='whether the receiver is selectable (but not editable)'),
            Property(name='textColor', Type='appkit.Color', description='the color used to draw the receiver’s text'),
            Property(name='backgroundColor', Type='appkit.Color', description='the color of the background that the receiver’s cell draws behind the text.'),
        ],
        methods=[ 
            init_method(name="initWithFrame", Type="appkit.TextField", params=[Param(name='frame', Type='foundation.Rect')]),
        ],
        delegate_methods=[
            DelegateMethod(
                name='controlTextDidChange',
                params=[Param(name='notification', Type='foundation.Notification')],
                description=''
            ),
            DelegateMethod(
                name='controlTextDidEndEditing',
                params=[Param(name='notification', Type='foundation.Notification')],
                description=''
            ),
            DelegateMethod(
                name='controlTextDidBeginEditing',
                params=[Param(name='notification', Type='foundation.Notification')],
                description=''
            ),
        ],
        override_methods=[
'''- (BOOL)performKeyEquivalent:(NSEvent *)event {
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
