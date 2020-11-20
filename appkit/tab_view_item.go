package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include "tab_view_item.h"
import "C"

import (
	"github.com/hsiafan/cocoa/foundation"
	"unsafe"
)

// TabViewItem is an item in a tab view
type TabViewItem interface {
	foundation.Object

	// Label return the label text for the receiver to label
	Label() string
	// SetLabel set the label text for the receiver to label
	SetLabel(label string)
	// ToolTip return the tooltip displayed for the tab view item
	ToolTip() string
	// SetToolTip set the tooltip displayed for the tab view item
	SetToolTip(toolTip string)
	// TabState return the current display state of the tab associated with the receiver
	TabState() TabState
	// Identifier return the receiver’s optional identifier object to identifier
	Identifier() foundation.Object
	// SetIdentifier set the receiver’s optional identifier object to identifier
	SetIdentifier(identifier foundation.Object)
	// Color return the background color for content in the view
	Color() Color
	// SetColor set the background color for content in the view
	SetColor(color Color)
	// View return the view associated with the receiver to view
	View() View
	// SetView set the view associated with the receiver to view
	SetView(view View)
	// InitialFirstResponder return the initial first responder for the view associated with the receiver (the view that is displayed when a user clicks on the tab) to view
	InitialFirstResponder() View
	// SetInitialFirstResponder set the initial first responder for the view associated with the receiver (the view that is displayed when a user clicks on the tab) to view
	SetInitialFirstResponder(initialFirstResponder View)
	// TabView return the parent tab view for the receiver
	TabView() TabView
}

var _ TabViewItem = (*NSTabViewItem)(nil)

type NSTabViewItem struct {
	foundation.NSObject
}

// MakeTabViewItem create a TabViewItem from native pointer
func MakeTabViewItem(ptr unsafe.Pointer) *NSTabViewItem {
	if ptr == nil {
		return nil
	}
	return &NSTabViewItem{
		NSObject: *foundation.MakeObject(ptr),
	}
}

// NewTabViewItem create new TabViewItem
func NewTabViewItem(identifier foundation.Object) TabViewItem {
	id := resources.NextId()
	ptr := C.TabViewItem_initWithIdentifier(C.long(id), toPointer(identifier))
	v := &NSTabViewItem{
		NSObject: *foundation.MakeObject(ptr),
	}
	resources.Store(id, v)
	foundation.AddDeallocHook(v, func() {
		resources.Delete(id)
	})
	return v
}

func (t *NSTabViewItem) Label() string {
	return C.GoString(C.TabViewItem_Label(t.Ptr()))
}

func (t *NSTabViewItem) SetLabel(label string) {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))
	C.TabViewItem_SetLabel(t.Ptr(), c_label)
}

func (t *NSTabViewItem) ToolTip() string {
	return C.GoString(C.TabViewItem_ToolTip(t.Ptr()))
}

func (t *NSTabViewItem) SetToolTip(toolTip string) {
	c_toolTip := C.CString(toolTip)
	defer C.free(unsafe.Pointer(c_toolTip))
	C.TabViewItem_SetToolTip(t.Ptr(), c_toolTip)
}

func (t *NSTabViewItem) TabState() TabState {
	return TabState(C.TabViewItem_TabState(t.Ptr()))
}

func (t *NSTabViewItem) Identifier() foundation.Object {
	return foundation.MakeObject(C.TabViewItem_Identifier(t.Ptr()))
}

func (t *NSTabViewItem) SetIdentifier(identifier foundation.Object) {
	C.TabViewItem_SetIdentifier(t.Ptr(), toPointer(identifier))
}

func (t *NSTabViewItem) Color() Color {
	return MakeColor(C.TabViewItem_Color(t.Ptr()))
}

func (t *NSTabViewItem) SetColor(color Color) {
	C.TabViewItem_SetColor(t.Ptr(), toPointer(color))
}

func (t *NSTabViewItem) View() View {
	return MakeView(C.TabViewItem_View(t.Ptr()))
}

func (t *NSTabViewItem) SetView(view View) {
	C.TabViewItem_SetView(t.Ptr(), toPointer(view))
}

func (t *NSTabViewItem) InitialFirstResponder() View {
	return MakeView(C.TabViewItem_InitialFirstResponder(t.Ptr()))
}

func (t *NSTabViewItem) SetInitialFirstResponder(initialFirstResponder View) {
	C.TabViewItem_SetInitialFirstResponder(t.Ptr(), toPointer(initialFirstResponder))
}

func (t *NSTabViewItem) TabView() TabView {
	return MakeTabView(C.TabViewItem_TabView(t.Ptr()))
}
