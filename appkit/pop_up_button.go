package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include "pop_up_button.h"
import "C"

import (
	"github.com/hsiafan/cocoa/foundation"
	"unsafe"
)

// PopUpButton is a control for selecting an item from a list
type PopUpButton interface {
	Button

	// PullsDown return whether the button displays a pull-down or pop-up menu
	PullsDown() bool
	// SetPullsDown set whether the button displays a pull-down or pop-up menu
	SetPullsDown(pullsDown bool)
	// AutoenablesItems return whether the button enables and disables its items every time a user event occurs
	AutoenablesItems() bool
	// SetAutoenablesItems set whether the button enables and disables its items every time a user event occurs
	SetAutoenablesItems(autoenablesItems bool)
	// SelectedItem return the menu item that was last selected by the user
	SelectedItem() MenuItem
	// TitleOfSelectedItem return the title of the item that was last selected by the user
	TitleOfSelectedItem() string
	// SelectedTag return the tag of the menu item that was last selected by the user
	SelectedTag() int
	// Menu return the menu associated with the pop-up button
	Menu() Menu
	// SetMenu set the menu associated with the pop-up button
	SetMenu(menu Menu)
	// NumberOfItems return the number of items in the menu
	NumberOfItems() int
	// ItemArray return the array of menu item objects associated with the button
	ItemArray() []MenuItem
	// ItemTitles return an array of strings corresponding to the titles of the items in the menu
	ItemTitles() []string
	// LastItem return the last item in the menu
	LastItem() MenuItem
	// PreferredEdge return the edge of the button on which to display the menu when screen space is constrained
	PreferredEdge() foundation.RectEdge
	// SetPreferredEdge set the edge of the button on which to display the menu when screen space is constrained
	SetPreferredEdge(preferredEdge foundation.RectEdge)
	// AddItemWithTitle adds an item with the specified title to the end of the menu
	AddItemWithTitle(title string)
	// AddItemsWithTitles adds multiple items to the end of the menu
	AddItemsWithTitles(itemTitles []string)
	// InsertItemWithTitle inserts an item at the specified position in the menu
	InsertItemWithTitle(title string, index int)
	// RemoveAllItems removes all items in the receiver’s item menu
	RemoveAllItems()
	// RemoveItemWithTitle removes the item with the specified title from the menu
	RemoveItemWithTitle(title string)
	// RemoveItemAtIndex removes the item at the specified index
	RemoveItemAtIndex(index int)
	// SelectItem selects the specified menu item
	SelectItem(item MenuItem)
	// SelectItemAtIndex selects the item in the menu at the specified index
	SelectItemAtIndex(index int)
	// SelectItemWithTag selects the menu item with the specified tag
	SelectItemWithTag(tag int) bool
	// SelectItemWithTitle selects the item with the specified title
	SelectItemWithTitle(title string)
	// ItemAtIndex returns the menu item at the specified index
	ItemAtIndex(index int) MenuItem
	// ItemTitleAtIndex returns the title of the item at the specified index
	ItemTitleAtIndex(index int) string
	// ItemWithTitle returns the menu item with the specified title
	ItemWithTitle(title string) MenuItem
	// IndexOfItem returns the index of the specified menu item
	IndexOfItem(item MenuItem) int
	// IndexOfItemWithTag returns the index of the menu item with the specified tag
	IndexOfItemWithTag(tag int) int
	// IndexOfItemWithTitle returns the index of the item with the specified title
	IndexOfItemWithTitle(title string) int
	// SynchronizeTitleAndSelectedItem ensures that the item being displayed by the receiver agrees with the selected item
	SynchronizeTitleAndSelectedItem()
}

var _ PopUpButton = (*NSPopUpButton)(nil)

type NSPopUpButton struct {
	NSButton
}

// MakePopUpButton create a PopUpButton from native pointer
func MakePopUpButton(ptr unsafe.Pointer) *NSPopUpButton {
	if ptr == nil {
		return nil
	}
	return &NSPopUpButton{
		NSButton: *MakeButton(ptr),
	}
}

func (p *NSPopUpButton) PullsDown() bool {
	return bool(C.PopUpButton_PullsDown(p.Ptr()))
}

func (p *NSPopUpButton) SetPullsDown(pullsDown bool) {
	C.PopUpButton_SetPullsDown(p.Ptr(), C.bool(pullsDown))
}

func (p *NSPopUpButton) AutoenablesItems() bool {
	return bool(C.PopUpButton_AutoenablesItems(p.Ptr()))
}

func (p *NSPopUpButton) SetAutoenablesItems(autoenablesItems bool) {
	C.PopUpButton_SetAutoenablesItems(p.Ptr(), C.bool(autoenablesItems))
}

func (p *NSPopUpButton) SelectedItem() MenuItem {
	return MakeMenuItem(C.PopUpButton_SelectedItem(p.Ptr()))
}

func (p *NSPopUpButton) TitleOfSelectedItem() string {
	return C.GoString(C.PopUpButton_TitleOfSelectedItem(p.Ptr()))
}

func (p *NSPopUpButton) SelectedTag() int {
	return int(C.PopUpButton_SelectedTag(p.Ptr()))
}

func (p *NSPopUpButton) Menu() Menu {
	return MakeMenu(C.PopUpButton_Menu(p.Ptr()))
}

func (p *NSPopUpButton) SetMenu(menu Menu) {
	C.PopUpButton_SetMenu(p.Ptr(), toPointer(menu))
}

func (p *NSPopUpButton) NumberOfItems() int {
	return int(C.PopUpButton_NumberOfItems(p.Ptr()))
}

func (p *NSPopUpButton) ItemArray() []MenuItem {
	var cArray C.Array = C.PopUpButton_ItemArray(p.Ptr())
	defer C.free(cArray.data)
	result := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(cArray.data))[:cArray.len:cArray.len]
	var goResult = make([]MenuItem, len(result))
	for idx, r := range result {
		goResult[idx] = MakeMenuItem(r)
	}
	return goResult
}

func (p *NSPopUpButton) ItemTitles() []string {
	var cArray C.Array = C.PopUpButton_ItemTitles(p.Ptr())
	defer C.free(cArray.data)
	result := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(cArray.data))[:cArray.len:cArray.len]
	var goResult = make([]string, len(result))
	for idx, r := range result {
		goResult[idx] = C.GoString((*C.char)(r))
	}
	return goResult
}

func (p *NSPopUpButton) LastItem() MenuItem {
	return MakeMenuItem(C.PopUpButton_LastItem(p.Ptr()))
}

func (p *NSPopUpButton) PreferredEdge() foundation.RectEdge {
	return foundation.RectEdge(C.PopUpButton_PreferredEdge(p.Ptr()))
}

func (p *NSPopUpButton) SetPreferredEdge(preferredEdge foundation.RectEdge) {
	C.PopUpButton_SetPreferredEdge(p.Ptr(), C.long(preferredEdge))
}

func NewPopUpButton(buttonFrame foundation.Rect, flag bool) PopUpButton {
	return MakePopUpButton(C.PopUpButton_NewPopUpButton(toNSRect(buttonFrame), C.bool(flag)))
}

func (p *NSPopUpButton) AddItemWithTitle(title string) {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	C.PopUpButton_AddItemWithTitle(p.Ptr(), cTitle)
}

func (p *NSPopUpButton) AddItemsWithTitles(itemTitles []string) {
	cItemTitlesData := make([]unsafe.Pointer, len(itemTitles))
	for idx, v := range itemTitles {
		cItemTitlesData[idx] = unsafe.Pointer(C.CString(v))
	}
	cItemTitles := C.Array{data: unsafe.Pointer(&cItemTitlesData[0]), len: C.int(len(itemTitles))}
	defer func() {
		for _, p := range cItemTitlesData {
			C.free(p)
		}
	}()
	C.PopUpButton_AddItemsWithTitles(p.Ptr(), cItemTitles)
}

func (p *NSPopUpButton) InsertItemWithTitle(title string, index int) {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	C.PopUpButton_InsertItemWithTitle(p.Ptr(), cTitle, C.long(index))
}

func (p *NSPopUpButton) RemoveAllItems() {
	C.PopUpButton_RemoveAllItems(p.Ptr())
}

func (p *NSPopUpButton) RemoveItemWithTitle(title string) {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	C.PopUpButton_RemoveItemWithTitle(p.Ptr(), cTitle)
}

func (p *NSPopUpButton) RemoveItemAtIndex(index int) {
	C.PopUpButton_RemoveItemAtIndex(p.Ptr(), C.long(index))
}

func (p *NSPopUpButton) SelectItem(item MenuItem) {
	C.PopUpButton_SelectItem(p.Ptr(), toPointer(item))
}

func (p *NSPopUpButton) SelectItemAtIndex(index int) {
	C.PopUpButton_SelectItemAtIndex(p.Ptr(), C.long(index))
}

func (p *NSPopUpButton) SelectItemWithTag(tag int) bool {
	return bool(C.PopUpButton_SelectItemWithTag(p.Ptr(), C.long(tag)))
}

func (p *NSPopUpButton) SelectItemWithTitle(title string) {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	C.PopUpButton_SelectItemWithTitle(p.Ptr(), cTitle)
}

func (p *NSPopUpButton) ItemAtIndex(index int) MenuItem {
	return MakeMenuItem(C.PopUpButton_ItemAtIndex(p.Ptr(), C.long(index)))
}

func (p *NSPopUpButton) ItemTitleAtIndex(index int) string {
	return C.GoString(C.PopUpButton_ItemTitleAtIndex(p.Ptr(), C.long(index)))
}

func (p *NSPopUpButton) ItemWithTitle(title string) MenuItem {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	return MakeMenuItem(C.PopUpButton_ItemWithTitle(p.Ptr(), cTitle))
}

func (p *NSPopUpButton) IndexOfItem(item MenuItem) int {
	return int(C.PopUpButton_IndexOfItem(p.Ptr(), toPointer(item)))
}

func (p *NSPopUpButton) IndexOfItemWithTag(tag int) int {
	return int(C.PopUpButton_IndexOfItemWithTag(p.Ptr(), C.long(tag)))
}

func (p *NSPopUpButton) IndexOfItemWithTitle(title string) int {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	return int(C.PopUpButton_IndexOfItemWithTitle(p.Ptr(), cTitle))
}

func (p *NSPopUpButton) SynchronizeTitleAndSelectedItem() {
	C.PopUpButton_SynchronizeTitleAndSelectedItem(p.Ptr())
}
