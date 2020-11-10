package widget

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include "text_view.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
)

// TextView is a view that draws text and handles user interactions with that text.
type TextView interface {
	Text
}

var _ TextView = (*NSTextView)(nil)

type NSTextView struct {
	NSText
}

// NewTextView create new TextView
func NewTextView(frame foundation.Rect) TextView {
	id := resources.NextId()
	ptr := C.TextView_New(C.long(id), toNSRect(frame))

	v := &NSTextView{
		NSText: *MakeText(ptr),
	}

	resources.Store(id, v)

	foundation.AddDeallocHook(v, func() {
		resources.Delete(id)
	})

	return v
}
