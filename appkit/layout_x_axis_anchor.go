package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include "layout_x_axis_anchor.h"
import "C"

import (
	"unsafe"
)

// LayoutXAxisAnchor is a factory class for creating horizontal layout constraint objects using a fluent API
type LayoutXAxisAnchor interface {
	LayoutAnchor

	// ConstraintEqualToSystemSpacingAfterAnchor returns a constraint that defines by how much the current anchor trails the specified anchor
	ConstraintEqualToSystemSpacingAfterAnchor(anchor LayoutXAxisAnchor, multiplier float64) LayoutConstraint
	// ConstraintGreaterThanOrEqualToSystemSpacingAfterAnchor returns a constraint that defines the minimum amount by which the current anchor trails the specified anchor
	ConstraintGreaterThanOrEqualToSystemSpacingAfterAnchor(anchor LayoutXAxisAnchor, multiplier float64) LayoutConstraint
	// ConstraintLessThanOrEqualToSystemSpacingAfterAnchor returns a constraint that defines the maximum amount by which the current anchor trails the specified anchor
	ConstraintLessThanOrEqualToSystemSpacingAfterAnchor(anchor LayoutXAxisAnchor, multiplier float64) LayoutConstraint
	// AnchorWithOffsetToAnchor creates a layout dimension object from two anchors
	AnchorWithOffsetToAnchor(otherAnchor LayoutXAxisAnchor) LayoutDimension
}

var _ LayoutXAxisAnchor = (*NSLayoutXAxisAnchor)(nil)

type NSLayoutXAxisAnchor struct {
	NSLayoutAnchor
}

// MakeLayoutXAxisAnchor create a LayoutXAxisAnchor from native pointer
func MakeLayoutXAxisAnchor(ptr unsafe.Pointer) *NSLayoutXAxisAnchor {
	if ptr == nil {
		return nil
	}
	return &NSLayoutXAxisAnchor{
		NSLayoutAnchor: *MakeLayoutAnchor(ptr),
	}
}

func (l *NSLayoutXAxisAnchor) ConstraintEqualToSystemSpacingAfterAnchor(anchor LayoutXAxisAnchor, multiplier float64) LayoutConstraint {
	return MakeLayoutConstraint(C.LayoutXAxisAnchor_ConstraintEqualToSystemSpacingAfterAnchor(l.Ptr(), toPointer(anchor), C.double(multiplier)))
}

func (l *NSLayoutXAxisAnchor) ConstraintGreaterThanOrEqualToSystemSpacingAfterAnchor(anchor LayoutXAxisAnchor, multiplier float64) LayoutConstraint {
	return MakeLayoutConstraint(C.LayoutXAxisAnchor_ConstraintGreaterThanOrEqualToSystemSpacingAfterAnchor(l.Ptr(), toPointer(anchor), C.double(multiplier)))
}

func (l *NSLayoutXAxisAnchor) ConstraintLessThanOrEqualToSystemSpacingAfterAnchor(anchor LayoutXAxisAnchor, multiplier float64) LayoutConstraint {
	return MakeLayoutConstraint(C.LayoutXAxisAnchor_ConstraintLessThanOrEqualToSystemSpacingAfterAnchor(l.Ptr(), toPointer(anchor), C.double(multiplier)))
}

func (l *NSLayoutXAxisAnchor) AnchorWithOffsetToAnchor(otherAnchor LayoutXAxisAnchor) LayoutDimension {
	return MakeLayoutDimension(C.LayoutXAxisAnchor_AnchorWithOffsetToAnchor(l.Ptr(), toPointer(otherAnchor)))
}
