#import <Appkit/NSControl.h>
        #import "control.h"
        #include "_cgo_export.h"


void* Control_New(long id, NSRect frame) {
    NSControl* control = [[[NSControl alloc] initWithFrame:frame] autorelease];
    return (void*)control;
}
bool Control_IsEnabled(void* ptr) {
	NSControl* control = (NSControl*)ptr;
	return [control isEnabled];
}
void Control_SetEnabled(void* ptr, bool enabled) {
	NSControl* control = (NSControl*)ptr;
	[control setEnabled:enabled];
}
double Control_DoubleValue(void* ptr) {
	NSControl* control = (NSControl*)ptr;
	return [control doubleValue];
}
void Control_SetDoubleValue(void* ptr, double doubleValue) {
	NSControl* control = (NSControl*)ptr;
	[control setDoubleValue:doubleValue];
}
float Control_FloatValue(void* ptr) {
	NSControl* control = (NSControl*)ptr;
	return [control floatValue];
}
void Control_SetFloatValue(void* ptr, float floatValue) {
	NSControl* control = (NSControl*)ptr;
	[control setFloatValue:floatValue];
}
int Control_IntValue(void* ptr) {
	NSControl* control = (NSControl*)ptr;
	return [control intValue];
}
void Control_SetIntValue(void* ptr, int intValue) {
	NSControl* control = (NSControl*)ptr;
	[control setIntValue:intValue];
}
long Control_IntegerValue(void* ptr) {
	NSControl* control = (NSControl*)ptr;
	return [control integerValue];
}
void Control_SetIntegerValue(void* ptr, long integerValue) {
	NSControl* control = (NSControl*)ptr;
	[control setIntegerValue:integerValue];
}
const char* Control_StringValue(void* ptr) {
	NSControl* control = (NSControl*)ptr;
	return [[control stringValue] UTF8String];
}
void Control_SetStringValue(void* ptr, const char* stringValue) {
	NSControl* control = (NSControl*)ptr;
	[control setStringValue:[NSString stringWithUTF8String:stringValue]];
}
