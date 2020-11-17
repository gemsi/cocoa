#import <Appkit/NSButton.h>
#import "_cgo_export.h"

@interface NSButtonHandler : NSObject
@property (assign) long goID;
@end

@implementation NSButtonHandler

- (void)onAction:(NSObject*)sender {
	return Button_Target_Action([self goID], sender);
}

@end
