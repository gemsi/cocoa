package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include "text.h"
import "C"

import (
	"github.com/hsiafan/cocoa/foundation"
	"unsafe"
)

// Text is The most general programmatic interface for objects that manage text.
type Text interface {
	View

	// String return the characters of the receiver’s text
	String() string
	// SetString set the characters of the receiver’s text
	SetString(string string)
	// IsEditable return whether the receiver allows the user to edit its text
	IsEditable() bool
	// SetEditable set whether the receiver allows the user to edit its text
	SetEditable(editable bool)
	// IsSelectable return whether the receiver allows the user to select its text
	IsSelectable() bool
	// SetSelectable set whether the receiver allows the user to select its text
	SetSelectable(selectable bool)
	// IsFieldEditor return whether the receiver interprets Tab, Shift-Tab, and Return (Enter) as cues to end editing and possibly to change the first responder
	IsFieldEditor() bool
	// SetFieldEditor set whether the receiver interprets Tab, Shift-Tab, and Return (Enter) as cues to end editing and possibly to change the first responder
	SetFieldEditor(fieldEditor bool)
	// IsRichText return whether the receiver allows the user to apply attributes to specific ranges of the text
	IsRichText() bool
	// SetRichText set whether the receiver allows the user to apply attributes to specific ranges of the text
	SetRichText(richText bool)
	// MinSize return the receiver’s minimum size
	MinSize() foundation.Size
	// SetMinSize set the receiver’s minimum size
	SetMinSize(minSize foundation.Size)
	// MaxSize return the receiver’s maximum size
	MaxSize() foundation.Size
	// SetMaxSize set the receiver’s maximum size
	SetMaxSize(maxSize foundation.Size)
	// IsVerticallyResizable return whether the receiver changes its height to fit the height of its text.
	IsVerticallyResizable() bool
	// SetVerticallyResizable set whether the receiver changes its height to fit the height of its text.
	SetVerticallyResizable(verticallyResizable bool)
	// IsHorizontallyResizable return whether the receiver changes its width to fit the width of its text
	IsHorizontallyResizable() bool
	// SetHorizontallyResizable set whether the receiver changes its width to fit the width of its text
	SetHorizontallyResizable(horizontallyResizable bool)
}

var _ Text = (*NSText)(nil)

type NSText struct {
	NSView
}

// MakeText create a Text from native pointer
func MakeText(ptr unsafe.Pointer) *NSText {
	if ptr == nil {
		return nil
	}
	return &NSText{
		NSView: *MakeView(ptr),
	}
}

func (t *NSText) String() string {
	return C.GoString(C.Text_String(t.Ptr()))
}

func (t *NSText) SetString(string string) {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))
	C.Text_SetString(t.Ptr(), c_string)
}

func (t *NSText) IsEditable() bool {
	return bool(C.Text_IsEditable(t.Ptr()))
}

func (t *NSText) SetEditable(editable bool) {
	C.Text_SetEditable(t.Ptr(), C.bool(editable))
}

func (t *NSText) IsSelectable() bool {
	return bool(C.Text_IsSelectable(t.Ptr()))
}

func (t *NSText) SetSelectable(selectable bool) {
	C.Text_SetSelectable(t.Ptr(), C.bool(selectable))
}

func (t *NSText) IsFieldEditor() bool {
	return bool(C.Text_IsFieldEditor(t.Ptr()))
}

func (t *NSText) SetFieldEditor(fieldEditor bool) {
	C.Text_SetFieldEditor(t.Ptr(), C.bool(fieldEditor))
}

func (t *NSText) IsRichText() bool {
	return bool(C.Text_IsRichText(t.Ptr()))
}

func (t *NSText) SetRichText(richText bool) {
	C.Text_SetRichText(t.Ptr(), C.bool(richText))
}

func (t *NSText) MinSize() foundation.Size {
	return toSize(C.Text_MinSize(t.Ptr()))
}

func (t *NSText) SetMinSize(minSize foundation.Size) {
	C.Text_SetMinSize(t.Ptr(), toNSSize(minSize))
}

func (t *NSText) MaxSize() foundation.Size {
	return toSize(C.Text_MaxSize(t.Ptr()))
}

func (t *NSText) SetMaxSize(maxSize foundation.Size) {
	C.Text_SetMaxSize(t.Ptr(), toNSSize(maxSize))
}

func (t *NSText) IsVerticallyResizable() bool {
	return bool(C.Text_IsVerticallyResizable(t.Ptr()))
}

func (t *NSText) SetVerticallyResizable(verticallyResizable bool) {
	C.Text_SetVerticallyResizable(t.Ptr(), C.bool(verticallyResizable))
}

func (t *NSText) IsHorizontallyResizable() bool {
	return bool(C.Text_IsHorizontallyResizable(t.Ptr()))
}

func (t *NSText) SetHorizontallyResizable(horizontallyResizable bool) {
	C.Text_SetHorizontallyResizable(t.Ptr(), C.bool(horizontallyResizable))
}
