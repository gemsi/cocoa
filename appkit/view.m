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

void* View_BottomAnchor(void* ptr) {
	NSView* view = (NSView*)ptr;
	return [view bottomAnchor];
}

void* View_CenterXAnchor(void* ptr) {
	NSView* view = (NSView*)ptr;
	return [view centerXAnchor];
}

void* View_CenterYAnchor(void* ptr) {
	NSView* view = (NSView*)ptr;
	return [view centerYAnchor];
}

void* View_FirstBaselineAnchor(void* ptr) {
	NSView* view = (NSView*)ptr;
	return [view firstBaselineAnchor];
}

void* View_HeightAnchor(void* ptr) {
	NSView* view = (NSView*)ptr;
	return [view heightAnchor];
}

void* View_LastBaselineAnchor(void* ptr) {
	NSView* view = (NSView*)ptr;
	return [view lastBaselineAnchor];
}

void* View_LeadingAnchor(void* ptr) {
	NSView* view = (NSView*)ptr;
	return [view leadingAnchor];
}

void* View_LeftAnchor(void* ptr) {
	NSView* view = (NSView*)ptr;
	return [view leftAnchor];
}

void* View_RightAnchor(void* ptr) {
	NSView* view = (NSView*)ptr;
	return [view rightAnchor];
}

void* View_TopAnchor(void* ptr) {
	NSView* view = (NSView*)ptr;
	return [view topAnchor];
}

void* View_TrailingAnchor(void* ptr) {
	NSView* view = (NSView*)ptr;
	return [view trailingAnchor];
}

void* View_WidthAnchor(void* ptr) {
	NSView* view = (NSView*)ptr;
	return [view widthAnchor];
}

void View_AddSubview(void* ptr, void* subView) {
	NSView* view = (NSView*)ptr;
	[view addSubview:subView];
}
