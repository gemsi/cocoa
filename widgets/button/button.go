package button

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #import "button.h"
import "C"
import (
	"github.com/hsiafan/cocoa/internal"
	"unsafe"

	"github.com/hsiafan/cocoa"

	"github.com/hsiafan/cocoa/widgets/control"
)

type BezelStyle int

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
	// SetAction set the action when trigger the button
	SetAction(fn func())
	// SetTitle sets the title of the button, which is the text displayed on the button.
	SetTitle(title string)
	SizeToFit()
	SetBezelStyle(BezelStyle)
}

var _ Button = (*NSButton)(nil)

type NSButton struct {
	control.NSControl
	onClick func()
}

func (btn *NSButton) SetBezelStyle(style BezelStyle) {
	C.Button_SetBezelStyle(btn.Ptr(), C.int(style))
}

var resources = internal.NewResourceRegistry()

// NewButton constructs a new button
func New() Button {
	ptr := C.Button_New()
	btn := &NSButton{
		NSControl: *control.Make(ptr),
	}
	id := resources.NextId()
	resources.Store(id, btn)
	cocoa.AddDeallocHook(btn, func() {
		resources.Delete(id)
	})
	C.Button_SetTarget(btn.Ptr(), C.int(id))
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

func (btn *NSButton) SetAction(fn func()) {
	btn.onClick = fn
}

//export onButtonClicked
func onButtonClicked(id C.int) {
	buttonID := int64(id)
	button := resources.Get(buttonID).(*NSButton)

	onClick := button.onClick
	if onClick != nil {
		onClick()
	}
}
