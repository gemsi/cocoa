#import <Appkit/NSSecureTextField.h>
#import "secure_text_field.h"

void* SecureTextField_initWithFrame(long goID, NSRect frame) {
	NSSecureTextField* secure_text_field = [[[NSSecureTextField alloc] initWithFrame:frame] autorelease];
	return (void*)secure_text_field;
}
