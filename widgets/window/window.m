#import "window.h"
#import "window_delegate.h"
#include "_cgo_export.h"


void* Window_New(int x, int y, int width, int height, int goID) {
    NSRect windowRect = NSMakeRect(x, y, width, height);
    id window = [[NSWindow alloc] initWithContentRect:windowRect 
        styleMask:(NSWindowStyleMaskTitled | NSWindowStyleMaskClosable | NSWindowStyleMaskResizable | NSWindowStyleMaskMiniaturizable)
        backing:NSBackingStoreBuffered
        defer:NO];
    WindowDelegate* windowDelegate = [[WindowDelegate alloc] init];
    [windowDelegate setGoID:goID];
    [window setDelegate:windowDelegate];
    [window autorelease];
    return window;

}

void Window_SetTitle(void *ptr, const char* title) {
    NSWindow* window = (NSWindow*)ptr;
    [window setTitle:[NSString stringWithUTF8String:title]];
}

void Window_MakeKeyAndOrderFront(void *ptr) {
    NSWindow* window = (NSWindow*)ptr;
    [window makeKeyAndOrderFront:nil];
}

void Window_Center(void *ptr) {
    NSWindow* window = (NSWindow*)ptr;
    [window center];
}

void Window_AddView(void *ptr, void* viewPtr) {
    NSView* view = (NSView*)viewPtr;
    NSWindow* window = (NSWindow*)ptr;
    [[window contentView] addSubview:view];
}

void Window_Update(void *ptr) {
    NSWindow* window = (NSWindow*)ptr;
    [[window contentView] setNeedsDisplay:YES];
}

NSRect Window_Frame(void *ptr) {
    NSWindow* window = (NSWindow*)ptr;
    return [window frame];
}
