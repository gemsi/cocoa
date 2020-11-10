#import <Appkit/NSControl.h>
        #import "control.h"
        #include "_cgo_export.h"


void* Control_New(long id, NSRect frame) {
    NSControl* control = [[[NSControl alloc] initWithFrame:frame] autorelease];
    return (void*)control;
}

bool Control_Enabled(void* ptr) {
    NSControl* control = (NSControl*)ptr;
    return control.enabled;
}

void Control_SetEnabled(void* ptr, bool value) {
    NSControl* control = (NSControl*)ptr;
    return [control setEnabled:value];
}

double Control_DoubleValue(void* ptr) {
    NSControl* control = (NSControl*)ptr;
    return control.doubleValue;
}

void Control_SetDoubleValue(void* ptr, double value) {
    NSControl* control = (NSControl*)ptr;
    return [control setDoubleValue:value];
}

float Control_FloatValue(void* ptr) {
    NSControl* control = (NSControl*)ptr;
    return control.floatValue;
}

void Control_SetFloatValue(void* ptr, float value) {
    NSControl* control = (NSControl*)ptr;
    return [control setFloatValue:value];
}

int Control_IntValue(void* ptr) {
    NSControl* control = (NSControl*)ptr;
    return control.intValue;
}

void Control_SetIntValue(void* ptr, int value) {
    NSControl* control = (NSControl*)ptr;
    return [control setIntValue:value];
}

long Control_IntegerValue(void* ptr) {
    NSControl* control = (NSControl*)ptr;
    return control.integerValue;
}

void Control_SetIntegerValue(void* ptr, long value) {
    NSControl* control = (NSControl*)ptr;
    return [control setIntegerValue:value];
}

const char* Control_StringValue(void* ptr) {
    NSControl* control = (NSControl*)ptr;
    return [control.stringValue UTF8String];
}

void Control_SetStringValue(void* ptr, const char* value) {
    NSControl* control = (NSControl*)ptr;
    [control setStringValue:[NSString stringWithUTF8String:value]];
}
