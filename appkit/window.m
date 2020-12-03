#import <Appkit/NSWindow.h>
#import "window.h"
#import "window_delegate.h"

@implementation GoNSWindowDelegate

- (void)windowDidResize:(NSNotification*)notification {
	return Window_Delegate_WindowDidResize([self goID], notification);
}

- (void)windowDidMove:(NSNotification*)notification {
	return Window_Delegate_WindowDidMove([self goID], notification);
}

- (void)windowDidMiniaturize:(NSNotification*)notification {
	return Window_Delegate_WindowDidMiniaturize([self goID], notification);
}

- (void)windowDidDeminiaturize:(NSNotification*)notification {
	return Window_Delegate_WindowDidDeminiaturize([self goID], notification);
}

@end
void Window_RegisterDelegate(void *ptr, long goID) {
	NSWindow* window = (NSWindow*)ptr;
	GoNSWindowDelegate* delegate = [[GoNSWindowDelegate alloc] init];
	[delegate setGoID:goID];
	[window setDelegate:delegate];
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
	[window setContentView:(NSView*)contentView];
}

unsigned long Window_StyleMask(void* ptr) {
	NSWindow* window = (NSWindow*)ptr;
	return [window styleMask];
}

void Window_SetStyleMask(void* ptr, unsigned long styleMask) {
	NSWindow* window = (NSWindow*)ptr;
	[window setStyleMask:styleMask];
}

void* Window_NewWindow(NSRect rect, unsigned long styleMask, unsigned long backing, bool Defer) {
	NSWindow* window = [NSWindow alloc];
	return [[window initWithContentRect:rect styleMask:styleMask backing:backing defer:Defer] autorelease];
}

void Window_Center(void* ptr) {
	NSWindow* window = (NSWindow*)ptr;
	[window center];
}

void Window_MakeKeyAndOrderFront(void* ptr, void* sender) {
	NSWindow* window = (NSWindow*)ptr;
	[window makeKeyAndOrderFront:(NSObject*)sender];
}

bool Window_MakeFirstResponder(void* ptr, void* responder) {
	NSWindow* window = (NSWindow*)ptr;
	return [window makeFirstResponder:(NSResponder*)responder];
}
