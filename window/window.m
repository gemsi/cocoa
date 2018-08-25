#import "window.h"
#import "windowdelegate.h"
#include "_cgo_export.h"


void* Window_New(int x, int y, int width, int height)
{
    NSRect windowRect = NSMakeRect(x, y, width, height);
    id window = [[NSWindow alloc] initWithContentRect:windowRect 
        styleMask:(NSWindowStyleMaskTitled | NSWindowStyleMaskClosable | NSWindowStyleMaskResizable | NSWindowStyleMaskMiniaturizable)
        backing:NSBackingStoreBuffered
        defer:NO];
    [window autorelease];
    return window;

}

void Window_SetTitle(void *wndPtr, const char* title)
{
    NSWindow* window = (NSWindow*)wndPtr;
    [window setTitle:[NSString stringWithUTF8String:title]];
}

void Window_SetDelegate(void *wndPtr, int id)
{
    NSWindow* window = (NSWindow*)wndPtr;
    WindowDelegate* cocoa_windowDelegate = [[WindowDelegate alloc] init];
    [cocoa_windowDelegate setGoWindowID:id];
    [window setDelegate:cocoa_windowDelegate];
}

void Window_MakeKeyAndOrderFront(void *wndPtr)
{
    NSWindow* window = (NSWindow*)wndPtr;
    [window makeKeyAndOrderFront:nil];
}

void Window_Center(void *wndPtr)
{
    NSWindow* window = (NSWindow*)wndPtr;
    [window center];
}

void Window_AddView(void *wndPtr, void* viewPtr) 
{
    NSView* view = (NSView*)viewPtr;
    NSWindow* window = (NSWindow*)wndPtr;
    [[window contentView] addSubview:view];
}

void Window_Update(void *wndPtr)
{
    NSWindow* window = (NSWindow*)wndPtr;
    [[window contentView] setNeedsDisplay:YES];
}

NSRect Window_Frame(void *ptr)
{
    NSWindow* window = (NSWindow*)ptr;
    return [window frame];
}
