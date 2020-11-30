#import <Appkit/NSResponder.h>
#import "responder.h"

bool Responder_AcceptsFirstResponder(void* ptr) {
	NSResponder* responder = (NSResponder*)ptr;
	return [responder acceptsFirstResponder];
}

void* Responder_NextResponder(void* ptr) {
	NSResponder* responder = (NSResponder*)ptr;
	return [responder nextResponder];
}

void Responder_SetNextResponder(void* ptr, void* nextResponder) {
	NSResponder* responder = (NSResponder*)ptr;
	[responder setNextResponder:(NSResponder*)nextResponder];
}

bool Responder_BecomeFirstResponder(void* ptr) {
	NSResponder* responder = (NSResponder*)ptr;
	return [responder becomeFirstResponder];
}
