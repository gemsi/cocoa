#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>

void MenuItem_RegisterDelegate(void *ptr, long goID);
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
const char* MenuItem_KeyEquivalent(void* ptr);
void MenuItem_SetKeyEquivalent(void* ptr, const char* keyEquivalent);
unsigned long MenuItem_KeyEquivalentModifierMask(void* ptr);
void MenuItem_SetKeyEquivalentModifierMask(void* ptr, unsigned long keyEquivalentModifierMask);
const char* MenuItem_UserKeyEquivalent(void* ptr);
bool MenuItem_IsAlternate(void* ptr);
void MenuItem_SetAlternate(void* ptr, bool alternate);
long MenuItem_IndentationLevel(void* ptr);
void MenuItem_SetIndentationLevel(void* ptr, long indentationLevel);
void* MenuItem_View(void* ptr);
void MenuItem_SetView(void* ptr, void* view);
bool MenuItem_AllowsKeyEquivalentWhenHidden(void* ptr);
void MenuItem_SetAllowsKeyEquivalentWhenHidden(void* ptr, bool allowsKeyEquivalentWhenHidden);
bool MenuItem_UsesUserKeyEquivalents();
void MenuItem_SetUsesUserKeyEquivalents(bool usesUserKeyEquivalents);
long MenuItem_State(void* ptr);
void MenuItem_SetState(void* ptr, long state);
long MenuItem_Tag(void* ptr);
void MenuItem_SetTag(void* ptr, long tag);

void* MenuItem_NewSeparatorItem();

void* MenuItem_NewMenuItem(const char* title, const char* charCode, long goID);
