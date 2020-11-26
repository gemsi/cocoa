package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include <Foundation/NSGeometry.h>
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/internal"
	"unsafe"
)

var resources = internal.NewResourceRegistry()

func toNSRect(rect foundation.Rect) C.NSRect {
	return C.NSRect{
		C.NSPoint{C.double(rect.X), C.double(rect.Y)},
		C.NSSize{C.double(rect.Width), C.double(rect.Height)},
	}
}
func toRect(nsRect C.NSRect) foundation.Rect {
	return foundation.MakeRect(float64(nsRect.origin.x), float64(nsRect.origin.y),
		float64(nsRect.size.width), float64(nsRect.size.height))
}

func toNSPoint(point foundation.Point) C.NSPoint {
	return C.NSPoint{C.double(point.X), C.double(point.Y)}
}
func toPoint(point C.NSPoint) foundation.Point {
	return foundation.Point{X: float64(point.x), Y: float64(point.y)}
}

func toNSSize(size foundation.Size) C.NSSize {
	return C.NSSize{C.double(size.Width), C.double(size.Height)}
}
func toSize(size C.NSSize) foundation.Size {
	return foundation.Size{Width: float64(size.width), Height: float64(size.height)}
}

func toNSEdgeInsets(insets foundation.EdgeInsets) C.NSEdgeInsets {
	return C.NSEdgeInsets{C.double(insets.Top), C.double(insets.Left), C.double(insets.Bottom), C.double(insets.Right)}
}
func toEdgeInsets(insets C.NSEdgeInsets) foundation.EdgeInsets {
	return foundation.EdgeInsets{
		Top:    float64(insets.top),
		Left:   float64(insets.left),
		Bottom: float64(insets.bottom),
		Right:  float64(insets.right),
	}
}

func toPointer(o foundation.Object) unsafe.Pointer {
	if o == nil {
		return nil
	}
	return o.Ptr()
}
