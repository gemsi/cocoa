#import <Appkit/NSScrollView.h>
        #import "scroll_view.h"
        #include "_cgo_export.h"


void* ScrollView_New(long id, NSRect frame) {
    NSScrollView* scroll_view = [[[NSScrollView alloc] initWithFrame:frame] autorelease];
    return (void*)scroll_view;
}
bool ScrollView_HasVerticalScroller(void* ptr) {
	NSScrollView* scrollView = (NSScrollView*)ptr;
	return [scrollView hasVerticalScroller];
}
void ScrollView_SetHasVerticalScroller(void* ptr, bool hasVerticalScroller) {
	NSScrollView* scrollView = (NSScrollView*)ptr;
	[scrollView setHasVerticalScroller:hasVerticalScroller];
}
bool ScrollView_HasHorizontalScroller(void* ptr) {
	NSScrollView* scrollView = (NSScrollView*)ptr;
	return [scrollView hasHorizontalScroller];
}
void ScrollView_SetHasHorizontalScroller(void* ptr, bool hasHorizontalScroller) {
	NSScrollView* scrollView = (NSScrollView*)ptr;
	[scrollView setHasHorizontalScroller:hasHorizontalScroller];
}
void* ScrollView_DocumentView(void* ptr) {
	NSScrollView* scrollView = (NSScrollView*)ptr;
	return [scrollView documentView];
}
void ScrollView_SetDocumentView(void* ptr, void* documentView) {
	NSScrollView* scrollView = (NSScrollView*)ptr;
	[scrollView setDocumentView:documentView];
}
unsigned long ScrollView_BorderType(void* ptr) {
	NSScrollView* scrollView = (NSScrollView*)ptr;
	return [scrollView borderType];
}
void ScrollView_SetBorderType(void* ptr, unsigned long borderType) {
	NSScrollView* scrollView = (NSScrollView*)ptr;
	[scrollView setBorderType:borderType];
}
NSSize ScrollView_ContentSize(void* ptr) {
	NSScrollView* scrollView = (NSScrollView*)ptr;
	return [scrollView contentSize];
}
