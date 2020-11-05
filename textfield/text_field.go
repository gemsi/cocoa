package textfield

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include "text_field.h"
import "C"
import (
	"github.com/hsiafan/cocoa"
	c "github.com/hsiafan/cocoa/color"
	"github.com/hsiafan/cocoa/control"
	"github.com/hsiafan/cocoa/foundation/notification"
	"github.com/hsiafan/cocoa/internal"
	"unsafe"
)

// TextField wrap cocoa NSTextField
type TextField interface {
	control.Control
	SetBezeled(bezeled bool)
	SetDrawsBackground(draws bool)
	// SetEditable set if this field can be edited
	SetEditable(editable bool)
	// SetSelectable set if the text can be selectable
	SetSelectable(selectable bool)
	// SetAsLabel set the styles of this text-field, so it act as a Label
	SetAsLabel()
	// TextColor return text color
	TextColor() c.Color
	// SetTextColor set text color
	SetTextColor(color c.Color)
	// BackgroundColor return background color
	BackgroundColor() c.Color
	// SetBackgroundColor set background color
	SetBackgroundColor(color c.Color)
	// TextDidChange set handler for text change
	TextDidChange(handler notification.Handler)
	// TextDidEndEditing set handler for text edit finished
	TextDidEndEditing(handler notification.Handler)
}

var _ TextField = (*NSTextField)(nil)

type NSTextField struct {
	control.NSControl
	handlers *internal.HandlerRegistry
}

var resources = internal.NewResourceRegistry()

func New() TextField {
	id := resources.NextId()
	ptr := C.TextField_New(C.int(id))

	textField := &NSTextField{
		NSControl: *control.Make(ptr),
		handlers:  internal.NewHandlerRegistry(),
	}

	resources.Store(id, textField)

	cocoa.AddDeallocHook(textField, func() {
		resources.Delete(id)
	})

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

const (
	controlTextDidChange     internal.Event = 0
	controlTextDidEndEditing internal.Event = 1
)

func (f *NSTextField) TextDidChange(handler notification.Handler) {
	f.handlers.Add(controlTextDidChange, handler)
}

func (f *NSTextField) TextDidEndEditing(handler notification.Handler) {
	f.handlers.Add(controlTextDidEndEditing, handler)
}

//export onTextFieldEvent
func onTextFieldEvent(id C.int, notificationPtr unsafe.Pointer, eventType C.int) {
	event := internal.Event(eventType)
	textField := resources.Get(int64(id)).(*NSTextField)

	handlers := textField.handlers.Get(event)
	for _, handler := range handlers {
		handler(notification.MakeNotification(notificationPtr, textField))
	}
}
