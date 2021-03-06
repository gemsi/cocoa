package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include "window.h"
import "C"

import (
	"github.com/hsiafan/cocoa/foundation"
	"unsafe"
)

// Window is a window that an app displays on the screen.
type Window interface {
	Responder

	// Title return the string that appears in the title bar of the window or the path to the represented file
	Title() string
	// SetTitle set the string that appears in the title bar of the window or the path to the represented file
	SetTitle(title string)
	// ContentView return the window’s content view, the highest accessible NSView object in the window’s view hierarchy
	ContentView() View
	// SetContentView set the window’s content view, the highest accessible NSView object in the window’s view hierarchy
	SetContentView(contentView View)
	// StyleMask return the style mask
	StyleMask() WindowStyleMask
	// SetStyleMask set the style mask
	SetStyleMask(styleMask WindowStyleMask)
	// Center sets the window’s location to the center of the screen
	Center()
	// MakeKeyAndOrderFront moves the window to the front of the screen list, within its level, and makes it the key window; that is, it shows the window.
	MakeKeyAndOrderFront(sender foundation.Object)
	// MakeFirstResponder attempts to make a given responder the first responder for the window.nil makes the window its first responder.
	MakeFirstResponder(responder Responder) bool
	// WindowDidResize the window has been resized
	WindowDidResize(callback func(notification foundation.Notification))
	_windowDidResize() func(notification foundation.Notification)
	// WindowDidMove the window has moved
	WindowDidMove(callback func(notification foundation.Notification))
	_windowDidMove() func(notification foundation.Notification)
	// WindowDidMiniaturize the window has been minimized
	WindowDidMiniaturize(callback func(notification foundation.Notification))
	_windowDidMiniaturize() func(notification foundation.Notification)
	// WindowDidDeminiaturize the window has been deminimized
	WindowDidDeminiaturize(callback func(notification foundation.Notification))
	_windowDidDeminiaturize() func(notification foundation.Notification)
}

var _ Window = (*NSWindow)(nil)

type NSWindow struct {
	NSResponder
	windowDidResize        func(notification foundation.Notification)
	windowDidMove          func(notification foundation.Notification)
	windowDidMiniaturize   func(notification foundation.Notification)
	windowDidDeminiaturize func(notification foundation.Notification)
}

// MakeWindow create a Window from native pointer
func MakeWindow(ptr unsafe.Pointer) *NSWindow {
	if ptr == nil {
		return nil
	}
	return &NSWindow{
		NSResponder: *MakeResponder(ptr),
	}
}
func (w *NSWindow) setDelegate() {
	id := resources.NextId()
	C.Window_RegisterDelegate(w.Ptr(), C.long(id))
	resources.Store(id, w)
	foundation.AddDeallocHook(w, func() {
		resources.Delete(id)
	})
}

func (w *NSWindow) Title() string {
	return C.GoString(C.Window_Title(w.Ptr()))
}

func (w *NSWindow) SetTitle(title string) {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	C.Window_SetTitle(w.Ptr(), cTitle)
}

func (w *NSWindow) ContentView() View {
	return MakeView(C.Window_ContentView(w.Ptr()))
}

func (w *NSWindow) SetContentView(contentView View) {
	C.Window_SetContentView(w.Ptr(), toPointer(contentView))
}

func (w *NSWindow) StyleMask() WindowStyleMask {
	return WindowStyleMask(C.Window_StyleMask(w.Ptr()))
}

func (w *NSWindow) SetStyleMask(styleMask WindowStyleMask) {
	C.Window_SetStyleMask(w.Ptr(), C.ulong(styleMask))
}

func NewWindow(rect foundation.Rect, styleMask WindowStyleMask, backing BackingStoreType, Defer bool) Window {
	instance := MakeWindow(C.Window_NewWindow(toNSRect(rect), C.ulong(styleMask), C.ulong(backing), C.bool(Defer)))
	instance.setDelegate()
	return instance
}

func (w *NSWindow) Center() {
	C.Window_Center(w.Ptr())
}

func (w *NSWindow) MakeKeyAndOrderFront(sender foundation.Object) {
	C.Window_MakeKeyAndOrderFront(w.Ptr(), toPointer(sender))
}

func (w *NSWindow) MakeFirstResponder(responder Responder) bool {
	return bool(C.Window_MakeFirstResponder(w.Ptr(), toPointer(responder)))
}

func (w *NSWindow) WindowDidResize(callback func(notification foundation.Notification)) {
	w.windowDidResize = callback
}

func (w *NSWindow) _windowDidResize() func(notification foundation.Notification) {
	return w.windowDidResize
}

func (w *NSWindow) WindowDidMove(callback func(notification foundation.Notification)) {
	w.windowDidMove = callback
}

func (w *NSWindow) _windowDidMove() func(notification foundation.Notification) {
	return w.windowDidMove
}

func (w *NSWindow) WindowDidMiniaturize(callback func(notification foundation.Notification)) {
	w.windowDidMiniaturize = callback
}

func (w *NSWindow) _windowDidMiniaturize() func(notification foundation.Notification) {
	return w.windowDidMiniaturize
}

func (w *NSWindow) WindowDidDeminiaturize(callback func(notification foundation.Notification)) {
	w.windowDidDeminiaturize = callback
}

func (w *NSWindow) _windowDidDeminiaturize() func(notification foundation.Notification) {
	return w.windowDidDeminiaturize
}

//export Window_Delegate_WindowDidResize
func Window_Delegate_WindowDidResize(id int64, notification unsafe.Pointer) {
	window := resources.Get(id).(Window)
	callback := window._windowDidResize()
	if callback != nil {
		callback(foundation.MakeNotification(notification))
	}
}

//export Window_Delegate_WindowDidMove
func Window_Delegate_WindowDidMove(id int64, notification unsafe.Pointer) {
	window := resources.Get(id).(Window)
	callback := window._windowDidMove()
	if callback != nil {
		callback(foundation.MakeNotification(notification))
	}
}

//export Window_Delegate_WindowDidMiniaturize
func Window_Delegate_WindowDidMiniaturize(id int64, notification unsafe.Pointer) {
	window := resources.Get(id).(Window)
	callback := window._windowDidMiniaturize()
	if callback != nil {
		callback(foundation.MakeNotification(notification))
	}
}

//export Window_Delegate_WindowDidDeminiaturize
func Window_Delegate_WindowDidDeminiaturize(id int64, notification unsafe.Pointer) {
	window := resources.Get(id).(Window)
	callback := window._windowDidDeminiaturize()
	if callback != nil {
		callback(foundation.MakeNotification(notification))
	}
}
