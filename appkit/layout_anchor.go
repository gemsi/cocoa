package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include "layout_anchor.h"
import "C"

import (
	"github.com/hsiafan/cocoa/foundation"
	"unsafe"
)

// LayoutAnchor is a factory class for creating layout constraint objects using a fluent API
type LayoutAnchor interface {
	foundation.Object

	// ConstraintsAffectingLayout return the constraints that impact the layout of the anchor
	ConstraintsAffectingLayout() []LayoutConstraint
	// HasAmbiguousLayout return whether the constraints impacting the anchor specify its location ambiguously
	HasAmbiguousLayout() bool
	// Name return the name assigned to the anchor for debugging purposes
	Name() string
	// Item return the layout item used to calculate the anchor's position
	Item() foundation.Object
	// ConstraintEqualToAnchor returns a constraint that defines one item’s attribute as equal to another
	ConstraintEqualToAnchor(anchor LayoutAnchor) LayoutConstraint
	// ConstraintEqualToAnchor2 returns a constraint that defines one item’s attribute as equal to another item’s attribute plus a constant offset
	ConstraintEqualToAnchor2(anchor LayoutAnchor, constant float64) LayoutConstraint
	// ConstraintGreaterThanOrEqualToAnchor returns a constraint that defines one item’s attribute as greater than or equal to another
	ConstraintGreaterThanOrEqualToAnchor(anchor LayoutAnchor) LayoutConstraint
	// ConstraintGreaterThanOrEqualToAnchor2 returns a constraint that defines one item’s attribute as greater than or equal to another item’s attribute plus a constant offset
	ConstraintGreaterThanOrEqualToAnchor2(anchor LayoutAnchor, constant float64) LayoutConstraint
	// ConstraintLessThanOrEqualToAnchor returns a constraint that defines one item’s attribute as less than or equal to another
	ConstraintLessThanOrEqualToAnchor(anchor LayoutAnchor) LayoutConstraint
	// ConstraintLessThanOrEqualToAnchor2 returns a constraint that defines one item’s attribute as less than or equal to another item’s attribute plus a constant offset
	ConstraintLessThanOrEqualToAnchor2(anchor LayoutAnchor, constant float64) LayoutConstraint
}

var _ LayoutAnchor = (*NSLayoutAnchor)(nil)

type NSLayoutAnchor struct {
	foundation.NSObject
}

// MakeLayoutAnchor create a LayoutAnchor from native pointer
func MakeLayoutAnchor(ptr unsafe.Pointer) *NSLayoutAnchor {
	if ptr == nil {
		return nil
	}
	return &NSLayoutAnchor{
		NSObject: *foundation.MakeObject(ptr),
	}
}

func (l *NSLayoutAnchor) ConstraintsAffectingLayout() []LayoutConstraint {
	var cArray C.Array = C.LayoutAnchor_ConstraintsAffectingLayout(l.Ptr())
	defer C.free(cArray.data)
	result := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(cArray.data))[:cArray.len:cArray.len]
	var goResult = make([]LayoutConstraint, len(result))
	for idx, r := range result {
		goResult[idx] = MakeLayoutConstraint(r)
	}
	return goResult
}

func (l *NSLayoutAnchor) HasAmbiguousLayout() bool {
	return bool(C.LayoutAnchor_HasAmbiguousLayout(l.Ptr()))
}

func (l *NSLayoutAnchor) Name() string {
	return C.GoString(C.LayoutAnchor_Name(l.Ptr()))
}

func (l *NSLayoutAnchor) Item() foundation.Object {
	return foundation.MakeObject(C.LayoutAnchor_Item(l.Ptr()))
}

func (l *NSLayoutAnchor) ConstraintEqualToAnchor(anchor LayoutAnchor) LayoutConstraint {
	return MakeLayoutConstraint(C.LayoutAnchor_ConstraintEqualToAnchor(l.Ptr(), toPointer(anchor)))
}

func (l *NSLayoutAnchor) ConstraintEqualToAnchor2(anchor LayoutAnchor, constant float64) LayoutConstraint {
	return MakeLayoutConstraint(C.LayoutAnchor_ConstraintEqualToAnchor2(l.Ptr(), toPointer(anchor), C.double(constant)))
}

func (l *NSLayoutAnchor) ConstraintGreaterThanOrEqualToAnchor(anchor LayoutAnchor) LayoutConstraint {
	return MakeLayoutConstraint(C.LayoutAnchor_ConstraintGreaterThanOrEqualToAnchor(l.Ptr(), toPointer(anchor)))
}

func (l *NSLayoutAnchor) ConstraintGreaterThanOrEqualToAnchor2(anchor LayoutAnchor, constant float64) LayoutConstraint {
	return MakeLayoutConstraint(C.LayoutAnchor_ConstraintGreaterThanOrEqualToAnchor2(l.Ptr(), toPointer(anchor), C.double(constant)))
}

func (l *NSLayoutAnchor) ConstraintLessThanOrEqualToAnchor(anchor LayoutAnchor) LayoutConstraint {
	return MakeLayoutConstraint(C.LayoutAnchor_ConstraintLessThanOrEqualToAnchor(l.Ptr(), toPointer(anchor)))
}

func (l *NSLayoutAnchor) ConstraintLessThanOrEqualToAnchor2(anchor LayoutAnchor, constant float64) LayoutConstraint {
	return MakeLayoutConstraint(C.LayoutAnchor_ConstraintLessThanOrEqualToAnchor2(l.Ptr(), toPointer(anchor), C.double(constant)))
}
