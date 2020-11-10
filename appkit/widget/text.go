package widget

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include "text.h"
import "C"

import (
	"github.com/hsiafan/cocoa/foundation"
	"unsafe"
)

// The most general programmatic interface for objects that manage text.
type Text interface {
	View

	// String return the characters of the receiver’s text
	String() string
	// SetString set the characters of the receiver’s text
	SetString(value string)
	// Editable return whether the receiver allows the user to edit its text
	Editable() bool
	// SetEditable set whether the receiver allows the user to edit its text
	SetEditable(value bool)
	// Selectable return whether the receiver allows the user to select its text
	Selectable() bool
	// SetSelectable set whether the receiver allows the user to select its text
	SetSelectable(value bool)
	// FieldEditor return whether the receiver interprets Tab, Shift-Tab, and Return (Enter) as cues to end editing and possibly to change the first responder
	FieldEditor() bool
	// SetFieldEditor set whether the receiver interprets Tab, Shift-Tab, and Return (Enter) as cues to end editing and possibly to change the first responder
	SetFieldEditor(value bool)
	// RichText return whether the receiver allows the user to apply attributes to specific ranges of the text
	RichText() bool
	// SetRichText set whether the receiver allows the user to apply attributes to specific ranges of the text
	SetRichText(value bool)
	// MinSize return the receiver’s minimum size
	MinSize() foundation.Size
	// SetMinSize set the receiver’s minimum size
	SetMinSize(value foundation.Size)
	// MaxSize return the receiver’s maximum size
	MaxSize() foundation.Size
	// SetMaxSize set the receiver’s maximum size
	SetMaxSize(value foundation.Size)
	// VerticallyResizable return whether the receiver changes its height to fit the height of its text.
	VerticallyResizable() bool
	// SetVerticallyResizable set whether the receiver changes its height to fit the height of its text.
	SetVerticallyResizable(value bool)
	// HorizontallyResizable return whether the receiver changes its width to fit the width of its text
	HorizontallyResizable() bool
	// SetHorizontallyResizable set whether the receiver changes its width to fit the width of its text
	SetHorizontallyResizable(value bool)
}

var _ Text = (*NSText)(nil)

type NSText struct {
	NSView
}

// Make create a View from native pointer
func MakeText(ptr unsafe.Pointer) *NSText {
	return &NSText{*MakeView(ptr)}
}

// New create new Text
func NewText(frame foundation.Rect) Text {
	id := resources.NextId()
	ptr := C.Text_New(C.long(id), toNSRect(frame))

	v := &NSText{
		NSView: *MakeView(ptr),
	}

	resources.Store(id, v)

	foundation.AddDeallocHook(v, func() {
		resources.Delete(id)
	})

	return v
}

func (t *NSText) String() string {
	return C.GoString(C.Text_String(t.Ptr()))
}

func (t *NSText) SetString(value string) {
	cstr := C.CString(value)
	defer C.free(unsafe.Pointer(cstr))
	C.Text_SetString(t.Ptr(), cstr)
}

func (t *NSText) Editable() bool {
	return bool(C.Text_Editable(t.Ptr()))
}

func (t *NSText) SetEditable(value bool) {
	C.Text_SetEditable(t.Ptr(), C.bool(value))
}

func (t *NSText) Selectable() bool {
	return bool(C.Text_Selectable(t.Ptr()))
}

func (t *NSText) SetSelectable(value bool) {
	C.Text_SetSelectable(t.Ptr(), C.bool(value))
}

func (t *NSText) FieldEditor() bool {
	return bool(C.Text_FieldEditor(t.Ptr()))
}

func (t *NSText) SetFieldEditor(value bool) {
	C.Text_SetFieldEditor(t.Ptr(), C.bool(value))
}

func (t *NSText) RichText() bool {
	return bool(C.Text_RichText(t.Ptr()))
}

func (t *NSText) SetRichText(value bool) {
	C.Text_SetRichText(t.Ptr(), C.bool(value))
}

func (t *NSText) MinSize() foundation.Size {
	return toSize(C.Text_MinSize(t.Ptr()))
}

func (t *NSText) SetMinSize(value foundation.Size) {
	C.Text_SetMinSize(t.Ptr(), toNSSize(value))
}

func (t *NSText) MaxSize() foundation.Size {
	return toSize(C.Text_MaxSize(t.Ptr()))
}

func (t *NSText) SetMaxSize(value foundation.Size) {
	C.Text_SetMaxSize(t.Ptr(), toNSSize(value))
}

func (t *NSText) VerticallyResizable() bool {
	return bool(C.Text_VerticallyResizable(t.Ptr()))
}

func (t *NSText) SetVerticallyResizable(value bool) {
	C.Text_SetVerticallyResizable(t.Ptr(), C.bool(value))
}

func (t *NSText) HorizontallyResizable() bool {
	return bool(C.Text_HorizontallyResizable(t.Ptr()))
}

func (t *NSText) SetHorizontallyResizable(value bool) {
	C.Text_SetHorizontallyResizable(t.Ptr(), C.bool(value))
}
