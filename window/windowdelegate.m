#import "windowdelegate.h"
#include "_cgo_export.h"


const int DID_RESIZE_EVENT           = 0;
const int DID_MOVE_EVENT             = 1;
const int DID_MINIATURIZE_EVENT      = 2;
const int DID_DEMINIATURIZE_EVENT    = 3;

@implementation WindowDelegate

- (void)dealloc
{
    [super dealloc];
}

- (void)windowDidResize:(NSNotification *)notification
{
    triggerEvent([self goWindowID], notification, DID_RESIZE_EVENT);
}

- (void)windowDidMove:(NSNotification *)notification
{
    triggerEvent([self goWindowID], notification, DID_MOVE_EVENT);
}

- (void)windowDidMiniaturize:(NSNotification *)notification
{
    triggerEvent([self goWindowID], notification, DID_MINIATURIZE_EVENT);
}

- (void)windowDidDeminiaturize:(NSNotification *)notification
{
    triggerEvent([self goWindowID], notification, DID_DEMINIATURIZE_EVENT);
}

@end

void triggerEvent(int goWindowID, NSNotification *notification, int eventType)
{
    onWindowEvent(goWindowID, (void*)notification, eventType);
}