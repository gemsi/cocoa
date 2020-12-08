#import <AppKit/AppKit.h>
#import "layout_y_axis_anchor.h"

void* LayoutYAxisAnchor_ConstraintEqualToSystemSpacingBelowAnchor(void* ptr, void* anchor, double multiplier) {
	NSLayoutYAxisAnchor* layoutYAxisAnchor = (NSLayoutYAxisAnchor*)ptr;
	return [layoutYAxisAnchor constraintEqualToSystemSpacingBelowAnchor:(NSLayoutYAxisAnchor*)anchor multiplier:multiplier];
}

void* LayoutYAxisAnchor_ConstraintGreaterThanOrEqualToSystemSpacingBelowAnchor(void* ptr, void* anchor, double multiplier) {
	NSLayoutYAxisAnchor* layoutYAxisAnchor = (NSLayoutYAxisAnchor*)ptr;
	return [layoutYAxisAnchor constraintGreaterThanOrEqualToSystemSpacingBelowAnchor:(NSLayoutYAxisAnchor*)anchor multiplier:multiplier];
}

void* LayoutYAxisAnchor_ConstraintLessThanOrEqualToSystemSpacingBelowAnchor(void* ptr, void* anchor, double multiplier) {
	NSLayoutYAxisAnchor* layoutYAxisAnchor = (NSLayoutYAxisAnchor*)ptr;
	return [layoutYAxisAnchor constraintLessThanOrEqualToSystemSpacingBelowAnchor:(NSLayoutYAxisAnchor*)anchor multiplier:multiplier];
}

void* LayoutYAxisAnchor_AnchorWithOffsetToAnchor(void* ptr, void* otherAnchor) {
	NSLayoutYAxisAnchor* layoutYAxisAnchor = (NSLayoutYAxisAnchor*)ptr;
	return [layoutYAxisAnchor anchorWithOffsetToAnchor:(NSLayoutYAxisAnchor*)otherAnchor];
}
