package foundation

import "C"

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include "rect.h"
import "C"

// Point
type Point C.NSPoint

func (r *Point) SetX(x float64) {
	r.x = C.double(x)
}

func (r *Point) X() float64 {
	return float64(r.x)
}

func (r *Point) SetY(y float64) {
	r.y = C.double(y)
}

func (r *Point) Y() float64 {
	return float64(r.y)
}

// Size
type Size C.NSSize

func (r *Size) SetWidth(v float64) {
	r.width = C.double(v)
}

func (r *Size) Width() float64 {
	return float64(r.width)
}

func (r *Size) SetHeight(v float64) {
	r.height = C.double(v)
}

func (r *Size) Height() float64 {
	return float64(r.height)
}

// Rect
type Rect C.NSRect

func (r *Rect) SetX(x float64) {
	r.origin.x = C.double(x)
}

func (r *Rect) X() float64 {
	return float64(r.origin.x)
}

func (r *Rect) SetY(y float64) {
	r.origin.y = C.double(y)
}

func (r *Rect) Y() float64 {
	return float64(r.origin.y)
}

func (r *Rect) SetWidth(v float64) {
	r.size.width = C.double(v)
}

func (r *Rect) Width() float64 {
	return float64(r.size.width)
}

func (r *Rect) SetHeight(v float64) {
	r.size.height = C.double(v)
}

func (r *Rect) Height() float64 {
	return float64(r.size.height)
}

// MkRect create a Rect struct
func MkRect(x, y, width, height float64) Rect {
	return Rect{
		C.NSPoint(Point{C.double(x), C.double(y)}),
		C.NSSize(Size{C.double(width), C.double(height)}),
	}
}

// MkSize create a Size struct
func MkSize(width, height float64) Size {
	return Size{C.double(width), C.double(height)}
}

// MkPoint create a Point struct
func MkPoint(x, y float64) Point {
	return Point{C.double(x), C.double(y)}
}
