#import <Appkit/NSTextField.h>
#import "text_field.h"
#import "text_field_delegate.h"

void* TextField_initWithFrame(long goID, NSRect frame) {
	NSTextField* text_field = [[[NSTextField alloc] initWithFrame:frame] autorelease];
	GoNSTextFieldDelegate* delegate = [[GoNSTextFieldDelegate alloc] init];
	[delegate setGoID:goID];
	[text_field setDelegate:delegate];
	return (void*)text_field;
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
	[textField setTextColor:textColor];
}

void* TextField_BackgroundColor(void* ptr) {
	NSTextField* textField = (NSTextField*)ptr;
	return [textField backgroundColor];
}

void TextField_SetBackgroundColor(void* ptr, void* backgroundColor) {
	NSTextField* textField = (NSTextField*)ptr;
	[textField setBackgroundColor:backgroundColor];
}
