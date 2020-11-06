#import "view.h"
#include "_cgo_export.h"

void View_SetFrame(void* viewPtr, int x, int y, int w, int h) {
    NSView* view = (NSView*)viewPtr;
    [view setFrame:NSMakeRect(x, y, w, h)];
}
