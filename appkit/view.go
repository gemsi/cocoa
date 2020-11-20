package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include "view.h"
import "C"

import (
	"github.com/hsiafan/cocoa/foundation"
	"unsafe"
)

// View is The infrastructure for drawing, printing, and handling events in an app.
type View interface {
	Responder

	// Frame return the view’s frame rectangle, which defines its position and size in its superview’s coordinate system
	Frame() foundation.Rect
	// SetFrame set the view’s frame rectangle, which defines its position and size in its superview’s coordinate system
	SetFrame(frame foundation.Rect)
	// AutoresizingMask return the view’s frame rectangle, which defines its position and size in its superview’s coordinate system
	AutoresizingMask() AutoresizingMaskOptions
	// SetAutoresizingMask set the view’s frame rectangle, which defines its position and size in its superview’s coordinate system
	SetAutoresizingMask(autoresizingMask AutoresizingMaskOptions)
	// NeedsDisplay return whether the view needs to be redrawn before being displayed.
	NeedsDisplay() bool
	// SetNeedsDisplay set whether the view needs to be redrawn before being displayed.
	SetNeedsDisplay(needsDisplay bool)
	// AddSubview adds a view to the view’s subviews so it’s displayed above its siblings.
	AddSubview(subView View)
}

var _ View = (*NSView)(nil)

type NSView struct {
	NSResponder
}

// MakeView create a View from native pointer
func MakeView(ptr unsafe.Pointer) *NSView {
	if ptr == nil {
		return nil
	}
	return &NSView{
		NSResponder: *MakeResponder(ptr),
	}
}

func (v *NSView) Frame() foundation.Rect {
	return toRect(C.View_Frame(v.Ptr()))
}

func (v *NSView) SetFrame(frame foundation.Rect) {
	C.View_SetFrame(v.Ptr(), toNSRect(frame))
}

func (v *NSView) AutoresizingMask() AutoresizingMaskOptions {
	return AutoresizingMaskOptions(C.View_AutoresizingMask(v.Ptr()))
}

func (v *NSView) SetAutoresizingMask(autoresizingMask AutoresizingMaskOptions) {
	C.View_SetAutoresizingMask(v.Ptr(), C.ulong(autoresizingMask))
}

func (v *NSView) NeedsDisplay() bool {
	return bool(C.View_NeedsDisplay(v.Ptr()))
}

func (v *NSView) SetNeedsDisplay(needsDisplay bool) {
	C.View_SetNeedsDisplay(v.Ptr(), C.bool(needsDisplay))
}

func (v *NSView) AddSubview(subView View) {
	C.View_AddSubview(v.Ptr(), toPointer(subView))
}
