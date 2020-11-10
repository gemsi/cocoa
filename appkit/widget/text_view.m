#import <Appkit/NSTextView.h>
#import "text_view.h"
#include "_cgo_export.h"

void* TextView_New(long id, NSRect frame) {
    NSTextView* view = [[[NSTextView alloc] initWithFrame:frame] autorelease];
    return (void*)view;
}