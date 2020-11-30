package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include "tab_view.h"
import "C"

import (
	"github.com/hsiafan/cocoa/foundation"
	"unsafe"
)

// TabView is a multipage interface that displays one page at a time
type TabView interface {
	View

	// NumberOfTabViewItems return the number of items in the tab view’s array of tab view items
	NumberOfTabViewItems() int
	// TabViewType return
	TabViewType() TabViewType
	// SetTabViewType set
	SetTabViewType(tabViewType TabViewType)
	// TabPosition return
	TabPosition() TabPosition
	// SetTabPosition set
	SetTabPosition(tabPosition TabPosition)
	// TabViewBorderType return
	TabViewBorderType() TabViewBorderType
	// SetTabViewBorderType set
	SetTabViewBorderType(tabViewBorderType TabViewBorderType)
	// AllowsTruncatedLabels return if the tab view allows truncating for labels that don’t fit on a tab
	AllowsTruncatedLabels() bool
	// SetAllowsTruncatedLabels set if the tab view allows truncating for labels that don’t fit on a tab
	SetAllowsTruncatedLabels(allowsTruncatedLabels bool)
	// SelectedTabViewItem return the tab view item for the currently selected tab
	SelectedTabViewItem() TabViewItem
	// Font return the font used for the tab view’s label text
	Font() Font
	// SetFont set the font used for the tab view’s label text
	SetFont(font Font)
	// MinimumSize return the minimum size necessary for the tab view to display tabs in a useful way
	MinimumSize() foundation.Size
	// ControlSize return the size of the tab view
	ControlSize() ControlSize
	// SetControlSize set the size of the tab view
	SetControlSize(controlSize ControlSize)
	// AddTabViewItem adds the specified tab item
	AddTabViewItem(tabViewItem TabViewItem)
	// InsertTabViewItem inserts the specified item into the tab view’s array of tab view items at the specified index
	InsertTabViewItem(tabViewItem TabViewItem, index int)
	// RemoveTabViewItem removes the specified item from the tab view’s array of tab view items
	RemoveTabViewItem(tabViewItem TabViewItem)
	// IndexOfTabViewItem returns the index of the specified item in the tab view
	IndexOfTabViewItem(tabViewItem TabViewItem) int
	// TabViewItemAtIndex returns the tab view item at index in the tab view’s array of items
	TabViewItemAtIndex(index int) TabViewItem
	// SelectFirstTabViewItem selects the first tab view item
	SelectFirstTabViewItem(sender foundation.Object)
	// SelectLastTabViewItem selects the last tab view item
	SelectLastTabViewItem(sender foundation.Object)
	// SelectNextTabViewItem selects the next tab view item
	SelectNextTabViewItem(sender foundation.Object)
	// SelectPreviousTabViewItem selects the previous tab view item
	SelectPreviousTabViewItem(sender foundation.Object)
	// SelectTabViewItem selects the specified tab view item
	SelectTabViewItem(tabViewItem TabViewItem)
	// SelectTabViewItemAtIndex selects the tab view item specified by index
	SelectTabViewItemAtIndex(index int)
	// TabViewDidChangeNumberOfTabViewItems informs the delegate that the number of tab view items in tabView has changed
	TabViewDidChangeNumberOfTabViewItems(callback func(view TabView))
	_tabViewDidChangeNumberOfTabViewItems() func(view TabView)
	// ShouldSelectTabViewItem invoked just before tabViewItem in tabView is selected. Return true if the tab view item should be selected, otherwise no.
	ShouldSelectTabViewItem(callback func(view TabView, tabViewItem TabViewItem) bool)
	_shouldSelectTabViewItem() func(view TabView, tabViewItem TabViewItem) bool
	// WillSelectTabViewItem informs the delegate that tabView is about to select tabViewItem
	WillSelectTabViewItem(callback func(view TabView, tabViewItem TabViewItem))
	_willSelectTabViewItem() func(view TabView, tabViewItem TabViewItem)
	// DidSelectTabViewItem informs the delegate that tabView has selected tabViewItem
	DidSelectTabViewItem(callback func(view TabView, tabViewItem TabViewItem))
	_didSelectTabViewItem() func(view TabView, tabViewItem TabViewItem)
}

var _ TabView = (*NSTabView)(nil)

type NSTabView struct {
	NSView
	tabViewDidChangeNumberOfTabViewItems func(view TabView)
	shouldSelectTabViewItem              func(view TabView, tabViewItem TabViewItem) bool
	willSelectTabViewItem                func(view TabView, tabViewItem TabViewItem)
	didSelectTabViewItem                 func(view TabView, tabViewItem TabViewItem)
}

// MakeTabView create a TabView from native pointer
func MakeTabView(ptr unsafe.Pointer) *NSTabView {
	if ptr == nil {
		return nil
	}
	return &NSTabView{
		NSView: *MakeView(ptr),
	}
}
func (t *NSTabView) setDelegate() {
	id := resources.NextId()
	C.TabView_RegisterDelegate(t.Ptr(), C.long(id))
	resources.Store(id, t)
	foundation.AddDeallocHook(t, func() {
		resources.Delete(id)
	})
}

func (t *NSTabView) NumberOfTabViewItems() int {
	return int(C.TabView_NumberOfTabViewItems(t.Ptr()))
}

func (t *NSTabView) TabViewType() TabViewType {
	return TabViewType(C.TabView_TabViewType(t.Ptr()))
}

func (t *NSTabView) SetTabViewType(tabViewType TabViewType) {
	C.TabView_SetTabViewType(t.Ptr(), C.ulong(tabViewType))
}

func (t *NSTabView) TabPosition() TabPosition {
	return TabPosition(C.TabView_TabPosition(t.Ptr()))
}

func (t *NSTabView) SetTabPosition(tabPosition TabPosition) {
	C.TabView_SetTabPosition(t.Ptr(), C.ulong(tabPosition))
}

func (t *NSTabView) TabViewBorderType() TabViewBorderType {
	return TabViewBorderType(C.TabView_TabViewBorderType(t.Ptr()))
}

func (t *NSTabView) SetTabViewBorderType(tabViewBorderType TabViewBorderType) {
	C.TabView_SetTabViewBorderType(t.Ptr(), C.ulong(tabViewBorderType))
}

func (t *NSTabView) AllowsTruncatedLabels() bool {
	return bool(C.TabView_AllowsTruncatedLabels(t.Ptr()))
}

func (t *NSTabView) SetAllowsTruncatedLabels(allowsTruncatedLabels bool) {
	C.TabView_SetAllowsTruncatedLabels(t.Ptr(), C.bool(allowsTruncatedLabels))
}

func (t *NSTabView) SelectedTabViewItem() TabViewItem {
	return MakeTabViewItem(C.TabView_SelectedTabViewItem(t.Ptr()))
}

func (t *NSTabView) Font() Font {
	return MakeFont(C.TabView_Font(t.Ptr()))
}

func (t *NSTabView) SetFont(font Font) {
	C.TabView_SetFont(t.Ptr(), toPointer(font))
}

func (t *NSTabView) MinimumSize() foundation.Size {
	return toSize(C.TabView_MinimumSize(t.Ptr()))
}

func (t *NSTabView) ControlSize() ControlSize {
	return ControlSize(C.TabView_ControlSize(t.Ptr()))
}

func (t *NSTabView) SetControlSize(controlSize ControlSize) {
	C.TabView_SetControlSize(t.Ptr(), C.ulong(controlSize))
}

func NewTabView(frame foundation.Rect) TabView {
	instance := MakeTabView(C.TabView_NewTabView(toNSRect(frame)))
	instance.setDelegate()
	return instance
}

func (t *NSTabView) AddTabViewItem(tabViewItem TabViewItem) {
	C.TabView_AddTabViewItem(t.Ptr(), toPointer(tabViewItem))
}

func (t *NSTabView) InsertTabViewItem(tabViewItem TabViewItem, index int) {
	C.TabView_InsertTabViewItem(t.Ptr(), toPointer(tabViewItem), C.long(index))
}

func (t *NSTabView) RemoveTabViewItem(tabViewItem TabViewItem) {
	C.TabView_RemoveTabViewItem(t.Ptr(), toPointer(tabViewItem))
}

func (t *NSTabView) IndexOfTabViewItem(tabViewItem TabViewItem) int {
	return int(C.TabView_IndexOfTabViewItem(t.Ptr(), toPointer(tabViewItem)))
}

func (t *NSTabView) TabViewItemAtIndex(index int) TabViewItem {
	return MakeTabViewItem(C.TabView_TabViewItemAtIndex(t.Ptr(), C.long(index)))
}

func (t *NSTabView) SelectFirstTabViewItem(sender foundation.Object) {
	C.TabView_SelectFirstTabViewItem(t.Ptr(), toPointer(sender))
}

func (t *NSTabView) SelectLastTabViewItem(sender foundation.Object) {
	C.TabView_SelectLastTabViewItem(t.Ptr(), toPointer(sender))
}

func (t *NSTabView) SelectNextTabViewItem(sender foundation.Object) {
	C.TabView_SelectNextTabViewItem(t.Ptr(), toPointer(sender))
}

func (t *NSTabView) SelectPreviousTabViewItem(sender foundation.Object) {
	C.TabView_SelectPreviousTabViewItem(t.Ptr(), toPointer(sender))
}

func (t *NSTabView) SelectTabViewItem(tabViewItem TabViewItem) {
	C.TabView_SelectTabViewItem(t.Ptr(), toPointer(tabViewItem))
}

func (t *NSTabView) SelectTabViewItemAtIndex(index int) {
	C.TabView_SelectTabViewItemAtIndex(t.Ptr(), C.long(index))
}

func (t *NSTabView) TabViewDidChangeNumberOfTabViewItems(callback func(view TabView)) {
	t.tabViewDidChangeNumberOfTabViewItems = callback
}

func (t *NSTabView) _tabViewDidChangeNumberOfTabViewItems() func(view TabView) {
	return t.tabViewDidChangeNumberOfTabViewItems
}

func (t *NSTabView) ShouldSelectTabViewItem(callback func(view TabView, tabViewItem TabViewItem) bool) {
	t.shouldSelectTabViewItem = callback
}

func (t *NSTabView) _shouldSelectTabViewItem() func(view TabView, tabViewItem TabViewItem) bool {
	return t.shouldSelectTabViewItem
}

func (t *NSTabView) WillSelectTabViewItem(callback func(view TabView, tabViewItem TabViewItem)) {
	t.willSelectTabViewItem = callback
}

func (t *NSTabView) _willSelectTabViewItem() func(view TabView, tabViewItem TabViewItem) {
	return t.willSelectTabViewItem
}

func (t *NSTabView) DidSelectTabViewItem(callback func(view TabView, tabViewItem TabViewItem)) {
	t.didSelectTabViewItem = callback
}

func (t *NSTabView) _didSelectTabViewItem() func(view TabView, tabViewItem TabViewItem) {
	return t.didSelectTabViewItem
}

//export TabView_Delegate_TabViewDidChangeNumberOfTabViewItems
func TabView_Delegate_TabViewDidChangeNumberOfTabViewItems(id int64, view unsafe.Pointer) {
	tabView := resources.Get(id).(TabView)
	callback := tabView._tabViewDidChangeNumberOfTabViewItems()
	if callback != nil {
		callback(MakeTabView(view))
	}
}

//export TabView_Delegate_ShouldSelectTabViewItem
func TabView_Delegate_ShouldSelectTabViewItem(id int64, view unsafe.Pointer, tabViewItem unsafe.Pointer) bool {
	tabView := resources.Get(id).(TabView)
	callback := tabView._shouldSelectTabViewItem()
	if callback != nil {
		return callback(MakeTabView(view), MakeTabViewItem(tabViewItem))
	}
	return true
}

//export TabView_Delegate_WillSelectTabViewItem
func TabView_Delegate_WillSelectTabViewItem(id int64, view unsafe.Pointer, tabViewItem unsafe.Pointer) {
	tabView := resources.Get(id).(TabView)
	callback := tabView._willSelectTabViewItem()
	if callback != nil {
		callback(MakeTabView(view), MakeTabViewItem(tabViewItem))
	}
}

//export TabView_Delegate_DidSelectTabViewItem
func TabView_Delegate_DidSelectTabViewItem(id int64, view unsafe.Pointer, tabViewItem unsafe.Pointer) {
	tabView := resources.Get(id).(TabView)
	callback := tabView._didSelectTabViewItem()
	if callback != nil {
		callback(MakeTabView(view), MakeTabViewItem(tabViewItem))
	}
}
