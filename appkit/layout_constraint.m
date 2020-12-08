#import <AppKit/AppKit.h>
#import "layout_constraint.h"

bool LayoutConstraint_IsActive(void* ptr) {
	NSLayoutConstraint* layoutConstraint = (NSLayoutConstraint*)ptr;
	return [layoutConstraint isActive];
}

void LayoutConstraint_SetActive(void* ptr, bool active) {
	NSLayoutConstraint* layoutConstraint = (NSLayoutConstraint*)ptr;
	[layoutConstraint setActive:active];
}

void* LayoutConstraint_FirstItem(void* ptr) {
	NSLayoutConstraint* layoutConstraint = (NSLayoutConstraint*)ptr;
	return [layoutConstraint firstItem];
}

long LayoutConstraint_FirstAttribute(void* ptr) {
	NSLayoutConstraint* layoutConstraint = (NSLayoutConstraint*)ptr;
	return [layoutConstraint firstAttribute];
}

long LayoutConstraint_Relation(void* ptr) {
	NSLayoutConstraint* layoutConstraint = (NSLayoutConstraint*)ptr;
	return [layoutConstraint relation];
}

void* LayoutConstraint_SecondItem(void* ptr) {
	NSLayoutConstraint* layoutConstraint = (NSLayoutConstraint*)ptr;
	return [layoutConstraint secondItem];
}

long LayoutConstraint_SecondAttribute(void* ptr) {
	NSLayoutConstraint* layoutConstraint = (NSLayoutConstraint*)ptr;
	return [layoutConstraint secondAttribute];
}

double LayoutConstraint_Multiplier(void* ptr) {
	NSLayoutConstraint* layoutConstraint = (NSLayoutConstraint*)ptr;
	return [layoutConstraint multiplier];
}

double LayoutConstraint_Constant(void* ptr) {
	NSLayoutConstraint* layoutConstraint = (NSLayoutConstraint*)ptr;
	return [layoutConstraint constant];
}

void* LayoutConstraint_FirstAnchor(void* ptr) {
	NSLayoutConstraint* layoutConstraint = (NSLayoutConstraint*)ptr;
	return [layoutConstraint firstAnchor];
}

void* LayoutConstraint_SecondAnchor(void* ptr) {
	NSLayoutConstraint* layoutConstraint = (NSLayoutConstraint*)ptr;
	return [layoutConstraint secondAnchor];
}

float LayoutConstraint_Priority(void* ptr) {
	NSLayoutConstraint* layoutConstraint = (NSLayoutConstraint*)ptr;
	return [layoutConstraint priority];
}

void LayoutConstraint_SetPriority(void* ptr, float priority) {
	NSLayoutConstraint* layoutConstraint = (NSLayoutConstraint*)ptr;
	[layoutConstraint setPriority:priority];
}

const char* LayoutConstraint_Identifier(void* ptr) {
	NSLayoutConstraint* layoutConstraint = (NSLayoutConstraint*)ptr;
	return [[layoutConstraint identifier] UTF8String];
}

void LayoutConstraint_SetIdentifier(void* ptr, const char* identifier) {
	NSLayoutConstraint* layoutConstraint = (NSLayoutConstraint*)ptr;
	[layoutConstraint setIdentifier:[NSString stringWithUTF8String:identifier]];
}

bool LayoutConstraint_ShouldBeArchived(void* ptr) {
	NSLayoutConstraint* layoutConstraint = (NSLayoutConstraint*)ptr;
	return [layoutConstraint shouldBeArchived];
}

void LayoutConstraint_SetShouldBeArchived(void* ptr, bool shouldBeArchived) {
	NSLayoutConstraint* layoutConstraint = (NSLayoutConstraint*)ptr;
	[layoutConstraint setShouldBeArchived:shouldBeArchived];
}

void LayoutConstraint_ConstraintWithItem(void* view1, long attr1, long relation, void* view2, long attr2, double multiplier, double c) {
	[NSLayoutConstraint constraintWithItem:(NSObject*)view1 attribute:attr1 relatedBy:relation toItem:(NSObject*)view2 attribute:attr2 multiplier:multiplier constant:c];
}

void LayoutConstraint_ActivateConstraints(Array constraints) {
    NSMutableArray* objcConstraints = [[NSMutableArray alloc] init];
    void** constraintsData = (void**)constraints.data;
    for (int i = 0; i < constraints.len; i++) {
    	[objcConstraints addObject:(NSLayoutConstraint*)constraintsData[i]];
    }
	[NSLayoutConstraint activateConstraints:objcConstraints];
}

void LayoutConstraint_DeactivateConstraints(Array constraints) {
    NSMutableArray* objcConstraints = [[NSMutableArray alloc] init];
    void** constraintsData = (void**)constraints.data;
    for (int i = 0; i < constraints.len; i++) {
    	[objcConstraints addObject:(NSLayoutConstraint*)constraintsData[i]];
    }
	[NSLayoutConstraint deactivateConstraints:objcConstraints];
}
