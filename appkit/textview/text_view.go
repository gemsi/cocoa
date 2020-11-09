package textview

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include "text_view.h"
import "C"
import (
	"github.com/hsiafan/cocoa/appkit/text"
	"github.com/hsiafan/cocoa/foundation/geometry"
	"github.com/hsiafan/cocoa/foundation/notification"
	"github.com/hsiafan/cocoa/foundation/object"
	"github.com/hsiafan/cocoa/internal"
	"unsafe"
)

// TextView is a view that draws text and handles user interactions with that text.
type TextView interface {
	text.Text
}

var _ TextView = (*NSTextView)(nil)

type NSTextView struct {
	text.NSText
	textDidChange     func(notification.Notification)
	textDidEndEditing func(notification.Notification)
	onEnterKey        func(sender object.Object)
}

var resources = internal.NewResourceRegistry()

// New create new TextView
func New(frame geometry.Rect) TextView {
	id := resources.NextId()
	ptr := C.TextView_New(C.long(id), toNSRect(frame))

	v := &NSTextView{
		NSText: *text.Make(ptr),
	}

	resources.Store(id, v)

	object.AddDeallocHook(v, func() {
		resources.Delete(id)
	})

	return v
}

func toNSRect(rect geometry.Rect) C.NSRect {
	return *(*C.NSRect)(unsafe.Pointer(&rect))
}
func toRect(nsRect C.NSRect) geometry.Rect {
	return *(*geometry.Rect)(unsafe.Pointer(&nsRect))
}
