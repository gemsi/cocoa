package widget

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include "view.h"
import "C"

import (
	"github.com/hsiafan/cocoa/foundation"
	"unsafe"
)

// The infrastructure for drawing, printing, and handling events in an app.
type View interface {
	Responder

	// Frame return the view’s frame rectangle, which defines its position and size in its superview’s coordinate system
	Frame() foundation.Rect
	// SetFrame set the view’s frame rectangle, which defines its position and size in its superview’s coordinate system
	SetFrame(value foundation.Rect)
	// AutoresizingMask return the view’s frame rectangle, which defines its position and size in its superview’s coordinate system
	AutoresizingMask() AutoresizingMaskOptions
	// SetAutoresizingMask set the view’s frame rectangle, which defines its position and size in its superview’s coordinate system
	SetAutoresizingMask(value AutoresizingMaskOptions)
}

var _ View = (*NSView)(nil)

type NSView struct {
	NSResponder
}

// Make create a View from native pointer
func MakeView(ptr unsafe.Pointer) *NSView {
	return &NSView{*MakeResponder(ptr)}
}

// New create new View
func NewView(frame foundation.Rect) View {
	id := resources.NextId()
	ptr := C.View_New(C.long(id), toNSRect(frame))

	v := &NSView{
		NSResponder: *MakeResponder(ptr),
	}

	resources.Store(id, v)

	foundation.AddDeallocHook(v, func() {
		resources.Delete(id)
	})

	return v
}

func (v *NSView) Frame() foundation.Rect {
	return toRect(C.View_Frame(v.Ptr()))
}

func (v *NSView) SetFrame(value foundation.Rect) {
	C.View_SetFrame(v.Ptr(), toNSRect(value))
}

func (v *NSView) AutoresizingMask() AutoresizingMaskOptions {
	return AutoresizingMaskOptions(C.View_AutoresizingMask(v.Ptr()))
}

func (v *NSView) SetAutoresizingMask(value AutoresizingMaskOptions) {
	C.View_SetAutoresizingMask(v.Ptr(), C.ulong(value))
}
