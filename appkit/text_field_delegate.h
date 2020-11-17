#import <Appkit/NSTextField.h>
#import "_cgo_export.h"

@interface GoNSTextFieldDelegate : NSObject <NSTextFieldDelegate>
@property (assign) long goID;
@end

@implementation GoNSTextFieldDelegate

- (void)controlTextDidChange:(NSNotification*)notification {
	return TextField_Delegate_ControlTextDidChange([self goID], notification);
}

- (void)controlTextDidEndEditing:(NSNotification*)notification {
	return TextField_Delegate_ControlTextDidEndEditing([self goID], notification);
}

- (void)controlTextDidBeginEditing:(NSNotification*)notification {
	return TextField_Delegate_ControlTextDidBeginEditing([self goID], notification);
}

@end
