#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>

bool LayoutConstraint_IsActive(void* ptr);
void LayoutConstraint_SetActive(void* ptr, bool active);
void* LayoutConstraint_FirstItem(void* ptr);
long LayoutConstraint_FirstAttribute(void* ptr);
long LayoutConstraint_Relation(void* ptr);
void* LayoutConstraint_SecondItem(void* ptr);
long LayoutConstraint_SecondAttribute(void* ptr);
double LayoutConstraint_Multiplier(void* ptr);
double LayoutConstraint_Constant(void* ptr);
void* LayoutConstraint_FirstAnchor(void* ptr);
void* LayoutConstraint_SecondAnchor(void* ptr);
float LayoutConstraint_Priority(void* ptr);
void LayoutConstraint_SetPriority(void* ptr, float priority);
const char* LayoutConstraint_Identifier(void* ptr);
void LayoutConstraint_SetIdentifier(void* ptr, const char* identifier);
bool LayoutConstraint_ShouldBeArchived(void* ptr);
void LayoutConstraint_SetShouldBeArchived(void* ptr, bool shouldBeArchived);

void LayoutConstraint_ConstraintWithItem(void* view1, long attr1, long relation, void* view2, long attr2, double multiplier, double c);
void LayoutConstraint_ActivateConstraints(Array constraints);
void LayoutConstraint_DeactivateConstraints(Array constraints);
