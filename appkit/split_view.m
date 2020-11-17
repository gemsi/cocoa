#import <Appkit/NSSplitView.h>
#import "split_view.h"

void* SplitView_initWithFrame(long goID, NSRect frame) {
	NSSplitView* split_view = [[[NSSplitView alloc] initWithFrame:frame] autorelease];
	return (void*)split_view;
}

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

double SplitView_MinPossiblePositionOfDividerAtIndex(void* ptr, int dividerIndex) {
	NSSplitView* splitView = (NSSplitView*)ptr;
	return [splitView minPossiblePositionOfDividerAtIndex:dividerIndex];
}
