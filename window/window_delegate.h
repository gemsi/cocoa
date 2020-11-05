#import <Cocoa/Cocoa.h>

@interface WindowDelegate : NSObject <NSWindowDelegate>

@property (assign) int goID;

@end

void triggerWindowEvent(int goID, NSNotification *notification, int eventType);