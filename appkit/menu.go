package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include "menu.h"
import "C"

import (
	"github.com/hsiafan/cocoa/foundation"
	"unsafe"
)

// Menu is an object that manages an app’s menus
type Menu interface {
	foundation.Object

	// MenuBarHeight return the menu bar height for the main menu in pixels
	MenuBarHeight() float64
	// Font return the font of the menu and its submenus
	Font() Font
	// SetFont set the font of the menu and its submenus
	SetFont(font Font)
	// AutoenablesItems return whether the menu automatically enables and disables its menu items
	AutoenablesItems() bool
	// SetAutoenablesItems set whether the menu automatically enables and disables its menu items
	SetAutoenablesItems(autoenablesItems bool)
	// Title return the title of the menu
	Title() string
	// SetTitle set the title of the menu
	SetTitle(title string)
	// MinimumWidth return the minimum width of the menu in screen coordinates
	MinimumWidth() float64
	// SetMinimumWidth set the minimum width of the menu in screen coordinates
	SetMinimumWidth(minimumWidth float64)
	// Size return the size of the menu in screen coordinates
	Size() foundation.Size
	// PropertiesToUpdate return the size of the menu in screen coordinates
	PropertiesToUpdate() MenuProperties
	// AllowsContextMenuPlugIns return whether the pop-up menu allows appending of contextual menu plug-in items
	AllowsContextMenuPlugIns() bool
	// SetAllowsContextMenuPlugIns set whether the pop-up menu allows appending of contextual menu plug-in items
	SetAllowsContextMenuPlugIns(allowsContextMenuPlugIns bool)
	// ShowsStateColumn return whether the menu displays the state column
	ShowsStateColumn() bool
	// SetShowsStateColumn set whether the menu displays the state column
	SetShowsStateColumn(showsStateColumn bool)
	// HighlightedItem return indicates the currently highlighted item in the menu
	HighlightedItem() MenuItem
	// UserInterfaceLayoutDirection return the layout direction of menu items in the menu
	UserInterfaceLayoutDirection() UserInterfaceLayoutDirection
	// SetUserInterfaceLayoutDirection set the layout direction of menu items in the menu
	SetUserInterfaceLayoutDirection(userInterfaceLayoutDirection UserInterfaceLayoutDirection)
}

var _ Menu = (*NSMenu)(nil)

type NSMenu struct {
	foundation.NSObject
}

// MakeMenu create a Menu from native pointer
func MakeMenu(ptr unsafe.Pointer) *NSMenu {
	if ptr == nil {
		return nil
	}
	return &NSMenu{
		NSObject: *foundation.MakeObject(ptr),
	}
}

func (m *NSMenu) MenuBarHeight() float64 {
	return float64(C.Menu_MenuBarHeight(m.Ptr()))
}

func (m *NSMenu) Font() Font {
	return MakeFont(C.Menu_Font(m.Ptr()))
}

func (m *NSMenu) SetFont(font Font) {
	C.Menu_SetFont(m.Ptr(), toPointer(font))
}

func (m *NSMenu) AutoenablesItems() bool {
	return bool(C.Menu_AutoenablesItems(m.Ptr()))
}

func (m *NSMenu) SetAutoenablesItems(autoenablesItems bool) {
	C.Menu_SetAutoenablesItems(m.Ptr(), C.bool(autoenablesItems))
}

func (m *NSMenu) Title() string {
	return C.GoString(C.Menu_Title(m.Ptr()))
}

func (m *NSMenu) SetTitle(title string) {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	C.Menu_SetTitle(m.Ptr(), cTitle)
}

func (m *NSMenu) MinimumWidth() float64 {
	return float64(C.Menu_MinimumWidth(m.Ptr()))
}

func (m *NSMenu) SetMinimumWidth(minimumWidth float64) {
	C.Menu_SetMinimumWidth(m.Ptr(), C.double(minimumWidth))
}

func (m *NSMenu) Size() foundation.Size {
	return toSize(C.Menu_Size(m.Ptr()))
}

func (m *NSMenu) PropertiesToUpdate() MenuProperties {
	return MenuProperties(C.Menu_PropertiesToUpdate(m.Ptr()))
}

func (m *NSMenu) AllowsContextMenuPlugIns() bool {
	return bool(C.Menu_AllowsContextMenuPlugIns(m.Ptr()))
}

func (m *NSMenu) SetAllowsContextMenuPlugIns(allowsContextMenuPlugIns bool) {
	C.Menu_SetAllowsContextMenuPlugIns(m.Ptr(), C.bool(allowsContextMenuPlugIns))
}

func (m *NSMenu) ShowsStateColumn() bool {
	return bool(C.Menu_ShowsStateColumn(m.Ptr()))
}

func (m *NSMenu) SetShowsStateColumn(showsStateColumn bool) {
	C.Menu_SetShowsStateColumn(m.Ptr(), C.bool(showsStateColumn))
}

func (m *NSMenu) HighlightedItem() MenuItem {
	return MakeMenuItem(C.Menu_HighlightedItem(m.Ptr()))
}

func (m *NSMenu) UserInterfaceLayoutDirection() UserInterfaceLayoutDirection {
	return UserInterfaceLayoutDirection(C.Menu_UserInterfaceLayoutDirection(m.Ptr()))
}

func (m *NSMenu) SetUserInterfaceLayoutDirection(userInterfaceLayoutDirection UserInterfaceLayoutDirection) {
	C.Menu_SetUserInterfaceLayoutDirection(m.Ptr(), C.long(userInterfaceLayoutDirection))
}

func NewMenu(title string) Menu {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	return MakeMenu(C.Menu_InitWithTitle(cTitle))
}
