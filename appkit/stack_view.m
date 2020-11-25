#import <Appkit/NSStackView.h>
#import "stack_view.h"

void* StackView_stackViewWithViews(long goID, Array views) {
	NSMutableArray* objcViews = [[NSMutableArray alloc] init];
	NSView** viewsData = (NSView**)views.data;
	for (int i = 0; i < views.len; i++) {
	    [objcViews addObject:viewsData[i]];
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
