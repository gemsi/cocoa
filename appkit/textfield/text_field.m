#import <Appkit/NSTextField.h>
#import <Appkit/NSSecureTextField.h>
#import "text_field.h"
#include "_cgo_export.h"

@interface GoTextFieldDelegate : NSObject <NSTextFieldDelegate>

@property (assign) long goID;

@end

@implementation GoTextFieldDelegate

- (void)controlTextDidChange:(NSNotification*)notification {
	onTextDidChange([self goID], notification);
}

- (void)controlTextDidEndEditing:(NSNotification*)notification {
	onTextDidEndEditing([self goID], notification);
}

- (void)onEnterKey:(id)sender {
	onEnterKey([self goID], sender);
}

@end

void* TextField_New(long id) {
    NSTextField* textField = [[[NSTextField alloc] init] autorelease];
    [textField.cell setWraps:NO];
    [textField.cell setScrollable:YES];
    GoTextFieldDelegate* delegate = [[GoTextFieldDelegate alloc] init];
    [delegate setGoID:id];
    [textField setDelegate:delegate];
    return (void*)textField;
}

void* SecureTextField_New(long id) {
    NSSecureTextField* textField = [[[NSSecureTextField alloc] init] autorelease];
    [textField.cell setWraps:NO];
    [textField.cell setScrollable:YES];
    GoTextFieldDelegate* delegate = [[GoTextFieldDelegate alloc] init];
    [delegate setGoID:id];
    [textField setDelegate:delegate];
    return (void*)textField;
}

void TextField_SetBezeled(void* ptr, bool bezeled) {
    NSTextField* textField = (NSTextField*)ptr;
    [textField setBezeled:bezeled];
}

void TextField_SetEditable(void* ptr, bool editable) {
    NSTextField* textField = (NSTextField*)ptr;
    [textField setEditable:editable];
}

void TextField_SetSelectable(void* ptr, bool selectable) {
    NSTextField* textField = (NSTextField*)ptr;
    [textField setSelectable:selectable];
}

void TextField_SetDrawsBackground(void* ptr, bool draws) {
    NSTextField* textField = (NSTextField*)ptr;
    [textField setDrawsBackground:draws];
}

void* TextField_TextColor(void* ptr) {
    NSTextField* textField = (NSTextField*)ptr;
    return [textField textColor];
}

void TextField_SetTextColor(void* ptr, void* colorPtr) {
    NSTextField* textField = (NSTextField*)ptr;
    NSColor* color = (NSColor*)colorPtr;
    [textField setTextColor:color];
}

void* TextField_BackgroundColor(void* ptr) {
    NSTextField* textField = (NSTextField*)ptr;
    return [textField backgroundColor];
}

void TextField_SetBackgroundColor(void* ptr, void* colorPtr) {
    NSTextField* textField = (NSTextField*)ptr;
    NSColor* color = (NSColor*)colorPtr;
    [textField setBackgroundColor:color];
}