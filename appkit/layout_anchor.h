#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>

void* LayoutAnchor_ConstraintEqualTo(void* ptr, void* anchorPtr);
void* LayoutAnchor_ConstraintEqualTo2(void* ptr, void* anchorPtr, double constant);
void* LayoutAnchor_ConstraintGreaterThanOrEqualTo(void* ptr, void* anchorPtr);
void* LayoutAnchor_ConstraintGreaterThanOrEqualTo2(void* ptr, void* anchorPtr, double constant);
void* LayoutAnchor_ConstraintLessThanOrEqualTo(void* ptr, void* anchorPtr);
void* LayoutAnchor_ConstraintLessThanOrEqualTo2(void* ptr, void* anchorPtr, double constant);

void* LayoutDimension_ConstraintEqualTo3(void* ptr, void* anchorPtr, double multiplier, double constant);
void* LayoutDimension_ConstraintEqualToConstant(void* ptr, double constant);
void* LayoutDimension_ConstraintGreaterThanOrEqualTo3(void* ptr, void* anchorPtr, double multiplier, double constant);
void* LayoutDimension_ConstraintGreaterThanOrEqualToConstant(void* ptr, double constant);
void* LayoutDimension_ConstraintLessThanOrEqualTo3(void* ptr, void* anchorPtr, double multiplier, double constant);
void* LayoutDimension_ConstraintLessThanOrEqualToConstant(void* ptr, double constant);

void* LayoutXAxisAnchor_ConstraintEqualToSystemSpacingAfter(void* ptr, void* anchorPtr, double multiplier);
void* LayoutXAxisAnchor_ConstraintGreaterThanOrEqualToSystemSpacingAfter(void* ptr, void* anchorPtr, double multiplier);
void* LayoutXAxisAnchor_ConstraintLessThanOrEqualToSystemSpacingAfter(void* ptr, void* anchorPtr, double multiplier);
void* LayoutXAxisAnchor_AnchorWithOffsetTo(void* ptr, void* anchorPtr);

void* LayoutYAxisAnchor_ConstraintEqualToSystemSpacingBelow(void* ptr, void* anchorPtr, double multiplier);
void* LayoutYAxisAnchor_ConstraintGreaterThanOrEqualToSystemSpacingBelow(void* ptr, void* anchorPtr, double multiplier);
void* LayoutYAxisAnchor_ConstraintLessThanOrEqualToSystemSpacingBelow(void* ptr, void* anchorPtr, double multiplier);
void* LayoutYAxisAnchor_AnchorWithOffsetTo(void* ptr, void* anchorPtr);
