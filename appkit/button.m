#import <Appkit/NSButton.h>
#import "button.h"
#import "button_delegate.h"

@implementation NSButtonHandler

- (void)onAction:(NSObject*)sender {
	return Button_Target_Action([self goID], sender);
}

@end
void Button_RegisterDelegate(void *ptr, long goID) {
	NSButton* button = (NSButton*)ptr;
	NSButtonHandler* handler = [[[NSButtonHandler alloc] init] autorelease];
	[handler setGoID:goID];
	[button setTarget:handler];
	[button setAction:@selector(onAction:)];
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

void* Button_InitWithFrame(NSRect frame) {
	NSButton* button = [NSButton alloc];
	return [[button initWithFrame:frame] autorelease];
}

void Button_SetButtonType(void* ptr, unsigned long buttonType) {
	NSButton* button = (NSButton*)ptr;
	[button setButtonType:buttonType];
}
