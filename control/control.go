package control

import (
	"unsafe"

	"github.com/hsiafan/cocoa/view"
)

// Control wrap a cocoa NSControl
type Control interface {
	view.View
	// IsEnabled return if this control is enabled
	IsEnabled() bool
}

type NSControl struct {
	view.NSView
}

// Make create a Control from native pointer
func Make(ptr unsafe.Pointer) *NSControl {
	return &NSControl{*view.Make(ptr)}
}

func (c *NSControl) IsEnabled() bool {
	return false
}
