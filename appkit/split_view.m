#import <AppKit/AppKit.h>
#import "split_view.h"

bool SplitView_ArrangesAllSubviews(void* ptr) {
	NSSplitView* splitView = (NSSplitView*)ptr;
	return [splitView arrangesAllSubviews];
}

void SplitView_SetArrangesAllSubviews(void* ptr, bool arrangesAllSubviews) {
	NSSplitView* splitView = (NSSplitView*)ptr;
	[splitView setArrangesAllSubviews:arrangesAllSubviews];
}

bool SplitView_IsVertical(void* ptr) {
	NSSplitView* splitView = (NSSplitView*)ptr;
	return [splitView isVertical];
}

void SplitView_SetVertical(void* ptr, bool vertical) {
	NSSplitView* splitView = (NSSplitView*)ptr;
	[splitView setVertical:vertical];
}

void* SplitView_DividerColor(void* ptr) {
	NSSplitView* splitView = (NSSplitView*)ptr;
	return [splitView dividerColor];
}

double SplitView_DividerThickness(void* ptr) {
	NSSplitView* splitView = (NSSplitView*)ptr;
	return [splitView dividerThickness];
}

long SplitView_DividerStyle(void* ptr) {
	NSSplitView* splitView = (NSSplitView*)ptr;
	return [splitView dividerStyle];
}

void SplitView_SetDividerStyle(void* ptr, long dividerStyle) {
	NSSplitView* splitView = (NSSplitView*)ptr;
	[splitView setDividerStyle:dividerStyle];
}

void* SplitView_NewSplitView(NSRect frame) {
	NSSplitView* splitView = [NSSplitView alloc];
	return [[splitView initWithFrame:frame] autorelease];
}

double SplitView_MinPossiblePositionOfDividerAtIndex(void* ptr, long dividerIndex) {
	NSSplitView* splitView = (NSSplitView*)ptr;
	return [splitView minPossiblePositionOfDividerAtIndex:dividerIndex];
}

double SplitView_MaxPossiblePositionOfDividerAtIndex(void* ptr, long dividerIndex) {
	NSSplitView* splitView = (NSSplitView*)ptr;
	return [splitView maxPossiblePositionOfDividerAtIndex:dividerIndex];
}

void SplitView_SetPosition(void* ptr, double position, long dividerIndex) {
	NSSplitView* splitView = (NSSplitView*)ptr;
	[splitView setPosition:position ofDividerAtIndex:dividerIndex];
}
