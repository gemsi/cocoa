package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #import "layout_anchor.h"
import "C"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"unsafe"
)

//LayoutAnchor is a factory class for creating layout constraint objects using a fluent API.
type LayoutAnchor interface {
	foundation.Object
	ConstraintEqualTo(anchor LayoutAnchor) LayoutConstraint
	ConstraintEqualTo2(anchor LayoutAnchor, constant float64) LayoutConstraint
	ConstraintGreaterThanOrEqualTo(anchor LayoutAnchor) LayoutConstraint
	ConstraintGreaterThanOrEqualTo2(anchor LayoutAnchor, constant float64) LayoutConstraint
	ConstraintLessThanOrEqualTo(anchor LayoutAnchor) LayoutConstraint
	ConstraintLessThanOrEqualTo2(anchor LayoutAnchor, constant float64) LayoutConstraint

	// The constraints that impact the layout of the anchor.
	ConstraintsAffectingLayout() []LayoutConstraint
	// whether the constraints impacting the anchor specify its location ambiguously
	HasAmbiguousLayout() bool
	// The layout item used to calculate the anchor's position.
	Item() foundation.Object
	// The name assigned to the anchor for debugging purposes.
	Name() string
}

var _ LayoutAnchor = (*NSLayoutAnchor)(nil)

type NSLayoutAnchor struct {
	foundation.NSObject
}

func (a *NSLayoutAnchor) ConstraintEqualTo(anchor LayoutAnchor) LayoutConstraint {
	return MakeLayoutConstraint(C.LayoutAnchor_ConstraintEqualTo(a.Ptr(), anchor.Ptr()))
}

func (a *NSLayoutAnchor) ConstraintEqualTo2(anchor LayoutAnchor, constant float64) LayoutConstraint {
	return MakeLayoutConstraint(C.LayoutAnchor_ConstraintEqualTo2(a.Ptr(), anchor.Ptr(), C.double(constant)))
}

func (a *NSLayoutAnchor) ConstraintGreaterThanOrEqualTo(anchor LayoutAnchor) LayoutConstraint {
	return MakeLayoutConstraint(C.LayoutAnchor_ConstraintGreaterThanOrEqualTo(a.Ptr(), anchor.Ptr()))
}

func (a *NSLayoutAnchor) ConstraintGreaterThanOrEqualTo2(anchor LayoutAnchor, constant float64) LayoutConstraint {
	return MakeLayoutConstraint(C.LayoutAnchor_ConstraintGreaterThanOrEqualTo2(a.Ptr(), anchor.Ptr(), C.double(constant)))
}

func (a *NSLayoutAnchor) ConstraintLessThanOrEqualTo(anchor LayoutAnchor) LayoutConstraint {
	return MakeLayoutConstraint(C.LayoutAnchor_ConstraintLessThanOrEqualTo(a.Ptr(), anchor.Ptr()))
}

func (a *NSLayoutAnchor) ConstraintLessThanOrEqualTo2(anchor LayoutAnchor, constant float64) LayoutConstraint {
	return MakeLayoutConstraint(C.LayoutAnchor_ConstraintLessThanOrEqualTo2(a.Ptr(), anchor.Ptr(), C.double(constant)))
}

func (a *NSLayoutAnchor) ConstraintsAffectingLayout() []LayoutConstraint {
	panic("implement me")
}

func (a *NSLayoutAnchor) HasAmbiguousLayout() bool {
	panic("implement me")
}

func (a *NSLayoutAnchor) Item() foundation.Object {
	panic("implement me")
}

func (a *NSLayoutAnchor) Name() string {
	panic("implement me")
}

func MakeLayoutAnchor(ptr unsafe.Pointer) *NSLayoutAnchor {
	return &NSLayoutAnchor{
		NSObject: *foundation.MakeObject(ptr),
	}
}

type LayoutDimension interface {
	LayoutAnchor
	// Returns a constraint that defines the anchor’s size attribute as equal to the specified size attribute multiplied by a constant plus an offset.
	ConstraintEqualTo3(anchor LayoutDimension, multiplier float64, constant float64) LayoutConstraint
	// Returns a constraint that defines a constant size for the anchor’s size attribute.
	ConstraintEqualToConstant(constant float64) LayoutConstraint
	// Returns a constraint that defines the anchor’s size attribute as greater than or equal to the specified anchor multiplied by the constant.
	ConstraintGreaterThanOrEqualTo3(anchor LayoutDimension, multiplier float64, constant float64) LayoutConstraint
	ConstraintGreaterThanOrEqualToConstant(constant float64) LayoutConstraint
	ConstraintLessThanOrEqualTo3(anchor LayoutDimension, multiplier float64, constant float64) LayoutConstraint
	ConstraintLessThanOrEqualToConstant(constant float64) LayoutConstraint
}

var _ LayoutDimension = (*NSLayoutDimension)(nil)

type NSLayoutDimension struct {
	NSLayoutAnchor
}

func (d *NSLayoutDimension) ConstraintEqualTo3(anchor LayoutDimension, multiplier float64, constant float64) LayoutConstraint {
	return MakeLayoutConstraint(C.LayoutDimension_ConstraintEqualTo3(d.Ptr(), anchor.Ptr(), C.double(multiplier), C.double(constant)))
}

func (d *NSLayoutDimension) ConstraintEqualToConstant(constant float64) LayoutConstraint {
	return MakeLayoutConstraint(C.LayoutDimension_ConstraintEqualToConstant(d.Ptr(), C.double(constant)))
}

func (d *NSLayoutDimension) ConstraintGreaterThanOrEqualTo3(anchor LayoutDimension, multiplier float64, constant float64) LayoutConstraint {
	return MakeLayoutConstraint(C.LayoutDimension_ConstraintGreaterThanOrEqualTo3(d.Ptr(), anchor.Ptr(), C.double(multiplier), C.double(constant)))
}

func (d *NSLayoutDimension) ConstraintGreaterThanOrEqualToConstant(constant float64) LayoutConstraint {
	return MakeLayoutConstraint(C.LayoutDimension_ConstraintGreaterThanOrEqualToConstant(d.Ptr(), C.double(constant)))
}

func (d *NSLayoutDimension) ConstraintLessThanOrEqualTo3(anchor LayoutDimension, multiplier float64, constant float64) LayoutConstraint {
	return MakeLayoutConstraint(C.LayoutDimension_ConstraintLessThanOrEqualTo3(d.Ptr(), anchor.Ptr(), C.double(multiplier), C.double(constant)))
}

func (d *NSLayoutDimension) ConstraintLessThanOrEqualToConstant(constant float64) LayoutConstraint {
	return MakeLayoutConstraint(C.LayoutDimension_ConstraintLessThanOrEqualToConstant(d.Ptr(), C.double(constant)))
}

func MakeLayoutDimension(ptr unsafe.Pointer) *NSLayoutDimension {
	return &NSLayoutDimension{
		NSLayoutAnchor: *MakeLayoutAnchor(ptr),
	}
}

type LayoutXAxisAnchor interface {
	LayoutAnchor
	// Returns a constraint that defines by how much the current anchor trails the specified anchor.
	ConstraintEqualToSystemSpacingAfter(anchor LayoutXAxisAnchor, multiplier float64) LayoutConstraint
	ConstraintGreaterThanOrEqualToSystemSpacingAfter(anchor LayoutXAxisAnchor, multiplier float64) LayoutConstraint
	ConstraintLessThanOrEqualToSystemSpacingAfter(anchor LayoutXAxisAnchor, multiplier float64) LayoutConstraint
	AnchorWithOffsetTo(anchor LayoutXAxisAnchor) LayoutDimension
}

var _ LayoutXAxisAnchor = (*NSLayoutXAxisAnchor)(nil)

type NSLayoutXAxisAnchor struct {
	NSLayoutAnchor
}

func (x *NSLayoutXAxisAnchor) ConstraintEqualToSystemSpacingAfter(anchor LayoutXAxisAnchor, multiplier float64) LayoutConstraint {
	return MakeLayoutConstraint(C.LayoutXAxisAnchor_ConstraintEqualToSystemSpacingAfter(x.Ptr(), anchor.Ptr(), C.double(multiplier)))
}

func (x *NSLayoutXAxisAnchor) ConstraintGreaterThanOrEqualToSystemSpacingAfter(anchor LayoutXAxisAnchor, multiplier float64) LayoutConstraint {
	return MakeLayoutConstraint(C.LayoutXAxisAnchor_ConstraintGreaterThanOrEqualToSystemSpacingAfter(x.Ptr(), anchor.Ptr(), C.double(multiplier)))
}

func (x *NSLayoutXAxisAnchor) ConstraintLessThanOrEqualToSystemSpacingAfter(anchor LayoutXAxisAnchor, multiplier float64) LayoutConstraint {
	return MakeLayoutConstraint(C.LayoutXAxisAnchor_ConstraintLessThanOrEqualToSystemSpacingAfter(x.Ptr(), anchor.Ptr(), C.double(multiplier)))
}

func (x *NSLayoutXAxisAnchor) AnchorWithOffsetTo(anchor LayoutXAxisAnchor) LayoutDimension {
	return MakeLayoutDimension(C.LayoutXAxisAnchor_AnchorWithOffsetTo(x.Ptr(), anchor.Ptr()))
}

func MakeLayoutXAxisAnchor(ptr unsafe.Pointer) *NSLayoutXAxisAnchor {
	return &NSLayoutXAxisAnchor{
		NSLayoutAnchor: *MakeLayoutAnchor(ptr),
	}
}

type LayoutYAxisAnchor interface {
	LayoutAnchor
	// Returns a constraint that defines by how much the current anchor trails the specified anchor.
	ConstraintEqualToSystemSpacingBelow(anchor LayoutYAxisAnchor, multiplier float64) LayoutConstraint
	ConstraintGreaterThanOrEqualToSystemSpacingBelow(anchor LayoutYAxisAnchor, multiplier float64) LayoutConstraint
	ConstraintLessThanOrEqualToSystemSpacingBelow(anchor LayoutYAxisAnchor, multiplier float64) LayoutConstraint
	AnchorWithOffsetTo(anchor LayoutYAxisAnchor) LayoutDimension
}

var _ LayoutYAxisAnchor = (*NSLayoutYAxisAnchor)(nil)

type NSLayoutYAxisAnchor struct {
	NSLayoutAnchor
}

func (y *NSLayoutYAxisAnchor) ConstraintEqualToSystemSpacingBelow(anchor LayoutYAxisAnchor, multiplier float64) LayoutConstraint {
	return MakeLayoutConstraint(C.LayoutYAxisAnchor_ConstraintEqualToSystemSpacingBelow(y.Ptr(), anchor.Ptr(), C.double(multiplier)))
}

func (y *NSLayoutYAxisAnchor) ConstraintGreaterThanOrEqualToSystemSpacingBelow(anchor LayoutYAxisAnchor, multiplier float64) LayoutConstraint {
	return MakeLayoutConstraint(C.LayoutYAxisAnchor_ConstraintGreaterThanOrEqualToSystemSpacingBelow(y.Ptr(), anchor.Ptr(), C.double(multiplier)))
}

func (y *NSLayoutYAxisAnchor) ConstraintLessThanOrEqualToSystemSpacingBelow(anchor LayoutYAxisAnchor, multiplier float64) LayoutConstraint {
	return MakeLayoutConstraint(C.LayoutYAxisAnchor_ConstraintLessThanOrEqualToSystemSpacingBelow(y.Ptr(), anchor.Ptr(), C.double(multiplier)))
}

func (y *NSLayoutYAxisAnchor) AnchorWithOffsetTo(anchor LayoutYAxisAnchor) LayoutDimension {
	return MakeLayoutDimension(C.LayoutYAxisAnchor_AnchorWithOffsetTo(y.Ptr(), anchor.Ptr()))
}

func MakeLayoutYAxisAnchor(ptr unsafe.Pointer) *NSLayoutYAxisAnchor {
	return &NSLayoutYAxisAnchor{
		NSLayoutAnchor: *MakeLayoutAnchor(ptr),
	}
}
