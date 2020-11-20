package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include "split_view.h"
import "C"

import (
	"github.com/hsiafan/cocoa/foundation"
	"unsafe"
)

// SplitView is a view that arranges two or more views in a linear stack running horizontally or vertically
type SplitView interface {
	View

	// ArrangesAllSubviews return whether the split view arranges all of its subviews as split panes
	ArrangesAllSubviews() bool
	// SetArrangesAllSubviews set whether the split view arranges all of its subviews as split panes
	SetArrangesAllSubviews(arrangesAllSubviews bool)
	// IsVertical return the geometric orientation of the split view's dividers
	IsVertical() bool
	// SetVertical set the geometric orientation of the split view's dividers
	SetVertical(vertical bool)
	// DividerColor return the color of the dividers that the split view draws between subviews
	DividerColor() Color
	// DividerThickness return the thickness of the dividers for the split view
	DividerThickness() float64
	// DividerStyle return the style of divider between views
	DividerStyle() SplitViewDividerStyle
	// SetDividerStyle set the style of divider between views
	SetDividerStyle(dividerStyle SplitViewDividerStyle)
	// MinPossiblePositionOfDividerAtIndex returns the minimum possible position of the divider at the specified index
	MinPossiblePositionOfDividerAtIndex(dividerIndex int) float64
	// MaxPossiblePositionOfDividerAtIndex returns the maximum possible position of the divider at the specified index
	MaxPossiblePositionOfDividerAtIndex(dividerIndex int) float64
	// SetPosition sets the position of the divider at the specified index
	SetPosition(position float64, dividerIndex int)
}

var _ SplitView = (*NSSplitView)(nil)

type NSSplitView struct {
	NSView
}

// MakeSplitView create a SplitView from native pointer
func MakeSplitView(ptr unsafe.Pointer) *NSSplitView {
	if ptr == nil {
		return nil
	}
	return &NSSplitView{
		NSView: *MakeView(ptr),
	}
}

// NewSplitView create new SplitView
func NewSplitView(frame foundation.Rect) SplitView {
	id := resources.NextId()
	ptr := C.SplitView_initWithFrame(C.long(id), toNSRect(frame))
	v := &NSSplitView{
		NSView: *MakeView(ptr),
	}
	resources.Store(id, v)
	foundation.AddDeallocHook(v, func() {
		resources.Delete(id)
	})
	return v
}

func (s *NSSplitView) ArrangesAllSubviews() bool {
	return bool(C.SplitView_ArrangesAllSubviews(s.Ptr()))
}

func (s *NSSplitView) SetArrangesAllSubviews(arrangesAllSubviews bool) {
	C.SplitView_SetArrangesAllSubviews(s.Ptr(), C.bool(arrangesAllSubviews))
}

func (s *NSSplitView) IsVertical() bool {
	return bool(C.SplitView_IsVertical(s.Ptr()))
}

func (s *NSSplitView) SetVertical(vertical bool) {
	C.SplitView_SetVertical(s.Ptr(), C.bool(vertical))
}

func (s *NSSplitView) DividerColor() Color {
	return MakeColor(C.SplitView_DividerColor(s.Ptr()))
}

func (s *NSSplitView) DividerThickness() float64 {
	return float64(C.SplitView_DividerThickness(s.Ptr()))
}

func (s *NSSplitView) DividerStyle() SplitViewDividerStyle {
	return SplitViewDividerStyle(C.SplitView_DividerStyle(s.Ptr()))
}

func (s *NSSplitView) SetDividerStyle(dividerStyle SplitViewDividerStyle) {
	C.SplitView_SetDividerStyle(s.Ptr(), C.long(dividerStyle))
}

func (s *NSSplitView) MinPossiblePositionOfDividerAtIndex(dividerIndex int) float64 {
	return float64(C.SplitView_MinPossiblePositionOfDividerAtIndex(s.Ptr(), C.long(dividerIndex)))
}

func (s *NSSplitView) MaxPossiblePositionOfDividerAtIndex(dividerIndex int) float64 {
	return float64(C.SplitView_MaxPossiblePositionOfDividerAtIndex(s.Ptr(), C.long(dividerIndex)))
}

func (s *NSSplitView) SetPosition(position float64, dividerIndex int) {
	C.SplitView_SetPosition(s.Ptr(), C.double(position), C.long(dividerIndex))
}
