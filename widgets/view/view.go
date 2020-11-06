package view

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include "view.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation/geometry"
	"unsafe"

	"github.com/hsiafan/cocoa/interaction/responder"
)

// View is interface for all NSView type
type View interface {
	responder.Responder
	// SuperView return super view of this view
	SuperView() View
	// SetFrame set frame for this view
	SetFrame(frame geometry.Rect)
	// Frame return the frame of this View
	Frame() geometry.Rect
}

var _ View = (*NSView)(nil)

// NSView is cocoa NSView
type NSView struct {
	responder.NSResponder
}

// Make create a View from native pointer
func Make(ptr unsafe.Pointer) *NSView {
	return &NSView{*responder.Make(ptr)}
}

func (*NSView) SuperView() View {
	return nil
}

func (btn *NSView) SetFrame(frame geometry.Rect) {
	C.View_SetFrame(btn.Ptr(), toNSRect(frame))
}

func (btn *NSView) Frame() geometry.Rect {
	nsFrame := C.View_Frame(btn.Ptr())
	return toRect(nsFrame)
}

func toNSRect(rect geometry.Rect) C.NSRect {
	return *(*C.NSRect)(unsafe.Pointer(&rect))
}
func toRect(nsRect C.NSRect) geometry.Rect {
	return *(*geometry.Rect)(unsafe.Pointer(&nsRect))
}
