#import <Appkit/NSStackView.h>
#import "stack_view.h"

void* StackView_stackViewWithViews(long goID, void** views, size_t views_len) {
	NSMutableArray* objcViews = [[NSMutableArray alloc] init];;
	for (int i = 0; i < views_len; i++) {
		[objcViews addObject:(NSView*)views[i]];
	}
	NSStackView* stack_view = [[NSStackView stackViewWithViews:objcViews] autorelease];
	return (void*)stack_view;
}

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
