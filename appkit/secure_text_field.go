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
func (s *NSSecureTextField) setDelegate() {
	id := resources.NextId()
	C.SecureTextField_RegisterDelegate(s.Ptr(), C.long(id))
	resources.Store(id, s)
	foundation.AddDeallocHook(s, func() {
		resources.Delete(id)
	})
}

func NewSecureTextField(frame foundation.Rect) SecureTextField {
	instance := MakeSecureTextField(C.SecureTextField_NewSecureTextField(toNSRect(frame)))
	instance.setDelegate()
	return instance
}
