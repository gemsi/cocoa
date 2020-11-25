package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include "menu_item.h"
import "C"

import (
	"github.com/hsiafan/cocoa/foundation"
	"unsafe"
)

// MenuItem is a command item in an app menu
type MenuItem interface {
	foundation.Object

	// IsEnabled return whether the menu item is enabled
	IsEnabled() bool
	// SetEnabled set whether the menu item is enabled
	SetEnabled(enabled bool)
	// IsHidden return whether the menu item is hidden
	IsHidden() bool
	// SetHidden set whether the menu item is hidden
	SetHidden(hidden bool)
	// Title return the menu item's title
	Title() string
	// SetTitle set the menu item's title
	SetTitle(title string)
	// Submenu return the submenu of the menu item
	Submenu() Mennu
	// SetSubmenu set the submenu of the menu item
	SetSubmenu(submenu Mennu)
	// HasSubmenu return whether the menu item has a submenu
	HasSubmenu() bool
	// IsSeparatorItem return a menu item that is used to separate logical groups of menu commands
	IsSeparatorItem() bool
	// Menu return the menu item’s menu
	Menu() Mennu
	// SetMenu set the menu item’s menu
	SetMenu(menu Mennu)
	// ToolTip return a help tag for the menu item
	ToolTip() string
	// SetToolTip set a help tag for the menu item
	SetToolTip(toolTip string)
	// IsHighlighted return whether the menu item should be drawn highlighted
	IsHighlighted() bool
}

var _ MenuItem = (*NSMenuItem)(nil)

type NSMenuItem struct {
	foundation.NSObject
}

// MakeMenuItem create a MenuItem from native pointer
func MakeMenuItem(ptr unsafe.Pointer) *NSMenuItem {
	if ptr == nil {
		return nil
	}
	return &NSMenuItem{
		NSObject: *foundation.MakeObject(ptr),
	}
}

func (m *NSMenuItem) IsEnabled() bool {
	return bool(C.MenuItem_IsEnabled(m.Ptr()))
}

func (m *NSMenuItem) SetEnabled(enabled bool) {
	C.MenuItem_SetEnabled(m.Ptr(), C.bool(enabled))
}

func (m *NSMenuItem) IsHidden() bool {
	return bool(C.MenuItem_IsHidden(m.Ptr()))
}

func (m *NSMenuItem) SetHidden(hidden bool) {
	C.MenuItem_SetHidden(m.Ptr(), C.bool(hidden))
}

func (m *NSMenuItem) Title() string {
	return C.GoString(C.MenuItem_Title(m.Ptr()))
}

func (m *NSMenuItem) SetTitle(title string) {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	C.MenuItem_SetTitle(m.Ptr(), cTitle)
}

func (m *NSMenuItem) Submenu() Mennu {
	return MakeMennu(C.MenuItem_Submenu(m.Ptr()))
}

func (m *NSMenuItem) SetSubmenu(submenu Mennu) {
	C.MenuItem_SetSubmenu(m.Ptr(), toPointer(submenu))
}

func (m *NSMenuItem) HasSubmenu() bool {
	return bool(C.MenuItem_HasSubmenu(m.Ptr()))
}

func (m *NSMenuItem) IsSeparatorItem() bool {
	return bool(C.MenuItem_IsSeparatorItem(m.Ptr()))
}

func (m *NSMenuItem) Menu() Mennu {
	return MakeMennu(C.MenuItem_Menu(m.Ptr()))
}

func (m *NSMenuItem) SetMenu(menu Mennu) {
	C.MenuItem_SetMenu(m.Ptr(), toPointer(menu))
}

func (m *NSMenuItem) ToolTip() string {
	return C.GoString(C.MenuItem_ToolTip(m.Ptr()))
}

func (m *NSMenuItem) SetToolTip(toolTip string) {
	cToolTip := C.CString(toolTip)
	defer C.free(unsafe.Pointer(cToolTip))
	C.MenuItem_SetToolTip(m.Ptr(), cToolTip)
}

func (m *NSMenuItem) IsHighlighted() bool {
	return bool(C.MenuItem_IsHighlighted(m.Ptr()))
}
