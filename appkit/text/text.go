package text

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include "text.h"
import "C"
import (
	"github.com/hsiafan/cocoa/appkit/view"
	"unsafe"
)

// Wrap NSText. The most general programmatic interface for objects that manage text.
type Text interface {
	view.View
	// StringValue return the characters of the text.
	StringValue() string
	// Editable return  whether the receiver allows the user to edit its text.
	SetEditable(value bool)
	// SetSelectable set whether the receiver allows the user to select its text.
	SetSelectable(value bool)
	// SetFieldEditor set whether the receiver interprets Tab, Shift-Tab, and Return (Enter) as cues to end editing and possibly to change the first responder.
	SetFieldEditor(value bool)
	// SetRichText set whether the receiver allows the user to apply attributes to specific ranges of the text.
	SetRichText(value bool)
}

var _ Text = (*NSText)(nil)

// NSText
type NSText struct {
	view.NSView
}

// Make create a NSText from native pointer
func Make(ptr unsafe.Pointer) *NSText {
	return &NSText{
		NSView: *view.Make(ptr),
	}
}

func (t *NSText) StringValue() string {
	str := C.Text_String(t.Ptr())
	return C.GoString(str)
}

func (t *NSText) SetEditable(value bool) {
	C.Text_SetEditable(t.Ptr(), C.bool(value))
}

func (t *NSText) SetSelectable(value bool) {
	C.Text_SetSelectable(t.Ptr(), C.bool(value))
}

func (t *NSText) SetFieldEditor(value bool) {
	C.Text_SetFieldEditor(t.Ptr(), C.bool(value))
}

func (t *NSText) SetRichText(value bool) {
	C.Text_SetRichText(t.Ptr(), C.bool(value))
}
