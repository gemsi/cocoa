#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>

bool MenuItem_IsEnabled(void* ptr);
void MenuItem_SetEnabled(void* ptr, bool enabled);
bool MenuItem_IsHidden(void* ptr);
void MenuItem_SetHidden(void* ptr, bool hidden);
const char* MenuItem_Title(void* ptr);
void MenuItem_SetTitle(void* ptr, const char* title);
void* MenuItem_Submenu(void* ptr);
void MenuItem_SetSubmenu(void* ptr, void* submenu);
bool MenuItem_HasSubmenu(void* ptr);
bool MenuItem_IsSeparatorItem(void* ptr);
void* MenuItem_Menu(void* ptr);
void MenuItem_SetMenu(void* ptr, void* menu);
const char* MenuItem_ToolTip(void* ptr);
void MenuItem_SetToolTip(void* ptr, const char* toolTip);
bool MenuItem_IsHighlighted(void* ptr);

