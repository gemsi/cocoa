package textfield

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include "text_field.h"
import "C"
import (
	"github.com/hsiafan/cocoa"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/foundation/notification"
	c "github.com/hsiafan/cocoa/graphics/color"
	"github.com/hsiafan/cocoa/internal"
	"github.com/hsiafan/cocoa/widgets/control"
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
}

var _ TextField = (*NSTextField)(nil)

type NSTextField struct {
	control.NSControl
	textDidChange        func(notification.Notification)
	textDidEndEditing    func(notification.Notification)
	becameFirstResponder func()
	onEnterKey           func(sender foundation.Object)
}

var resources = internal.NewResourceRegistry()

func New() TextField {
	id := resources.NextId()
	ptr := C.TextField_New(C.long(id))

	textField := &NSTextField{
		NSControl: *control.Make(ptr),
	}

	resources.Store(id, textField)

	cocoa.AddDeallocHook(textField, func() {
		resources.Delete(id)
	})

	return textField
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

func (f *NSTextField) TextDidChange(handler func(notification.Notification)) {
	f.textDidChange = handler
}

func (f *NSTextField) TextDidEndEditing(handler func(notification.Notification)) {
	f.textDidEndEditing = handler
}

func (f *NSTextField) onBecameFirstResponder(handler func()) {
	f.becameFirstResponder = handler
}

func (f *NSTextField) OnEnterKey(handler func(sender foundation.Object)) {
	f.onEnterKey = handler
}

//export onTextDidChange
func onTextDidChange(id int64, notificationPtr unsafe.Pointer) {
	f := resources.Get(id).(*NSTextField)
	if f.textDidChange != nil {
		f.textDidChange(notification.Make(notificationPtr, f))
	}
}

//export onTextDidEndEditing
func onTextDidEndEditing(id int64, notificationPtr unsafe.Pointer) {
	f := resources.Get(id).(*NSTextField)
	if f.textDidEndEditing != nil {
		f.textDidEndEditing(notification.Make(notificationPtr, f))
	}
}

//export onBecameFirstResponder
func onBecameFirstResponder(id int64) {
	f := resources.Get(id).(*NSTextField)
	if f.becameFirstResponder != nil {
		f.becameFirstResponder()
	}
}

//export onEnterKey
func onEnterKey(id int64, sender unsafe.Pointer) {
	f := resources.Get(id).(*NSTextField)
	if f.onEnterKey != nil {
		f.onEnterKey(foundation.MakeObject(sender))
	}
}
