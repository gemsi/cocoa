#import <Appkit/NSTabView.h>
#import "tab_view.h"
#import "tab_view_delegate.h"

@implementation GoNSTabViewDelegate

- (void)tabViewDidChangeNumberOfTabViewItems:(NSTabView*)view {
	return TabView_Delegate_TabViewDidChangeNumberOfTabViewItems([self goID], view);
}

- (BOOL)tabView:(NSTabView*)view shouldSelectTabViewItem:(NSTabViewItem*)tabViewItem {
	return TabView_Delegate_ShouldSelectTabViewItem([self goID], view, tabViewItem);
}

- (void)tabView:(NSTabView*)view willSelectTabViewItem:(NSTabViewItem*)tabViewItem {
	return TabView_Delegate_WillSelectTabViewItem([self goID], view, tabViewItem);
}

- (void)tabView:(NSTabView*)view didSelectTabViewItem:(NSTabViewItem*)tabViewItem {
	return TabView_Delegate_DidSelectTabViewItem([self goID], view, tabViewItem);
}

@end
void TabView_RegisterDelegate(void *ptr, long goID) {
	NSTabView* tabView = (NSTabView*)ptr;
	GoNSTabViewDelegate* delegate = [[[GoNSTabViewDelegate alloc] init] autorelease];
	[delegate setGoID:goID];
	[tabView setDelegate:delegate];
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
	[tabView setFont:(NSFont*)font];
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

void* TabView_NewTabView(NSRect frame) {
	NSTabView* tabView = [NSTabView alloc];
	return [[tabView initWithFrame:frame] autorelease];
}

void TabView_AddTabViewItem(void* ptr, void* tabViewItem) {
	NSTabView* tabView = (NSTabView*)ptr;
	[tabView addTabViewItem:(NSTabViewItem*)tabViewItem];
}

void TabView_InsertTabViewItem(void* ptr, void* tabViewItem, long index) {
	NSTabView* tabView = (NSTabView*)ptr;
	[tabView insertTabViewItem:(NSTabViewItem*)tabViewItem atIndex:index];
}

void TabView_RemoveTabViewItem(void* ptr, void* tabViewItem) {
	NSTabView* tabView = (NSTabView*)ptr;
	[tabView removeTabViewItem:(NSTabViewItem*)tabViewItem];
}

long TabView_IndexOfTabViewItem(void* ptr, void* tabViewItem) {
	NSTabView* tabView = (NSTabView*)ptr;
	return [tabView indexOfTabViewItem:(NSTabViewItem*)tabViewItem];
}

void* TabView_TabViewItemAtIndex(void* ptr, long index) {
	NSTabView* tabView = (NSTabView*)ptr;
	return [tabView tabViewItemAtIndex:index];
}

void TabView_SelectFirstTabViewItem(void* ptr, void* sender) {
	NSTabView* tabView = (NSTabView*)ptr;
	[tabView selectFirstTabViewItem:(NSObject*)sender];
}

void TabView_SelectLastTabViewItem(void* ptr, void* sender) {
	NSTabView* tabView = (NSTabView*)ptr;
	[tabView selectLastTabViewItem:(NSObject*)sender];
}

void TabView_SelectNextTabViewItem(void* ptr, void* sender) {
	NSTabView* tabView = (NSTabView*)ptr;
	[tabView selectNextTabViewItem:(NSObject*)sender];
}

void TabView_SelectPreviousTabViewItem(void* ptr, void* sender) {
	NSTabView* tabView = (NSTabView*)ptr;
	[tabView selectPreviousTabViewItem:(NSObject*)sender];
}

void TabView_SelectTabViewItem(void* ptr, void* tabViewItem) {
	NSTabView* tabView = (NSTabView*)ptr;
	[tabView selectTabViewItem:(NSTabViewItem*)tabViewItem];
}

void TabView_SelectTabViewItemAtIndex(void* ptr, long index) {
	NSTabView* tabView = (NSTabView*)ptr;
	[tabView selectTabViewItemAtIndex:index];
}
