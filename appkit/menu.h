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

void* Menu_InitWithTitle(const char* title);
