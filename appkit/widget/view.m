#import <Appkit/NSView.h>
        #import "view.h"
        #include "_cgo_export.h"


void* View_New(long id, NSRect frame) {
    NSView* view = [[[NSView alloc] initWithFrame:frame] autorelease];
    return (void*)view;
}
NSRect View_Frame(void* ptr) {
	NSView* view = (NSView*)ptr;
	return [view frame];
}
void View_SetFrame(void* ptr, NSRect frame) {
	NSView* view = (NSView*)ptr;
	[view setFrame:frame];
}
unsigned long View_AutoresizingMask(void* ptr) {
	NSView* view = (NSView*)ptr;
	return [view autoresizingMask];
}
void View_SetAutoresizingMask(void* ptr, unsigned long autoresizingMask) {
	NSView* view = (NSView*)ptr;
	[view setAutoresizingMask:autoresizingMask];
}
