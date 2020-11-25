package foundation

import "C"

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include "geometry.h"
import "C"
import "math"

const SIZE_MAX = math.MaxFloat64

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

// MakeRect create a Rect struct
func MakeRect(x, y, width, height float64) Rect {
	return Rect{
		C.NSPoint(Point{C.double(x), C.double(y)}),
		C.NSSize(Size{C.double(width), C.double(height)}),
	}
}

// MakeSize create a Size struct
func MakeSize(width, height float64) Size {
	return Size{C.double(width), C.double(height)}
}

// MakePoint create a Point struct
func MakePoint(x, y float64) Point {
	return Point{C.double(x), C.double(y)}
}

// EdgeInsets is a description of the distance between the edges of two rectangles.
type EdgeInsets C.NSEdgeInsets

func (e *EdgeInsets) Top() float64 {
	return float64(e.top)
}

func (e *EdgeInsets) Left() float64 {
	return float64(e.left)
}

func (e *EdgeInsets) Bottom() float64 {
	return float64(e.bottom)
}

func (e *EdgeInsets) Right() float64 {
	return float64(e.right)
}

func (e *EdgeInsets) SetTop(value float64) {
	e.top = C.double(value)
}

func (e *EdgeInsets) SetLeft(value float64) {
	e.left = C.double(value)
}

func (e *EdgeInsets) SetBottom(value float64) {
	e.bottom = C.double(value)
}

func (e *EdgeInsets) SetRight(value float64) {
	e.right = C.double(value)
}

func MakeEdgeInsets(top, left, bottom, right float64) EdgeInsets {
	return EdgeInsets{
		top:    C.double(top),
		left:   C.double(left),
		bottom: C.double(bottom),
		right:  C.double(right),
	}
}
