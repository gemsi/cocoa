#import "text_field_delegate.h"
#include "_cgo_export.h"


const int CONTROL_TEXT_DID_CHANGE           = 0;
const int CONTROL_TEXT_DID_END_EDITING      = 1;

@implementation TextFieldDelegate

- (void)controlTextDidChange:(NSNotification*)notification {
	triggerTextFieldEvent([self goID], notification, CONTROL_TEXT_DID_CHANGE);
}

// - (BOOL)becomeFirstResponder {
// 	BOOL rc = [super becomeFirstResponder];
// 	if (rc) {
// 		textfieldOnFocus(self);
// 	}
// 	return rc;
// }

- (void)controlTextDidEndEditing:(NSNotification*)notification {
	triggerTextFieldEvent([self goID], notification, CONTROL_TEXT_DID_END_EDITING);
}

// - (void)onEnterKey:(id)sender {
// 	NSString* v = [self stringValue];
// 	textfieldOnEnterKey(self, (char*)[v cStringUsingEncoding:NSUTF8StringEncoding]);
// }

@end

void triggerTextFieldEvent(int goID, NSNotification *notification, int eventType) {
    onTextFieldEvent(goID, (void*)notification, eventType);
}