#import <Appkit/NSTabViewItem.h>
#import "tab_view_item.h"

const char* TabViewItem_Label(void* ptr) {
	NSTabViewItem* tabViewItem = (NSTabViewItem*)ptr;
	return [[tabViewItem label] UTF8String];
}

void TabViewItem_SetLabel(void* ptr, const char* label) {
	NSTabViewItem* tabViewItem = (NSTabViewItem*)ptr;
	[tabViewItem setLabel:[NSString stringWithUTF8String:label]];
}

const char* TabViewItem_ToolTip(void* ptr) {
	NSTabViewItem* tabViewItem = (NSTabViewItem*)ptr;
	return [[tabViewItem toolTip] UTF8String];
}

void TabViewItem_SetToolTip(void* ptr, const char* toolTip) {
	NSTabViewItem* tabViewItem = (NSTabViewItem*)ptr;
	[tabViewItem setToolTip:[NSString stringWithUTF8String:toolTip]];
}

unsigned long TabViewItem_TabState(void* ptr) {
	NSTabViewItem* tabViewItem = (NSTabViewItem*)ptr;
	return [tabViewItem tabState];
}

void* TabViewItem_Identifier(void* ptr) {
	NSTabViewItem* tabViewItem = (NSTabViewItem*)ptr;
	return [tabViewItem identifier];
}

void TabViewItem_SetIdentifier(void* ptr, void* identifier) {
	NSTabViewItem* tabViewItem = (NSTabViewItem*)ptr;
	[tabViewItem setIdentifier:identifier];
}

void* TabViewItem_Color(void* ptr) {
	NSTabViewItem* tabViewItem = (NSTabViewItem*)ptr;
	return [tabViewItem color];
}

void TabViewItem_SetColor(void* ptr, void* color) {
	NSTabViewItem* tabViewItem = (NSTabViewItem*)ptr;
	[tabViewItem setColor:color];
}

void* TabViewItem_View(void* ptr) {
	NSTabViewItem* tabViewItem = (NSTabViewItem*)ptr;
	return [tabViewItem view];
}

void TabViewItem_SetView(void* ptr, void* view) {
	NSTabViewItem* tabViewItem = (NSTabViewItem*)ptr;
	[tabViewItem setView:view];
}

void* TabViewItem_InitialFirstResponder(void* ptr) {
	NSTabViewItem* tabViewItem = (NSTabViewItem*)ptr;
	return [tabViewItem initialFirstResponder];
}

void TabViewItem_SetInitialFirstResponder(void* ptr, void* initialFirstResponder) {
	NSTabViewItem* tabViewItem = (NSTabViewItem*)ptr;
	[tabViewItem setInitialFirstResponder:initialFirstResponder];
}

void* TabViewItem_TabView(void* ptr) {
	NSTabViewItem* tabViewItem = (NSTabViewItem*)ptr;
	return [tabViewItem tabView];
}

void* TabViewItem_InitWithIdentifier(void* identifier) {
	NSTabViewItem* tabViewItem = [NSTabViewItem alloc];
	return [[tabViewItem initWithIdentifier:identifier] autorelease];
}
