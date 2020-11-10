#import <Appkit/NSTextView.h>
        #import "text_view.h"
        #include "_cgo_export.h"


void* TextView_New(long id, NSRect frame) {
    NSTextView* text_view = [[[NSTextView alloc] initWithFrame:frame] autorelease];
    return (void*)text_view;
}

void* TextView_TextContainer(void* ptr) {
    NSTextView* text_view = (NSTextView*)ptr;
    return text_view.textContainer;
}

void TextView_SetTextContainer(void* ptr, void* valuePtr) {
    NSTextView* text_view = (NSTextView*)ptr;
    NSTextContainer* value = (NSTextContainer*)valuePtr;
    [text_view setTextContainer:value];
}
