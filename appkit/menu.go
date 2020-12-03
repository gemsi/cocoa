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
	// NumberOfItems return the number of menu items in the menu, including separator items
	NumberOfItems() int
	// ItemArray return an array containing the menu items in the menu
	ItemArray() []MenuItem
	// SetItemArray set an array containing the menu items in the menu
	SetItemArray(itemArray []MenuItem)
	// InsertItem inserts a menu item into the menu at a specific location
	InsertItem(newItem MenuItem, index int)
	// AddItem adds a menu item to the end of the menu
	AddItem(newItem MenuItem)
	// RemoveItem removes a menu item from the menu
	RemoveItem(newItem MenuItem)
	// RemoveItemAtIndex removes the menu item at a specified location in the menu
	RemoveItemAtIndex(index int)
	// RemoveAllItems removes all the menu items in the menu
	RemoveAllItems()
	// ItemAtIndex returns the menu item at a specific location of the menu
	ItemAtIndex(index int) MenuItem
	// ItemWithTitle returns the first menu item in the menu with a specified title
	ItemWithTitle(title string) MenuItem
	// ItemWithTag returns the first menu item in the menu with the specified tag
	ItemWithTag(tag int) MenuItem
	// IndexOfItem returns the index identifying the location of a specified menu item in the menu
	IndexOfItem(item MenuItem) int
	// IndexOfItemWithTitle returns the index of the first menu item in the menu that has a specified title
	IndexOfItemWithTitle(title string) int
	// IndexOfItemWithTag returns the index of the first menu item in the menu identified by a tag
	IndexOfItemWithTag(tag int) int
	// IndexOfItemWithSubmenu returns the index of the menu item in the menu with the given submenu
	IndexOfItemWithSubmenu(subMenu Menu) int
	// SetSubmenu assigns a menu to be a submenu of the menu controlled by a given menu item
	SetSubmenu(subMenu Menu, item MenuItem)
	// Update enables or disables the menu items of the menu based on the NSMenuValidation informal protocol and sizes the menu to fit its current menu items if necessary
	Update()
	// CancelTracking dismisses the menu and ends all menu tracking
	CancelTracking()
	// CancelTrackingWithoutAnimation dismisses the menu and ends all menu tracking without displaying the associated animation
	CancelTrackingWithoutAnimation()
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

func (m *NSMenu) NumberOfItems() int {
	return int(C.Menu_NumberOfItems(m.Ptr()))
}

func (m *NSMenu) ItemArray() []MenuItem {
	var cArray C.Array = C.Menu_ItemArray(m.Ptr())
	defer C.free(cArray.data)
	result := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(cArray.data))[:cArray.len:cArray.len]
	var goResult = make([]MenuItem, len(result))
	for idx, r := range result {
		goResult[idx] = MakeMenuItem(r)
	}
	return goResult
}

func (m *NSMenu) SetItemArray(itemArray []MenuItem) {
	cItemArrayData := make([]unsafe.Pointer, len(itemArray))
	for idx, v := range itemArray {
		cItemArrayData[idx] = v.Ptr()
	}
	cItemArray := C.Array{data: unsafe.Pointer(&cItemArrayData[0]), len: C.int(len(itemArray))}
	C.Menu_SetItemArray(m.Ptr(), cItemArray)
}

func NewMenu(title string) Menu {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	return MakeMenu(C.Menu_NewMenu(cTitle))
}

func MenuBarVisible() bool {
	return bool(C.Menu_MenuBarVisible())
}

func SetMenuBarVisible(visible bool) {
	C.Menu_SetMenuBarVisible(C.bool(visible))
}

func (m *NSMenu) InsertItem(newItem MenuItem, index int) {
	C.Menu_InsertItem(m.Ptr(), toPointer(newItem), C.long(index))
}

func (m *NSMenu) AddItem(newItem MenuItem) {
	C.Menu_AddItem(m.Ptr(), toPointer(newItem))
}

func (m *NSMenu) RemoveItem(newItem MenuItem) {
	C.Menu_RemoveItem(m.Ptr(), toPointer(newItem))
}

func (m *NSMenu) RemoveItemAtIndex(index int) {
	C.Menu_RemoveItemAtIndex(m.Ptr(), C.long(index))
}

func (m *NSMenu) RemoveAllItems() {
	C.Menu_RemoveAllItems(m.Ptr())
}

func (m *NSMenu) ItemAtIndex(index int) MenuItem {
	return MakeMenuItem(C.Menu_ItemAtIndex(m.Ptr(), C.long(index)))
}

func (m *NSMenu) ItemWithTitle(title string) MenuItem {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	return MakeMenuItem(C.Menu_ItemWithTitle(m.Ptr(), cTitle))
}

func (m *NSMenu) ItemWithTag(tag int) MenuItem {
	return MakeMenuItem(C.Menu_ItemWithTag(m.Ptr(), C.long(tag)))
}

func (m *NSMenu) IndexOfItem(item MenuItem) int {
	return int(C.Menu_IndexOfItem(m.Ptr(), toPointer(item)))
}

func (m *NSMenu) IndexOfItemWithTitle(title string) int {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	return int(C.Menu_IndexOfItemWithTitle(m.Ptr(), cTitle))
}

func (m *NSMenu) IndexOfItemWithTag(tag int) int {
	return int(C.Menu_IndexOfItemWithTag(m.Ptr(), C.long(tag)))
}

func (m *NSMenu) IndexOfItemWithSubmenu(subMenu Menu) int {
	return int(C.Menu_IndexOfItemWithSubmenu(m.Ptr(), toPointer(subMenu)))
}

func (m *NSMenu) SetSubmenu(subMenu Menu, item MenuItem) {
	C.Menu_SetSubmenu(m.Ptr(), toPointer(subMenu), toPointer(item))
}

func (m *NSMenu) Update() {
	C.Menu_Update(m.Ptr())
}

func (m *NSMenu) CancelTracking() {
	C.Menu_CancelTracking(m.Ptr())
}

func (m *NSMenu) CancelTrackingWithoutAnimation() {
	C.Menu_CancelTrackingWithoutAnimation(m.Ptr())
}
