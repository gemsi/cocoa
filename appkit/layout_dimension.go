package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include "layout_dimension.h"
import "C"

import (
	"unsafe"
)

// LayoutDimension is a factory class for creating size-based layout constraint objects using a fluent API
type LayoutDimension interface {
	LayoutAnchor

	// ConstraintEqualToConstant returns a constraint that defines a constant size for the anchor’s size attribute.
	ConstraintEqualToConstant(constant float64) LayoutConstraint
	// ConstraintEqualToAnchor3 returns a constraint that defines the anchor’s size attribute as equal to the specified size attribute multiplied by a constant plus an offset
	ConstraintEqualToAnchor3(anchor LayoutDimension, multiplier float64, constant float64) LayoutConstraint
	// ConstraintGreaterThanOrEqualToConstant returns a constraint that defines the minimum size for the anchor’s size attribute
	ConstraintGreaterThanOrEqualToConstant(constant float64) LayoutConstraint
	// ConstraintGreaterThanOrEqualToAnchor3 returns a constraint that defines the anchor’s size attribute as greater than or equal to the specified anchor multiplied by the constant
	ConstraintGreaterThanOrEqualToAnchor3(anchor LayoutDimension, multiplier float64, constant float64) LayoutConstraint
	// ConstraintLessThanOrEqualToConstant returns a constraint that defines the maximum size for the anchor’s size attribute
	ConstraintLessThanOrEqualToConstant(constant float64) LayoutConstraint
	// ConstraintLessThanOrEqualToAnchor3 returns a constraint that defines the anchor’s size attribute as greater than or equal to the specified anchor multiplied by the constant plus an offset
	ConstraintLessThanOrEqualToAnchor3(anchor LayoutDimension, multiplier float64, constant float64) LayoutConstraint
}

var _ LayoutDimension = (*NSLayoutDimension)(nil)

type NSLayoutDimension struct {
	NSLayoutAnchor
}

// MakeLayoutDimension create a LayoutDimension from native pointer
func MakeLayoutDimension(ptr unsafe.Pointer) *NSLayoutDimension {
	if ptr == nil {
		return nil
	}
	return &NSLayoutDimension{
		NSLayoutAnchor: *MakeLayoutAnchor(ptr),
	}
}

func (l *NSLayoutDimension) ConstraintEqualToConstant(constant float64) LayoutConstraint {
	return MakeLayoutConstraint(C.LayoutDimension_ConstraintEqualToConstant(l.Ptr(), C.double(constant)))
}

func (l *NSLayoutDimension) ConstraintEqualToAnchor3(anchor LayoutDimension, multiplier float64, constant float64) LayoutConstraint {
	return MakeLayoutConstraint(C.LayoutDimension_ConstraintEqualToAnchor3(l.Ptr(), toPointer(anchor), C.double(multiplier), C.double(constant)))
}

func (l *NSLayoutDimension) ConstraintGreaterThanOrEqualToConstant(constant float64) LayoutConstraint {
	return MakeLayoutConstraint(C.LayoutDimension_ConstraintGreaterThanOrEqualToConstant(l.Ptr(), C.double(constant)))
}

func (l *NSLayoutDimension) ConstraintGreaterThanOrEqualToAnchor3(anchor LayoutDimension, multiplier float64, constant float64) LayoutConstraint {
	return MakeLayoutConstraint(C.LayoutDimension_ConstraintGreaterThanOrEqualToAnchor3(l.Ptr(), toPointer(anchor), C.double(multiplier), C.double(constant)))
}

func (l *NSLayoutDimension) ConstraintLessThanOrEqualToConstant(constant float64) LayoutConstraint {
	return MakeLayoutConstraint(C.LayoutDimension_ConstraintLessThanOrEqualToConstant(l.Ptr(), C.double(constant)))
}

func (l *NSLayoutDimension) ConstraintLessThanOrEqualToAnchor3(anchor LayoutDimension, multiplier float64, constant float64) LayoutConstraint {
	return MakeLayoutConstraint(C.LayoutDimension_ConstraintLessThanOrEqualToAnchor3(l.Ptr(), toPointer(anchor), C.double(multiplier), C.double(constant)))
}
