package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include "layout_constraint.h"
import "C"

import (
	"github.com/hsiafan/cocoa/foundation"
	"unsafe"
)

// LayoutConstraint is the relationship between two user interface objects that must be satisfied by the constraint-based layout system
type LayoutConstraint interface {
	foundation.Object

	// IsActive return the active state of the constraint
	IsActive() bool
	// SetActive set the active state of the constraint
	SetActive(active bool)
	// FirstItem return the first object participating in the constraint
	FirstItem() foundation.Object
	// FirstAttribute return the attribute of the first object participating in the constraint
	FirstAttribute() LayoutAttribute
	// Relation return the relation between the two attributes in the constraint
	Relation() LayoutRelation
	// SecondItem return the second object participating in the constraint
	SecondItem() foundation.Object
	// SecondAttribute return the attribute of the second object participating in the constraint
	SecondAttribute() LayoutAttribute
	// Multiplier return the multiplier applied to the second attribute participating in the constraint
	Multiplier() float64
	// Constant return the constant added to the multiplied second attribute participating in the constraint
	Constant() float64
	// FirstAnchor return the first anchor that defines the constraint
	FirstAnchor() LayoutAnchor
	// SecondAnchor return the second anchor that defines the constraint
	SecondAnchor() LayoutAnchor
	// Priority return the priority of the constraint
	Priority() LayoutPriority
	// SetPriority set the priority of the constraint
	SetPriority(priority LayoutPriority)
	// Identifier return the name that identifies the constraint
	Identifier() string
	// SetIdentifier set the name that identifies the constraint
	SetIdentifier(identifier string)
	// ShouldBeArchived return whether the constraint should be archived by its owning view
	ShouldBeArchived() bool
	// SetShouldBeArchived set whether the constraint should be archived by its owning view
	SetShouldBeArchived(shouldBeArchived bool)
}

var _ LayoutConstraint = (*NSLayoutConstraint)(nil)

type NSLayoutConstraint struct {
	foundation.NSObject
}

// MakeLayoutConstraint create a LayoutConstraint from native pointer
func MakeLayoutConstraint(ptr unsafe.Pointer) *NSLayoutConstraint {
	if ptr == nil {
		return nil
	}
	return &NSLayoutConstraint{
		NSObject: *foundation.MakeObject(ptr),
	}
}

func (l *NSLayoutConstraint) IsActive() bool {
	return bool(C.LayoutConstraint_IsActive(l.Ptr()))
}

func (l *NSLayoutConstraint) SetActive(active bool) {
	C.LayoutConstraint_SetActive(l.Ptr(), C.bool(active))
}

func (l *NSLayoutConstraint) FirstItem() foundation.Object {
	return foundation.MakeObject(C.LayoutConstraint_FirstItem(l.Ptr()))
}

func (l *NSLayoutConstraint) FirstAttribute() LayoutAttribute {
	return LayoutAttribute(C.LayoutConstraint_FirstAttribute(l.Ptr()))
}

func (l *NSLayoutConstraint) Relation() LayoutRelation {
	return LayoutRelation(C.LayoutConstraint_Relation(l.Ptr()))
}

func (l *NSLayoutConstraint) SecondItem() foundation.Object {
	return foundation.MakeObject(C.LayoutConstraint_SecondItem(l.Ptr()))
}

func (l *NSLayoutConstraint) SecondAttribute() LayoutAttribute {
	return LayoutAttribute(C.LayoutConstraint_SecondAttribute(l.Ptr()))
}

func (l *NSLayoutConstraint) Multiplier() float64 {
	return float64(C.LayoutConstraint_Multiplier(l.Ptr()))
}

func (l *NSLayoutConstraint) Constant() float64 {
	return float64(C.LayoutConstraint_Constant(l.Ptr()))
}

func (l *NSLayoutConstraint) FirstAnchor() LayoutAnchor {
	return MakeLayoutAnchor(C.LayoutConstraint_FirstAnchor(l.Ptr()))
}

func (l *NSLayoutConstraint) SecondAnchor() LayoutAnchor {
	return MakeLayoutAnchor(C.LayoutConstraint_SecondAnchor(l.Ptr()))
}

func (l *NSLayoutConstraint) Priority() LayoutPriority {
	return LayoutPriority(C.LayoutConstraint_Priority(l.Ptr()))
}

func (l *NSLayoutConstraint) SetPriority(priority LayoutPriority) {
	C.LayoutConstraint_SetPriority(l.Ptr(), C.float(priority))
}

func (l *NSLayoutConstraint) Identifier() string {
	return C.GoString(C.LayoutConstraint_Identifier(l.Ptr()))
}

func (l *NSLayoutConstraint) SetIdentifier(identifier string) {
	cIdentifier := C.CString(identifier)
	defer C.free(unsafe.Pointer(cIdentifier))
	C.LayoutConstraint_SetIdentifier(l.Ptr(), cIdentifier)
}

func (l *NSLayoutConstraint) ShouldBeArchived() bool {
	return bool(C.LayoutConstraint_ShouldBeArchived(l.Ptr()))
}

func (l *NSLayoutConstraint) SetShouldBeArchived(shouldBeArchived bool) {
	C.LayoutConstraint_SetShouldBeArchived(l.Ptr(), C.bool(shouldBeArchived))
}

func ConstraintWithItem(view1 foundation.Object, attr1 LayoutAttribute, relation LayoutRelation, view2 foundation.Object, attr2 LayoutAttribute, multiplier float64, c float64) {
	C.LayoutConstraint_ConstraintWithItem(toPointer(view1), C.long(attr1), C.long(relation), toPointer(view2), C.long(attr2), C.double(multiplier), C.double(c))
}

func ActivateConstraints(constraints []LayoutConstraint) {
	cConstraintsData := make([]unsafe.Pointer, len(constraints))
	for idx, v := range constraints {
		cConstraintsData[idx] = v.Ptr()
	}
	cConstraints := C.Array{data: unsafe.Pointer(&cConstraintsData[0]), len: C.int(len(constraints))}
	C.LayoutConstraint_ActivateConstraints(cConstraints)
}

func DeactivateConstraints(constraints []LayoutConstraint) {
	cConstraintsData := make([]unsafe.Pointer, len(constraints))
	for idx, v := range constraints {
		cConstraintsData[idx] = v.Ptr()
	}
	cConstraints := C.Array{data: unsafe.Pointer(&cConstraintsData[0]), len: C.int(len(constraints))}
	C.LayoutConstraint_DeactivateConstraints(cConstraints)
}
