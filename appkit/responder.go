package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include "responder.h"
import "C"

import (
	"github.com/hsiafan/cocoa/foundation"
	"unsafe"
)

// Responder is the basis of event and command processing in AppKit
type Responder interface {
	foundation.Object

	// AcceptsFirstResponder return whether the responder accepts first responder status
	AcceptsFirstResponder() bool
	// NextResponder return the next responder after this one, or nil if it has none
	NextResponder() Responder
	// SetNextResponder set the next responder after this one, or nil if it has none
	SetNextResponder(nextResponder Responder)
}

var _ Responder = (*NSResponder)(nil)

type NSResponder struct {
	foundation.NSObject
}

// MakeResponder create a Responder from native pointer
func MakeResponder(ptr unsafe.Pointer) *NSResponder {
	return &NSResponder{
		NSObject: *foundation.MakeObject(ptr),
	}
}

func (r *NSResponder) AcceptsFirstResponder() bool {
	return bool(C.Responder_AcceptsFirstResponder(r.Ptr()))
}

func (r *NSResponder) NextResponder() Responder {
	return MakeResponder(C.Responder_NextResponder(r.Ptr()))
}

func (r *NSResponder) SetNextResponder(nextResponder Responder) {
	C.Responder_SetNextResponder(r.Ptr(), nextResponder.Ptr())
}
