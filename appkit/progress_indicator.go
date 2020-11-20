package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include "progress_indicator.h"
import "C"

import (
	"github.com/hsiafan/cocoa/foundation"
	"unsafe"
)

// ProgressIndicator is an interface that provides visual feedback to the user about the status of an ongoing task
type ProgressIndicator interface {
	View

	// UsesThreadedAnimation return whether the progress indicator implements animation in a separate thread
	UsesThreadedAnimation() bool
	// SetUsesThreadedAnimation set whether the progress indicator implements animation in a separate thread
	SetUsesThreadedAnimation(usesThreadedAnimation bool)
	// DoubleValue return the value that indicates the current extent of the progress indicator
	DoubleValue() float64
	// SetDoubleValue set the value that indicates the current extent of the progress indicator
	SetDoubleValue(doubleValue float64)
	// MinValue return the minimum value for the progress indicator
	MinValue() float64
	// SetMinValue set the minimum value for the progress indicator
	SetMinValue(minValue float64)
	// MaxValue return the maximum value for the progress indicator
	MaxValue() float64
	// SetMaxValue set the maximum value for the progress indicator
	SetMaxValue(maxValue float64)
	// IsIndeterminate return whether the progress indicator is indeterminate
	IsIndeterminate() bool
	// SetIndeterminate set whether the progress indicator is indeterminate
	SetIndeterminate(indeterminate bool)
	// IsBezeled return whether the progress indicator’s frame has a three-dimensional bezel
	IsBezeled() bool
	// SetBezeled set whether the progress indicator’s frame has a three-dimensional bezel
	SetBezeled(bezeled bool)
	// IsDisplayedWhenStopped return whether the progress indicator hides itself when it isn’t animating
	IsDisplayedWhenStopped() bool
	// SetDisplayedWhenStopped set whether the progress indicator hides itself when it isn’t animating
	SetDisplayedWhenStopped(displayedWhenStopped bool)
	// StartAnimation starts the animation of an indeterminate progress indicator
	StartAnimation(sender foundation.Object)
	// StopAnimation stops the animation of an indeterminate progress indicator
	StopAnimation(sender foundation.Object)
	// IncrementBy advances the progress bar of a determinate progress indicator by the specified amount
	IncrementBy(delta float64)
}

var _ ProgressIndicator = (*NSProgressIndicator)(nil)

type NSProgressIndicator struct {
	NSView
}

// MakeProgressIndicator create a ProgressIndicator from native pointer
func MakeProgressIndicator(ptr unsafe.Pointer) *NSProgressIndicator {
	if ptr == nil {
		return nil
	}
	return &NSProgressIndicator{
		NSView: *MakeView(ptr),
	}
}

// NewProgressIndicator create new ProgressIndicator
func NewProgressIndicator(frame foundation.Rect) ProgressIndicator {
	id := resources.NextId()
	ptr := C.ProgressIndicator_initWithFrame(C.long(id), toNSRect(frame))
	v := &NSProgressIndicator{
		NSView: *MakeView(ptr),
	}
	resources.Store(id, v)
	foundation.AddDeallocHook(v, func() {
		resources.Delete(id)
	})
	return v
}

func (p *NSProgressIndicator) UsesThreadedAnimation() bool {
	return bool(C.ProgressIndicator_UsesThreadedAnimation(p.Ptr()))
}

func (p *NSProgressIndicator) SetUsesThreadedAnimation(usesThreadedAnimation bool) {
	C.ProgressIndicator_SetUsesThreadedAnimation(p.Ptr(), C.bool(usesThreadedAnimation))
}

func (p *NSProgressIndicator) DoubleValue() float64 {
	return float64(C.ProgressIndicator_DoubleValue(p.Ptr()))
}

func (p *NSProgressIndicator) SetDoubleValue(doubleValue float64) {
	C.ProgressIndicator_SetDoubleValue(p.Ptr(), C.double(doubleValue))
}

func (p *NSProgressIndicator) MinValue() float64 {
	return float64(C.ProgressIndicator_MinValue(p.Ptr()))
}

func (p *NSProgressIndicator) SetMinValue(minValue float64) {
	C.ProgressIndicator_SetMinValue(p.Ptr(), C.double(minValue))
}

func (p *NSProgressIndicator) MaxValue() float64 {
	return float64(C.ProgressIndicator_MaxValue(p.Ptr()))
}

func (p *NSProgressIndicator) SetMaxValue(maxValue float64) {
	C.ProgressIndicator_SetMaxValue(p.Ptr(), C.double(maxValue))
}

func (p *NSProgressIndicator) IsIndeterminate() bool {
	return bool(C.ProgressIndicator_IsIndeterminate(p.Ptr()))
}

func (p *NSProgressIndicator) SetIndeterminate(indeterminate bool) {
	C.ProgressIndicator_SetIndeterminate(p.Ptr(), C.bool(indeterminate))
}

func (p *NSProgressIndicator) IsBezeled() bool {
	return bool(C.ProgressIndicator_IsBezeled(p.Ptr()))
}

func (p *NSProgressIndicator) SetBezeled(bezeled bool) {
	C.ProgressIndicator_SetBezeled(p.Ptr(), C.bool(bezeled))
}

func (p *NSProgressIndicator) IsDisplayedWhenStopped() bool {
	return bool(C.ProgressIndicator_IsDisplayedWhenStopped(p.Ptr()))
}

func (p *NSProgressIndicator) SetDisplayedWhenStopped(displayedWhenStopped bool) {
	C.ProgressIndicator_SetDisplayedWhenStopped(p.Ptr(), C.bool(displayedWhenStopped))
}

func (p *NSProgressIndicator) StartAnimation(sender foundation.Object) {
	C.ProgressIndicator_StartAnimation(p.Ptr(), toPointer(sender))
}

func (p *NSProgressIndicator) StopAnimation(sender foundation.Object) {
	C.ProgressIndicator_StopAnimation(p.Ptr(), toPointer(sender))
}

func (p *NSProgressIndicator) IncrementBy(delta float64) {
	C.ProgressIndicator_IncrementBy(p.Ptr(), C.double(delta))
}
