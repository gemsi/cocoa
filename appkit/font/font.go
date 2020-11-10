package font

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #import "font.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"unsafe"
)

// Font
type Font interface {
	foundation.Object
	// FixedPitch returns whether all glyphs in the font have the same advancement
	FixedPitch() bool
	// PointSize returns the point size of the font
	PointSize() float64
	// DisplayName returns the name of the font, including family and face names, to use when displaying the font information to the user.
	DisplayName() string
	// FamilyName returns the family name of the font—for example, “Times” or “Helvetica.”
	FamilyName() string
	// FontName return the full name of the font, as used in PostScript language code—for example, “Times-Roman” or “Helvetica-Oblique.”
	FontName() string
	// Vertical returns whether the font is a vertical font
	Vertical() bool
	// VerticalFont returns a vertical version of the font
	VerticalFont() Font
}

var _ Font = (*NSFont)(nil)

type NSFont struct {
	foundation.NSObject
}

func MakeFont(ptr unsafe.Pointer) *NSFont {
	return &NSFont{
		NSObject: *foundation.MakeObject(ptr),
	}
}

// NewFontWithName creates a font object for the specified font name and font size
func NewFontWithName(name string, size float64) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	C.Font_FontWithName(cname, C.double(size))
}

func (f *NSFont) FixedPitch() bool {
	return bool(C.Font_FixedPitch(f.Ptr()))
}

func (f *NSFont) PointSize() float64 {
	return float64(C.Font_PointSize(f.Ptr()))
}

func (f *NSFont) DisplayName() string {
	return C.GoString(C.Font_DisplayName(f.Ptr()))
}

func (f *NSFont) FamilyName() string {
	return C.GoString(C.Font_FamilyName(f.Ptr()))
}

func (f *NSFont) FontName() string {
	return C.GoString(C.Font_FontName(f.Ptr()))
}

func (f *NSFont) Vertical() bool {
	return bool(C.Font_Vertical(f.Ptr()))
}

func (f *NSFont) VerticalFont() Font {
	panic("implement me")
}

// NewUserFontOfSize returns the font used by default for documents and other text under the user’s control (that is, text whose font the user can normally change), in the specified size.
func NewUserFontOfSize(size float64) {
	C.Font_UserFontOfSize(C.double(size))
}

// NewUserFixedPitchFontOfSize returns the font used by default for documents and other text under the user’s control (that is, text whose font the user can normally change), when that font should be fixed-pitch, in the specified size.
func NewUserFixedPitchFontOfSize(size float64) {
	C.Font_UserFixedPitchFontOfSize(C.double(size))
}

// NewSystemFontOfSize returns the standard system font with the specified size
func NewSystemFontOfSize(size float64) {
	C.Font_SystemFontOfSize(C.double(size))
}

// NewSystemFontOfSize2 returns the standard system font with the specified size and weight
func NewSystemFontOfSize2(size float64, weight Weight) {
	C.Font_SystemFontOfSize2(C.double(size), C.double(weight))
}

// NewBoldSystemFontOfSize returns the standard system font in boldface type with the specified size
func NewBoldSystemFontOfSize(size float64) {
	C.Font_BoldSystemFontOfSize(C.double(size))
}

// NewLabelFontOfSize returns the font used for standard interface labels in the specified size
func NewLabelFontOfSize(size float64) {
	C.Font_LabelFontOfSize(C.double(size))
}

// NewMessageFontOfSize returns the font used for standard interface items, such as button labels, menu items, and so on, in the specified size.
func NewMessageFontOfSize(size float64) {
	C.Font_MessageFontOfSize(C.double(size))
}

// NewMenuBarFontOfSize returns the font used for menu bar items, in the specified size
func NewMenuBarFontOfSize(size float64) {
	C.Font_MenuBarFontOfSize(C.double(size))
}

// NewMenuFontOfSize returns the font used for menu items, in the specified size
func NewMenuFontOfSize(size float64) {
	C.Font_MenuFontOfSize(C.double(size))
}

// NewControlContentFontOfSize returns the font used for the content of controls in the specified size
func NewControlContentFontOfSize(size float64) {
	C.Font_ControlContentFontOfSize(C.double(size))
}

// NewTitleBarFontOfSize returns the font used for window title bars, in the specified size
func NewTitleBarFontOfSize(size float64) {
	C.Font_TitleBarFontOfSize(C.double(size))
}

// NewPaletteFontOfSize returns the font used for palette window title bars, in the specified size
func NewPaletteFontOfSize(size float64) {
	C.Font_PaletteFontOfSize(C.double(size))
}

// NewToolTipsFontOfSize returns the font used for tool tips labels, in the specified size
func NewToolTipsFontOfSize(size float64) {
	C.Font_ToolTipsFontOfSize(C.double(size))
}

// NewMonospacedSystemFontOfSize returns a monospace version of the system font with the specified size and weight.
func NewMonospacedSystemFontOfSize(size float64, weight Weight) {
	C.Font_MonospacedSystemFontOfSize(C.double(size), C.double(weight))
}

// NewMonospacedDigitSystemFontOfSize returns a version of the standard system font that contains monospaced digit glyphs.
func NewMonospacedDigitSystemFontOfSize(size float64, weight Weight) {
	C.Font_MonospacedDigitSystemFontOfSize(C.double(size), C.double(weight))
}

// SystemFontSize returns the size of the standard system font.
func SystemFontSize() float64 {
	return float64(C.Font_SystemFontSize())
}

// SmallSystemFontSize returns the size of the standard small system font.
func SmallSystemFontSize() float64 {
	return float64(C.Font_SmallSystemFontSize())
}

// LabelFontSize returns the size of the standard label font.
func LabelFontSize() float64 {
	return float64(C.Font_LabelFontSize())
}

// Weight is font weight
type Weight float64

var WeightUltraLight = Weight(C.Font_FontWeightUltraLight())

var WeightThin = Weight(C.Font_FontWeightThin())

var WeightLight = Weight(C.Font_FontWeightLight())

var WeightRegular = Weight(C.Font_FontWeightRegular())

var WeightMedium = Weight(C.Font_FontWeightMedium())

var WeightSemibold = Weight(C.Font_FontWeightSemibold())

var WeightBold = Weight(C.Font_FontWeightBold())

var WeightHeavy = Weight(C.Font_FontWeightHeavy())

var WeightBlack = Weight(C.Font_FontWeightBlack())
