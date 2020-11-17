#import <Appkit/NSTextView.h>
#import "text_view.h"

void* TextView_initWithFrame(long goID, NSRect frame) {
	NSTextView* text_view = [[[NSTextView alloc] initWithFrame:frame] autorelease];
	return (void*)text_view;
}

void* TextView_TextContainer(void* ptr) {
	NSTextView* textView = (NSTextView*)ptr;
	return [textView textContainer];
}

void TextView_SetTextContainer(void* ptr, void* textContainer) {
	NSTextView* textView = (NSTextView*)ptr;
	[textView setTextContainer:textContainer];
}
