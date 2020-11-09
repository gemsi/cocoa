package textfield

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include "text_field.h"
import "C"
import (
	c "github.com/hsiafan/cocoa/appkit/color"
	"github.com/hsiafan/cocoa/appkit/control"
	"github.com/hsiafan/cocoa/foundation/notification"
	"github.com/hsiafan/cocoa/foundation/object"
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
	// TextColor return text color
	TextColor() c.Color
	// SetTextColor set text color
	SetTextColor(color c.Color)
	// BackgroundColor return background color
	BackgroundColor() c.Color
	// SetBackgroundColor set background color
	SetBackgroundColor(color c.Color)
	// TextDidChange set handler for text change
	TextDidChange(handler func(notification.Notification))
	// TextDidEndEditing set handler for text edit finished
	TextDidEndEditing(handler func(notification.Notification))

	getTextDidChange() func(notification.Notification)
	getTextDidEndEditing() func(notification.Notification)
	getOnEnterKey() func(sender object.Object)
}

// SecureTextField wrap for NSSecureTextField
type SecureTextField interface {
	TextField
}

var _ TextField = (*NSTextField)(nil)

type NSTextField struct {
	control.NSControl
	textDidChange     func(notification.Notification)
	textDidEndEditing func(notification.Notification)
	onEnterKey        func(sender object.Object)
}

type NSSecureTextField struct {
	NSTextField
}

var resources = internal.NewResourceRegistry()

// New create new TextField
func New() TextField {
	id := resources.NextId()
	ptr := C.TextField_New(C.long(id))

	textField := &NSTextField{
		NSControl: *control.Make(ptr),
	}

	resources.Store(id, textField)

	object.AddDeallocHook(textField, func() {
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
			NSControl: *control.Make(ptr),
		},
	}

	resources.Store(id, textField)

	object.AddDeallocHook(textField, func() {
		resources.Delete(id)
	})

	return textField
}

func (f *NSTextField) getTextDidChange() func(notification.Notification) {
	return f.textDidChange
}

func (f *NSTextField) getTextDidEndEditing() func(notification.Notification) {
	return f.textDidEndEditing
}

func (f *NSTextField) getOnEnterKey() func(sender object.Object) {
	return f.onEnterKey
}

// NewLabel create a text field, which looks like a Label
func NewLabel() TextField {
	tf := New()
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

func (f *NSTextField) TextDidChange(handler func(notification.Notification)) {
	f.textDidChange = handler
}

func (f *NSTextField) TextDidEndEditing(handler func(notification.Notification)) {
	f.textDidEndEditing = handler
}

func (f *NSTextField) OnEnterKey(handler func(sender object.Object)) {
	f.onEnterKey = handler
}

//export onTextDidChange
func onTextDidChange(id int64, notificationPtr unsafe.Pointer) {
	f := resources.Get(id).(TextField)
	if f.getTextDidChange() != nil {
		f.getTextDidChange()(notification.Make(notificationPtr, f))
	}
}

//export onTextDidEndEditing
func onTextDidEndEditing(id int64, notificationPtr unsafe.Pointer) {
	f := resources.Get(id).(TextField)
	if f.getTextDidEndEditing() != nil {
		f.getTextDidEndEditing()(notification.Make(notificationPtr, f))
	}
}

//export onEnterKey
func onEnterKey(id int64, sender unsafe.Pointer) {
	f := resources.Get(id).(TextField)
	if f.getOnEnterKey() != nil {
		f.getOnEnterKey()(object.Make(sender))
	}
}
