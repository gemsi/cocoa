package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #import "layout_constraint.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"unsafe"
)

// LayoutConstraint is the relationship between two user interface objects that must be satisfied by the constraint-based layout system.
type LayoutConstraint interface {
	foundation.Object
	IsActive() bool
	SetActive(active bool)
	// active this LayoutConstraint
	Active()
	FirstItem() foundation.Object
	FirstAttribute() LayoutAttribute
	Relation() LayoutRelation
	SecondItem() foundation.Object
	SecondAttribute() LayoutAttribute
	Multiplier() float64
	Constant() float64
	SetConstant(constant float64)
	// The first anchor that defines the constraint.
	FirstAnchor() LayoutAnchor
	// The second anchor that defines the constraint.
	SecondAnchor() LayoutAnchor
	Priority() LayoutPriority
	SetPriority(priority LayoutPriority)
	Identifier() string
	// whether the constraint should be archived by its owning view
	ShouldBeArchived() bool
	// set whether the constraint should be archived by its owning view
	SetShouldBeArchived(value bool)
}

var _ LayoutConstraint = (*NSLayoutConstraint)(nil)

type NSLayoutConstraint struct {
	foundation.NSObject
}

func (c *NSLayoutConstraint) Active() {
	c.SetActive(true)
}

func (c *NSLayoutConstraint) IsActive() bool {
	return bool(C.LayoutConstraint_IsActive(c.Ptr()))
}

func (c *NSLayoutConstraint) SetActive(active bool) {
	C.LayoutConstraint_SetActive(c.Ptr(), C.bool(active))
}

func (c *NSLayoutConstraint) FirstItem() foundation.Object {
	return foundation.MakeObject(C.LayoutConstraint_FirstItem(c.Ptr()))
}

func (c *NSLayoutConstraint) FirstAttribute() LayoutAttribute {
	return LayoutAttribute(C.LayoutConstraint_FirstAttribute(c.Ptr()))
}

func (c *NSLayoutConstraint) Relation() LayoutRelation {
	return LayoutRelation(C.LayoutConstraint_Relation(c.Ptr()))
}

func (c *NSLayoutConstraint) SecondItem() foundation.Object {
	return foundation.MakeObject(C.LayoutConstraint_SecondItem(c.Ptr()))
}

func (c *NSLayoutConstraint) SecondAttribute() LayoutAttribute {
	return LayoutAttribute(C.LayoutConstraint_SecondAttribute(c.Ptr()))
}

func (c *NSLayoutConstraint) Multiplier() float64 {
	return float64(C.LayoutConstraint_Multiplier(c.Ptr()))
}

func (c *NSLayoutConstraint) Constant() float64 {
	return float64(C.LayoutConstraint_Constant(c.Ptr()))
}

func (c *NSLayoutConstraint) SetConstant(constant float64) {
	C.LayoutConstraint_SetConstant(c.Ptr(), C.double(constant))
}

func (c *NSLayoutConstraint) FirstAnchor() LayoutAnchor {
	panic("implement me")
}

func (c *NSLayoutConstraint) SecondAnchor() LayoutAnchor {
	panic("implement me")
}

func (c *NSLayoutConstraint) Priority() LayoutPriority {
	return LayoutPriority(C.LayoutConstraint_Priority(c.Ptr()))
}

func (c *NSLayoutConstraint) SetPriority(priority LayoutPriority) {
	C.LayoutConstraint_SetPriority(c.Ptr(), C.float(priority))
}

func (c *NSLayoutConstraint) Identifier() string {
	cstr := C.LayoutConstraint_Identifier(c.Ptr())
	return C.GoString(cstr)
}

func (c *NSLayoutConstraint) ShouldBeArchived() bool {
	return bool(C.LayoutConstraint_ShouldBeArchived(c.Ptr()))
}

func (c *NSLayoutConstraint) SetShouldBeArchived(value bool) {
	C.LayoutConstraint_SetShouldBeArchived(c.Ptr(), C.bool(value))
}

// NewLayoutConstraint creates a constraint that defines the relationship between the specified attributes of the given views.
func NewLayoutConstraint(view1 foundation.Object, attr1 LayoutAttribute, relation LayoutRelation,
	view2 foundation.Object, attr2 LayoutAttribute, multiplier float64, constant float64) LayoutConstraint {
	ptr := C.LayoutConstraint_New(view1.Ptr(), C.int(attr1), C.int(relation),
		view2.Ptr(), C.int(attr2), C.double(multiplier), C.double(constant))
	return MakeLayoutConstraint(ptr)
}

func MakeLayoutConstraint(ptr unsafe.Pointer) LayoutConstraint {
	return &NSLayoutConstraint{
		NSObject: *foundation.MakeObject(ptr),
	}
}

// ActivateConstraints convenience method that activates each constraint in the contained array, in the same manner as setting active=YES.
// This is often more efficient than activating each constraint individually.
func ActivateConstraints(constraints ...LayoutConstraint) {
	pointers := make([]unsafe.Pointer, len(constraints))
	for i, constraint := range constraints {
		pointers[i] = constraint.Ptr()
	}
	var array = C.Array{
		data: unsafe.Pointer(&pointers[0]),
		len:  C.int(len(pointers)),
	}
	C.LayoutConstraint_ActivateConstraints(array)
}

// DeactivateConstraints is convenience method that deactivates each constraint in the contained array, in the same manner as setting active=NO.
// This is often more efficient than deactivating each constraint individually. */
func DeactivateConstraints(constraints ...LayoutConstraint) {
	pointers := make([]unsafe.Pointer, len(constraints))
	for i, constraint := range constraints {
		pointers[i] = constraint.Ptr()
	}
	var array = C.Array{
		data: unsafe.Pointer(&pointers[0]),
		len:  C.int(len(pointers)),
	}
	C.LayoutConstraint_DeactivateConstraints(array)
}
