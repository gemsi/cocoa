#import <AppKit/NSLayoutAnchor.h>
#import "layout_constraint.h"

void* LayoutAnchor_ConstraintEqualTo(void* ptr, void* anchorPtr) {
    NSLayoutAnchor* self = (NSLayoutAnchor*)ptr;
    NSLayoutAnchor* anchor = (NSLayoutAnchor*)anchorPtr;
    return [self constraintEqualToAnchor:anchor];
}

void* LayoutAnchor_ConstraintEqualTo2(void* ptr, void* anchorPtr, double constant) {
    NSLayoutAnchor* self = (NSLayoutAnchor*)ptr;
    NSLayoutAnchor* anchor = (NSLayoutAnchor*)anchorPtr;
    return [self constraintEqualToAnchor:anchor constant:constant];
}

void* LayoutAnchor_ConstraintGreaterThanOrEqualTo(void* ptr, void* anchorPtr) {
    NSLayoutAnchor* self = (NSLayoutAnchor*)ptr;
    NSLayoutAnchor* anchor = (NSLayoutAnchor*)anchorPtr;
    return [self constraintGreaterThanOrEqualToAnchor:anchor];
}

void* LayoutAnchor_ConstraintGreaterThanOrEqualTo2(void* ptr, void* anchorPtr, double constant) {
    NSLayoutAnchor* self = (NSLayoutAnchor*)ptr;
    NSLayoutAnchor* anchor = (NSLayoutAnchor*)anchorPtr;
    return [self constraintGreaterThanOrEqualToAnchor:anchor constant:constant];
}

void* LayoutAnchor_ConstraintLessThanOrEqualTo(void* ptr, void* anchorPtr) {
    NSLayoutAnchor* self = (NSLayoutAnchor*)ptr;
    NSLayoutAnchor* anchor = (NSLayoutAnchor*)anchorPtr;
    return [self constraintLessThanOrEqualToAnchor:anchor];
}

void* LayoutAnchor_ConstraintLessThanOrEqualTo2(void* ptr, void* anchorPtr, double constant) {
    NSLayoutAnchor* self = (NSLayoutAnchor*)ptr;
    NSLayoutAnchor* anchor = (NSLayoutAnchor*)anchorPtr;
    return [self constraintLessThanOrEqualToAnchor:anchor constant:constant];
}

void* LayoutDimension_ConstraintEqualTo3(void* ptr, void* anchorPtr, double multiplier, double constant) {
    NSLayoutDimension* self = (NSLayoutDimension*)ptr;
    NSLayoutDimension* anchor = (NSLayoutDimension*)anchorPtr;
    return [self constraintEqualToAnchor:anchor multiplier:multiplier constant:constant];
}

void* LayoutDimension_ConstraintEqualToConstant(void* ptr, double constant) {
    NSLayoutDimension* self = (NSLayoutDimension*)ptr;
    return [self constraintEqualToConstant:constant];
}

void* LayoutDimension_ConstraintGreaterThanOrEqualTo3(void* ptr, void* anchorPtr, double multiplier, double constant) {
    NSLayoutDimension* self = (NSLayoutDimension*)ptr;
    NSLayoutDimension* anchor = (NSLayoutDimension*)anchorPtr;
    return [self constraintGreaterThanOrEqualToAnchor:anchor multiplier:multiplier constant:constant];
}

void* LayoutDimension_ConstraintGreaterThanOrEqualToConstant(void* ptr, double constant) {
    NSLayoutDimension* self = (NSLayoutDimension*)ptr;
    return [self constraintGreaterThanOrEqualToConstant:constant];
}

void* LayoutDimension_ConstraintLessThanOrEqualTo3(void* ptr, void* anchorPtr, double multiplier, double constant) {
    NSLayoutDimension* self = (NSLayoutDimension*)ptr;
    NSLayoutDimension* anchor = (NSLayoutDimension*)anchorPtr;
    return [self constraintLessThanOrEqualToAnchor:anchor multiplier:multiplier constant:constant];
}

void* LayoutDimension_ConstraintLessThanOrEqualToConstant(void* ptr, double constant) {
    NSLayoutDimension* self = (NSLayoutDimension*)ptr;
    return [self constraintLessThanOrEqualToConstant:constant];
}

void* LayoutXAxisAnchor_ConstraintEqualToSystemSpacingAfter(void* ptr, void* anchorPtr, double multiplier) {
    NSLayoutXAxisAnchor* self = (NSLayoutXAxisAnchor*)ptr;
    NSLayoutXAxisAnchor* anchor = (NSLayoutXAxisAnchor*)anchorPtr;
    return [self constraintEqualToSystemSpacingAfterAnchor:anchor multiplier:multiplier];
}

void* LayoutXAxisAnchor_ConstraintGreaterThanOrEqualToSystemSpacingAfter(void* ptr, void* anchorPtr, double multiplier) {
    NSLayoutXAxisAnchor* self = (NSLayoutXAxisAnchor*)ptr;
    NSLayoutXAxisAnchor* anchor = (NSLayoutXAxisAnchor*)anchorPtr;
    return [self constraintGreaterThanOrEqualToSystemSpacingAfterAnchor:anchor multiplier:multiplier];
}

void* LayoutXAxisAnchor_ConstraintLessThanOrEqualToSystemSpacingAfter(void* ptr, void* anchorPtr, double multiplier) {
    NSLayoutXAxisAnchor* self = (NSLayoutXAxisAnchor*)ptr;
    NSLayoutXAxisAnchor* anchor = (NSLayoutXAxisAnchor*)anchorPtr;
    return [self constraintLessThanOrEqualToSystemSpacingAfterAnchor:anchor multiplier:multiplier];
}
void* LayoutXAxisAnchor_AnchorWithOffsetTo(void* ptr, void* anchorPtr) {
    NSLayoutXAxisAnchor* self = (NSLayoutXAxisAnchor*)ptr;
    NSLayoutXAxisAnchor* anchor = (NSLayoutXAxisAnchor*)anchorPtr;
    return [self anchorWithOffsetToAnchor:anchor];
}

void* LayoutYAxisAnchor_ConstraintEqualToSystemSpacingBelow(void* ptr, void* anchorPtr, double multiplier) {
    NSLayoutYAxisAnchor* self = (NSLayoutYAxisAnchor*)ptr;
    NSLayoutYAxisAnchor* anchor = (NSLayoutYAxisAnchor*)anchorPtr;
    return [self constraintEqualToSystemSpacingBelowAnchor:anchor multiplier:multiplier];
}
void* LayoutYAxisAnchor_ConstraintGreaterThanOrEqualToSystemSpacingBelow(void* ptr, void* anchorPtr, double multiplier) {
    NSLayoutYAxisAnchor* self = (NSLayoutYAxisAnchor*)ptr;
    NSLayoutYAxisAnchor* anchor = (NSLayoutYAxisAnchor*)anchorPtr;
    return [self constraintGreaterThanOrEqualToSystemSpacingBelowAnchor:anchor multiplier:multiplier];
}
void* LayoutYAxisAnchor_ConstraintLessThanOrEqualToSystemSpacingBelow(void* ptr, void* anchorPtr, double multiplier) {
    NSLayoutYAxisAnchor* self = (NSLayoutYAxisAnchor*)ptr;
    NSLayoutYAxisAnchor* anchor = (NSLayoutYAxisAnchor*)anchorPtr;
    return [self constraintLessThanOrEqualToSystemSpacingBelowAnchor:anchor multiplier:multiplier];
}
void* LayoutYAxisAnchor_AnchorWithOffsetTo(void* ptr, void* anchorPtr) {
    NSLayoutYAxisAnchor* self = (NSLayoutYAxisAnchor*)ptr;
    NSLayoutYAxisAnchor* anchor = (NSLayoutYAxisAnchor*)anchorPtr;
    return [self anchorWithOffsetToAnchor:anchor];
}