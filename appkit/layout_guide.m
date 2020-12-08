#import <AppKit/AppKit.h>
#import "layout_guide.h"

NSRect LayoutGuide_Frame(void* ptr) {
	NSLayoutGuide* layoutGuide = (NSLayoutGuide*)ptr;
	return [layoutGuide frame];
}

void* LayoutGuide_BottomAnchor(void* ptr) {
	NSLayoutGuide* layoutGuide = (NSLayoutGuide*)ptr;
	return [layoutGuide bottomAnchor];
}

void* LayoutGuide_CenterXAnchor(void* ptr) {
	NSLayoutGuide* layoutGuide = (NSLayoutGuide*)ptr;
	return [layoutGuide centerXAnchor];
}

void* LayoutGuide_CenterYAnchor(void* ptr) {
	NSLayoutGuide* layoutGuide = (NSLayoutGuide*)ptr;
	return [layoutGuide centerYAnchor];
}

void* LayoutGuide_HeightAnchor(void* ptr) {
	NSLayoutGuide* layoutGuide = (NSLayoutGuide*)ptr;
	return [layoutGuide heightAnchor];
}

void* LayoutGuide_LeadingAnchor(void* ptr) {
	NSLayoutGuide* layoutGuide = (NSLayoutGuide*)ptr;
	return [layoutGuide leadingAnchor];
}

void* LayoutGuide_LeftAnchor(void* ptr) {
	NSLayoutGuide* layoutGuide = (NSLayoutGuide*)ptr;
	return [layoutGuide leftAnchor];
}

void* LayoutGuide_RightAnchor(void* ptr) {
	NSLayoutGuide* layoutGuide = (NSLayoutGuide*)ptr;
	return [layoutGuide rightAnchor];
}

void* LayoutGuide_TopAnchor(void* ptr) {
	NSLayoutGuide* layoutGuide = (NSLayoutGuide*)ptr;
	return [layoutGuide topAnchor];
}

void* LayoutGuide_TrailingAnchor(void* ptr) {
	NSLayoutGuide* layoutGuide = (NSLayoutGuide*)ptr;
	return [layoutGuide trailingAnchor];
}

void* LayoutGuide_WidthAnchor(void* ptr) {
	NSLayoutGuide* layoutGuide = (NSLayoutGuide*)ptr;
	return [layoutGuide widthAnchor];
}

void* LayoutGuide_OwningView(void* ptr) {
	NSLayoutGuide* layoutGuide = (NSLayoutGuide*)ptr;
	return [layoutGuide owningView];
}

void* LayoutGuide_NewLayoutGuide() {
	NSLayoutGuide* layoutGuide = [NSLayoutGuide alloc];
	return [[layoutGuide init] autorelease];
}
