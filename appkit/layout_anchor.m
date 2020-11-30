#import <Appkit/NSLayoutAnchor.h>
#import <Appkit/NSLayoutConstraint.h>
#import "layout_anchor.h"

Array LayoutAnchor_ConstraintsAffectingLayout(void* ptr) {
	NSLayoutAnchor* layoutAnchor = (NSLayoutAnchor*)ptr;
	NSArray* array = [layoutAnchor constraintsAffectingLayout];
	int count = [array count];
	void** data = malloc(count * sizeof(void*));
	for (int i = 0; i < count; i++) {
		 data[i] = [array objectAtIndex:i];
	}
	Array result;
	result.data = data;
	result.len = count;
	return result;
}

bool LayoutAnchor_HasAmbiguousLayout(void* ptr) {
	NSLayoutAnchor* layoutAnchor = (NSLayoutAnchor*)ptr;
	return [layoutAnchor hasAmbiguousLayout];
}

const char* LayoutAnchor_Name(void* ptr) {
	NSLayoutAnchor* layoutAnchor = (NSLayoutAnchor*)ptr;
	return [[layoutAnchor name] UTF8String];
}

void* LayoutAnchor_Item(void* ptr) {
	NSLayoutAnchor* layoutAnchor = (NSLayoutAnchor*)ptr;
	return [layoutAnchor item];
}

void* LayoutAnchor_ConstraintEqualToAnchor(void* ptr, void* anchor) {
	NSLayoutAnchor* layoutAnchor = (NSLayoutAnchor*)ptr;
	return [layoutAnchor constraintEqualToAnchor:(NSLayoutAnchor*)anchor];
}

void* LayoutAnchor_ConstraintEqualToAnchor2(void* ptr, void* anchor, double constant) {
	NSLayoutAnchor* layoutAnchor = (NSLayoutAnchor*)ptr;
	return [layoutAnchor constraintEqualToAnchor:(NSLayoutAnchor*)anchor constant:constant];
}

void* LayoutAnchor_ConstraintGreaterThanOrEqualToAnchor(void* ptr, void* anchor) {
	NSLayoutAnchor* layoutAnchor = (NSLayoutAnchor*)ptr;
	return [layoutAnchor constraintGreaterThanOrEqualToAnchor:(NSLayoutAnchor*)anchor];
}

void* LayoutAnchor_ConstraintGreaterThanOrEqualToAnchor2(void* ptr, void* anchor, double constant) {
	NSLayoutAnchor* layoutAnchor = (NSLayoutAnchor*)ptr;
	return [layoutAnchor constraintGreaterThanOrEqualToAnchor:(NSLayoutAnchor*)anchor constant:constant];
}

void* LayoutAnchor_ConstraintLessThanOrEqualToAnchor(void* ptr, void* anchor) {
	NSLayoutAnchor* layoutAnchor = (NSLayoutAnchor*)ptr;
	return [layoutAnchor constraintLessThanOrEqualToAnchor:(NSLayoutAnchor*)anchor];
}

void* LayoutAnchor_ConstraintLessThanOrEqualToAnchor2(void* ptr, void* anchor, double constant) {
	NSLayoutAnchor* layoutAnchor = (NSLayoutAnchor*)ptr;
	return [layoutAnchor constraintLessThanOrEqualToAnchor:(NSLayoutAnchor*)anchor constant:constant];
}
