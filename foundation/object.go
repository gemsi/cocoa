package foundation

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #import "object.h"
import "C"
import (
	"unsafe"
)

// interface for all NSObject type
type Object interface {
	// return the delegate pointer of this object
	Ptr() unsafe.Pointer
	Dealloc()
}

var _ Object = (*NSObject)(nil)

// wrapper for NSObject
type NSObject struct {
	ptr unsafe.Pointer
}

func (o *NSObject) Dealloc() {

}

func (o *NSObject) Ptr() unsafe.Pointer {
	return o.ptr
}

func MakeObject(ptr unsafe.Pointer) *NSObject {
	return &NSObject{ptr}
}
