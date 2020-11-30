#import <stdlib.h>
#import <utils.h>


void* LayoutDimension_ConstraintEqualToConstant(void* ptr, double constant);
void* LayoutDimension_ConstraintEqualToAnchor3(void* ptr, void* anchor, double multiplier, double constant);
void* LayoutDimension_ConstraintGreaterThanOrEqualToConstant(void* ptr, double constant);
void* LayoutDimension_ConstraintGreaterThanOrEqualToAnchor3(void* ptr, void* anchor, double multiplier, double constant);
void* LayoutDimension_ConstraintLessThanOrEqualToConstant(void* ptr, double constant);
void* LayoutDimension_ConstraintLessThanOrEqualToAnchor3(void* ptr, void* anchor, double multiplier, double constant);
