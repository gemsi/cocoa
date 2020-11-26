#import <Foundation/NSGeometry.h>
#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>

Array StackView_Views(void* ptr);
Array StackView_DetachedViews(void* ptr);
long StackView_Orientation(void* ptr);
void StackView_SetOrientation(void* ptr, long orientation);
long StackView_Alignment(void* ptr);
void StackView_SetAlignment(void* ptr, long alignment);
double StackView_Spacing(void* ptr);
void StackView_SetSpacing(void* ptr, double spacing);
NSEdgeInsets StackView_EdgeInsets(void* ptr);
void StackView_SetEdgeInsets(void* ptr, NSEdgeInsets edgeInsets);
long StackView_Distribution(void* ptr);
void StackView_SetDistribution(void* ptr, long distribution);
Array StackView_ArrangedSubviews(void* ptr);
bool StackView_DetachesHiddenViews(void* ptr);
void StackView_SetDetachesHiddenViews(void* ptr, bool detachesHiddenViews);

void* StackView_StackViewWithViews(Array views);
Array StackView_ViewsInGravity(void* ptr, long gravity);
void StackView_AddView(void* ptr, void* view, long gravity);
void StackView_InsertView(void* ptr, void* view, unsigned long index, long gravity);
void StackView_SetViews(void* ptr, Array views, long gravity);
void StackView_RemoveView(void* ptr, void* view);
void StackView_AddArrangedSubview(void* ptr, void* view);
void StackView_InsertArrangedSubview(void* ptr, void* view, unsigned long index);
void StackView_RemoveArrangedSubview(void* ptr, void* view);
float StackView_ClippingResistancePriorityForOrientation(void* ptr, long orientation);
void StackView_SetClippingResistancePriority(void* ptr, float clippingResistancePriority, long orientation);
float StackView_HuggingPriorityForOrientation(void* ptr, long orientation);
void StackView_SetHuggingPriority(void* ptr, float huggingPriority, long orientation);
double StackView_CustomSpacingAfterView(void* ptr, void* view);
void StackView_SetCustomSpacing(void* ptr, double spacing, void* view);
float StackView_VisibilityPriorityForView(void* ptr, void* view);
void StackView_SetVisibilityPriority(void* ptr, float priority, void* view);
