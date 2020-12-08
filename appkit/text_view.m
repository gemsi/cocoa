#import <AppKit/AppKit.h>
#import "text_view.h"

@interface MyTextView : NSTextView {
}
@end

@implementation MyTextView


- (BOOL)performKeyEquivalent:(NSEvent *)event {
    if ([event type] == NSEventTypeKeyDown) {
        if (([event modifierFlags] & NSEventModifierFlagDeviceIndependentFlagsMask) == NSEventModifierFlagCommand) {
            if ([[event charactersIgnoringModifiers] isEqualToString:@"x"]) {
                if ([NSApp sendAction:@selector(cut:) to:nil from:self]) {
                    return YES;
                }
            } else if ([[event charactersIgnoringModifiers] isEqualToString:@"c"]) {
                if ([NSApp sendAction:@selector(copy:) to:nil from:self]) {
                    return YES;
                }
            } else if ([[event charactersIgnoringModifiers] isEqualToString:@"v"]) {
                if ([NSApp sendAction:@selector(paste:) to:nil from:self]) {
                    return YES;
                }
            } else if ([[event charactersIgnoringModifiers] isEqualToString:@"z"]) {
                if ([NSApp sendAction:@selector(undo:) to:nil from:self]) {
                    return YES;
                }
            } else if ([[event charactersIgnoringModifiers] isEqualToString:@"a"]) {
                if ([NSApp sendAction:@selector(selectAll:) to:nil from:self]) {
                    return YES;
                }
            }
        } else if (([event modifierFlags] & NSEventModifierFlagDeviceIndependentFlagsMask) == (NSEventModifierFlagCommand | NSEventModifierFlagShift)) {
            if ([[event charactersIgnoringModifiers] isEqualToString:@"Z"]) {
                if ([NSApp sendAction:@selector(redo:) to:nil from:self]) {
                    return YES;
                }
            }
        }
    }
    return [super performKeyEquivalent:event];
}
@end

bool TextView_AllowsUndo(void* ptr) {
	NSTextView* textView = (NSTextView*)ptr;
	return [textView allowsUndo];
}

void TextView_SetAllowsUndo(void* ptr, bool allowsUndo) {
	NSTextView* textView = (NSTextView*)ptr;
	[textView setAllowsUndo:allowsUndo];
}

void* TextView_TextContainer(void* ptr) {
	NSTextView* textView = (NSTextView*)ptr;
	return [textView textContainer];
}

void TextView_SetTextContainer(void* ptr, void* textContainer) {
	NSTextView* textView = (NSTextView*)ptr;
	[textView setTextContainer:(NSTextContainer*)textContainer];
}

void* TextView_NewTextView(NSRect frame) {
	NSTextView* textView = [MyTextView alloc];
	return [[textView initWithFrame:frame] autorelease];
}

void* TextView_ScrollableTextView() {
	return [NSTextView scrollableTextView];
}
