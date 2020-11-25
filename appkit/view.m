#import <Appkit/NSView.h>
#import "view.h"
#import <Appkit/NSLayoutConstraint.h>

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

bool View_NeedsDisplay(void* ptr) {
	NSView* view = (NSView*)ptr;
	return [view needsDisplay];
}

void View_SetNeedsDisplay(void* ptr, bool needsDisplay) {
	NSView* view = (NSView*)ptr;
	[view setNeedsDisplay:needsDisplay];
}

bool View_TranslatesAutoresizingMaskIntoConstraints(void* ptr) {
	NSView* view = (NSView*)ptr;
	return [view translatesAutoresizingMaskIntoConstraints];
}

void View_SetTranslatesAutoresizingMaskIntoConstraints(void* ptr, bool translatesAutoresizingMaskIntoConstraints) {
	NSView* view = (NSView*)ptr;
	[view setTranslatesAutoresizingMaskIntoConstraints:translatesAutoresizingMaskIntoConstraints];
}

void View_AddSubview(void* ptr, void* subView) {
	NSView* view = (NSView*)ptr;
	[view addSubview:subView];
}
