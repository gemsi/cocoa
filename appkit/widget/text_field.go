package widget

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include "text_field.h"
import "C"
import (
	c "github.com/hsiafan/cocoa/appkit/color"
	"github.com/hsiafan/cocoa/foundation"
	"unsafe"
)

// TextField wrap cocoa NSTextField
type TextField interface {
	Control
	SetBezeled(bezeled bool)
	SetDrawsBackground(draws bool)
	// SetEditable set if this field can be edited
	SetEditable(editable bool)
	// SetSelectable set if the text can be selectable
	SetSelectable(selectable bool)
	// TextColor return text color
	TextColor() c.Color
	// SetTextColor set text color
	SetTextColor(color c.Color)
	// BackgroundColor return background color
	BackgroundColor() c.Color
	// SetBackgroundColor set background color
	SetBackgroundColor(color c.Color)
	// TextDidChange set handler for text change
	TextDidChange(handler func(foundation.Notification))
	// TextDidEndEditing set handler for text edit finished
	TextDidEndEditing(handler func(foundation.Notification))

	getTextDidChange() func(foundation.Notification)
	getTextDidEndEditing() func(foundation.Notification)
	getOnEnterKey() func(sender foundation.Object)
}

// SecureTextField wrap for NSSecureTextField
type SecureTextField interface {
	TextField
}

var _ TextField = (*NSTextField)(nil)

type NSTextField struct {
	NSControl
	textDidChange     func(foundation.Notification)
	textDidEndEditing func(foundation.Notification)
	onEnterKey        func(sender foundation.Object)
}

type NSSecureTextField struct {
	NSTextField
}

// NewTextField create new TextField
func NewTextField() TextField {
	id := resources.NextId()
	ptr := C.TextField_New(C.long(id))

	textField := &NSTextField{
		NSControl: *MakeControl(ptr),
	}

	resources.Store(id, textField)

	foundation.AddDeallocHook(textField, func() {
		resources.Delete(id)
	})

	return textField
}

// NewSecure create new TextField
func NewSecure() SecureTextField {
	id := resources.NextId()
	ptr := C.SecureTextField_New(C.long(id))

	textField := &NSSecureTextField{
		NSTextField: NSTextField{
			NSControl: *MakeControl(ptr),
		},
	}

	resources.Store(id, textField)

	foundation.AddDeallocHook(textField, func() {
		resources.Delete(id)
	})

	return textField
}

func (f *NSTextField) getTextDidChange() func(foundation.Notification) {
	return f.textDidChange
}

func (f *NSTextField) getTextDidEndEditing() func(foundation.Notification) {
	return f.textDidEndEditing
}

func (f *NSTextField) getOnEnterKey() func(sender foundation.Object) {
	return f.onEnterKey
}

// NewLabel create a text field, which looks like a Label
func NewLabel() TextField {
	tf := NewTextField()
	tf.SetBezeled(false)
	tf.SetDrawsBackground(false)
	tf.SetEditable(false)
	tf.SetSelectable(false)
	return tf
}

func (f *NSTextField) SetDrawsBackground(draws bool) {
	C.TextField_SetDrawsBackground(f.Ptr(), C.bool(draws))
}

func (f *NSTextField) SetEditable(editable bool) {
	C.TextField_SetEditable(f.Ptr(), C.bool(editable))
}

func (f *NSTextField) SetSelectable(selectable bool) {
	C.TextField_SetSelectable(f.Ptr(), C.bool(selectable))
}

func (f *NSTextField) SetBezeled(bezeled bool) {
	C.TextField_SetBezeled(f.Ptr(), C.bool(bezeled))
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

func (f *NSTextField) TextDidChange(handler func(foundation.Notification)) {
	f.textDidChange = handler
}

func (f *NSTextField) TextDidEndEditing(handler func(foundation.Notification)) {
	f.textDidEndEditing = handler
}

func (f *NSTextField) OnEnterKey(handler func(sender foundation.Object)) {
	f.onEnterKey = handler
}

//export onTextDidChange
func onTextDidChange(id int64, notificationPtr unsafe.Pointer) {
	f := resources.Get(id).(TextField)
	if f.getTextDidChange() != nil {
		f.getTextDidChange()(foundation.MakeNotification(notificationPtr, f))
	}
}

//export onTextDidEndEditing
func onTextDidEndEditing(id int64, notificationPtr unsafe.Pointer) {
	f := resources.Get(id).(TextField)
	if f.getTextDidEndEditing() != nil {
		f.getTextDidEndEditing()(foundation.MakeNotification(notificationPtr, f))
	}
}

//export onEnterKey
func onEnterKey(id int64, sender unsafe.Pointer) {
	f := resources.Get(id).(TextField)
	if f.getOnEnterKey() != nil {
		f.getOnEnterKey()(foundation.MakeObject(sender))
	}
}
