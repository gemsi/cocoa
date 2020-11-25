#import <AppKit/NSLayoutConstraint.h>
#import "layout_constraint.h"


void* LayoutConstraint_New(void* view1, int attr1, int relation, void* view2, int attr2, double multiplier, double constant) {
    return [NSLayoutConstraint
             constraintWithItem:view1
             attribute:attr1
             relatedBy:relation
             toItem:view2
             attribute:attr2
             multiplier:multiplier
             constant:constant];
}

bool LayoutConstraint_IsActive(void* ptr) {
    NSLayoutConstraint* c = (NSLayoutConstraint*)ptr;
    return [c isActive];
}

void LayoutConstraint_SetActive(void* ptr, bool active) {
    NSLayoutConstraint* c = (NSLayoutConstraint*)ptr;
    [c setActive:active];
}

void* LayoutConstraint_FirstItem(void* ptr) {
    NSLayoutConstraint* c = (NSLayoutConstraint*)ptr;
    return [c firstItem];
}

int LayoutConstraint_FirstAttribute(void* ptr) {
    NSLayoutConstraint* c = (NSLayoutConstraint*)ptr;
    return [c firstAttribute];
}

int LayoutConstraint_Relation(void* ptr) {
    NSLayoutConstraint* c = (NSLayoutConstraint*)ptr;
    return [c relation];
}

void* LayoutConstraint_SecondItem(void* ptr) {
    NSLayoutConstraint* c = (NSLayoutConstraint*)ptr;
    return [c secondItem];
}

int LayoutConstraint_SecondAttribute(void* ptr) {
    NSLayoutConstraint* c = (NSLayoutConstraint*)ptr;
    return [c secondAttribute];
}

double LayoutConstraint_Multiplier(void* ptr) {
    NSLayoutConstraint* c = (NSLayoutConstraint*)ptr;
    return [c multiplier];
}

double LayoutConstraint_Constant(void* ptr) {
    NSLayoutConstraint* c = (NSLayoutConstraint*)ptr;
    return [c constant];
}

void LayoutConstraint_SetConstant(void* ptr, double constant) {
    NSLayoutConstraint* c = (NSLayoutConstraint*)ptr;
    [c setConstant:constant];
}

float LayoutConstraint_Priority(void* ptr) {
    NSLayoutConstraint* c = (NSLayoutConstraint*)ptr;
    return [c priority];
}

void LayoutConstraint_SetPriority(void* ptr, float priority) {
    NSLayoutConstraint* c = (NSLayoutConstraint*)ptr;
    [c setPriority:priority];
}

const char* LayoutConstraint_Identifier(void* ptr) {
    NSLayoutConstraint* c = (NSLayoutConstraint*)ptr;
    return [[c identifier] UTF8String];
}

bool LayoutConstraint_ShouldBeArchived(void* ptr) {
    NSLayoutConstraint* c = (NSLayoutConstraint*)ptr;
    return [c shouldBeArchived];
}

void LayoutConstraint_SetShouldBeArchived(void* ptr, bool value) {
    NSLayoutConstraint* c = (NSLayoutConstraint*)ptr;
    [c setShouldBeArchived:value];
}

void LayoutConstraint_ActivateConstraints(Array constraints) {
    NSMutableArray* objcArray = [[NSMutableArray alloc] init];
    NSLayoutConstraint** data = (NSLayoutConstraint**)constraints.data;
    for (int i = 0; i < constraints.len; i++) {
        [objcArray addObject:data[i]];
    }
    return [NSLayoutConstraint activateConstraints:objcArray];
}

void LayoutConstraint_DeactivateConstraints(Array constraints) {
    NSMutableArray* objcArray = [[NSMutableArray alloc] init];
    NSLayoutConstraint** data = (NSLayoutConstraint**)constraints.data;
    for (int i = 0; i < constraints.len; i++) {
        [objcArray addObject:data[i]];
    }
    return [NSLayoutConstraint deactivateConstraints:objcArray];
}
