package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include "layout_guide.h"
import "C"

import (
	"github.com/hsiafan/cocoa/foundation"
	"unsafe"
)

// LayoutGuide is a rectangular area that can interact with Auto Layout
type LayoutGuide interface {
	foundation.Object

	// Frame return the layout guide’s frame in its owning view’s coordinate system
	Frame() foundation.Rect
	// BottomAnchor return a layout anchor representing the bottom edge of the view’s frame
	BottomAnchor() LayoutYAxisAnchor
	// CenterXAnchor return a layout anchor representing the horizontal center of the view’s frame
	CenterXAnchor() LayoutXAxisAnchor
	// CenterYAnchor return a layout anchor representing the vertical center of the view’s frame
	CenterYAnchor() LayoutYAxisAnchor
	// HeightAnchor return
	HeightAnchor() LayoutDimension
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
	// OwningView return
	OwningView() View
}

var _ LayoutGuide = (*NSLayoutGuide)(nil)

type NSLayoutGuide struct {
	foundation.NSObject
}

// MakeLayoutGuide create a LayoutGuide from native pointer
func MakeLayoutGuide(ptr unsafe.Pointer) *NSLayoutGuide {
	if ptr == nil {
		return nil
	}
	return &NSLayoutGuide{
		NSObject: *foundation.MakeObject(ptr),
	}
}

func (l *NSLayoutGuide) Frame() foundation.Rect {
	return toRect(C.LayoutGuide_Frame(l.Ptr()))
}

func (l *NSLayoutGuide) BottomAnchor() LayoutYAxisAnchor {
	return MakeLayoutYAxisAnchor(C.LayoutGuide_BottomAnchor(l.Ptr()))
}

func (l *NSLayoutGuide) CenterXAnchor() LayoutXAxisAnchor {
	return MakeLayoutXAxisAnchor(C.LayoutGuide_CenterXAnchor(l.Ptr()))
}

func (l *NSLayoutGuide) CenterYAnchor() LayoutYAxisAnchor {
	return MakeLayoutYAxisAnchor(C.LayoutGuide_CenterYAnchor(l.Ptr()))
}

func (l *NSLayoutGuide) HeightAnchor() LayoutDimension {
	return MakeLayoutDimension(C.LayoutGuide_HeightAnchor(l.Ptr()))
}

func (l *NSLayoutGuide) LeadingAnchor() LayoutXAxisAnchor {
	return MakeLayoutXAxisAnchor(C.LayoutGuide_LeadingAnchor(l.Ptr()))
}

func (l *NSLayoutGuide) LeftAnchor() LayoutXAxisAnchor {
	return MakeLayoutXAxisAnchor(C.LayoutGuide_LeftAnchor(l.Ptr()))
}

func (l *NSLayoutGuide) RightAnchor() LayoutXAxisAnchor {
	return MakeLayoutXAxisAnchor(C.LayoutGuide_RightAnchor(l.Ptr()))
}

func (l *NSLayoutGuide) TopAnchor() LayoutYAxisAnchor {
	return MakeLayoutYAxisAnchor(C.LayoutGuide_TopAnchor(l.Ptr()))
}

func (l *NSLayoutGuide) TrailingAnchor() LayoutXAxisAnchor {
	return MakeLayoutXAxisAnchor(C.LayoutGuide_TrailingAnchor(l.Ptr()))
}

func (l *NSLayoutGuide) WidthAnchor() LayoutDimension {
	return MakeLayoutDimension(C.LayoutGuide_WidthAnchor(l.Ptr()))
}

func (l *NSLayoutGuide) OwningView() View {
	return MakeView(C.LayoutGuide_OwningView(l.Ptr()))
}

func NewLayoutGuide() LayoutGuide {
	return MakeLayoutGuide(C.LayoutGuide_NewLayoutGuide())
}
