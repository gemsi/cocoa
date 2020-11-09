package button

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #import "button.h"
import "C"
import (
	"github.com/hsiafan/cocoa/appkit/responder"
	"github.com/hsiafan/cocoa/foundation/object"
	"github.com/hsiafan/cocoa/internal"
	"unsafe"

	"github.com/hsiafan/cocoa/appkit/control"
)

type BezelStyle uint32

const (
	BezelStyleRounded           BezelStyle = 1
	BezelStyleRegularSquare     BezelStyle = 2
	BezelStyleDisclosure        BezelStyle = 5
	BezelStyleShadowlessSquare  BezelStyle = 6
	BezelStyleCircular          BezelStyle = 7
	BezelStyleTexturedSquare    BezelStyle = 8
	BezelStyleHelpButton        BezelStyle = 9
	BezelStyleSmallSquare       BezelStyle = 10
	BezelStyleTexturedRounded   BezelStyle = 11
	BezelStyleRoundRect         BezelStyle = 12
	BezelStyleRecessed          BezelStyle = 13
	BezelStyleRoundedDisclosure BezelStyle = 14
	BezelStyleInline            BezelStyle = 15
)

// Button represents a button control that can trigger actions.
type Button interface {
	control.Control
	// Clicked set the action when trigger the button
	SetAction(handler responder.Action)
	// SetTitle sets the title of the button, which is the text displayed on the button.
	SetTitle(title string)
	SizeToFit()
	SetBezelStyle(BezelStyle)
}

var _ Button = (*NSButton)(nil)

type NSButton struct {
	control.NSControl
	action responder.Action
}

func (btn *NSButton) SetBezelStyle(style BezelStyle) {
	C.Button_SetBezelStyle(btn.Ptr(), C.uint(style))
}

var resources = internal.NewResourceRegistry()

// NewButton constructs a new button
func New() Button {
	id := resources.NextId()
	ptr := C.Button_New(C.long(id))
	btn := &NSButton{
		NSControl: *control.Make(ptr),
	}
	resources.Store(id, btn)
	object.AddDeallocHook(btn, func() {
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

func (btn *NSButton) SetAction(action responder.Action) {
	btn.action = action
}

//export onButtonAction
func onButtonAction(id int64, sender unsafe.Pointer) {
	button := resources.Get(id).(*NSButton)
	if button.action != nil {
		button.action(object.Make(sender))
	}
}
