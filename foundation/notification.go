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
func MakeNotification(ptr unsafe.Pointer, obj Object) *NSNotification {
	return &NSNotification{
		NSObject: *MakeObject(ptr),
		object:   obj,
	}
}

type NSNotification struct {
	NSObject
	name   *string
	object Object
}

func (n *NSNotification) Name() string {
	if n.name == nil {
		cstr := C.Notification_Name(n.Ptr())
		name := C.GoString(cstr)
		n.name = &name
	}
	return *n.name
}

func (n *NSNotification) Object() Object {
	return n.object
}
