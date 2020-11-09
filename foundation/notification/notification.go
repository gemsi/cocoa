package notification

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #import "notification.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation/object"
	"unsafe"
)

// Notification wrap cocoa NSNotification
type Notification interface {
	object.Object
	// Name return the notification message name
	Name() string
	// Object return the sender of this notification
	Object() object.Object
}

var _ Notification = (*NSNotification)(nil)

// Make create new Notification, from native pointer of NSNotification, and the sender object
func Make(ptr unsafe.Pointer, obj object.Object) *NSNotification {
	return &NSNotification{
		NSObject: *object.Make(ptr),
		object:   obj,
	}
}

type NSNotification struct {
	object.NSObject
	name   *string
	object object.Object
}

func (n *NSNotification) Name() string {
	if n.name == nil {
		cstr := C.Notification_Name(n.Ptr())
		name := C.GoString(cstr)
		n.name = &name
	}
	return *n.name
}

func (n *NSNotification) Object() object.Object {
	return n.object
}
