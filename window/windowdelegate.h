#import <Cocoa/Cocoa.h>

@interface WindowDelegate : NSObject <NSWindowDelegate>

@property (assign) int goWindowID;

@end

void triggerEvent(int goWindowID, NSNotification *notification, int eventType);