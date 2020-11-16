package widget

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include "control.h"
import "C"

import (
	"github.com/hsiafan/cocoa/foundation"
	"unsafe"
)

// A definition of the fundamental behavior for controls, which are specialized views that notify your app of relevant events by using the target-action design pattern.
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
	IntegerValue() int64
	// SetIntegerValue set the value of the receiver’s cell as an NSInteger value.
	SetIntegerValue(integerValue int64)
	// StringValue return the value of the receiver’s cell as an NSString object
	StringValue() string
	// SetStringValue set the value of the receiver’s cell as an NSString object
	SetStringValue(stringValue string)
}

var _ Control = (*NSControl)(nil)

type NSControl struct {
	NSView
}

// Make create a Control from native pointer
func MakeControl(ptr unsafe.Pointer) *NSControl {
	return &NSControl{*MakeView(ptr)}
}

// New create new Control
func NewControl(frame foundation.Rect) Control {
	id := resources.NextId()
	ptr := C.Control_New(C.long(id), toNSRect(frame))

	v := &NSControl{
		NSView: *MakeView(ptr),
	}

	resources.Store(id, v)

	foundation.AddDeallocHook(v, func() {
		resources.Delete(id)
	})

	return v
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
	C.Control_SetIntValue(c.Ptr(), C.int(intValue))
}

func (c *NSControl) IntegerValue() int64 {
	return int64(C.Control_IntegerValue(c.Ptr()))
}

func (c *NSControl) SetIntegerValue(integerValue int64) {
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
