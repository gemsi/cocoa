#import <Appkit/NSTextContainer.h>
#import "text_container.h"

NSSize TextContainer_Size(void* ptr) {
	NSTextContainer* textContainer = (NSTextContainer*)ptr;
	return [textContainer size];
}

void TextContainer_SetSize(void* ptr, NSSize size) {
	NSTextContainer* textContainer = (NSTextContainer*)ptr;
	[textContainer setSize:size];
}

bool TextContainer_WidthTracksTextView(void* ptr) {
	NSTextContainer* textContainer = (NSTextContainer*)ptr;
	return [textContainer widthTracksTextView];
}

void TextContainer_SetWidthTracksTextView(void* ptr, bool widthTracksTextView) {
	NSTextContainer* textContainer = (NSTextContainer*)ptr;
	[textContainer setWidthTracksTextView:widthTracksTextView];
}

bool TextContainer_HeightTracksTextView(void* ptr) {
	NSTextContainer* textContainer = (NSTextContainer*)ptr;
	return [textContainer heightTracksTextView];
}

void TextContainer_SetHeightTracksTextView(void* ptr, bool heightTracksTextView) {
	NSTextContainer* textContainer = (NSTextContainer*)ptr;
	[textContainer setHeightTracksTextView:heightTracksTextView];
}

void* TextContainer_InitWithSize(NSSize size) {
	NSTextContainer* textContainer = [NSTextContainer alloc];
	return [[textContainer initWithSize:size] autorelease];
}
