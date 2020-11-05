package control

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include "control.h"
import "C"
import (
	"unsafe"

	"github.com/hsiafan/cocoa/view"
)

// Control wrap a cocoa NSControl
type Control interface {
	view.View
	// IsEnabled return if this control is enabled
	IsEnabled() bool
	// SetStringValue set the string value
	SetStringValue(value string)
	// StringValue return the string value
	StringValue() string
}

type NSControl struct {
	view.NSView
}

// Make create a Control from native pointer
func Make(ptr unsafe.Pointer) *NSControl {
	return &NSControl{*view.Make(ptr)}
}

func (c *NSControl) IsEnabled() bool {
	return false
}

func (c *NSControl) SetStringValue(value string) {
	cstr := C.CString(value)
	defer C.free(unsafe.Pointer(cstr))
	C.Control_SetStringValue(c.Ptr(), cstr)
}

func (c *NSControl) StringValue() string {
	cstr := C.Control_StringValue(c.Ptr())
	return C.GoString(cstr)
}
