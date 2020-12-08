#import <AppKit/AppKit.h>
#import "pop_up_button.h"

bool PopUpButton_PullsDown(void* ptr) {
	NSPopUpButton* popUpButton = (NSPopUpButton*)ptr;
	return [popUpButton pullsDown];
}

void PopUpButton_SetPullsDown(void* ptr, bool pullsDown) {
	NSPopUpButton* popUpButton = (NSPopUpButton*)ptr;
	[popUpButton setPullsDown:pullsDown];
}

bool PopUpButton_AutoenablesItems(void* ptr) {
	NSPopUpButton* popUpButton = (NSPopUpButton*)ptr;
	return [popUpButton autoenablesItems];
}

void PopUpButton_SetAutoenablesItems(void* ptr, bool autoenablesItems) {
	NSPopUpButton* popUpButton = (NSPopUpButton*)ptr;
	[popUpButton setAutoenablesItems:autoenablesItems];
}

void* PopUpButton_SelectedItem(void* ptr) {
	NSPopUpButton* popUpButton = (NSPopUpButton*)ptr;
	return [popUpButton selectedItem];
}

const char* PopUpButton_TitleOfSelectedItem(void* ptr) {
	NSPopUpButton* popUpButton = (NSPopUpButton*)ptr;
	return [[popUpButton titleOfSelectedItem] UTF8String];
}

long PopUpButton_SelectedTag(void* ptr) {
	NSPopUpButton* popUpButton = (NSPopUpButton*)ptr;
	return [popUpButton selectedTag];
}

void* PopUpButton_Menu(void* ptr) {
	NSPopUpButton* popUpButton = (NSPopUpButton*)ptr;
	return [popUpButton menu];
}

void PopUpButton_SetMenu(void* ptr, void* menu) {
	NSPopUpButton* popUpButton = (NSPopUpButton*)ptr;
	[popUpButton setMenu:(NSMenu*)menu];
}

long PopUpButton_NumberOfItems(void* ptr) {
	NSPopUpButton* popUpButton = (NSPopUpButton*)ptr;
	return [popUpButton numberOfItems];
}

Array PopUpButton_ItemArray(void* ptr) {
	NSPopUpButton* popUpButton = (NSPopUpButton*)ptr;
	NSArray* array = [popUpButton itemArray];
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

Array PopUpButton_ItemTitles(void* ptr) {
	NSPopUpButton* popUpButton = (NSPopUpButton*)ptr;
	NSArray* array = [popUpButton itemTitles];
	int count = [array count];
	void** data = malloc(count * sizeof(void*));
	for (int i = 0; i < count; i++) {
		 data[i] = (void*)[[array objectAtIndex:i] UTF8String];
	}
	Array result;
	result.data = data;
	result.len = count;
	return result;
}

void* PopUpButton_LastItem(void* ptr) {
	NSPopUpButton* popUpButton = (NSPopUpButton*)ptr;
	return [popUpButton lastItem];
}

long PopUpButton_PreferredEdge(void* ptr) {
	NSPopUpButton* popUpButton = (NSPopUpButton*)ptr;
	return [popUpButton preferredEdge];
}

void PopUpButton_SetPreferredEdge(void* ptr, long preferredEdge) {
	NSPopUpButton* popUpButton = (NSPopUpButton*)ptr;
	[popUpButton setPreferredEdge:preferredEdge];
}

void* PopUpButton_NewPopUpButton(NSRect buttonFrame, bool flag) {
	NSPopUpButton* popUpButton = [NSPopUpButton alloc];
	return [[popUpButton initWithFrame:buttonFrame pullsDown:flag] autorelease];
}

void PopUpButton_AddItemWithTitle(void* ptr, const char* title) {
	NSPopUpButton* popUpButton = (NSPopUpButton*)ptr;
	[popUpButton addItemWithTitle:[NSString stringWithUTF8String:title]];
}

void PopUpButton_AddItemsWithTitles(void* ptr, Array itemTitles) {
	NSPopUpButton* popUpButton = (NSPopUpButton*)ptr;
    NSMutableArray* objcItemTitles = [[NSMutableArray alloc] init];
    void** itemTitlesData = (void**)itemTitles.data;
    for (int i = 0; i < itemTitles.len; i++) {
    	[objcItemTitles addObject:[NSString stringWithUTF8String:(const char*)itemTitlesData[i]]];
    }
	[popUpButton addItemsWithTitles:objcItemTitles];
}

void PopUpButton_InsertItemWithTitle(void* ptr, const char* title, long index) {
	NSPopUpButton* popUpButton = (NSPopUpButton*)ptr;
	[popUpButton insertItemWithTitle:[NSString stringWithUTF8String:title] atIndex:index];
}

void PopUpButton_RemoveAllItems(void* ptr) {
	NSPopUpButton* popUpButton = (NSPopUpButton*)ptr;
	[popUpButton removeAllItems];
}

void PopUpButton_RemoveItemWithTitle(void* ptr, const char* title) {
	NSPopUpButton* popUpButton = (NSPopUpButton*)ptr;
	[popUpButton removeItemWithTitle:[NSString stringWithUTF8String:title]];
}

void PopUpButton_RemoveItemAtIndex(void* ptr, long index) {
	NSPopUpButton* popUpButton = (NSPopUpButton*)ptr;
	[popUpButton removeItemAtIndex:index];
}

void PopUpButton_SelectItem(void* ptr, void* item) {
	NSPopUpButton* popUpButton = (NSPopUpButton*)ptr;
	[popUpButton selectItem:(NSMenuItem*)item];
}

void PopUpButton_SelectItemAtIndex(void* ptr, long index) {
	NSPopUpButton* popUpButton = (NSPopUpButton*)ptr;
	[popUpButton selectItemAtIndex:index];
}

bool PopUpButton_SelectItemWithTag(void* ptr, long tag) {
	NSPopUpButton* popUpButton = (NSPopUpButton*)ptr;
	return [popUpButton selectItemWithTag:tag];
}

void PopUpButton_SelectItemWithTitle(void* ptr, const char* title) {
	NSPopUpButton* popUpButton = (NSPopUpButton*)ptr;
	[popUpButton selectItemWithTitle:[NSString stringWithUTF8String:title]];
}

void* PopUpButton_ItemAtIndex(void* ptr, long index) {
	NSPopUpButton* popUpButton = (NSPopUpButton*)ptr;
	return [popUpButton itemAtIndex:index];
}

const char* PopUpButton_ItemTitleAtIndex(void* ptr, long index) {
	NSPopUpButton* popUpButton = (NSPopUpButton*)ptr;
	return [[popUpButton itemTitleAtIndex:index] UTF8String];
}

void* PopUpButton_ItemWithTitle(void* ptr, const char* title) {
	NSPopUpButton* popUpButton = (NSPopUpButton*)ptr;
	return [popUpButton itemWithTitle:[NSString stringWithUTF8String:title]];
}

long PopUpButton_IndexOfItem(void* ptr, void* item) {
	NSPopUpButton* popUpButton = (NSPopUpButton*)ptr;
	return [popUpButton indexOfItem:(NSMenuItem*)item];
}

long PopUpButton_IndexOfItemWithTag(void* ptr, long tag) {
	NSPopUpButton* popUpButton = (NSPopUpButton*)ptr;
	return [popUpButton indexOfItemWithTag:tag];
}

long PopUpButton_IndexOfItemWithTitle(void* ptr, const char* title) {
	NSPopUpButton* popUpButton = (NSPopUpButton*)ptr;
	return [popUpButton indexOfItemWithTitle:[NSString stringWithUTF8String:title]];
}

void PopUpButton_SynchronizeTitleAndSelectedItem(void* ptr) {
	NSPopUpButton* popUpButton = (NSPopUpButton*)ptr;
	[popUpButton synchronizeTitleAndSelectedItem];
}
