package foundation

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include "geometry.h"
import "C"
import "math"

const SIZE_MAX = math.MaxFloat64

// ZeroRect for convenient
var ZeroRect Rect

// ZeroSize for convenient
var ZeroSize Size

// Point
type Point struct {
	X float64
	Y float64
}

// Size
type Size struct {
	Width  float64
	Height float64
}

// Rect
type Rect struct {
	Point
	Size
}

// MakeRect create a Rect struct
func MakeRect(x, y, width, height float64) Rect {
	return Rect{
		Point: Point{x, y},
		Size:  Size{width, height},
	}
}

// EdgeInsets is a description of the distance between the edges of two rectangles.
type EdgeInsets struct {
	Top    float64
	Left   float64
	Bottom float64
	Right  float64
}
