package notification

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #import "notification.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"unsafe"
)

// Notification wrap cocoa NSNotification
type Notification interface {
	foundation.Object
	// Name return the notification message name
	Name() string
	// Object return the sender of this notification
	Object() foundation.Object
}

var _ Notification = (*NSNotification)(nil)

// Make create new Notification, from native pointer of NSNotification, and the sender object
func Make(ptr unsafe.Pointer, object foundation.Object) *NSNotification {
	return &NSNotification{
		NSObject: *foundation.MakeObject(ptr),
		object:   object,
	}
}

type NSNotification struct {
	foundation.NSObject
	name   *string
	object foundation.Object
}

func (n *NSNotification) Name() string {
	if n.name == nil {
		cstr := C.Notification_Name(n.Ptr())
		name := C.GoString(cstr)
		n.name = &name
	}
	return *n.name
}

func (n *NSNotification) Object() foundation.Object {
	return n.object
}
