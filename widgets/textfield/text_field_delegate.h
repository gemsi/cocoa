#import <Cocoa/Cocoa.h>

@interface TextFieldDelegate : NSObject <NSTextFieldDelegate>

@property (assign) int goID;

@end

void triggerTextFieldEvent(int goID, NSNotification *notification, int eventType);