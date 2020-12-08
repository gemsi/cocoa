#import <AppKit/AppKit.h>
#import "popover.h"

long Popover_Behavior(void* ptr) {
	NSPopover* popover = (NSPopover*)ptr;
	return [popover behavior];
}

void Popover_SetBehavior(void* ptr, long behavior) {
	NSPopover* popover = (NSPopover*)ptr;
	[popover setBehavior:behavior];
}

NSRect Popover_PositioningRect(void* ptr) {
	NSPopover* popover = (NSPopover*)ptr;
	return [popover positioningRect];
}

void Popover_SetPositioningRect(void* ptr, NSRect positioningRect) {
	NSPopover* popover = (NSPopover*)ptr;
	[popover setPositioningRect:positioningRect];
}

bool Popover_Animates(void* ptr) {
	NSPopover* popover = (NSPopover*)ptr;
	return [popover animates];
}

void Popover_SetAnimates(void* ptr, bool animates) {
	NSPopover* popover = (NSPopover*)ptr;
	[popover setAnimates:animates];
}

NSSize Popover_ContentSize(void* ptr) {
	NSPopover* popover = (NSPopover*)ptr;
	return [popover contentSize];
}

void Popover_SetContentSize(void* ptr, NSSize contentSize) {
	NSPopover* popover = (NSPopover*)ptr;
	[popover setContentSize:contentSize];
}

bool Popover_IsShown(void* ptr) {
	NSPopover* popover = (NSPopover*)ptr;
	return [popover isShown];
}

bool Popover_IsDetached(void* ptr) {
	NSPopover* popover = (NSPopover*)ptr;
	return [popover isDetached];
}

void* Popover_Appearance(void* ptr) {
	NSPopover* popover = (NSPopover*)ptr;
	return [popover appearance];
}

void Popover_SetAppearance(void* ptr, void* appearance) {
	NSPopover* popover = (NSPopover*)ptr;
	[popover setAppearance:(NSAppearance*)appearance];
}

void* Popover_EffectiveAppearance(void* ptr) {
	NSPopover* popover = (NSPopover*)ptr;
	return [popover effectiveAppearance];
}

void* Popover_NewPopover() {
	NSPopover* popover = [NSPopover alloc];
	return [[popover init] autorelease];
}

void Popover_PerformClose(void* ptr, void* sender) {
	NSPopover* popover = (NSPopover*)ptr;
	[popover performClose:(NSObject*)sender];
}

void Popover_Close(void* ptr) {
	NSPopover* popover = (NSPopover*)ptr;
	[popover close];
}

void Popover_ShowRelativeTo(void* ptr, NSRect positioningRect, void* positioningView, long preferredEdge) {
	NSPopover* popover = (NSPopover*)ptr;
	[popover showRelativeToRect:positioningRect ofView:(NSView*)positioningView preferredEdge:preferredEdge];
}
