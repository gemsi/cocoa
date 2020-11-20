#import <Appkit/NSWindow.h>
#import "window.h"
#import "window_delegate.h"

void* Window_initWithContentRect(long goID, NSRect rect, unsigned long styleMask, unsigned long backing, bool Defer) {
	NSWindow* window = [[[NSWindow alloc] initWithContentRect:rect styleMask:styleMask backing:backing defer:Defer] autorelease];
	GoNSWindowDelegate* delegate = [[GoNSWindowDelegate alloc] init];
	[delegate setGoID:goID];
	[window setDelegate:delegate];
	return (void*)window;
}

const char* Window_Title(void* ptr) {
	NSWindow* window = (NSWindow*)ptr;
	return [[window title] UTF8String];
}

void Window_SetTitle(void* ptr, const char* title) {
	NSWindow* window = (NSWindow*)ptr;
	[window setTitle:[NSString stringWithUTF8String:title]];
}

void* Window_ContentView(void* ptr) {
	NSWindow* window = (NSWindow*)ptr;
	return [window contentView];
}

void Window_SetContentView(void* ptr, void* contentView) {
	NSWindow* window = (NSWindow*)ptr;
	[window setContentView:contentView];
}

unsigned long Window_StyleMask(void* ptr) {
	NSWindow* window = (NSWindow*)ptr;
	return [window styleMask];
}

void Window_SetStyleMask(void* ptr, unsigned long styleMask) {
	NSWindow* window = (NSWindow*)ptr;
	[window setStyleMask:styleMask];
}

void Window_Center(void* ptr) {
	NSWindow* window = (NSWindow*)ptr;
	[window center];
}

void Window_MakeKeyAndOrderFront(void* ptr, void* sender) {
	NSWindow* window = (NSWindow*)ptr;
	[window makeKeyAndOrderFront:sender];
}