#import <AppKit/AppKit.h>
#import "control.h"

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

long Control_IntValue(void* ptr) {
	NSControl* control = (NSControl*)ptr;
	return [control intValue];
}

void Control_SetIntValue(void* ptr, long intValue) {
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

void* Control_Cell(void* ptr) {
	NSControl* control = (NSControl*)ptr;
	return [control cell];
}

void Control_SetCell(void* ptr, void* cell) {
	NSControl* control = (NSControl*)ptr;
	[control setCell:(NSCell*)cell];
}

void Control_SizeToFit(void* ptr) {
	NSControl* control = (NSControl*)ptr;
	[control sizeToFit];
}
