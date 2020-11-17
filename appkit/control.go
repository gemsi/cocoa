package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include "control.h"
import "C"

import (
	"unsafe"
)

// Control is A definition of the fundamental behavior for controls, which are specialized views that notify your app of relevant events by using the target-action design pattern.
type Control interface {
	View

	// IsEnabled return whether the receiver reacts to mouse events
	IsEnabled() bool
	// SetEnabled set whether the receiver reacts to mouse events
	SetEnabled(enabled bool)
	// DoubleValue return the value of the receiver’s cell as a double-precision floating-point number
	DoubleValue() float64
	// SetDoubleValue set the value of the receiver’s cell as a double-precision floating-point number
	SetDoubleValue(doubleValue float64)
	// FloatValue return the value of the receiver’s cell as a single-precision floating-point number
	FloatValue() float32
	// SetFloatValue set the value of the receiver’s cell as a single-precision floating-point number
	SetFloatValue(floatValue float32)
	// IntValue return the value of the receiver’s cell as an integer
	IntValue() int
	// SetIntValue set the value of the receiver’s cell as an integer
	SetIntValue(intValue int)
	// IntegerValue return the value of the receiver’s cell as an NSInteger value.
	IntegerValue() int
	// SetIntegerValue set the value of the receiver’s cell as an NSInteger value.
	SetIntegerValue(integerValue int)
	// StringValue return the value of the receiver’s cell as an NSString object
	StringValue() string
	// SetStringValue set the value of the receiver’s cell as an NSString object
	SetStringValue(stringValue string)
	// Cell return the receiver’s cell object
	Cell() Cell
	// SetCell set the receiver’s cell object
	SetCell(cell Cell)
	// SizeToFit resizes the receiver’s frame so that it is the minimum size needed to contain its cell
	SizeToFit()
}

var _ Control = (*NSControl)(nil)

type NSControl struct {
	NSView
}

// MakeControl create a Control from native pointer
func MakeControl(ptr unsafe.Pointer) *NSControl {
	return &NSControl{
		NSView: *MakeView(ptr),
	}
}

func (c *NSControl) IsEnabled() bool {
	return bool(C.Control_IsEnabled(c.Ptr()))
}

func (c *NSControl) SetEnabled(enabled bool) {
	C.Control_SetEnabled(c.Ptr(), C.bool(enabled))
}

func (c *NSControl) DoubleValue() float64 {
	return float64(C.Control_DoubleValue(c.Ptr()))
}

func (c *NSControl) SetDoubleValue(doubleValue float64) {
	C.Control_SetDoubleValue(c.Ptr(), C.double(doubleValue))
}

func (c *NSControl) FloatValue() float32 {
	return float32(C.Control_FloatValue(c.Ptr()))
}

func (c *NSControl) SetFloatValue(floatValue float32) {
	C.Control_SetFloatValue(c.Ptr(), C.float(floatValue))
}

func (c *NSControl) IntValue() int {
	return int(C.Control_IntValue(c.Ptr()))
}

func (c *NSControl) SetIntValue(intValue int) {
	C.Control_SetIntValue(c.Ptr(), C.long(intValue))
}

func (c *NSControl) IntegerValue() int {
	return int(C.Control_IntegerValue(c.Ptr()))
}

func (c *NSControl) SetIntegerValue(integerValue int) {
	C.Control_SetIntegerValue(c.Ptr(), C.long(integerValue))
}

func (c *NSControl) StringValue() string {
	return C.GoString(C.Control_StringValue(c.Ptr()))
}

func (c *NSControl) SetStringValue(stringValue string) {
	c_stringValue := C.CString(stringValue)
	defer C.free(unsafe.Pointer(c_stringValue))
	C.Control_SetStringValue(c.Ptr(), c_stringValue)
}

func (c *NSControl) Cell() Cell {
	return MakeCell(C.Control_Cell(c.Ptr()))
}

func (c *NSControl) SetCell(cell Cell) {
	C.Control_SetCell(c.Ptr(), cell.Ptr())
}

func (c *NSControl) SizeToFit() {
	C.Control_SizeToFit(c.Ptr())
}
