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
	return *(*C.NSRect)(unsafe.Pointer(&rect))
}
func toRect(nsRect C.NSRect) foundation.Rect {
	return *(*foundation.Rect)(unsafe.Pointer(&nsRect))
}

func toNSSize(size foundation.Size) C.NSSize {
	return *(*C.NSSize)(unsafe.Pointer(&size))
}
func toSize(size C.NSSize) foundation.Size {
	return *(*foundation.Size)(unsafe.Pointer(&size))
}

func toNSPoint(pointer foundation.Point) C.NSPoint {
	return *(*C.NSPoint)(unsafe.Pointer(&pointer))
}
func toPoint(pointer C.NSPoint) foundation.Point {
	return *(*foundation.Point)(unsafe.Pointer(&pointer))
}

func toPointer(o foundation.Object) unsafe.Pointer {
	if o == nil {
		return nil
	}
	return o.Ptr()
}
