package window

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include "window.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation/notification"
	"sync"
	"unsafe"

	"github.com/hsiafan/cocoa"
	"github.com/hsiafan/cocoa/foundation"

	"github.com/hsiafan/cocoa/responder"
	"github.com/hsiafan/cocoa/view"
)

// event - different window delegate Events
type event int

const (
	didResize        event = 0
	didMove          event = 1
	didMiniaturize   event = 2
	didDeminiaturize event = 3
)

// EventHandler - handler functions that accepts the updated window as parameter
type EventHandler func(notification notification.Notification)

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
	OnDidResize(fn EventHandler)
	OnDidMiniaturize(fn EventHandler)
	OnDidDeminiaturize(fn EventHandler)
	OnDidMove(fn EventHandler)
	Frame() foundation.Rect
}

var _ Window = (*NSWindow)(nil)

var currentId int64
var windows = make(map[int64]*NSWindow)
var lock sync.RWMutex

// New constructs and returns a new window.
func New(x, y, width, height int) Window {
	ptr := C.Window_New(C.int(x), C.int(y), C.int(width), C.int(height))
	wnd := &NSWindow{
		callbacks:   make(map[event]EventHandler),
		NSResponder: *responder.Make(ptr),
	}
	lock.Lock()
	currentId++
	id := currentId
	windows[id] = wnd
	lock.Unlock()
	C.Window_SetDelegate(wnd.Ptr(), C.int(id))
	cocoa.AddDeallocHook(wnd, func() {
		lock.Lock()
		delete(windows, id)
		lock.Unlock()
	})
	return wnd
}

// NewRect constructs and returns a new window.
func NewRect(r foundation.Rect) Window {
	return New(r.X, r.Y, r.Width, r.Height)
}

type NSWindow struct {
	title     string
	x         int
	y         int
	w         int
	h         int
	callbacks map[event]EventHandler
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

func (w *NSWindow) OnDidResize(fn EventHandler) {
	w.callbacks[didResize] = fn
}

func (w *NSWindow) OnDidMiniaturize(fn EventHandler) {
	w.callbacks[didMiniaturize] = fn
}

func (w *NSWindow) OnDidDeminiaturize(fn EventHandler) {
	w.callbacks[didDeminiaturize] = fn
}

func (w *NSWindow) OnDidMove(fn EventHandler) {
	w.callbacks[didMove] = fn
}

//export onWindowEvent
func onWindowEvent(id C.int, notificationPtr unsafe.Pointer, eventType C.int) {
	event := event(eventType)
	lock.Lock()
	window := windows[int64(id)]
	lock.Unlock()

	handler := window.callbacks[event]
	if handler != nil {
		handler(notification.MakeNotification(notificationPtr, window))
	}
}
