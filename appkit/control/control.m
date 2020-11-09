#import <Appkit/NSControl.h>
#import "control.h"

void Control_SetStringValue(void* ptr, const char* value) {
    NSControl* control = (NSControl*)ptr;
    [control setStringValue:[NSString stringWithUTF8String:value]];
}

const char* Control_StringValue(void* ptr) {
    NSControl* control = (NSControl*)ptr;
    return [[control stringValue] UTF8String];
}