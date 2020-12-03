package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include "text_field.h"
import "C"

import (
	"github.com/hsiafan/cocoa"
	"github.com/hsiafan/cocoa/foundation"
	"unsafe"
)

// TextField is text that the user can select or edit and that sends its action message to its target when the user presses the Return key
type TextField interface {
	Control

	// IsBezeled return whether the receiver draws a bezeled border around its contents
	IsBezeled() bool
	// SetBezeled set whether the receiver draws a bezeled border around its contents
	SetBezeled(bezeled bool)
	// DrawsBackground return whether the receiver’s cell draws its background color behind its text
	DrawsBackground() bool
	// SetDrawsBackground set whether the receiver’s cell draws its background color behind its text
	SetDrawsBackground(drawsBackground bool)
	// IsEditable return whether the user can edit the value in the text field
	IsEditable() bool
	// SetEditable set whether the user can edit the value in the text field
	SetEditable(editable bool)
	// IsSelectable return whether the receiver is selectable (but not editable)
	IsSelectable() bool
	// SetSelectable set whether the receiver is selectable (but not editable)
	SetSelectable(selectable bool)
	// TextColor return the color used to draw the receiver’s text
	TextColor() Color
	// SetTextColor set the color used to draw the receiver’s text
	SetTextColor(textColor Color)
	// BackgroundColor return the color of the background that the receiver’s cell draws behind the text.
	BackgroundColor() Color
	// SetBackgroundColor set the color of the background that the receiver’s cell draws behind the text.
	SetBackgroundColor(backgroundColor Color)
	ControlTextDidChange(callback func(notification foundation.Notification))
	_controlTextDidChange() func(notification foundation.Notification)
	ControlTextDidEndEditing(callback func(notification foundation.Notification))
	_controlTextDidEndEditing() func(notification foundation.Notification)
	ControlTextDidBeginEditing(callback func(notification foundation.Notification))
	_controlTextDidBeginEditing() func(notification foundation.Notification)
}

var _ TextField = (*NSTextField)(nil)

type NSTextField struct {
	NSControl
	controlTextDidChange       func(notification foundation.Notification)
	controlTextDidEndEditing   func(notification foundation.Notification)
	controlTextDidBeginEditing func(notification foundation.Notification)
}

// MakeTextField create a TextField from native pointer
func MakeTextField(ptr unsafe.Pointer) *NSTextField {
	if ptr == nil {
		return nil
	}
	return &NSTextField{
		NSControl: *MakeControl(ptr),
	}
}
func (t *NSTextField) setDelegate() {
	id := resources.NextId()
	C.TextField_RegisterDelegate(t.Ptr(), C.long(id))
	resources.Store(id, t)
	cocoa.AddDeallocHook(t, func() {
		resources.Delete(id)
	})
}

func (t *NSTextField) IsBezeled() bool {
	return bool(C.TextField_IsBezeled(t.Ptr()))
}

func (t *NSTextField) SetBezeled(bezeled bool) {
	C.TextField_SetBezeled(t.Ptr(), C.bool(bezeled))
}

func (t *NSTextField) DrawsBackground() bool {
	return bool(C.TextField_DrawsBackground(t.Ptr()))
}

func (t *NSTextField) SetDrawsBackground(drawsBackground bool) {
	C.TextField_SetDrawsBackground(t.Ptr(), C.bool(drawsBackground))
}

func (t *NSTextField) IsEditable() bool {
	return bool(C.TextField_IsEditable(t.Ptr()))
}

func (t *NSTextField) SetEditable(editable bool) {
	C.TextField_SetEditable(t.Ptr(), C.bool(editable))
}

func (t *NSTextField) IsSelectable() bool {
	return bool(C.TextField_IsSelectable(t.Ptr()))
}

func (t *NSTextField) SetSelectable(selectable bool) {
	C.TextField_SetSelectable(t.Ptr(), C.bool(selectable))
}

func (t *NSTextField) TextColor() Color {
	return MakeColor(C.TextField_TextColor(t.Ptr()))
}

func (t *NSTextField) SetTextColor(textColor Color) {
	C.TextField_SetTextColor(t.Ptr(), toPointer(textColor))
}

func (t *NSTextField) BackgroundColor() Color {
	return MakeColor(C.TextField_BackgroundColor(t.Ptr()))
}

func (t *NSTextField) SetBackgroundColor(backgroundColor Color) {
	C.TextField_SetBackgroundColor(t.Ptr(), toPointer(backgroundColor))
}

func NewTextField(frame foundation.Rect) TextField {
	instance := MakeTextField(C.TextField_NewTextField(toNSRect(frame)))
	instance.setDelegate()
	return instance
}

func (t *NSTextField) ControlTextDidChange(callback func(notification foundation.Notification)) {
	t.controlTextDidChange = callback
}

func (t *NSTextField) _controlTextDidChange() func(notification foundation.Notification) {
	return t.controlTextDidChange
}

func (t *NSTextField) ControlTextDidEndEditing(callback func(notification foundation.Notification)) {
	t.controlTextDidEndEditing = callback
}

func (t *NSTextField) _controlTextDidEndEditing() func(notification foundation.Notification) {
	return t.controlTextDidEndEditing
}

func (t *NSTextField) ControlTextDidBeginEditing(callback func(notification foundation.Notification)) {
	t.controlTextDidBeginEditing = callback
}

func (t *NSTextField) _controlTextDidBeginEditing() func(notification foundation.Notification) {
	return t.controlTextDidBeginEditing
}

//export TextField_Delegate_ControlTextDidChange
func TextField_Delegate_ControlTextDidChange(id int64, notification unsafe.Pointer) {
	textField := resources.Get(id).(TextField)
	callback := textField._controlTextDidChange()
	if callback != nil {
		callback(foundation.MakeNotification(notification))
	}
}

//export TextField_Delegate_ControlTextDidEndEditing
func TextField_Delegate_ControlTextDidEndEditing(id int64, notification unsafe.Pointer) {
	textField := resources.Get(id).(TextField)
	callback := textField._controlTextDidEndEditing()
	if callback != nil {
		callback(foundation.MakeNotification(notification))
	}
}

//export TextField_Delegate_ControlTextDidBeginEditing
func TextField_Delegate_ControlTextDidBeginEditing(id int64, notification unsafe.Pointer) {
	textField := resources.Get(id).(TextField)
	callback := textField._controlTextDidBeginEditing()
	if callback != nil {
		callback(foundation.MakeNotification(notification))
	}
}
