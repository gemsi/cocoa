#import <Cocoa/Cocoa.h>
#import "view.h"
#include "_cgo_export.h"

void View_SetFrame(void* viewPtr, NSRect rect) {
    NSView* view = (NSView*)viewPtr;
    [view setFrame:rect];
}

NSRect View_Frame(void* viewPtr) {
    NSView* view = (NSView*)viewPtr;
    return [view frame];
}
