#import <Foundation/NSGeometry.h>
#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>

bool PopUpButton_PullsDown(void* ptr);
void PopUpButton_SetPullsDown(void* ptr, bool pullsDown);
bool PopUpButton_AutoenablesItems(void* ptr);
void PopUpButton_SetAutoenablesItems(void* ptr, bool autoenablesItems);
void* PopUpButton_SelectedItem(void* ptr);
const char* PopUpButton_TitleOfSelectedItem(void* ptr);
long PopUpButton_SelectedTag(void* ptr);
void* PopUpButton_Menu(void* ptr);
void PopUpButton_SetMenu(void* ptr, void* menu);
long PopUpButton_NumberOfItems(void* ptr);
Array PopUpButton_ItemArray(void* ptr);
Array PopUpButton_ItemTitles(void* ptr);
void* PopUpButton_LastItem(void* ptr);
long PopUpButton_PreferredEdge(void* ptr);
void PopUpButton_SetPreferredEdge(void* ptr, long preferredEdge);

void* PopUpButton_NewPopUpButton(NSRect buttonFrame, bool flag);
void PopUpButton_AddItemWithTitle(void* ptr, const char* title);
void PopUpButton_AddItemsWithTitles(void* ptr, Array itemTitles);
void PopUpButton_InsertItemWithTitle(void* ptr, const char* title, long index);
void PopUpButton_RemoveAllItems(void* ptr);
void PopUpButton_RemoveItemWithTitle(void* ptr, const char* title);
void PopUpButton_RemoveItemAtIndex(void* ptr, long index);
void PopUpButton_SelectItem(void* ptr, void* item);
void PopUpButton_SelectItemAtIndex(void* ptr, long index);
bool PopUpButton_SelectItemWithTag(void* ptr, long tag);
void PopUpButton_SelectItemWithTitle(void* ptr, const char* title);
void* PopUpButton_ItemAtIndex(void* ptr, long index);
const char* PopUpButton_ItemTitleAtIndex(void* ptr, long index);
void* PopUpButton_ItemWithTitle(void* ptr, const char* title);
long PopUpButton_IndexOfItem(void* ptr, void* item);
long PopUpButton_IndexOfItemWithTag(void* ptr, long tag);
long PopUpButton_IndexOfItemWithTitle(void* ptr, const char* title);
void PopUpButton_SynchronizeTitleAndSelectedItem(void* ptr);
