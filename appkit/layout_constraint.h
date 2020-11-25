#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>

void* LayoutConstraint_New(void* view1, int attr1, int relation, void* view2, int attr2, double multiplier, double constant);

bool LayoutConstraint_IsActive(void* ptr);
void LayoutConstraint_SetActive(void* ptr, bool active);

void* LayoutConstraint_FirstItem(void* ptr);
int LayoutConstraint_FirstAttribute(void* ptr);
int LayoutConstraint_Relation(void* ptr);
void* LayoutConstraint_SecondItem(void* ptr);
int LayoutConstraint_SecondAttribute(void* ptr);
double LayoutConstraint_Multiplier(void* ptr);
double LayoutConstraint_Constant(void* ptr);
void LayoutConstraint_SetConstant(void* ptr, double constant);
float LayoutConstraint_Priority(void* ptr);
void LayoutConstraint_SetPriority(void* ptr, float priority);
const char* LayoutConstraint_Identifier(void* ptr);
bool LayoutConstraint_ShouldBeArchived(void* ptr);
void LayoutConstraint_SetShouldBeArchived(void* ptr, bool value);

void LayoutConstraint_ActivateConstraints(Array constraints);
void LayoutConstraint_DeactivateConstraints(Array constraints);

