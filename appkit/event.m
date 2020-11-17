#import <Appkit/NSEvent.h>
#import "event.h"

NSPoint Event_LocationInWindow(void* ptr) {
	NSEvent* event = (NSEvent*)ptr;
	return [event locationInWindow];
}

void* Event_Window(void* ptr) {
	NSEvent* event = (NSEvent*)ptr;
	return [event window];
}

long Event_WindowNumber(void* ptr) {
	NSEvent* event = (NSEvent*)ptr;
	return [event windowNumber];
}

double Event_Timestamp(void* ptr) {
	NSEvent* event = (NSEvent*)ptr;
	return [event timestamp];
}
