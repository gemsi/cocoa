#import <AppKit/AppKit.h>
#import "secure_text_field.h"
#import "secure_text_field_delegate.h"

@implementation GoNSSecureTextFieldDelegate

@end

@interface MySecureTextField : NSSecureTextField {
}
@end

@implementation MySecureTextField

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
}
@end
void SecureTextField_RegisterDelegate(void *ptr, long goID) {
	NSSecureTextField* secureTextField = (NSSecureTextField*)ptr;
	GoNSSecureTextFieldDelegate* delegate = [[GoNSSecureTextFieldDelegate alloc] init];
	[delegate setGoID:goID];
	[secureTextField setDelegate:delegate];
}

void* SecureTextField_NewSecureTextField(NSRect frame) {
	NSSecureTextField* secureTextField = [MySecureTextField alloc];
	return [[secureTextField initWithFrame:frame] autorelease];
}
