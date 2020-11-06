package view

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include "view.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"unsafe"

	"github.com/hsiafan/cocoa/interaction/responder"
)

// View is interface for all NSView type
type View interface {
	responder.Responder
	// SuperView return super view of this view
	SuperView() View
	// SetFrameRect set frame for this view
	SetFrameRect(frame foundation.Rect)
	// SetFrame set frame for this view
	SetFrame(x, y, width, height int)
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

func (btn *NSView) SetFrameRect(frame foundation.Rect) {
	C.View_SetFrame(btn.Ptr(), C.int(frame.X), C.int(frame.Y), C.int(frame.Width), C.int(frame.Height))
}

func (btn *NSView) SetFrame(x, y, width, height int) {
	C.View_SetFrame(btn.Ptr(), C.int(x), C.int(y), C.int(width), C.int(height))
}
