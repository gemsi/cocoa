#import <AppKit/AppKit.h>
#import "progress_indicator.h"

bool ProgressIndicator_UsesThreadedAnimation(void* ptr) {
	NSProgressIndicator* progressIndicator = (NSProgressIndicator*)ptr;
	return [progressIndicator usesThreadedAnimation];
}

void ProgressIndicator_SetUsesThreadedAnimation(void* ptr, bool usesThreadedAnimation) {
	NSProgressIndicator* progressIndicator = (NSProgressIndicator*)ptr;
	[progressIndicator setUsesThreadedAnimation:usesThreadedAnimation];
}

double ProgressIndicator_DoubleValue(void* ptr) {
	NSProgressIndicator* progressIndicator = (NSProgressIndicator*)ptr;
	return [progressIndicator doubleValue];
}

void ProgressIndicator_SetDoubleValue(void* ptr, double doubleValue) {
	NSProgressIndicator* progressIndicator = (NSProgressIndicator*)ptr;
	[progressIndicator setDoubleValue:doubleValue];
}

double ProgressIndicator_MinValue(void* ptr) {
	NSProgressIndicator* progressIndicator = (NSProgressIndicator*)ptr;
	return [progressIndicator minValue];
}

void ProgressIndicator_SetMinValue(void* ptr, double minValue) {
	NSProgressIndicator* progressIndicator = (NSProgressIndicator*)ptr;
	[progressIndicator setMinValue:minValue];
}

double ProgressIndicator_MaxValue(void* ptr) {
	NSProgressIndicator* progressIndicator = (NSProgressIndicator*)ptr;
	return [progressIndicator maxValue];
}

void ProgressIndicator_SetMaxValue(void* ptr, double maxValue) {
	NSProgressIndicator* progressIndicator = (NSProgressIndicator*)ptr;
	[progressIndicator setMaxValue:maxValue];
}

bool ProgressIndicator_IsIndeterminate(void* ptr) {
	NSProgressIndicator* progressIndicator = (NSProgressIndicator*)ptr;
	return [progressIndicator isIndeterminate];
}

void ProgressIndicator_SetIndeterminate(void* ptr, bool indeterminate) {
	NSProgressIndicator* progressIndicator = (NSProgressIndicator*)ptr;
	[progressIndicator setIndeterminate:indeterminate];
}

bool ProgressIndicator_IsBezeled(void* ptr) {
	NSProgressIndicator* progressIndicator = (NSProgressIndicator*)ptr;
	return [progressIndicator isBezeled];
}

void ProgressIndicator_SetBezeled(void* ptr, bool bezeled) {
	NSProgressIndicator* progressIndicator = (NSProgressIndicator*)ptr;
	[progressIndicator setBezeled:bezeled];
}

bool ProgressIndicator_IsDisplayedWhenStopped(void* ptr) {
	NSProgressIndicator* progressIndicator = (NSProgressIndicator*)ptr;
	return [progressIndicator isDisplayedWhenStopped];
}

void ProgressIndicator_SetDisplayedWhenStopped(void* ptr, bool displayedWhenStopped) {
	NSProgressIndicator* progressIndicator = (NSProgressIndicator*)ptr;
	[progressIndicator setDisplayedWhenStopped:displayedWhenStopped];
}

void* ProgressIndicator_NewProgressIndicator(NSRect frame) {
	NSProgressIndicator* progressIndicator = [NSProgressIndicator alloc];
	return [[progressIndicator initWithFrame:frame] autorelease];
}

void ProgressIndicator_StartAnimation(void* ptr, void* sender) {
	NSProgressIndicator* progressIndicator = (NSProgressIndicator*)ptr;
	[progressIndicator startAnimation:(NSObject*)sender];
}

void ProgressIndicator_StopAnimation(void* ptr, void* sender) {
	NSProgressIndicator* progressIndicator = (NSProgressIndicator*)ptr;
	[progressIndicator stopAnimation:(NSObject*)sender];
}

void ProgressIndicator_IncrementBy(void* ptr, double delta) {
	NSProgressIndicator* progressIndicator = (NSProgressIndicator*)ptr;
	[progressIndicator incrementBy:delta];
}
