package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include "stack_view.h"
import "C"

import (
	"github.com/hsiafan/cocoa/foundation"
	"unsafe"
)

// StackView is a view that arranges an array of views horizontally or vertically and updates their placement and sizing when the window size changes
type StackView interface {
	View

	// Views return the array of views owned by the stack view
	Views() []View
	// DetachedViews return the detached views from all the stack view’s gravity areas
	DetachedViews() []View
	// Orientation return the horizontal or vertical layout direction of the stack view
	Orientation() UserInterfaceLayoutOrientation
	// SetOrientation set the horizontal or vertical layout direction of the stack view
	SetOrientation(orientation UserInterfaceLayoutOrientation)
	// Alignment return the view alignment within the stack view
	Alignment() LayoutAttribute
	// SetAlignment set the view alignment within the stack view
	SetAlignment(alignment LayoutAttribute)
	// Spacing return the minimum spacing, in points, between adjacent views in the stack view
	Spacing() float64
	// SetSpacing set the minimum spacing, in points, between adjacent views in the stack view
	SetSpacing(spacing float64)
	// EdgeInsets return the geometric padding, in points, inside the stack view, surrounding its views
	EdgeInsets() foundation.EdgeInsets
	// SetEdgeInsets set the geometric padding, in points, inside the stack view, surrounding its views
	SetEdgeInsets(edgeInsets foundation.EdgeInsets)
	// Distribution return distribution
	Distribution() StackViewDistribution
	// SetDistribution set distribution
	SetDistribution(distribution StackViewDistribution)
	// ArrangedSubviews return the array of views arranged by the stack view
	ArrangedSubviews() []View
	// DetachesHiddenViews return whether the stack view removes hidden views from its view hierarchy
	DetachesHiddenViews() bool
	// SetDetachesHiddenViews set whether the stack view removes hidden views from its view hierarchy
	SetDetachesHiddenViews(detachesHiddenViews bool)
	// ViewsInGravity returns the array of views in the specified gravity area in the stack view
	ViewsInGravity(gravity StackViewGravity) []View
	// AddView adds a view to the end of the stack view gravity area
	AddView(view View, gravity StackViewGravity)
	// InsertView adds a view to a stack view gravity area at a specified index position
	InsertView(view View, index uint, gravity StackViewGravity)
	// SetViews specifies an array of views for a specified gravity area in the stack view, replacing any previous views in that area
	SetViews(views []View, gravity StackViewGravity)
	// RemoveView removes a specified view from the stack view
	RemoveView(view View)
	// AddArrangedSubview adds the specified view to the end of the arranged subviews list
	AddArrangedSubview(view View)
	// InsertArrangedSubview adds the provided view to the array of arranged subviews at the specified index
	InsertArrangedSubview(view View, index uint)
	// RemoveArrangedSubview removes the provided view from the stack’s array of arranged subviews
	RemoveArrangedSubview(view View)
	// ClippingResistancePriorityForOrientation returns the Auto Layout priority for resisting clipping of views in the stack view when Auto Layout attempts to reduce the stack view’s size
	ClippingResistancePriorityForOrientation(orientation LayoutConstraintOrientation) LayoutPriority
	// SetClippingResistancePriority sets the Auto Layout priority for resisting clipping of views in the stack view when Auto Layout attempts to reduce the stack view’s size
	SetClippingResistancePriority(clippingResistancePriority LayoutPriority, orientation LayoutConstraintOrientation)
	// HuggingPriorityForOrientation returns the Auto Layout priority for the stack view to minimize its size to fit its contained views as closely as possible, for a specified user interface axis
	HuggingPriorityForOrientation(orientation LayoutConstraintOrientation) LayoutPriority
	// SetHuggingPriority sets the Auto Layout priority for the stack view to minimize its size, for a specified user interface axis
	SetHuggingPriority(huggingPriority LayoutPriority, orientation LayoutConstraintOrientation)
	// CustomSpacingAfterView returns the custom spacing, in points, between a specified view in the stack view and the view that follows it
	CustomSpacingAfterView(view View) float64
	// SetCustomSpacing specifies the custom spacing, in points, between a specified view and the view that follows it in the stack view
	SetCustomSpacing(spacing float64, view View)
	// VisibilityPriorityForView returns the visibility priority for a specified view in the stack view
	VisibilityPriorityForView(view View) StackViewVisibilityPriority
	// SetVisibilityPriority sets the Auto Layout priority for a view to remain attached to the stack view when Auto Layout reduces the stack view’s size
	SetVisibilityPriority(priority StackViewVisibilityPriority, view View)
}

var _ StackView = (*NSStackView)(nil)

type NSStackView struct {
	NSView
}

// MakeStackView create a StackView from native pointer
func MakeStackView(ptr unsafe.Pointer) *NSStackView {
	if ptr == nil {
		return nil
	}
	return &NSStackView{
		NSView: *MakeView(ptr),
	}
}

func (s *NSStackView) Views() []View {
	var cArray C.Array = C.StackView_Views(s.Ptr())
	defer C.free(cArray.data)
	result := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(cArray.data))[:cArray.len:cArray.len]
	var goResult = make([]View, len(result))
	for idx, r := range result {
		goResult[idx] = MakeView(r)
	}
	return goResult
}

func (s *NSStackView) DetachedViews() []View {
	var cArray C.Array = C.StackView_DetachedViews(s.Ptr())
	defer C.free(cArray.data)
	result := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(cArray.data))[:cArray.len:cArray.len]
	var goResult = make([]View, len(result))
	for idx, r := range result {
		goResult[idx] = MakeView(r)
	}
	return goResult
}

func (s *NSStackView) Orientation() UserInterfaceLayoutOrientation {
	return UserInterfaceLayoutOrientation(C.StackView_Orientation(s.Ptr()))
}

func (s *NSStackView) SetOrientation(orientation UserInterfaceLayoutOrientation) {
	C.StackView_SetOrientation(s.Ptr(), C.long(orientation))
}

func (s *NSStackView) Alignment() LayoutAttribute {
	return LayoutAttribute(C.StackView_Alignment(s.Ptr()))
}

func (s *NSStackView) SetAlignment(alignment LayoutAttribute) {
	C.StackView_SetAlignment(s.Ptr(), C.long(alignment))
}

func (s *NSStackView) Spacing() float64 {
	return float64(C.StackView_Spacing(s.Ptr()))
}

func (s *NSStackView) SetSpacing(spacing float64) {
	C.StackView_SetSpacing(s.Ptr(), C.double(spacing))
}

func (s *NSStackView) EdgeInsets() foundation.EdgeInsets {
	return toEdgeInsets(C.StackView_EdgeInsets(s.Ptr()))
}

func (s *NSStackView) SetEdgeInsets(edgeInsets foundation.EdgeInsets) {
	C.StackView_SetEdgeInsets(s.Ptr(), toNSEdgeInsets(edgeInsets))
}

func (s *NSStackView) Distribution() StackViewDistribution {
	return StackViewDistribution(C.StackView_Distribution(s.Ptr()))
}

func (s *NSStackView) SetDistribution(distribution StackViewDistribution) {
	C.StackView_SetDistribution(s.Ptr(), C.long(distribution))
}

func (s *NSStackView) ArrangedSubviews() []View {
	var cArray C.Array = C.StackView_ArrangedSubviews(s.Ptr())
	defer C.free(cArray.data)
	result := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(cArray.data))[:cArray.len:cArray.len]
	var goResult = make([]View, len(result))
	for idx, r := range result {
		goResult[idx] = MakeView(r)
	}
	return goResult
}

func (s *NSStackView) DetachesHiddenViews() bool {
	return bool(C.StackView_DetachesHiddenViews(s.Ptr()))
}

func (s *NSStackView) SetDetachesHiddenViews(detachesHiddenViews bool) {
	C.StackView_SetDetachesHiddenViews(s.Ptr(), C.bool(detachesHiddenViews))
}

func StackViewWithViews(views []View) StackView {
	cViewsData := make([]unsafe.Pointer, len(views))
	for idx, v := range views {
		cViewsData[idx] = v.Ptr()
	}
	cViews := C.Array{data: unsafe.Pointer(&cViewsData[0]), len: C.int(len(views))}
	return MakeStackView(C.StackView_StackViewWithViews(cViews))
}

func (s *NSStackView) ViewsInGravity(gravity StackViewGravity) []View {
	var cArray C.Array = C.StackView_ViewsInGravity(s.Ptr(), C.long(gravity))
	defer C.free(cArray.data)
	result := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(cArray.data))[:cArray.len:cArray.len]
	var goResult = make([]View, len(result))
	for idx, r := range result {
		goResult[idx] = MakeView(r)
	}
	return goResult
}

func (s *NSStackView) AddView(view View, gravity StackViewGravity) {
	C.StackView_AddView(s.Ptr(), toPointer(view), C.long(gravity))
}

func (s *NSStackView) InsertView(view View, index uint, gravity StackViewGravity) {
	C.StackView_InsertView(s.Ptr(), toPointer(view), C.ulong(index), C.long(gravity))
}

func (s *NSStackView) SetViews(views []View, gravity StackViewGravity) {
	cViewsData := make([]unsafe.Pointer, len(views))
	for idx, v := range views {
		cViewsData[idx] = v.Ptr()
	}
	cViews := C.Array{data: unsafe.Pointer(&cViewsData[0]), len: C.int(len(views))}
	C.StackView_SetViews(s.Ptr(), cViews, C.long(gravity))
}

func (s *NSStackView) RemoveView(view View) {
	C.StackView_RemoveView(s.Ptr(), toPointer(view))
}

func (s *NSStackView) AddArrangedSubview(view View) {
	C.StackView_AddArrangedSubview(s.Ptr(), toPointer(view))
}

func (s *NSStackView) InsertArrangedSubview(view View, index uint) {
	C.StackView_InsertArrangedSubview(s.Ptr(), toPointer(view), C.ulong(index))
}

func (s *NSStackView) RemoveArrangedSubview(view View) {
	C.StackView_RemoveArrangedSubview(s.Ptr(), toPointer(view))
}

func (s *NSStackView) ClippingResistancePriorityForOrientation(orientation LayoutConstraintOrientation) LayoutPriority {
	return LayoutPriority(C.StackView_ClippingResistancePriorityForOrientation(s.Ptr(), C.long(orientation)))
}

func (s *NSStackView) SetClippingResistancePriority(clippingResistancePriority LayoutPriority, orientation LayoutConstraintOrientation) {
	C.StackView_SetClippingResistancePriority(s.Ptr(), C.float(clippingResistancePriority), C.long(orientation))
}

func (s *NSStackView) HuggingPriorityForOrientation(orientation LayoutConstraintOrientation) LayoutPriority {
	return LayoutPriority(C.StackView_HuggingPriorityForOrientation(s.Ptr(), C.long(orientation)))
}

func (s *NSStackView) SetHuggingPriority(huggingPriority LayoutPriority, orientation LayoutConstraintOrientation) {
	C.StackView_SetHuggingPriority(s.Ptr(), C.float(huggingPriority), C.long(orientation))
}

func (s *NSStackView) CustomSpacingAfterView(view View) float64 {
	return float64(C.StackView_CustomSpacingAfterView(s.Ptr(), toPointer(view)))
}

func (s *NSStackView) SetCustomSpacing(spacing float64, view View) {
	C.StackView_SetCustomSpacing(s.Ptr(), C.double(spacing), toPointer(view))
}

func (s *NSStackView) VisibilityPriorityForView(view View) StackViewVisibilityPriority {
	return StackViewVisibilityPriority(C.StackView_VisibilityPriorityForView(s.Ptr(), toPointer(view)))
}

func (s *NSStackView) SetVisibilityPriority(priority StackViewVisibilityPriority, view View) {
	C.StackView_SetVisibilityPriority(s.Ptr(), C.float(priority), toPointer(view))
}
