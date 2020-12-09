package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include "menu_items.h"
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
	Submenu() Menu
	// SetSubmenu set the submenu of the menu item
	SetSubmenu(submenu Menu)
	// HasSubmenu return whether the menu item has a submenu
	HasSubmenu() bool
	// IsSeparatorItem return a menu item that is used to separate logical groups of menu commands
	IsSeparatorItem() bool
	// Menu return the menu item’s menu
	Menu() Menu
	// SetMenu set the menu item’s menu
	SetMenu(menu Menu)
	// ToolTip return a help tag for the menu item
	ToolTip() string
	// SetToolTip set a help tag for the menu item
	SetToolTip(toolTip string)
	// IsHighlighted return whether the menu item should be drawn highlighted
	IsHighlighted() bool

	// KeyEquivalent return the menu item’s unmodified key equivalent
	KeyEquivalent() string
	// SetKeyEquivalent set the menu item’s unmodified key equivalent
	SetKeyEquivalent(keyEquivalent string)
	// KeyEquivalentModifierMask return the menu item’s unmodified key equivalent
	KeyEquivalentModifierMask() EventModifierFlags
	// SetKeyEquivalentModifierMask set the menu item’s unmodified key equivalent
	SetKeyEquivalentModifierMask(keyEquivalentModifierMask EventModifierFlags)
	// UserKeyEquivalent return Tte user-assigned key equivalent for the menu item
	UserKeyEquivalent() string
	// IsAlternate return a Boolean value that marks the menu item as an alternate to the previous menu item
	IsAlternate() bool
	// SetAlternate set a Boolean value that marks the menu item as an alternate to the previous menu item
	SetAlternate(alternate bool)
	// IndentationLevel return the menu item indentation level for the menu item
	IndentationLevel() int
	// SetIndentationLevel set the menu item indentation level for the menu item
	SetIndentationLevel(indentationLevel int)
	// View return the content view for the menu item
	View() View
	// SetView set the content view for the menu item
	SetView(view View)
	// AllowsKeyEquivalentWhenHidden return
	AllowsKeyEquivalentWhenHidden() bool
	// SetAllowsKeyEquivalentWhenHidden set
	SetAllowsKeyEquivalentWhenHidden(allowsKeyEquivalentWhenHidden bool)
}

var _ MenuItem = (*NSMenuItem)(nil)

type NSMenuItem struct {
	foundation.NSObject
	action ActionHandler
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
func (m *NSMenuItem) setDelegate() {
	id := resources.NextId()
	C.MenuItem_RegisterDelegate(m.Ptr(), C.long(id))
	resources.Store(id, m)
	foundation.AddDeallocHook(m, func() {
		resources.Delete(id)
	})
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

func (m *NSMenuItem) Submenu() Menu {
	return MakeMenu(C.MenuItem_Submenu(m.Ptr()))
}

func (m *NSMenuItem) SetSubmenu(submenu Menu) {
	C.MenuItem_SetSubmenu(m.Ptr(), toPointer(submenu))
}

func (m *NSMenuItem) HasSubmenu() bool {
	return bool(C.MenuItem_HasSubmenu(m.Ptr()))
}

func (m *NSMenuItem) IsSeparatorItem() bool {
	return bool(C.MenuItem_IsSeparatorItem(m.Ptr()))
}

func (m *NSMenuItem) Menu() Menu {
	return MakeMenu(C.MenuItem_Menu(m.Ptr()))
}

func (m *NSMenuItem) SetMenu(menu Menu) {
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

func (m *NSMenuItem) SetAction(handler ActionHandler) {
	m.action = handler
}

func (m *NSMenuItem) KeyEquivalent() string {
	return C.GoString(C.MenuItem_KeyEquivalent(m.Ptr()))
}

func (m *NSMenuItem) SetKeyEquivalent(keyEquivalent string) {
	cKeyEquivalent := C.CString(keyEquivalent)
	defer C.free(unsafe.Pointer(cKeyEquivalent))
	C.MenuItem_SetKeyEquivalent(m.Ptr(), cKeyEquivalent)
}

func (m *NSMenuItem) KeyEquivalentModifierMask() EventModifierFlags {
	return EventModifierFlags(C.MenuItem_KeyEquivalentModifierMask(m.Ptr()))
}

func (m *NSMenuItem) SetKeyEquivalentModifierMask(keyEquivalentModifierMask EventModifierFlags) {
	C.MenuItem_SetKeyEquivalentModifierMask(m.Ptr(), C.ulong(keyEquivalentModifierMask))
}

func (m *NSMenuItem) UserKeyEquivalent() string {
	return C.GoString(C.MenuItem_UserKeyEquivalent(m.Ptr()))
}

func (m *NSMenuItem) IsAlternate() bool {
	return bool(C.MenuItem_IsAlternate(m.Ptr()))
}

func (m *NSMenuItem) SetAlternate(alternate bool) {
	C.MenuItem_SetAlternate(m.Ptr(), C.bool(alternate))
}

func (m *NSMenuItem) IndentationLevel() int {
	return int(C.MenuItem_IndentationLevel(m.Ptr()))
}

func (m *NSMenuItem) SetIndentationLevel(indentationLevel int) {
	C.MenuItem_SetIndentationLevel(m.Ptr(), C.long(indentationLevel))
}

func (m *NSMenuItem) View() View {
	return MakeView(C.MenuItem_View(m.Ptr()))
}

func (m *NSMenuItem) SetView(view View) {
	C.MenuItem_SetView(m.Ptr(), toPointer(view))
}

func (m *NSMenuItem) AllowsKeyEquivalentWhenHidden() bool {
	return bool(C.MenuItem_AllowsKeyEquivalentWhenHidden(m.Ptr()))
}

func (m *NSMenuItem) SetAllowsKeyEquivalentWhenHidden(allowsKeyEquivalentWhenHidden bool) {
	C.MenuItem_SetAllowsKeyEquivalentWhenHidden(m.Ptr(), C.bool(allowsKeyEquivalentWhenHidden))
}

func UsesUserKeyEquivalents() bool {
	return bool(C.MenuItem_UsesUserKeyEquivalents())
}

func SetUsesUserKeyEquivalents(usesUserKeyEquivalents bool) {
	C.MenuItem_SetUsesUserKeyEquivalents(C.bool(usesUserKeyEquivalents))
}

func (m *NSMenuItem) State() ControlStateValue {
	return ControlStateValue(C.MenuItem_State(m.Ptr()))
}

func (m *NSMenuItem) SetState(state ControlStateValue) {
	C.MenuItem_SetState(m.Ptr(), C.long(state))
}

func (m *NSMenuItem) Tag() int {
	return int(C.MenuItem_Tag(m.Ptr()))
}

func (m *NSMenuItem) SetTag(tag int) {
	C.MenuItem_SetTag(m.Ptr(), C.long(tag))
}

func NewSeparatorItem() MenuItem {
	return MakeMenuItem(C.MenuItem_NewSeparatorItem())
}

func NewMenuItem(title string, charCode string, handler ActionHandler) MenuItem {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	cCharCode := C.CString(charCode)
	defer C.free(unsafe.Pointer(cCharCode))

	id := resources.NextId()

	instance := MakeMenuItem(C.MenuItem_NewMenuItem(cTitle, cCharCode, C.long(id)))
	instance.action = handler
	resources.Store(id, instance)
	foundation.AddDeallocHook(instance, func() {
		resources.Delete(id)
	})
	//instance.setDelegate()
	return instance
}

//export MenuItem_Target_Action
func MenuItem_Target_Action(id int64, sender unsafe.Pointer) {
	menuItem := resources.Get(id).(*NSMenuItem)
	if menuItem.action != nil {
		menuItem.action(foundation.MakeObject(sender))
	}
}
