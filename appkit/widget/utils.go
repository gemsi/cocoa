package widget

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
