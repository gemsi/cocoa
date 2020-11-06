#import "control.h"

void Control_SetStringValue(void* ptr, const char* value) {
    NSTextField* textField = (NSTextField*)ptr;
    [textField setStringValue:[NSString stringWithUTF8String:value]];
}

const char* Control_StringValue(void* ptr) {
    NSTextField* textField = (NSTextField*)ptr;
    return [[textField stringValue] UTF8String];
}