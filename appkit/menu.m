#import <Appkit/NSMenu.h>
#import "menu.h"

void* Menu_initWithTitle(long goID, const char* title) {
	NSMenu* menu = [[[NSMenu alloc] initWithTitle:[NSString stringWithUTF8String:title]] autorelease];
	return (void*)menu;
}

double Menu_MenuBarHeight(void* ptr) {
	NSMenu* menu = (NSMenu*)ptr;
	return [menu menuBarHeight];
}

void* Menu_Font(void* ptr) {
	NSMenu* menu = (NSMenu*)ptr;
	return [menu font];
}

void Menu_SetFont(void* ptr, void* font) {
	NSMenu* menu = (NSMenu*)ptr;
	[menu setFont:font];
}

bool Menu_AutoenablesItems(void* ptr) {
	NSMenu* menu = (NSMenu*)ptr;
	return [menu autoenablesItems];
}

void Menu_SetAutoenablesItems(void* ptr, bool autoenablesItems) {
	NSMenu* menu = (NSMenu*)ptr;
	[menu setAutoenablesItems:autoenablesItems];
}

const char* Menu_Title(void* ptr) {
	NSMenu* menu = (NSMenu*)ptr;
	return [[menu title] UTF8String];
}

void Menu_SetTitle(void* ptr, const char* title) {
	NSMenu* menu = (NSMenu*)ptr;
	[menu setTitle:[NSString stringWithUTF8String:title]];
}

double Menu_MinimumWidth(void* ptr) {
	NSMenu* menu = (NSMenu*)ptr;
	return [menu minimumWidth];
}

void Menu_SetMinimumWidth(void* ptr, double minimumWidth) {
	NSMenu* menu = (NSMenu*)ptr;
	[menu setMinimumWidth:minimumWidth];
}

NSSize Menu_Size(void* ptr) {
	NSMenu* menu = (NSMenu*)ptr;
	return [menu size];
}

unsigned long Menu_PropertiesToUpdate(void* ptr) {
	NSMenu* menu = (NSMenu*)ptr;
	return [menu propertiesToUpdate];
}

bool Menu_AllowsContextMenuPlugIns(void* ptr) {
	NSMenu* menu = (NSMenu*)ptr;
	return [menu allowsContextMenuPlugIns];
}

void Menu_SetAllowsContextMenuPlugIns(void* ptr, bool allowsContextMenuPlugIns) {
	NSMenu* menu = (NSMenu*)ptr;
	[menu setAllowsContextMenuPlugIns:allowsContextMenuPlugIns];
}

bool Menu_ShowsStateColumn(void* ptr) {
	NSMenu* menu = (NSMenu*)ptr;
	return [menu showsStateColumn];
}

void Menu_SetShowsStateColumn(void* ptr, bool showsStateColumn) {
	NSMenu* menu = (NSMenu*)ptr;
	[menu setShowsStateColumn:showsStateColumn];
}

void* Menu_HighlightedItem(void* ptr) {
	NSMenu* menu = (NSMenu*)ptr;
	return [menu highlightedItem];
}

long Menu_UserInterfaceLayoutDirection(void* ptr) {
	NSMenu* menu = (NSMenu*)ptr;
	return [menu userInterfaceLayoutDirection];
}

void Menu_SetUserInterfaceLayoutDirection(void* ptr, long userInterfaceLayoutDirection) {
	NSMenu* menu = (NSMenu*)ptr;
	[menu setUserInterfaceLayoutDirection:userInterfaceLayoutDirection];
}
