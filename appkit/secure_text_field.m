#import <Appkit/NSSecureTextField.h>
#import "secure_text_field.h"
#import "secure_text_field_delegate.h"

@implementation GoNSSecureTextFieldDelegate

@end
void SecureTextField_RegisterDelegate(void *ptr, long goID) {
	NSSecureTextField* secureTextField = (NSSecureTextField*)ptr;
	GoNSSecureTextFieldDelegate* delegate = [[[GoNSSecureTextFieldDelegate alloc] init] autorelease];
	[delegate setGoID:goID];
	[secureTextField setDelegate:delegate];
}

void* SecureTextField_InitWithFrame(NSRect frame) {
	NSSecureTextField* secureTextField = [NSSecureTextField alloc];
	return [[secureTextField initWithFrame:frame] autorelease];
}
