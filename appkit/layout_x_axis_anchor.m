#import <AppKit/AppKit.h>
#import "layout_x_axis_anchor.h"

void* LayoutXAxisAnchor_ConstraintEqualToSystemSpacingAfterAnchor(void* ptr, void* anchor, double multiplier) {
	NSLayoutXAxisAnchor* layoutXAxisAnchor = (NSLayoutXAxisAnchor*)ptr;
	return [layoutXAxisAnchor constraintEqualToSystemSpacingAfterAnchor:(NSLayoutXAxisAnchor*)anchor multiplier:multiplier];
}

void* LayoutXAxisAnchor_ConstraintGreaterThanOrEqualToSystemSpacingAfterAnchor(void* ptr, void* anchor, double multiplier) {
	NSLayoutXAxisAnchor* layoutXAxisAnchor = (NSLayoutXAxisAnchor*)ptr;
	return [layoutXAxisAnchor constraintGreaterThanOrEqualToSystemSpacingAfterAnchor:(NSLayoutXAxisAnchor*)anchor multiplier:multiplier];
}

void* LayoutXAxisAnchor_ConstraintLessThanOrEqualToSystemSpacingAfterAnchor(void* ptr, void* anchor, double multiplier) {
	NSLayoutXAxisAnchor* layoutXAxisAnchor = (NSLayoutXAxisAnchor*)ptr;
	return [layoutXAxisAnchor constraintLessThanOrEqualToSystemSpacingAfterAnchor:(NSLayoutXAxisAnchor*)anchor multiplier:multiplier];
}

void* LayoutXAxisAnchor_AnchorWithOffsetToAnchor(void* ptr, void* otherAnchor) {
	NSLayoutXAxisAnchor* layoutXAxisAnchor = (NSLayoutXAxisAnchor*)ptr;
	return [layoutXAxisAnchor anchorWithOffsetToAnchor:(NSLayoutXAxisAnchor*)otherAnchor];
}
