package view

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include "view.h"
import "C"

import (
	"github.com/hsiafan/cocoa/appkit/responder"
	"github.com/hsiafan/cocoa/foundation/geometry"
	"github.com/hsiafan/cocoa/foundation/object"
	"github.com/hsiafan/cocoa/internal"
	"unsafe"
)

// The infrastructure for drawing, printing, and handling events in an app.
type View interface {
	responder.Responder

	// Frame return the view’s frame rectangle, which defines its position and size in its superview’s coordinate system
	Frame() geometry.Rect
	// SetFrame set the view’s frame rectangle, which defines its position and size in its superview’s coordinate system
	SetFrame(value geometry.Rect)
	// AutoresizingMask return the view’s frame rectangle, which defines its position and size in its superview’s coordinate system
	AutoresizingMask() AutoresizingMaskOptions
	// SetAutoresizingMask set the view’s frame rectangle, which defines its position and size in its superview’s coordinate system
	SetAutoresizingMask(value AutoresizingMaskOptions)
}

var _ View = (*NSView)(nil)

type NSView struct {
	responder.NSResponder
}

func (v *NSView) Frame() geometry.Rect {
	return toRect(C.View_Frame(v.Ptr()))
}

func (v *NSView) SetFrame(value geometry.Rect) {
	C.View_SetFrame(v.Ptr(), toNSRect(value))
}

func (v *NSView) AutoresizingMask() AutoresizingMaskOptions {
	return AutoresizingMaskOptions(C.View_AutoresizingMask(v.Ptr()))
}

func (v *NSView) SetAutoresizingMask(value AutoresizingMaskOptions) {
	C.View_SetAutoresizingMask(v.Ptr(), C.ulong(value))
}

var resources = internal.NewResourceRegistry()

// Make create a View from native pointer
func Make(ptr unsafe.Pointer) *NSView {
	return &NSView{*responder.Make(ptr)}
}

// New create new View
func New(frame geometry.Rect) View {
	id := resources.NextId()
	ptr := C.View_New(C.long(id), toNSRect(frame))

	v := &NSView{
		NSResponder: *responder.Make(ptr),
	}

	resources.Store(id, v)

	object.AddDeallocHook(v, func() {
		resources.Delete(id)
	})

	return v
}

func toNSRect(rect geometry.Rect) C.NSRect {
	return *(*C.NSRect)(unsafe.Pointer(&rect))
}
func toRect(nsRect C.NSRect) geometry.Rect {
	return *(*geometry.Rect)(unsafe.Pointer(&nsRect))
}
