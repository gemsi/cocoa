package widget

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include "text_view.h"
import "C"

import (
	"github.com/hsiafan/cocoa/foundation"
	"unsafe"
)

// A view that draws text and handles user interactions with that text.
type TextView interface {
	Text

	// TextContainer return the receiver’s text container
	TextContainer() TextContainer
	// SetTextContainer set the receiver’s text container
	SetTextContainer(value TextContainer)
}

var _ TextView = (*NSTextView)(nil)

type NSTextView struct {
	NSText
}

// Make create a View from native pointer
func MakeTextView(ptr unsafe.Pointer) *NSTextView {
	return &NSTextView{*MakeText(ptr)}
}

// New create new TextView
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

func (t *NSTextView) TextContainer() TextContainer {
	return MakeTextContainer(C.TextView_TextContainer(t.Ptr()))
}

func (t *NSTextView) SetTextContainer(value TextContainer) {
	C.TextView_SetTextContainer(t.Ptr(), value.Ptr())
}
