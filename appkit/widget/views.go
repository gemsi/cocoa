package widget

// BorderType specify the type of a view’s border.
type BorderType uint64

const (
	NoBorder     BorderType = 0 // No border
	LineBorder   BorderType = 1 // A black line border around the view
	BezelBorder  BorderType = 2 // A concave border that makes the view look sunken
	GrooveBorder BorderType = 3 // A thin border that looks etched around the image
)

// Constants used by the autoresizingMask property.
type AutoresizingMaskOptions uint64

const (
	ViewNotSizable    AutoresizingMaskOptions = 0
	ViewMinXMargin    AutoresizingMaskOptions = 1
	ViewWidthSizable  AutoresizingMaskOptions = 2
	ViewMaxXMargin    AutoresizingMaskOptions = 4
	ViewMinYMargin    AutoresizingMaskOptions = 8
	ViewHeightSizable AutoresizingMaskOptions = 16
	ViewMaxYMargin    AutoresizingMaskOptions = 32
)
