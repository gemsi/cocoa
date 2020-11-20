package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include "event.h"
import "C"

import (
	"github.com/hsiafan/cocoa/foundation"
	"unsafe"
)

// Event is an object that contains information about an input action, such as a mouse click or a key press
type Event interface {
	foundation.Object

	// LocationInWindow return the receiver’s location in the base coordinate system of the associated window
	LocationInWindow() foundation.Point
	// Window return the window object associated with the event
	Window() Window
	// WindowNumber return the identifier for the window device associated with the event
	WindowNumber() int64
	// Timestamp return the time when the event occurred in seconds since system startup.
	Timestamp() float64
}

var _ Event = (*NSEvent)(nil)

type NSEvent struct {
	foundation.NSObject
}

// MakeEvent create a Event from native pointer
func MakeEvent(ptr unsafe.Pointer) *NSEvent {
	if ptr == nil {
		return nil
	}
	return &NSEvent{
		NSObject: *foundation.MakeObject(ptr),
	}
}

func (e *NSEvent) LocationInWindow() foundation.Point {
	return toPoint(C.Event_LocationInWindow(e.Ptr()))
}

func (e *NSEvent) Window() Window {
	return MakeWindow(C.Event_Window(e.Ptr()))
}

func (e *NSEvent) WindowNumber() int64 {
	return int64(C.Event_WindowNumber(e.Ptr()))
}

func (e *NSEvent) Timestamp() float64 {
	return float64(C.Event_Timestamp(e.Ptr()))
}
