#import <Appkit/NSStackView.h>
#import "stack_view.h"

Array StackView_Views(void* ptr) {
	NSStackView* stackView = (NSStackView*)ptr;
	NSArray* array = [stackView views];
	int count = [array count];
	void** data = malloc(count * sizeof(void*));
	for (int i = 0; i < count; i++) {
		 data[i] = [array objectAtIndex:i];
	}
	Array result;
	result.data = data;
	result.len = count;
	return result;
}

Array StackView_DetachedViews(void* ptr) {
	NSStackView* stackView = (NSStackView*)ptr;
	NSArray* array = [stackView detachedViews];
	int count = [array count];
	void** data = malloc(count * sizeof(void*));
	for (int i = 0; i < count; i++) {
		 data[i] = [array objectAtIndex:i];
	}
	Array result;
	result.data = data;
	result.len = count;
	return result;
}

long StackView_Orientation(void* ptr) {
	NSStackView* stackView = (NSStackView*)ptr;
	return [stackView orientation];
}

void StackView_SetOrientation(void* ptr, long orientation) {
	NSStackView* stackView = (NSStackView*)ptr;
	[stackView setOrientation:orientation];
}

long StackView_Alignment(void* ptr) {
	NSStackView* stackView = (NSStackView*)ptr;
	return [stackView alignment];
}

void StackView_SetAlignment(void* ptr, long alignment) {
	NSStackView* stackView = (NSStackView*)ptr;
	[stackView setAlignment:alignment];
}

double StackView_Spacing(void* ptr) {
	NSStackView* stackView = (NSStackView*)ptr;
	return [stackView spacing];
}

void StackView_SetSpacing(void* ptr, double spacing) {
	NSStackView* stackView = (NSStackView*)ptr;
	[stackView setSpacing:spacing];
}

NSEdgeInsets StackView_EdgeInsets(void* ptr) {
	NSStackView* stackView = (NSStackView*)ptr;
	return [stackView edgeInsets];
}

void StackView_SetEdgeInsets(void* ptr, NSEdgeInsets edgeInsets) {
	NSStackView* stackView = (NSStackView*)ptr;
	[stackView setEdgeInsets:edgeInsets];
}

long StackView_Distribution(void* ptr) {
	NSStackView* stackView = (NSStackView*)ptr;
	return [stackView distribution];
}

void StackView_SetDistribution(void* ptr, long distribution) {
	NSStackView* stackView = (NSStackView*)ptr;
	[stackView setDistribution:distribution];
}

Array StackView_ArrangedSubviews(void* ptr) {
	NSStackView* stackView = (NSStackView*)ptr;
	NSArray* array = [stackView arrangedSubviews];
	int count = [array count];
	void** data = malloc(count * sizeof(void*));
	for (int i = 0; i < count; i++) {
		 data[i] = [array objectAtIndex:i];
	}
	Array result;
	result.data = data;
	result.len = count;
	return result;
}

bool StackView_DetachesHiddenViews(void* ptr) {
	NSStackView* stackView = (NSStackView*)ptr;
	return [stackView detachesHiddenViews];
}

void StackView_SetDetachesHiddenViews(void* ptr, bool detachesHiddenViews) {
	NSStackView* stackView = (NSStackView*)ptr;
	[stackView setDetachesHiddenViews:detachesHiddenViews];
}

void* StackView_StackViewWithViews(Array views) {
    NSMutableArray* objcViews = [[NSMutableArray alloc] init];
    void** viewsData = (void**)views.data;
    for (int i = 0; i < views.len; i++) {
    	[objcViews addObject:(NSView*)viewsData[i]];
    }
	return [NSStackView stackViewWithViews:objcViews];
}

Array StackView_ViewsInGravity(void* ptr, long gravity) {
	NSStackView* stackView = (NSStackView*)ptr;
	NSArray* array = [stackView viewsInGravity:gravity];
	int count = [array count];
	void** data = malloc(count * sizeof(void*));
	for (int i = 0; i < count; i++) {
		 data[i] = [array objectAtIndex:i];
	}
	Array result;
	result.data = data;
	result.len = count;
	return result;
}

void StackView_AddView(void* ptr, void* view, long gravity) {
	NSStackView* stackView = (NSStackView*)ptr;
	[stackView addView:(NSView*)view inGravity:gravity];
}

void StackView_InsertView(void* ptr, void* view, unsigned long index, long gravity) {
	NSStackView* stackView = (NSStackView*)ptr;
	[stackView insertView:(NSView*)view atIndex:index inGravity:gravity];
}

void StackView_SetViews(void* ptr, Array views, long gravity) {
	NSStackView* stackView = (NSStackView*)ptr;
    NSMutableArray* objcViews = [[NSMutableArray alloc] init];
    void** viewsData = (void**)views.data;
    for (int i = 0; i < views.len; i++) {
    	[objcViews addObject:(NSView*)viewsData[i]];
    }
	[stackView setViews:objcViews inGravity:gravity];
}

void StackView_RemoveView(void* ptr, void* view) {
	NSStackView* stackView = (NSStackView*)ptr;
	[stackView removeView:(NSView*)view];
}

void StackView_AddArrangedSubview(void* ptr, void* view) {
	NSStackView* stackView = (NSStackView*)ptr;
	[stackView addArrangedSubview:(NSView*)view];
}

void StackView_InsertArrangedSubview(void* ptr, void* view, unsigned long index) {
	NSStackView* stackView = (NSStackView*)ptr;
	[stackView insertArrangedSubview:(NSView*)view atIndex:index];
}

void StackView_RemoveArrangedSubview(void* ptr, void* view) {
	NSStackView* stackView = (NSStackView*)ptr;
	[stackView removeArrangedSubview:(NSView*)view];
}

float StackView_ClippingResistancePriorityForOrientation(void* ptr, long orientation) {
	NSStackView* stackView = (NSStackView*)ptr;
	return [stackView clippingResistancePriorityForOrientation:orientation];
}

void StackView_SetClippingResistancePriority(void* ptr, float clippingResistancePriority, long orientation) {
	NSStackView* stackView = (NSStackView*)ptr;
	[stackView setClippingResistancePriority:clippingResistancePriority forOrientation:orientation];
}

float StackView_HuggingPriorityForOrientation(void* ptr, long orientation) {
	NSStackView* stackView = (NSStackView*)ptr;
	return [stackView huggingPriorityForOrientation:orientation];
}

void StackView_SetHuggingPriority(void* ptr, float huggingPriority, long orientation) {
	NSStackView* stackView = (NSStackView*)ptr;
	[stackView setHuggingPriority:huggingPriority forOrientation:orientation];
}

double StackView_CustomSpacingAfterView(void* ptr, void* view) {
	NSStackView* stackView = (NSStackView*)ptr;
	return [stackView customSpacingAfterView:(NSView*)view];
}

void StackView_SetCustomSpacing(void* ptr, double spacing, void* view) {
	NSStackView* stackView = (NSStackView*)ptr;
	[stackView setCustomSpacing:spacing afterView:(NSView*)view];
}

float StackView_VisibilityPriorityForView(void* ptr, void* view) {
	NSStackView* stackView = (NSStackView*)ptr;
	return [stackView visibilityPriorityForView:(NSView*)view];
}

void StackView_SetVisibilityPriority(void* ptr, float priority, void* view) {
	NSStackView* stackView = (NSStackView*)ptr;
	[stackView setVisibilityPriority:priority forView:(NSView*)view];
}
