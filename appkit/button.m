#import <Appkit/NSButton.h>
#import "button.h"
#import "button_delegate.h"

void* Button_initWithFrame(long goID, NSRect frame) {
	NSButton* button = [[[NSButton alloc] initWithFrame:frame] autorelease];
	NSButtonHandler* handler = [[[NSButtonHandler alloc] init] autorelease];
	[handler setGoID:goID];
	[button setTarget:handler];
	[button setAction:@selector(onAction:)];
	return (void*)button;
}

const char* Button_Title(void* ptr) {
	NSButton* button = (NSButton*)ptr;
	return [[button title] UTF8String];
}

void Button_SetTitle(void* ptr, const char* title) {
	NSButton* button = (NSButton*)ptr;
	[button setTitle:[NSString stringWithUTF8String:title]];
}

unsigned long Button_BezelStyle(void* ptr) {
	NSButton* button = (NSButton*)ptr;
	return [button bezelStyle];
}

void Button_SetBezelStyle(void* ptr, unsigned long bezelStyle) {
	NSButton* button = (NSButton*)ptr;
	[button setBezelStyle:bezelStyle];
}

long Button_State(void* ptr) {
	NSButton* button = (NSButton*)ptr;
	return [button state];
}

void Button_SetState(void* ptr, long state) {
	NSButton* button = (NSButton*)ptr;
	[button setState:state];
}

void Button_SetButtonType(void* ptr, unsigned long buttonType) {
	NSButton* button = (NSButton*)ptr;
	[button setButtonType:buttonType];
}
