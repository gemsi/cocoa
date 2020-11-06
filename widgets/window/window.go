package window

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include "window.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation/geometry"
	"github.com/hsiafan/cocoa/foundation/notification"
	"github.com/hsiafan/cocoa/internal"
	"unsafe"

	"github.com/hsiafan/cocoa"
	"github.com/hsiafan/cocoa/interaction/responder"
	"github.com/hsiafan/cocoa/widgets/view"
)

// Window is interface for NSWindow
type Window interface {
	responder.Responder
	// MakeKeyAndOrderFront moves the window to the front of the screen list, shows the window.
	MakeKeyAndOrderFront()
	// SetTitle set title of window
	SetTitle(title string)
	// Center adjust window to the center of screen
	Center()
	// AddView adds a view to the window.
	AddView(view view.View)
	// Update forces the whole window to repaint
	Update()
	DidResize(fn func(notification.Notification))
	DidMiniaturize(fn func(notification.Notification))
	DidDeminiaturize(fn func(notification.Notification))
	DidMove(fn func(notification.Notification))
	// SetFrame set frame for this window
	SetFrame(rect geometry.Rect, display bool)
	// Frame return the frame of this window
	Frame() geometry.Rect
}

var _ Window = (*NSWindow)(nil)

var resource = internal.NewResourceRegistry()

// New constructs and returns a new window.
func New(frame geometry.Rect) Window {
	id := resource.NextId()

	ptr := C.Window_New(C.long(id), toNSRect(frame))
	window := &NSWindow{
		NSResponder: *responder.Make(ptr),
	}

	resource.Store(id, window)

	cocoa.AddDeallocHook(window, func() {
		resource.Delete(id)
	})
	return window
}

type NSWindow struct {
	responder.NSResponder
	didResize        func(notification.Notification)
	didMiniaturize   func(notification.Notification)
	didDeminiaturize func(notification.Notification)
	didMove          func(notification.Notification)
}

func (w *NSWindow) SetFrame(rect geometry.Rect, display bool) {
	C.Window_SetFrame(w.Ptr(), toNSRect(rect), C.int(internal.BoolToInt(display)))
}

func (w *NSWindow) Frame() geometry.Rect {
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

func (w *NSWindow) AddView(view view.View) {
	C.Window_AddView(w.Ptr(), view.Ptr())
}

func (w *NSWindow) Update() {
	C.Window_Update(w.Ptr())
}

func (w *NSWindow) DidResize(handler func(notification.Notification)) {
	w.didResize = handler
}

func (w *NSWindow) DidMiniaturize(handler func(notification.Notification)) {
	w.didMiniaturize = handler
}

func (w *NSWindow) DidDeminiaturize(handler func(notification.Notification)) {
	w.didDeminiaturize = handler
}

func (w *NSWindow) DidMove(handler func(notification.Notification)) {
	w.didMove = handler
}

//export onWindowDidResize
func onWindowDidResize(id C.int, notificationPtr unsafe.Pointer) {
	window := resource.Get(int64(id)).(*NSWindow)
	if window.didResize != nil {
		window.didResize(notification.Make(notificationPtr, window))
	}
}

//export onWindowDidMiniaturize
func onWindowDidMiniaturize(id C.int, notificationPtr unsafe.Pointer) {
	window := resource.Get(int64(id)).(*NSWindow)
	if window.didMiniaturize != nil {
		window.didMiniaturize(notification.Make(notificationPtr, window))
	}
}

//export onWindowDidDeminiaturize
func onWindowDidDeminiaturize(id C.int, notificationPtr unsafe.Pointer) {
	window := resource.Get(int64(id)).(*NSWindow)
	if window.didDeminiaturize != nil {
		window.didDeminiaturize(notification.Make(notificationPtr, window))
	}
}

//export onWindowDidMove
func onWindowDidMove(id C.int, notificationPtr unsafe.Pointer) {
	window := resource.Get(int64(id)).(*NSWindow)
	if window.didMove != nil {
		window.didMove(notification.Make(notificationPtr, window))
	}
}

func toNSRect(rect geometry.Rect) C.NSRect {
	return *(*C.NSRect)(unsafe.Pointer(&rect))
}
func toRect(nsRect C.NSRect) geometry.Rect {
	return *(*geometry.Rect)(unsafe.Pointer(&nsRect))
}
