#import <Foundation/NSGeometry.h>
#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>

double Menu_MenuBarHeight(void* ptr);
void* Menu_Font(void* ptr);
void Menu_SetFont(void* ptr, void* font);
bool Menu_AutoenablesItems(void* ptr);
void Menu_SetAutoenablesItems(void* ptr, bool autoenablesItems);
const char* Menu_Title(void* ptr);
void Menu_SetTitle(void* ptr, const char* title);
double Menu_MinimumWidth(void* ptr);
void Menu_SetMinimumWidth(void* ptr, double minimumWidth);
NSSize Menu_Size(void* ptr);
unsigned long Menu_PropertiesToUpdate(void* ptr);
bool Menu_AllowsContextMenuPlugIns(void* ptr);
void Menu_SetAllowsContextMenuPlugIns(void* ptr, bool allowsContextMenuPlugIns);
bool Menu_ShowsStateColumn(void* ptr);
void Menu_SetShowsStateColumn(void* ptr, bool showsStateColumn);
void* Menu_HighlightedItem(void* ptr);
long Menu_UserInterfaceLayoutDirection(void* ptr);
void Menu_SetUserInterfaceLayoutDirection(void* ptr, long userInterfaceLayoutDirection);
long Menu_NumberOfItems(void* ptr);
Array Menu_ItemArray(void* ptr);
void Menu_SetItemArray(void* ptr, Array itemArray);

void* Menu_NewMenu(const char* title);
bool Menu_MenuBarVisible();
void Menu_SetMenuBarVisible(bool visible);
void Menu_InsertItem(void* ptr, void* newItem, long index);
void Menu_AddItem(void* ptr, void* newItem);
void Menu_RemoveItem(void* ptr, void* newItem);
void Menu_RemoveItemAtIndex(void* ptr, long index);
void Menu_RemoveAllItems(void* ptr);
void* Menu_ItemAtIndex(void* ptr, long index);
void* Menu_ItemWithTitle(void* ptr, const char* title);
void* Menu_ItemWithTag(void* ptr, long tag);
long Menu_IndexOfItem(void* ptr, void* item);
long Menu_IndexOfItemWithTitle(void* ptr, const char* title);
long Menu_IndexOfItemWithTag(void* ptr, long tag);
long Menu_IndexOfItemWithSubmenu(void* ptr, void* subMenu);
void Menu_SetSubmenu(void* ptr, void* subMenu, void* item);
void Menu_Update(void* ptr);
void Menu_CancelTracking(void* ptr);
void Menu_CancelTrackingWithoutAnimation(void* ptr);
