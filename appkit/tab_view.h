#import <Foundation/NSGeometry.h>
#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>

void* TabView_initWithFrame(long goID, NSRect frame);

long TabView_NumberOfTabViewItems(void* ptr);
unsigned long TabView_TabViewType(void* ptr);
void TabView_SetTabViewType(void* ptr, unsigned long tabViewType);
unsigned long TabView_TabPosition(void* ptr);
void TabView_SetTabPosition(void* ptr, unsigned long tabPosition);
unsigned long TabView_TabViewBorderType(void* ptr);
void TabView_SetTabViewBorderType(void* ptr, unsigned long tabViewBorderType);
bool TabView_AllowsTruncatedLabels(void* ptr);
void TabView_SetAllowsTruncatedLabels(void* ptr, bool allowsTruncatedLabels);
void* TabView_SelectedTabViewItem(void* ptr);
void* TabView_Font(void* ptr);
void TabView_SetFont(void* ptr, void* font);
NSSize TabView_MinimumSize(void* ptr);
unsigned long TabView_ControlSize(void* ptr);
void TabView_SetControlSize(void* ptr, unsigned long controlSize);

void TabView_AddTabViewItem(void* ptr, void* tabViewItem);
void TabView_InsertTabViewItem(void* ptr, void* tabViewItem, long index);
void TabView_RemoveTabViewItem(void* ptr, void* tabViewItem);
long TabView_IndexOfTabViewItem(void* ptr, void* tabViewItem);
void* TabView_TabViewItemAtIndex(void* ptr, long index);
void TabView_SelectFirstTabViewItem(void* ptr, void* sender);
void TabView_SelectLastTabViewItem(void* ptr, void* sender);
void TabView_SelectNextTabViewItem(void* ptr, void* sender);
void TabView_SelectPreviousTabViewItem(void* ptr, void* sender);
void TabView_SelectTabViewItem(void* ptr, void* tabViewItem);
void TabView_SelectTabViewItemAtIndex(void* ptr, long index);
