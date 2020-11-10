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

	// Enabled return whether the receiver reacts to mouse events
	Enabled() bool
	// SetEnabled set whether the receiver reacts to mouse events
	SetEnabled(value bool)
	// DoubleValue return the value of the receiver’s cell as a double-precision floating-point number
	DoubleValue() float64
	// SetDoubleValue set the value of the receiver’s cell as a double-precision floating-point number
	SetDoubleValue(value float64)
	// FloatValue return the value of the receiver’s cell as a single-precision floating-point number
	FloatValue() float32
	// SetFloatValue set the value of the receiver’s cell as a single-precision floating-point number
	SetFloatValue(value float32)
	// IntValue return the value of the receiver’s cell as an integer
	IntValue() int
	// SetIntValue set the value of the receiver’s cell as an integer
	SetIntValue(value int)
	// IntegerValue return the value of the receiver’s cell as an NSInteger value.
	IntegerValue() int64
	// SetIntegerValue set the value of the receiver’s cell as an NSInteger value.
	SetIntegerValue(value int64)
	// StringValue return the value of the receiver’s cell as an NSString object
	StringValue() string
	// SetStringValue set the value of the receiver’s cell as an NSString object
	SetStringValue(value string)
}

var _ Control = (*NSControl)(nil)

type NSControl struct {
	NSView
}

// Make create a View from native pointer
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

func (c *NSControl) Enabled() bool {
	return bool(C.Control_Enabled(c.Ptr()))
}

func (c *NSControl) SetEnabled(value bool) {
	C.Control_SetEnabled(c.Ptr(), C.bool(value))
}

func (c *NSControl) DoubleValue() float64 {
	return float64(C.Control_DoubleValue(c.Ptr()))
}

func (c *NSControl) SetDoubleValue(value float64) {
	C.Control_SetDoubleValue(c.Ptr(), C.double(value))
}

func (c *NSControl) FloatValue() float32 {
	return float32(C.Control_FloatValue(c.Ptr()))
}

func (c *NSControl) SetFloatValue(value float32) {
	C.Control_SetFloatValue(c.Ptr(), C.float(value))
}

func (c *NSControl) IntValue() int {
	return int(C.Control_IntValue(c.Ptr()))
}

func (c *NSControl) SetIntValue(value int) {
	C.Control_SetIntValue(c.Ptr(), C.int(value))
}

func (c *NSControl) IntegerValue() int64 {
	return int64(C.Control_IntegerValue(c.Ptr()))
}

func (c *NSControl) SetIntegerValue(value int64) {
	C.Control_SetIntegerValue(c.Ptr(), C.long(value))
}

func (c *NSControl) StringValue() string {
	return C.GoString(C.Control_StringValue(c.Ptr()))
}

func (c *NSControl) SetStringValue(value string) {
	cstr := C.CString(value)
	defer C.free(unsafe.Pointer(cstr))
	C.Control_SetStringValue(c.Ptr(), cstr)
}
