#import <Appkit/NSMenuItem.h>
#import "menu_item.h"

bool MenuItem_IsEnabled(void* ptr) {
	NSMenuItem* menuItem = (NSMenuItem*)ptr;
	return [menuItem isEnabled];
}

void MenuItem_SetEnabled(void* ptr, bool enabled) {
	NSMenuItem* menuItem = (NSMenuItem*)ptr;
	[menuItem setEnabled:enabled];
}

bool MenuItem_IsHidden(void* ptr) {
	NSMenuItem* menuItem = (NSMenuItem*)ptr;
	return [menuItem isHidden];
}

void MenuItem_SetHidden(void* ptr, bool hidden) {
	NSMenuItem* menuItem = (NSMenuItem*)ptr;
	[menuItem setHidden:hidden];
}

const char* MenuItem_Title(void* ptr) {
	NSMenuItem* menuItem = (NSMenuItem*)ptr;
	return [[menuItem title] UTF8String];
}

void MenuItem_SetTitle(void* ptr, const char* title) {
	NSMenuItem* menuItem = (NSMenuItem*)ptr;
	[menuItem setTitle:[NSString stringWithUTF8String:title]];
}

void* MenuItem_Submenu(void* ptr) {
	NSMenuItem* menuItem = (NSMenuItem*)ptr;
	return [menuItem submenu];
}

void MenuItem_SetSubmenu(void* ptr, void* submenu) {
	NSMenuItem* menuItem = (NSMenuItem*)ptr;
	[menuItem setSubmenu:submenu];
}

bool MenuItem_HasSubmenu(void* ptr) {
	NSMenuItem* menuItem = (NSMenuItem*)ptr;
	return [menuItem hasSubmenu];
}

bool MenuItem_IsSeparatorItem(void* ptr) {
	NSMenuItem* menuItem = (NSMenuItem*)ptr;
	return [menuItem isSeparatorItem];
}

void* MenuItem_Menu(void* ptr) {
	NSMenuItem* menuItem = (NSMenuItem*)ptr;
	return [menuItem menu];
}

void MenuItem_SetMenu(void* ptr, void* menu) {
	NSMenuItem* menuItem = (NSMenuItem*)ptr;
	[menuItem setMenu:menu];
}

const char* MenuItem_ToolTip(void* ptr) {
	NSMenuItem* menuItem = (NSMenuItem*)ptr;
	return [[menuItem toolTip] UTF8String];
}

void MenuItem_SetToolTip(void* ptr, const char* toolTip) {
	NSMenuItem* menuItem = (NSMenuItem*)ptr;
	[menuItem setToolTip:[NSString stringWithUTF8String:toolTip]];
}

bool MenuItem_IsHighlighted(void* ptr) {
	NSMenuItem* menuItem = (NSMenuItem*)ptr;
	return [menuItem isHighlighted];
}
