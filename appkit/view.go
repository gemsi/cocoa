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
	// NeedsDisplay return whether the view needs to be redrawn before being displayed
	NeedsDisplay() bool
	// SetNeedsDisplay set whether the view needs to be redrawn before being displayed
	SetNeedsDisplay(needsDisplay bool)
	// TranslatesAutoresizingMaskIntoConstraints return whether the view’s autoresizing mask is translated into Auto Layout constraints
	TranslatesAutoresizingMaskIntoConstraints() bool
	// SetTranslatesAutoresizingMaskIntoConstraints set whether the view’s autoresizing mask is translated into Auto Layout constraints
	SetTranslatesAutoresizingMaskIntoConstraints(translatesAutoresizingMaskIntoConstraints bool)
	// BottomAnchor return a layout anchor representing the bottom edge of the view’s frame
	BottomAnchor() LayoutYAxisAnchor
	// CenterXAnchor return a layout anchor representing the horizontal center of the view’s frame
	CenterXAnchor() LayoutXAxisAnchor
	// CenterYAnchor return a layout anchor representing the vertical center of the view’s frame
	CenterYAnchor() LayoutYAxisAnchor
	// FirstBaselineAnchor return
	FirstBaselineAnchor() LayoutYAxisAnchor
	// HeightAnchor return
	HeightAnchor() LayoutDimension
	// LastBaselineAnchor return
	LastBaselineAnchor() LayoutYAxisAnchor
	// LeadingAnchor return
	LeadingAnchor() LayoutXAxisAnchor
	// LeftAnchor return
	LeftAnchor() LayoutXAxisAnchor
	// RightAnchor return
	RightAnchor() LayoutXAxisAnchor
	// TopAnchor return
	TopAnchor() LayoutYAxisAnchor
	// TrailingAnchor return
	TrailingAnchor() LayoutXAxisAnchor
	// WidthAnchor return
	WidthAnchor() LayoutDimension
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

func (v *NSView) TranslatesAutoresizingMaskIntoConstraints() bool {
	return bool(C.View_TranslatesAutoresizingMaskIntoConstraints(v.Ptr()))
}

func (v *NSView) SetTranslatesAutoresizingMaskIntoConstraints(translatesAutoresizingMaskIntoConstraints bool) {
	C.View_SetTranslatesAutoresizingMaskIntoConstraints(v.Ptr(), C.bool(translatesAutoresizingMaskIntoConstraints))
}

func (v *NSView) BottomAnchor() LayoutYAxisAnchor {
	return MakeLayoutYAxisAnchor(C.View_BottomAnchor(v.Ptr()))
}

func (v *NSView) CenterXAnchor() LayoutXAxisAnchor {
	return MakeLayoutXAxisAnchor(C.View_CenterXAnchor(v.Ptr()))
}

func (v *NSView) CenterYAnchor() LayoutYAxisAnchor {
	return MakeLayoutYAxisAnchor(C.View_CenterYAnchor(v.Ptr()))
}

func (v *NSView) FirstBaselineAnchor() LayoutYAxisAnchor {
	return MakeLayoutYAxisAnchor(C.View_FirstBaselineAnchor(v.Ptr()))
}

func (v *NSView) HeightAnchor() LayoutDimension {
	return MakeLayoutDimension(C.View_HeightAnchor(v.Ptr()))
}

func (v *NSView) LastBaselineAnchor() LayoutYAxisAnchor {
	return MakeLayoutYAxisAnchor(C.View_LastBaselineAnchor(v.Ptr()))
}

func (v *NSView) LeadingAnchor() LayoutXAxisAnchor {
	return MakeLayoutXAxisAnchor(C.View_LeadingAnchor(v.Ptr()))
}

func (v *NSView) LeftAnchor() LayoutXAxisAnchor {
	return MakeLayoutXAxisAnchor(C.View_LeftAnchor(v.Ptr()))
}

func (v *NSView) RightAnchor() LayoutXAxisAnchor {
	return MakeLayoutXAxisAnchor(C.View_RightAnchor(v.Ptr()))
}

func (v *NSView) TopAnchor() LayoutYAxisAnchor {
	return MakeLayoutYAxisAnchor(C.View_TopAnchor(v.Ptr()))
}

func (v *NSView) TrailingAnchor() LayoutXAxisAnchor {
	return MakeLayoutXAxisAnchor(C.View_TrailingAnchor(v.Ptr()))
}

func (v *NSView) WidthAnchor() LayoutDimension {
	return MakeLayoutDimension(C.View_WidthAnchor(v.Ptr()))
}

func (v *NSView) AddSubview(subView View) {
	C.View_AddSubview(v.Ptr(), toPointer(subView))
}
