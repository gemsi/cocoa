#import <Appkit/NSTabView.h>
#import "tab_view.h"

void* TabView_initWithFrame(long goID, NSRect frame) {
	NSTabView* tab_view = [[[NSTabView alloc] initWithFrame:frame] autorelease];
	return (void*)tab_view;
}

long TabView_NumberOfTabViewItems(void* ptr) {
	NSTabView* tabView = (NSTabView*)ptr;
	return [tabView numberOfTabViewItems];
}

unsigned long TabView_TabViewType(void* ptr) {
	NSTabView* tabView = (NSTabView*)ptr;
	return [tabView tabViewType];
}

void TabView_SetTabViewType(void* ptr, unsigned long tabViewType) {
	NSTabView* tabView = (NSTabView*)ptr;
	[tabView setTabViewType:tabViewType];
}

unsigned long TabView_TabPosition(void* ptr) {
	NSTabView* tabView = (NSTabView*)ptr;
	return [tabView tabPosition];
}

void TabView_SetTabPosition(void* ptr, unsigned long tabPosition) {
	NSTabView* tabView = (NSTabView*)ptr;
	[tabView setTabPosition:tabPosition];
}

unsigned long TabView_TabViewBorderType(void* ptr) {
	NSTabView* tabView = (NSTabView*)ptr;
	return [tabView tabViewBorderType];
}

void TabView_SetTabViewBorderType(void* ptr, unsigned long tabViewBorderType) {
	NSTabView* tabView = (NSTabView*)ptr;
	[tabView setTabViewBorderType:tabViewBorderType];
}

bool TabView_AllowsTruncatedLabels(void* ptr) {
	NSTabView* tabView = (NSTabView*)ptr;
	return [tabView allowsTruncatedLabels];
}

void TabView_SetAllowsTruncatedLabels(void* ptr, bool allowsTruncatedLabels) {
	NSTabView* tabView = (NSTabView*)ptr;
	[tabView setAllowsTruncatedLabels:allowsTruncatedLabels];
}

void* TabView_SelectedTabViewItem(void* ptr) {
	NSTabView* tabView = (NSTabView*)ptr;
	return [tabView selectedTabViewItem];
}

void* TabView_Font(void* ptr) {
	NSTabView* tabView = (NSTabView*)ptr;
	return [tabView font];
}

void TabView_SetFont(void* ptr, void* font) {
	NSTabView* tabView = (NSTabView*)ptr;
	[tabView setFont:font];
}

NSSize TabView_MinimumSize(void* ptr) {
	NSTabView* tabView = (NSTabView*)ptr;
	return [tabView minimumSize];
}

unsigned long TabView_ControlSize(void* ptr) {
	NSTabView* tabView = (NSTabView*)ptr;
	return [tabView controlSize];
}

void TabView_SetControlSize(void* ptr, unsigned long controlSize) {
	NSTabView* tabView = (NSTabView*)ptr;
	[tabView setControlSize:controlSize];
}

void TabView_AddTabViewItem(void* ptr, void* tabViewItem) {
	NSTabView* tabView = (NSTabView*)ptr;
	[tabView addTabViewItem:tabViewItem];
}

void TabView_InsertTabViewItem(void* ptr, void* tabViewItem, long index) {
	NSTabView* tabView = (NSTabView*)ptr;
	[tabView insertTabViewItem:tabViewItem atIndex:index];
}

void TabView_RemoveTabViewItem(void* ptr, void* tabViewItem) {
	NSTabView* tabView = (NSTabView*)ptr;
	[tabView removeTabViewItem:tabViewItem];
}

long TabView_IndexOfTabViewItem(void* ptr, void* tabViewItem) {
	NSTabView* tabView = (NSTabView*)ptr;
	return [tabView indexOfTabViewItem:tabViewItem];
}

void* TabView_TabViewItemAtIndex(void* ptr, long index) {
	NSTabView* tabView = (NSTabView*)ptr;
	return [tabView tabViewItemAtIndex:index];
}

void TabView_SelectFirstTabViewItem(void* ptr, void* sender) {
	NSTabView* tabView = (NSTabView*)ptr;
	[tabView selectFirstTabViewItem:sender];
}

void TabView_SelectLastTabViewItem(void* ptr, void* sender) {
	NSTabView* tabView = (NSTabView*)ptr;
	[tabView selectLastTabViewItem:sender];
}

void TabView_SelectNextTabViewItem(void* ptr, void* sender) {
	NSTabView* tabView = (NSTabView*)ptr;
	[tabView selectNextTabViewItem:sender];
}

void TabView_SelectPreviousTabViewItem(void* ptr, void* sender) {
	NSTabView* tabView = (NSTabView*)ptr;
	[tabView selectPreviousTabViewItem:sender];
}

void TabView_SelectTabViewItem(void* ptr, void* tabViewItem) {
	NSTabView* tabView = (NSTabView*)ptr;
	[tabView selectTabViewItem:tabViewItem];
}

void TabView_SelectTabViewItemAtIndex(void* ptr, long index) {
	NSTabView* tabView = (NSTabView*)ptr;
	[tabView selectTabViewItemAtIndex:index];
}
