#import <Cocoa/Cocoa.h>
#import "text_field.h"
#include "_cgo_export.h"

@interface GoTextField : NSTextField <NSTextFieldDelegate>

@property (assign) long goID;

@end

@implementation GoTextField

- (void)controlTextDidChange:(NSNotification*)notification {
	onTextDidChange([self goID], notification);
}

- (BOOL)becomeFirstResponder {
	BOOL rc = [super becomeFirstResponder];
	if (rc) {
	    onBecameFirstResponder([self goID]);
	}
	return rc;
}

- (void)controlTextDidEndEditing:(NSNotification*)notification {
	onTextDidEndEditing([self goID], notification);
}

- (void)onEnterKey:(id)sender {
	onEnterKey([self goID], sender);
}

@end

void* TextField_New(long id) {
    GoTextField* textField = [[[GoTextField alloc] init] autorelease];
    [textField setGoID:id];
    [textField setDelegate:textField];
    return (void*)textField;
}

void TextField_SetBezeled(void* ptr, int bezeled) {
    NSTextField* textField = (NSTextField*)ptr;
    [textField setBezeled:bezeled==0?NO:YES];
}

void TextField_SetEditable(void* ptr, int editable) {
    NSTextField* textField = (NSTextField*)ptr;
    [textField setEditable:editable==0?NO:YES];
}

void TextField_SetSelectable(void* ptr, int selectable) {
    NSTextField* textField = (NSTextField*)ptr;
    [textField setSelectable:selectable==0?NO:YES];
}

void TextField_SetDrawsBackground(void* ptr, int draws) {
    NSTextField* textField = (NSTextField*)ptr;
    [textField setDrawsBackground:draws==0?NO:YES];
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