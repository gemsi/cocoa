package foundation

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #import "notification.h"
import "C"
import (
	"unsafe"
)

// Notification wrap cocoa NSNotification
type Notification interface {
	Object
	// Name return the notification message name
	Name() string
	// Object return the sender of this notification
	Object() Object
}

var _ Notification = (*NSNotification)(nil)

// MakeObject create new Notification, from native pointer of NSNotification, and the sender object
func MakeNotification(ptr unsafe.Pointer) *NSNotification {
	return &NSNotification{
		NSObject: *MakeObject(ptr),
	}
}

type NSNotification struct {
	NSObject
}

func (n *NSNotification) Name() string {
	cstr := C.Notification_Name(n.Ptr())
	return C.GoString(cstr)
}

func (n *NSNotification) Object() Object {
	return MakeObject(unsafe.Pointer(C.Notification_Object(n.ptr)))
}
