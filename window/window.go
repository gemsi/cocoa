package window

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include "window.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation/notification"
	"github.com/hsiafan/cocoa/internal"
	"unsafe"

	"github.com/hsiafan/cocoa"
	"github.com/hsiafan/cocoa/foundation"

	"github.com/hsiafan/cocoa/responder"
	"github.com/hsiafan/cocoa/view"
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
	DidResize(fn notification.Handler)
	DidMiniaturize(fn notification.Handler)
	DidDeminiaturize(fn notification.Handler)
	DidMove(fn notification.Handler)
	Frame() foundation.Rect
}

var _ Window = (*NSWindow)(nil)

var resource = internal.NewResourceRegistry()

// New constructs and returns a new window.
func New(x, y, width, height int) Window {
	id := resource.NextId()

	ptr := C.Window_New(C.int(x), C.int(y), C.int(width), C.int(height), C.int(id))
	window := &NSWindow{
		handlers:    internal.NewHandlerRegistry(),
		NSResponder: *responder.Make(ptr),
	}

	resource.Store(id, window)

	cocoa.AddDeallocHook(window, func() {
		resource.Delete(id)
	})
	return window
}

// NewRect constructs and returns a new window.
func NewRect(r foundation.Rect) Window {
	return New(r.X, r.Y, r.Width, r.Height)
}

type NSWindow struct {
	title    string
	x        int
	y        int
	w        int
	h        int
	handlers *internal.HandlerRegistry
	responder.NSResponder
}

func (w *NSWindow) Frame() foundation.Rect {
	nsRect := C.Window_Frame(w.Ptr())
	//defer C.free(unsafe.Pointer(nsRect))
	return foundation.MakeRect(int(nsRect.origin.x), int(nsRect.origin.y),
		int(nsRect.size.width), int(nsRect.size.height))
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

const (
	didResize        internal.Event = 0
	didMove          internal.Event = 1
	didMiniaturize   internal.Event = 2
	didDeminiaturize internal.Event = 3
)

func (w *NSWindow) DidResize(handler notification.Handler) {
	w.handlers.Add(didResize, handler)
}

func (w *NSWindow) DidMiniaturize(handler notification.Handler) {
	w.handlers.Add(didMiniaturize, handler)
}

func (w *NSWindow) DidDeminiaturize(handler notification.Handler) {
	w.handlers.Add(didDeminiaturize, handler)
}

func (w *NSWindow) DidMove(handler notification.Handler) {
	w.handlers.Add(didMove, handler)
}

//export onWindowEvent
func onWindowEvent(id C.int, notificationPtr unsafe.Pointer, eventType C.int) {
	event := internal.Event(eventType)
	window := resource.Get(int64(id)).(*NSWindow)

	for _, handler := range window.handlers.Get(event) {
		handler(notification.MakeNotification(notificationPtr, window))
	}
}
