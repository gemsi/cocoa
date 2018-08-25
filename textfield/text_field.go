package textfield

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include "text_field.h"
import "C"
import (
	c "github.com/hsiafan/cocoa/color"
	"github.com/hsiafan/cocoa/control"
	"github.com/hsiafan/cocoa/internal"
	"unsafe"
)

// TextField wrap cocoa NSTextField
type TextField interface {
	control.Control
	SetBezeled(bezeled bool)
	SetDrawsBackground(draws bool)
	SetEditable(editable bool)
	SetSelectable(selectable bool)
	SetAsLabel()
	SetStringValue(value string)
	TextColor() c.Color
	SetTextColor(color c.Color)
	BackgroundColor() c.Color
	SetBackgroundColor(color c.Color)
}

var _ TextField = (*NSTextField)(nil)

type NSTextField struct {
	control.NSControl
}

var textFields []*NSTextField

func New() TextField {
	buttonID := len(textFields)
	ptr := C.TextField_New(C.int(buttonID))

	textField := &NSTextField{
		NSControl: *control.Make(ptr),
	}
	textFields = append(textFields, textField)
	return textField
}

func (f *NSTextField) SetAsLabel() {
	f.SetBezeled(false)
	f.SetDrawsBackground(false)
	f.SetEditable(false)
	f.SetSelectable(false)
}

func (f *NSTextField) SetDrawsBackground(draws bool) {
	C.TextField_SetDrawsBackground(f.Ptr(), C.int(internal.BoolToInt(draws)))
}

func (f *NSTextField) SetEditable(editable bool) {
	C.TextField_SetEditable(f.Ptr(), C.int(internal.BoolToInt(editable)))
}

func (f *NSTextField) SetSelectable(selectable bool) {
	C.TextField_SetSelectable(f.Ptr(), C.int(internal.BoolToInt(selectable)))
}

func (f *NSTextField) SetBezeled(bezeled bool) {
	C.TextField_SetBezeled(f.Ptr(), C.int(internal.BoolToInt(bezeled)))
}

func (f *NSTextField) TextColor() c.Color {
	colorPtr := C.TextField_TextColor(f.Ptr())
	return c.Make(colorPtr)
}

func (f *NSTextField) SetTextColor(color c.Color) {
	C.TextField_SetTextColor(f.Ptr(), color.Ptr())
}

func (f *NSTextField) BackgroundColor() c.Color {
	colorPtr := C.TextField_BackgroundColor(f.Ptr())
	return c.Make(colorPtr)
}

func (f *NSTextField) SetBackgroundColor(color c.Color) {
	C.TextField_SetBackgroundColor(f.Ptr(), color.Ptr())
}

func (f *NSTextField) SetStringValue(value string) {
	cstr := C.CString(value)
	defer C.free(unsafe.Pointer(cstr))
	C.TextField_SetStringValue(f.Ptr(), cstr)
}
