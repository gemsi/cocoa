#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>

Array LayoutAnchor_ConstraintsAffectingLayout(void* ptr);
bool LayoutAnchor_HasAmbiguousLayout(void* ptr);
const char* LayoutAnchor_Name(void* ptr);
void* LayoutAnchor_Item(void* ptr);

void* LayoutAnchor_ConstraintEqualToAnchor(void* ptr, void* anchor);
void* LayoutAnchor_ConstraintEqualToAnchor2(void* ptr, void* anchor, double constant);
void* LayoutAnchor_ConstraintGreaterThanOrEqualToAnchor(void* ptr, void* anchor);
void* LayoutAnchor_ConstraintGreaterThanOrEqualToAnchor2(void* ptr, void* anchor, double constant);
void* LayoutAnchor_ConstraintLessThanOrEqualToAnchor(void* ptr, void* anchor);
void* LayoutAnchor_ConstraintLessThanOrEqualToAnchor2(void* ptr, void* anchor, double constant);
