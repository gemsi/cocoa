#import "text_field.h"
#import "text_field_delegate.h"
#include "_cgo_export.h"

void* TextField_New(int id) {
    NSTextField* textField = [[[NSTextField alloc] init] autorelease];
    TextFieldDelegate* textFieldDelegate = [[TextFieldDelegate alloc] init];
    [textFieldDelegate setGoID:id];
    [textField setDelegate:textFieldDelegate];
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