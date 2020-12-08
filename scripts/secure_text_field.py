#!env python3

from generate import Component, Param, init_method

if __name__ == "__main__":
    w = Component(
        Type="appkit.SecureTextField",
        super_type='appkit.TextField',
        delegate_type='TextFieldDelegate',
        description="a text field that hides the typed text",
        properties=[],
        extend_delegate=True,
        methods=[
            init_method(name="initWithFrame", Type="appkit.SecureTextField",
                        params=[Param(name='frame', Type='foundation.Rect')]),
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
}'''])
    w.generate_code()
