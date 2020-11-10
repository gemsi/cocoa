package widget

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #import "button.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"unsafe"
)

// Button represents a button control that can trigger actions.
type Button interface {
	Control
	// Clicked set the action when trigger the button
	SetAction(handler Action)
	// SetTitle sets the title of the button, which is the text displayed on the button.
	SetTitle(title string)
	SizeToFit()
	SetBezelStyle(BezelStyle)
}

var _ Button = (*NSButton)(nil)

type NSButton struct {
	NSControl
	action Action
}

func (btn *NSButton) SetBezelStyle(style BezelStyle) {
	C.Button_SetBezelStyle(btn.Ptr(), C.ulong(style))
}

// NewButton constructs a new button
func NewButton() Button {
	id := resources.NextId()
	ptr := C.Button_New(C.long(id))
	btn := &NSButton{
		NSControl: *MakeControl(ptr),
	}
	resources.Store(id, btn)
	foundation.AddDeallocHook(btn, func() {
		resources.Delete(id)
	})
	btn.SetBezelStyle(BezelStyleRounded)
	return btn
}

func (btn *NSButton) SetTitle(title string) {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	C.Button_SetTitle(btn.Ptr(), cTitle)
}

func (btn *NSButton) SizeToFit() {
	C.Button_SizeToFit(btn.Ptr())
}

func (btn *NSButton) SetAction(action Action) {
	btn.action = action
}

//export onButtonAction
func onButtonAction(id int64, sender unsafe.Pointer) {
	button := resources.Get(id).(*NSButton)
	if button.action != nil {
		button.action(foundation.MakeObject(sender))
	}
}
