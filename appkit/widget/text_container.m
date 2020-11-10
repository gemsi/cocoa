#import <Appkit/NSTextContainer.h>
        #import "text_container.h"
        #include "_cgo_export.h"


void* TextContainer_New(long id, NSSize size) {
    NSTextContainer* text_container = [[[NSTextContainer alloc] initWithSize:size] autorelease];
    return (void*)text_container;
}

NSSize TextContainer_Size(void* ptr) {
    NSTextContainer* text_container = (NSTextContainer*)ptr;
    return text_container.size;
}

void TextContainer_SetSize(void* ptr, NSSize value) {
    NSTextContainer* text_container = (NSTextContainer*)ptr;
    [text_container setSize:value];
}

bool TextContainer_WidthTracksTextView(void* ptr) {
    NSTextContainer* text_container = (NSTextContainer*)ptr;
    return text_container.widthTracksTextView;
}

void TextContainer_SetWidthTracksTextView(void* ptr, bool value) {
    NSTextContainer* text_container = (NSTextContainer*)ptr;
    return [text_container setWidthTracksTextView:value];
}

bool TextContainer_HeightTracksTextView(void* ptr) {
    NSTextContainer* text_container = (NSTextContainer*)ptr;
    return text_container.heightTracksTextView;
}

void TextContainer_SetHeightTracksTextView(void* ptr, bool value) {
    NSTextContainer* text_container = (NSTextContainer*)ptr;
    return [text_container setHeightTracksTextView:value];
}
