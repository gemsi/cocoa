package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include "layout_y_axis_anchor.h"
import "C"

import (
	"unsafe"
)

// LayoutYAxisAnchor is a factory class for creating vertical layout constraint objects using a fluent API
type LayoutYAxisAnchor interface {
	LayoutAnchor

	// ConstraintEqualToSystemSpacingBelowAnchor returns a constraint that defines by how much the current anchor below the specified anchor
	ConstraintEqualToSystemSpacingBelowAnchor(anchor LayoutYAxisAnchor, multiplier float64) LayoutConstraint
	// ConstraintGreaterThanOrEqualToSystemSpacingBelowAnchor returns a constraint that defines the minimum amount by which the current anchor below the specified anchor
	ConstraintGreaterThanOrEqualToSystemSpacingBelowAnchor(anchor LayoutYAxisAnchor, multiplier float64) LayoutConstraint
	// ConstraintLessThanOrEqualToSystemSpacingBelowAnchor returns a constraint that defines the maximum amount by which the current anchor below the specified anchor
	ConstraintLessThanOrEqualToSystemSpacingBelowAnchor(anchor LayoutYAxisAnchor, multiplier float64) LayoutConstraint
	// AnchorWithOffsetToAnchor creates a layout dimension object from two anchors
	AnchorWithOffsetToAnchor(otherAnchor LayoutYAxisAnchor) LayoutYAxisAnchor
}

var _ LayoutYAxisAnchor = (*NSLayoutYAxisAnchor)(nil)

type NSLayoutYAxisAnchor struct {
	NSLayoutAnchor
}

// MakeLayoutYAxisAnchor create a LayoutYAxisAnchor from native pointer
func MakeLayoutYAxisAnchor(ptr unsafe.Pointer) *NSLayoutYAxisAnchor {
	if ptr == nil {
		return nil
	}
	return &NSLayoutYAxisAnchor{
		NSLayoutAnchor: *MakeLayoutAnchor(ptr),
	}
}

func (l *NSLayoutYAxisAnchor) ConstraintEqualToSystemSpacingBelowAnchor(anchor LayoutYAxisAnchor, multiplier float64) LayoutConstraint {
	return MakeLayoutConstraint(C.LayoutYAxisAnchor_ConstraintEqualToSystemSpacingBelowAnchor(l.Ptr(), toPointer(anchor), C.double(multiplier)))
}

func (l *NSLayoutYAxisAnchor) ConstraintGreaterThanOrEqualToSystemSpacingBelowAnchor(anchor LayoutYAxisAnchor, multiplier float64) LayoutConstraint {
	return MakeLayoutConstraint(C.LayoutYAxisAnchor_ConstraintGreaterThanOrEqualToSystemSpacingBelowAnchor(l.Ptr(), toPointer(anchor), C.double(multiplier)))
}

func (l *NSLayoutYAxisAnchor) ConstraintLessThanOrEqualToSystemSpacingBelowAnchor(anchor LayoutYAxisAnchor, multiplier float64) LayoutConstraint {
	return MakeLayoutConstraint(C.LayoutYAxisAnchor_ConstraintLessThanOrEqualToSystemSpacingBelowAnchor(l.Ptr(), toPointer(anchor), C.double(multiplier)))
}

func (l *NSLayoutYAxisAnchor) AnchorWithOffsetToAnchor(otherAnchor LayoutYAxisAnchor) LayoutYAxisAnchor {
	return MakeLayoutYAxisAnchor(C.LayoutYAxisAnchor_AnchorWithOffsetToAnchor(l.Ptr(), toPointer(otherAnchor)))
}
