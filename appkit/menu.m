#import <AppKit/AppKit.h>
#import "menu.h"

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
	[menu setFont:(NSFont*)font];
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

long Menu_NumberOfItems(void* ptr) {
	NSMenu* menu = (NSMenu*)ptr;
	return [menu numberOfItems];
}

Array Menu_ItemArray(void* ptr) {
	NSMenu* menu = (NSMenu*)ptr;
	NSArray* array = [menu itemArray];
	int count = [array count];
	void** data = malloc(count * sizeof(void*));
	for (int i = 0; i < count; i++) {
		 data[i] = [array objectAtIndex:i];
	}
	Array result;
	result.data = data;
	result.len = count;
	return result;
}

void Menu_SetItemArray(void* ptr, Array itemArray) {
	NSMenu* menu = (NSMenu*)ptr;
    NSMutableArray* objcItemArray = [[NSMutableArray alloc] init];
    void** itemArrayData = (void**)itemArray.data;
    for (int i = 0; i < itemArray.len; i++) {
    	[objcItemArray addObject:(NSMenuItem*)itemArrayData[i]];
    }
	[menu setItemArray:objcItemArray];
}

void* Menu_NewMenu(const char* title) {
	NSMenu* menu = [NSMenu alloc];
	return [[menu initWithTitle:[NSString stringWithUTF8String:title]] autorelease];
}

bool Menu_MenuBarVisible() {
	return [NSMenu menuBarVisible];
}

void Menu_SetMenuBarVisible(bool visible) {
	[NSMenu setMenuBarVisible:visible];
}

void Menu_InsertItem(void* ptr, void* newItem, long index) {
	NSMenu* menu = (NSMenu*)ptr;
	[menu insertItem:(NSMenuItem*)newItem atIndex:index];
}

void Menu_AddItem(void* ptr, void* newItem) {
	NSMenu* menu = (NSMenu*)ptr;
	[menu addItem:(NSMenuItem*)newItem];
}

void Menu_RemoveItem(void* ptr, void* newItem) {
	NSMenu* menu = (NSMenu*)ptr;
	[menu removeItem:(NSMenuItem*)newItem];
}

void Menu_RemoveItemAtIndex(void* ptr, long index) {
	NSMenu* menu = (NSMenu*)ptr;
	[menu removeItemAtIndex:index];
}

void Menu_RemoveAllItems(void* ptr) {
	NSMenu* menu = (NSMenu*)ptr;
	[menu removeAllItems];
}

void* Menu_ItemAtIndex(void* ptr, long index) {
	NSMenu* menu = (NSMenu*)ptr;
	return [menu itemAtIndex:index];
}

void* Menu_ItemWithTitle(void* ptr, const char* title) {
	NSMenu* menu = (NSMenu*)ptr;
	return [menu itemWithTitle:[NSString stringWithUTF8String:title]];
}

void* Menu_ItemWithTag(void* ptr, long tag) {
	NSMenu* menu = (NSMenu*)ptr;
	return [menu itemWithTag:tag];
}

long Menu_IndexOfItem(void* ptr, void* item) {
	NSMenu* menu = (NSMenu*)ptr;
	return [menu indexOfItem:(NSMenuItem*)item];
}

long Menu_IndexOfItemWithTitle(void* ptr, const char* title) {
	NSMenu* menu = (NSMenu*)ptr;
	return [menu indexOfItemWithTitle:[NSString stringWithUTF8String:title]];
}

long Menu_IndexOfItemWithTag(void* ptr, long tag) {
	NSMenu* menu = (NSMenu*)ptr;
	return [menu indexOfItemWithTag:tag];
}

long Menu_IndexOfItemWithSubmenu(void* ptr, void* subMenu) {
	NSMenu* menu = (NSMenu*)ptr;
	return [menu indexOfItemWithSubmenu:(NSMenu*)subMenu];
}

void Menu_SetSubmenu(void* ptr, void* subMenu, void* item) {
	NSMenu* menu = (NSMenu*)ptr;
	[menu setSubmenu:(NSMenu*)subMenu forItem:(NSMenuItem*)item];
}

void Menu_Update(void* ptr) {
	NSMenu* menu = (NSMenu*)ptr;
	[menu update];
}

void Menu_CancelTracking(void* ptr) {
	NSMenu* menu = (NSMenu*)ptr;
	[menu cancelTracking];
}

void Menu_CancelTrackingWithoutAnimation(void* ptr) {
	NSMenu* menu = (NSMenu*)ptr;
	[menu cancelTrackingWithoutAnimation];
}
