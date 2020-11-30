package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include "button.h"
import "C"

import (
	"github.com/hsiafan/cocoa/foundation"
	"unsafe"
)

// Button is a control that defines an area on the screen that can be used to trigger actions
type Button interface {
	Control

	// Title return the title displayed on the button when it’s in an off state
	Title() string
	// SetTitle set the title displayed on the button when it’s in an off state
	SetTitle(title string)
	// BezelStyle return the appearance of the button’s border
	BezelStyle() BezelStyle
	// SetBezelStyle set the appearance of the button’s border
	SetBezelStyle(bezelStyle BezelStyle)
	// State return the button’s state
	State() ControlStateValue
	// SetState set the button’s state
	SetState(state ControlStateValue)
	// SetButtonType sets the button’s type, which affects its user interface and behavior when clicked
	SetButtonType(buttonType ButtonType)
	// SetAction
	SetAction(callback func(sender foundation.Object))
}

var _ Button = (*NSButton)(nil)

type NSButton struct {
	NSControl
	action func(sender foundation.Object)
}

// MakeButton create a Button from native pointer
func MakeButton(ptr unsafe.Pointer) *NSButton {
	if ptr == nil {
		return nil
	}
	return &NSButton{
		NSControl: *MakeControl(ptr),
	}
}
func (b *NSButton) setDelegate() {
	id := resources.NextId()
	C.Button_RegisterDelegate(b.Ptr(), C.long(id))
	resources.Store(id, b)
	foundation.AddDeallocHook(b, func() {
		resources.Delete(id)
	})
}

func (b *NSButton) Title() string {
	return C.GoString(C.Button_Title(b.Ptr()))
}

func (b *NSButton) SetTitle(title string) {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	C.Button_SetTitle(b.Ptr(), cTitle)
}

func (b *NSButton) BezelStyle() BezelStyle {
	return BezelStyle(C.Button_BezelStyle(b.Ptr()))
}

func (b *NSButton) SetBezelStyle(bezelStyle BezelStyle) {
	C.Button_SetBezelStyle(b.Ptr(), C.ulong(bezelStyle))
}

func (b *NSButton) State() ControlStateValue {
	return ControlStateValue(C.Button_State(b.Ptr()))
}

func (b *NSButton) SetState(state ControlStateValue) {
	C.Button_SetState(b.Ptr(), C.long(state))
}

func NewButton(frame foundation.Rect) Button {
	instance := MakeButton(C.Button_NewButton(toNSRect(frame)))
	instance.setDelegate()
	return instance
}

func (b *NSButton) SetButtonType(buttonType ButtonType) {
	C.Button_SetButtonType(b.Ptr(), C.ulong(buttonType))
}

func (b *NSButton) SetAction(callback func(sender foundation.Object)) {
	b.action = callback
}

//export Button_Target_Action
func Button_Target_Action(id int64, sender unsafe.Pointer) {
	button := resources.Get(id).(*NSButton)
	if button.action != nil {
		button.action(foundation.MakeObject(sender))
	}
}
