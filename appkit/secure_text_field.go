package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include "secure_text_field.h"
import "C"

import (
	"github.com/hsiafan/cocoa/foundation"
	"unsafe"
)

// SecureTextField is a text field that hides the typed text
type SecureTextField interface {
	TextField
}

var _ SecureTextField = (*NSSecureTextField)(nil)

type NSSecureTextField struct {
	NSTextField
}

// MakeSecureTextField create a SecureTextField from native pointer
func MakeSecureTextField(ptr unsafe.Pointer) *NSSecureTextField {
	if ptr == nil {
		return nil
	}
	return &NSSecureTextField{
		NSTextField: *MakeTextField(ptr),
	}
}

// NewSecureTextField create new SecureTextField
func NewSecureTextField(frame foundation.Rect) SecureTextField {
	id := resources.NextId()
	ptr := C.SecureTextField_initWithFrame(C.long(id), toNSRect(frame))
	v := &NSSecureTextField{
		NSTextField: *MakeTextField(ptr),
	}
	resources.Store(id, v)
	foundation.AddDeallocHook(v, func() {
		resources.Delete(id)
	})
	return v
}
