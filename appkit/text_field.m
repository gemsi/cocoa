#import <AppKit/AppKit.h>
#import "text_field.h"
#import "text_field_delegate.h"

@implementation GoNSTextFieldDelegate

- (void)controlTextDidChange:(NSNotification*)notification {
	return TextField_Delegate_ControlTextDidChange([self goID], notification);
}

- (void)controlTextDidEndEditing:(NSNotification*)notification {
	return TextField_Delegate_ControlTextDidEndEditing([self goID], notification);
}

- (void)controlTextDidBeginEditing:(NSNotification*)notification {
	return TextField_Delegate_ControlTextDidBeginEditing([self goID], notification);
}

@end

@interface MyTextField : NSTextField {
}
@end

@implementation MyTextField

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
void TextField_RegisterDelegate(void *ptr, long goID) {
	NSTextField* textField = (NSTextField*)ptr;
	GoNSTextFieldDelegate* delegate = [[GoNSTextFieldDelegate alloc] init];
	[delegate setGoID:goID];
	[textField setDelegate:delegate];
}

bool TextField_IsBezeled(void* ptr) {
	NSTextField* textField = (NSTextField*)ptr;
	return [textField isBezeled];
}

void TextField_SetBezeled(void* ptr, bool bezeled) {
	NSTextField* textField = (NSTextField*)ptr;
	[textField setBezeled:bezeled];
}

bool TextField_DrawsBackground(void* ptr) {
	NSTextField* textField = (NSTextField*)ptr;
	return [textField drawsBackground];
}

void TextField_SetDrawsBackground(void* ptr, bool drawsBackground) {
	NSTextField* textField = (NSTextField*)ptr;
	[textField setDrawsBackground:drawsBackground];
}

bool TextField_IsEditable(void* ptr) {
	NSTextField* textField = (NSTextField*)ptr;
	return [textField isEditable];
}

void TextField_SetEditable(void* ptr, bool editable) {
	NSTextField* textField = (NSTextField*)ptr;
	[textField setEditable:editable];
}

bool TextField_IsSelectable(void* ptr) {
	NSTextField* textField = (NSTextField*)ptr;
	return [textField isSelectable];
}

void TextField_SetSelectable(void* ptr, bool selectable) {
	NSTextField* textField = (NSTextField*)ptr;
	[textField setSelectable:selectable];
}

void* TextField_TextColor(void* ptr) {
	NSTextField* textField = (NSTextField*)ptr;
	return [textField textColor];
}

void TextField_SetTextColor(void* ptr, void* textColor) {
	NSTextField* textField = (NSTextField*)ptr;
	[textField setTextColor:(NSColor*)textColor];
}

void* TextField_BackgroundColor(void* ptr) {
	NSTextField* textField = (NSTextField*)ptr;
	return [textField backgroundColor];
}

void TextField_SetBackgroundColor(void* ptr, void* backgroundColor) {
	NSTextField* textField = (NSTextField*)ptr;
	[textField setBackgroundColor:(NSColor*)backgroundColor];
}

void* TextField_NewTextField(NSRect frame) {
	NSTextField* textField = [MyTextField alloc];
	return [[textField initWithFrame:frame] autorelease];
}
