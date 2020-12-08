#import <AppKit/AppKit.h>
#import "layout_dimension.h"

void* LayoutDimension_ConstraintEqualToConstant(void* ptr, double constant) {
	NSLayoutDimension* layoutDimension = (NSLayoutDimension*)ptr;
	return [layoutDimension constraintEqualToConstant:constant];
}

void* LayoutDimension_ConstraintEqualToAnchor3(void* ptr, void* anchor, double multiplier, double constant) {
	NSLayoutDimension* layoutDimension = (NSLayoutDimension*)ptr;
	return [layoutDimension constraintEqualToAnchor:(NSLayoutDimension*)anchor multiplier:multiplier constant:constant];
}

void* LayoutDimension_ConstraintGreaterThanOrEqualToConstant(void* ptr, double constant) {
	NSLayoutDimension* layoutDimension = (NSLayoutDimension*)ptr;
	return [layoutDimension constraintGreaterThanOrEqualToConstant:constant];
}

void* LayoutDimension_ConstraintGreaterThanOrEqualToAnchor3(void* ptr, void* anchor, double multiplier, double constant) {
	NSLayoutDimension* layoutDimension = (NSLayoutDimension*)ptr;
	return [layoutDimension constraintGreaterThanOrEqualToAnchor:(NSLayoutDimension*)anchor multiplier:multiplier constant:constant];
}

void* LayoutDimension_ConstraintLessThanOrEqualToConstant(void* ptr, double constant) {
	NSLayoutDimension* layoutDimension = (NSLayoutDimension*)ptr;
	return [layoutDimension constraintLessThanOrEqualToConstant:constant];
}

void* LayoutDimension_ConstraintLessThanOrEqualToAnchor3(void* ptr, void* anchor, double multiplier, double constant) {
	NSLayoutDimension* layoutDimension = (NSLayoutDimension*)ptr;
	return [layoutDimension constraintLessThanOrEqualToAnchor:(NSLayoutDimension*)anchor multiplier:multiplier constant:constant];
}
