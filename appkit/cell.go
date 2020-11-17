package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include "cell.h"
import "C"

import (
	"github.com/hsiafan/cocoa/foundation"
	"unsafe"
)

// a mechanism for displaying text or images in a view object without the overhead of a full NSView subclass
type Cell interface {
	foundation.Object

	// Wraps return whether the cell wraps text whose length that exceeds the cell’s frame
	Wraps() bool
	// SetWraps set whether the cell wraps text whose length that exceeds the cell’s frame
	SetWraps(wraps bool)
	// IsScrollable return whether excess text scrolls past the cell’s bounds
	IsScrollable() bool
	// SetScrollable set whether excess text scrolls past the cell’s bounds
	SetScrollable(scrollable bool)
	// IsEditable return whether the cell is editable
	IsEditable() bool
	// SetEditable set whether the cell is editable
	SetEditable(editable bool)
	// IsSelectable return whether the cell’s text can be selected
	IsSelectable() bool
	// SetSelectable set whether the cell’s text can be selected
	SetSelectable(selectable bool)
}

var _ Cell = (*NSCell)(nil)

type NSCell struct {
	foundation.NSObject
}

// Make create a Cell from native pointer
func MakeCell(ptr unsafe.Pointer) *NSCell {
	return &NSCell{
		NSObject: *foundation.MakeObject(ptr),
	}
}

func (c *NSCell) Wraps() bool {
	return bool(C.Cell_Wraps(c.Ptr()))
}

func (c *NSCell) SetWraps(wraps bool) {
	C.Cell_SetWraps(c.Ptr(), C.bool(wraps))
}

func (c *NSCell) IsScrollable() bool {
	return bool(C.Cell_IsScrollable(c.Ptr()))
}

func (c *NSCell) SetScrollable(scrollable bool) {
	C.Cell_SetScrollable(c.Ptr(), C.bool(scrollable))
}

func (c *NSCell) IsEditable() bool {
	return bool(C.Cell_IsEditable(c.Ptr()))
}

func (c *NSCell) SetEditable(editable bool) {
	C.Cell_SetEditable(c.Ptr(), C.bool(editable))
}

func (c *NSCell) IsSelectable() bool {
	return bool(C.Cell_IsSelectable(c.Ptr()))
}

func (c *NSCell) SetSelectable(selectable bool) {
	C.Cell_SetSelectable(c.Ptr(), C.bool(selectable))
}
