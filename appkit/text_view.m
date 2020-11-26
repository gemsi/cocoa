#import <Appkit/NSTextView.h>
#import "text_view.h"

void* TextView_TextContainer(void* ptr) {
	NSTextView* textView = (NSTextView*)ptr;
	return [textView textContainer];
}

void TextView_SetTextContainer(void* ptr, void* textContainer) {
	NSTextView* textView = (NSTextView*)ptr;
	[textView setTextContainer:textContainer];
}

void* TextView_InitWithFrame(NSRect frame) {
	NSTextView* textView = [NSTextView alloc];
	return [[textView initWithFrame:frame] autorelease];
}
