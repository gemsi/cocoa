#import "window_delegate.h"
#include "_cgo_export.h"


const int DID_RESIZE_EVENT           = 0;
const int DID_MOVE_EVENT             = 1;
const int DID_MINIATURIZE_EVENT      = 2;
const int DID_DEMINIATURIZE_EVENT    = 3;

@implementation WindowDelegate

- (void)dealloc {
    [super dealloc];
}

- (void)windowDidResize:(NSNotification *)notification {
    triggerWindowEvent([self goID], notification, DID_RESIZE_EVENT);
}

- (void)windowDidMove:(NSNotification *)notification {
    triggerWindowEvent([self goID], notification, DID_MOVE_EVENT);
}

- (void)windowDidMiniaturize:(NSNotification *)notification {
    triggerWindowEvent([self goID], notification, DID_MINIATURIZE_EVENT);
}

- (void)windowDidDeminiaturize:(NSNotification *)notification {
    triggerWindowEvent([self goID], notification, DID_DEMINIATURIZE_EVENT);
}

@end

void triggerWindowEvent(int goID, NSNotification *notification, int eventType) {
    onWindowEvent(goID, (void*)notification, eventType);
}