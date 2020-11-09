package indicator

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #import "progress_indicator.h"
import "C"
import (
	"github.com/hsiafan/cocoa/appkit/view"
)

// ProgressIndicator wrap cocoa NSProcessIndicator
type ProgressIndicator interface {
	view.View
	// StartAnimation starts the animation of an indeterminate progress indicator.
	StartAnimation()
	// StopAnimation stops the animation of an indeterminate progress indicator.
	StopAnimation()
	MinValue() float64
	MaxValue() float64
	// SetLimits sets min and max values
	SetLimits(minValue float64, maxValue float64)
	// Value returns the current value of the indicator
	Value() float64
	// SetValue sets the value
	SetValue(value float64)
	// IncrementBy increments value
	IncrementBy(inc float64)
	// IsIndeterminate return if the progressbar is indeterminate
	IsIndeterminate() bool
	// SetIndeterminate sets if the progressbar is indeterminate. Default is true
	SetIndeterminate(value bool)
	// SetDisplayedWhenStopped set whether the progress indicator hides itself when it isn’t animating.
	SetDisplayedWhenStopped(value bool)
	// SetHidden hide the indicator
	SetHidden(hidden bool)
	// Remove removes the indicator from the superview
	Remove()
}

var _ ProgressIndicator = (*NSProgressIndicator)(nil)

// ProgressIndicator represents a indicator control that can trigger actions.
type NSProgressIndicator struct {
	view.NSView
}

// New constructs a new ProgressIndicator
func New() ProgressIndicator {
	ptr := C.ProgressIndicator_New()
	indicator := &NSProgressIndicator{
		NSView: *view.Make(ptr),
	}
	return indicator
}

func (indicator *NSProgressIndicator) StartAnimation() {
	C.ProgressIndicator_StartAnimation(indicator.Ptr())
}

func (indicator *NSProgressIndicator) StopAnimation() {
	C.ProgressIndicator_StopAnimation(indicator.Ptr())
}

func (indicator *NSProgressIndicator) SetLimits(minValue float64, maxValue float64) {
	C.ProgressIndicator_SetLimits(indicator.Ptr(), C.double(minValue), C.double(maxValue))
}

func (indicator *NSProgressIndicator) Value() float64 {
	return float64(C.ProgressIndicator_Value(indicator.Ptr()))
}

func (indicator *NSProgressIndicator) SetValue(value float64) {
	C.ProgressIndicator_SetValue(indicator.Ptr(), C.double(value))
}

func (indicator *NSProgressIndicator) IncrementBy(inc float64) {
	C.ProgressIndicator_IncrementBy(indicator.Ptr(), C.double(inc))
}

func (indicator *NSProgressIndicator) IsIndeterminate() bool {
	return bool(C.ProgressIndicator_IsIndeterminate(indicator.Ptr()))
}

func (indicator *NSProgressIndicator) SetIndeterminate(value bool) {
	C.ProgressIndicator_SetIndeterminate(indicator.Ptr(), C.bool(value))
}

func (indicator *NSProgressIndicator) SetDisplayedWhenStopped(value bool) {
	C.ProgressIndicator_SetDisplayedWhenStopped(indicator.Ptr(), C.bool(value))
}

func (indicator *NSProgressIndicator) SetHidden(hidden bool) {
	C.ProgressIndicator_SetHidden(indicator.Ptr(), C.bool(hidden))
}

func (indicator *NSProgressIndicator) Remove() {
	C.ProgressIndicator_Remove(indicator.Ptr())
}

func (indicator *NSProgressIndicator) MinValue() float64 {
	return float64(C.ProgressIndicator_MinValue(indicator.Ptr()))
}

func (indicator *NSProgressIndicator) MaxValue() float64 {
	return float64(C.ProgressIndicator_MaxValue(indicator.Ptr()))
}
