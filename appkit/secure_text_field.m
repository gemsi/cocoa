#import <Appkit/NSSecureTextField.h>
#import "secure_text_field.h"
#import "secure_text_field_delegate.h"

@implementation GoNSSecureTextFieldDelegate

@end

void* SecureTextField_initWithFrame(long goID, NSRect frame) {
	NSSecureTextField* secure_text_field = [[[NSSecureTextField alloc] initWithFrame:frame] autorelease];
	GoNSSecureTextFieldDelegate* delegate = [[GoNSSecureTextFieldDelegate alloc] init];
	[delegate setGoID:goID];
	[secure_text_field setDelegate:delegate];
	return (void*)secure_text_field;
}
