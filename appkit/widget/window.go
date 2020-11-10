package widget

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include "window.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/internal"
	"unsafe"
)

// Window is interface for NSWindow
type Window interface {
	Responder
	// MakeKeyAndOrderFront moves the window to the front of the screen list, shows the window.
	MakeKeyAndOrderFront()
	// SetTitle set title of window
	SetTitle(title string)
	// Center adjust window to the center of screen
	Center()
	// AddView adds a view to the window.
	AddView(view View)
	// Update forces the whole window to repaint
	Update()
	DidResize(fn func(foundation.Notification))
	DidMiniaturize(fn func(foundation.Notification))
	DidDeminiaturize(fn func(foundation.Notification))
	DidMove(fn func(foundation.Notification))
	// SetFrame set frame for this window
	SetFrame(rect foundation.Rect, display bool)
	// Frame return the frame of this window
	Frame() foundation.Rect
}

var _ Window = (*NSWindow)(nil)

var resource = internal.NewResourceRegistry()

// NewWindow constructs and returns a new window.
func NewWindow(frame foundation.Rect) Window {
	id := resource.NextId()

	ptr := C.Window_New(C.long(id), toNSRect(frame))
	window := &NSWindow{
		NSResponder: *MakeResponder(ptr),
	}

	resource.Store(id, window)

	foundation.AddDeallocHook(window, func() {
		resource.Delete(id)
	})
	return window
}

type NSWindow struct {
	NSResponder
	didResize        func(foundation.Notification)
	didMiniaturize   func(foundation.Notification)
	didDeminiaturize func(foundation.Notification)
	didMove          func(foundation.Notification)
}

func (w *NSWindow) SetFrame(rect foundation.Rect, display bool) {
	C.Window_SetFrame(w.Ptr(), toNSRect(rect), C.bool(display))
}

func (w *NSWindow) Frame() foundation.Rect {
	nsRect := C.Window_Frame(w.Ptr())
	return toRect(nsRect)
}

func (w *NSWindow) SetTitle(title string) {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	C.Window_SetTitle(w.Ptr(), cTitle)
}

func (w *NSWindow) Center() {
	C.Window_Center(w.Ptr())
}

func (w *NSWindow) MakeKeyAndOrderFront() {
	C.Window_MakeKeyAndOrderFront(w.Ptr())
}

func (w *NSWindow) AddView(view View) {
	C.Window_AddView(w.Ptr(), view.Ptr())
}

func (w *NSWindow) Update() {
	C.Window_Update(w.Ptr())
}

func (w *NSWindow) DidResize(handler func(foundation.Notification)) {
	w.didResize = handler
}

func (w *NSWindow) DidMiniaturize(handler func(foundation.Notification)) {
	w.didMiniaturize = handler
}

func (w *NSWindow) DidDeminiaturize(handler func(foundation.Notification)) {
	w.didDeminiaturize = handler
}

func (w *NSWindow) DidMove(handler func(foundation.Notification)) {
	w.didMove = handler
}

//export onWindowDidResize
func onWindowDidResize(id C.int, notificationPtr unsafe.Pointer) {
	window := resource.Get(int64(id)).(*NSWindow)
	if window.didResize != nil {
		window.didResize(foundation.MakeNotification(notificationPtr, window))
	}
}

//export onWindowDidMiniaturize
func onWindowDidMiniaturize(id C.int, notificationPtr unsafe.Pointer) {
	window := resource.Get(int64(id)).(*NSWindow)
	if window.didMiniaturize != nil {
		window.didMiniaturize(foundation.MakeNotification(notificationPtr, window))
	}
}

//export onWindowDidDeminiaturize
func onWindowDidDeminiaturize(id C.int, notificationPtr unsafe.Pointer) {
	window := resource.Get(int64(id)).(*NSWindow)
	if window.didDeminiaturize != nil {
		window.didDeminiaturize(foundation.MakeNotification(notificationPtr, window))
	}
}

//export onWindowDidMove
func onWindowDidMove(id C.int, notificationPtr unsafe.Pointer) {
	window := resource.Get(int64(id)).(*NSWindow)
	if window.didMove != nil {
		window.didMove(foundation.MakeNotification(notificationPtr, window))
	}
}
