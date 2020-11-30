package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include "popover.h"
import "C"

import (
	"github.com/hsiafan/cocoa/foundation"
	"unsafe"
)

// Popover is a means to display additional content related to existing content on the screen
type Popover interface {
	Responder

	// Behavior return the behavior of the popover
	Behavior() PopoverBehavior
	// SetBehavior set the behavior of the popover
	SetBehavior(behavior PopoverBehavior)
	// PositioningRect return the rectangle within the positioning view relative to which the popover should be positioned
	PositioningRect() foundation.Rect
	// SetPositioningRect set the rectangle within the positioning view relative to which the popover should be positioned
	SetPositioningRect(positioningRect foundation.Rect)
	// Animates return if the popover is to be animated
	Animates() bool
	// SetAnimates set if the popover is to be animated
	SetAnimates(animates bool)
	// ContentSize return the content size of the popover
	ContentSize() foundation.Size
	// SetContentSize set the content size of the popover
	SetContentSize(contentSize foundation.Size)
	// IsShown return the display state of the popover
	IsShown() bool
	// IsDetached return whether the window created by a popover's detachment is automatically created
	IsDetached() bool
	// Appearance return the appearance of the popover
	Appearance() Appearance
	// SetAppearance set the appearance of the popover
	SetAppearance(appearance Appearance)
	// EffectiveAppearance return the appearance that will be used when the popover is displayed onscreen
	EffectiveAppearance() Appearance
	// PerformClose attempts to close the popover
	PerformClose(sender foundation.Object)
	// Close forces the popover to close without consulting its delegate
	Close()
	// ShowRelativeTo shows the popover anchored to the specified view
	ShowRelativeTo(positioningRect foundation.Rect, positioningView View, preferredEdge foundation.RectEdge)
}

var _ Popover = (*NSPopover)(nil)

type NSPopover struct {
	NSResponder
}

// MakePopover create a Popover from native pointer
func MakePopover(ptr unsafe.Pointer) *NSPopover {
	if ptr == nil {
		return nil
	}
	return &NSPopover{
		NSResponder: *MakeResponder(ptr),
	}
}

func (p *NSPopover) Behavior() PopoverBehavior {
	return PopoverBehavior(C.Popover_Behavior(p.Ptr()))
}

func (p *NSPopover) SetBehavior(behavior PopoverBehavior) {
	C.Popover_SetBehavior(p.Ptr(), C.long(behavior))
}

func (p *NSPopover) PositioningRect() foundation.Rect {
	return toRect(C.Popover_PositioningRect(p.Ptr()))
}

func (p *NSPopover) SetPositioningRect(positioningRect foundation.Rect) {
	C.Popover_SetPositioningRect(p.Ptr(), toNSRect(positioningRect))
}

func (p *NSPopover) Animates() bool {
	return bool(C.Popover_Animates(p.Ptr()))
}

func (p *NSPopover) SetAnimates(animates bool) {
	C.Popover_SetAnimates(p.Ptr(), C.bool(animates))
}

func (p *NSPopover) ContentSize() foundation.Size {
	return toSize(C.Popover_ContentSize(p.Ptr()))
}

func (p *NSPopover) SetContentSize(contentSize foundation.Size) {
	C.Popover_SetContentSize(p.Ptr(), toNSSize(contentSize))
}

func (p *NSPopover) IsShown() bool {
	return bool(C.Popover_IsShown(p.Ptr()))
}

func (p *NSPopover) IsDetached() bool {
	return bool(C.Popover_IsDetached(p.Ptr()))
}

func (p *NSPopover) Appearance() Appearance {
	return MakeAppearance(C.Popover_Appearance(p.Ptr()))
}

func (p *NSPopover) SetAppearance(appearance Appearance) {
	C.Popover_SetAppearance(p.Ptr(), toPointer(appearance))
}

func (p *NSPopover) EffectiveAppearance() Appearance {
	return MakeAppearance(C.Popover_EffectiveAppearance(p.Ptr()))
}

func NewPopover() Popover {
	return MakePopover(C.Popover_NewPopover())
}

func (p *NSPopover) PerformClose(sender foundation.Object) {
	C.Popover_PerformClose(p.Ptr(), toPointer(sender))
}

func (p *NSPopover) Close() {
	C.Popover_Close(p.Ptr())
}

func (p *NSPopover) ShowRelativeTo(positioningRect foundation.Rect, positioningView View, preferredEdge foundation.RectEdge) {
	C.Popover_ShowRelativeTo(p.Ptr(), toNSRect(positioningRect), toPointer(positioningView), C.long(preferredEdge))
}
