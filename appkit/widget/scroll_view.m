#import <Appkit/NSScrollView.h>
        #import "scroll_view.h"
        #include "_cgo_export.h"


void* ScrollView_New(long id, NSRect frame) {
    NSScrollView* scroll_view = [[[NSScrollView alloc] initWithFrame:frame] autorelease];
    return (void*)scroll_view;
}

bool ScrollView_HasVerticalScroller(void* ptr) {
    NSScrollView* scroll_view = (NSScrollView*)ptr;
    return scroll_view.hasVerticalScroller;
}

void ScrollView_SetHasVerticalScroller(void* ptr, bool value) {
    NSScrollView* scroll_view = (NSScrollView*)ptr;
    return [scroll_view setHasVerticalScroller:value];
}

bool ScrollView_HasHorizontalScroller(void* ptr) {
    NSScrollView* scroll_view = (NSScrollView*)ptr;
    return scroll_view.hasHorizontalScroller;
}

void ScrollView_SetHasHorizontalScroller(void* ptr, bool value) {
    NSScrollView* scroll_view = (NSScrollView*)ptr;
    return [scroll_view setHasHorizontalScroller:value];
}

void* ScrollView_DocumentView(void* ptr) {
    NSScrollView* scroll_view = (NSScrollView*)ptr;
    return scroll_view.documentView;
}

void ScrollView_SetDocumentView(void* ptr, void* valuePtr) {
    NSScrollView* scroll_view = (NSScrollView*)ptr;
    NSView* value = (NSView*)valuePtr;
    [scroll_view setDocumentView:value];
}

unsigned long ScrollView_BorderType(void* ptr) {
    NSScrollView* scroll_view = (NSScrollView*)ptr;
    return scroll_view.borderType;
}

void ScrollView_SetBorderType(void* ptr, unsigned long value) {
    NSScrollView* scroll_view = (NSScrollView*)ptr;
    return [scroll_view setBorderType:value];
}

NSSize ScrollView_ContentSize(void* ptr) {
    NSScrollView* scroll_view = (NSScrollView*)ptr;
    return scroll_view.contentSize;
}
