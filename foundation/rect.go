package foundation

// Point
type Point struct {
	X int
	Y int
}

// Size
type Size struct {
	Width  int
	Height int
}

// Rect
type Rect struct {
	Point
	Size
}

func MakeRect(x, y, width, height int) Rect {
	return Rect{
		Point{x, y},
		Size{width, height},
	}
}
