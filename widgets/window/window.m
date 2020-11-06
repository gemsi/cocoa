#import <Cocoa/Cocoa.h>
#import "window.h"
#include "_cgo_export.h"

@interface GoNSWindow : NSWindow <NSWindowDelegate>

@property (assign) long goID;

@end

@implementation GoNSWindow

- (void)dealloc {
    [super dealloc];
}

- (void)windowDidResize:(NSNotification *)notification {
    onWindowDidResize([self goID], notification);
}

- (void)windowDidMove:(NSNotification *)notification {
    onWindowDidMove([self goID], notification);
}

- (void)windowDidMiniaturize:(NSNotification *)notification {
    onWindowDidMiniaturize([self goID], notification);
}

- (void)windowDidDeminiaturize:(NSNotification *)notification {
    onWindowDidDeminiaturize([self goID], notification);
}

@end


void* Window_New(long goID, NSRect rect) {
    GoNSWindow* window = [[GoNSWindow alloc] initWithContentRect:rect
        styleMask:(NSWindowStyleMaskTitled | NSWindowStyleMaskClosable | NSWindowStyleMaskResizable | NSWindowStyleMaskMiniaturizable)
        backing:NSBackingStoreBuffered
        defer:NO];
    [window setGoID:goID];
    [window setDelegate:window];
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

void Window_SetFrame(void *ptr, NSRect rect, int display) {
    NSWindow* window = (NSWindow*)ptr;
    [window setFrame:rect display:display==1?YES:NO];
}
