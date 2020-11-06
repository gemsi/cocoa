package color

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #import "color.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"unsafe"
)

// Color is wrap for cocoa NSColor
type Color interface {
	foundation.Object
}

var _ Color = (*NSColor)(nil)

// Make create a Color from native pointer
func Make(ptr unsafe.Pointer) *NSColor {
	return &NSColor{
		NSObject: *foundation.MakeObject(ptr),
	}
}

type NSColor struct {
	foundation.NSObject
}

// New Creates a color using the given opacity value and RGB components
func New(red, green, blue float64, alpha float64) Color {
	ptr := C.Color_NewColorWithDeviceRed(C.double(red), C.double(green), C.double(blue), C.double(alpha))
	return Make(ptr)
}

func BlackColor() Color {
	ptr := C.Color_BlackColor()
	return Make(ptr)
}

func DarkGrayColor() Color {
	ptr := C.Color_DarkGrayColor()
	return Make(ptr)
}

func LightGrayColor() Color {
	ptr := C.Color_LightGrayColor()
	return Make(ptr)
}

func WhiteColor() Color {
	ptr := C.Color_WhiteColor()
	return Make(ptr)
}

func GrayColor() Color {
	ptr := C.Color_GrayColor()
	return Make(ptr)
}

func RedColor() Color {
	ptr := C.Color_RedColor()
	return Make(ptr)
}

func GreenColor() Color {
	ptr := C.Color_GreenColor()
	return Make(ptr)
}

func BlueColor() Color {
	ptr := C.Color_BlueColor()
	return Make(ptr)
}

func CyanColor() Color {
	ptr := C.Color_CyanColor()
	return Make(ptr)
}

func YellowColor() Color {
	ptr := C.Color_YellowColor()
	return Make(ptr)
}

func MagentaColor() Color {
	ptr := C.Color_MagentaColor()
	return Make(ptr)
}

func OrangeColor() Color {
	ptr := C.Color_OrangeColor()
	return Make(ptr)
}

func PurpleColor() Color {
	ptr := C.Color_PurpleColor()
	return Make(ptr)
}

func BrownColor() Color {
	ptr := C.Color_BrownColor()
	return Make(ptr)
}

func ClearColor() Color {
	ptr := C.Color_ClearColor()
	return Make(ptr)
}

func LabelColor() Color {
	ptr := C.Color_LabelColor()
	return Make(ptr)
}

func SecondaryLabelColor() Color {
	ptr := C.Color_SecondaryLabelColor()
	return Make(ptr)
}

func TertiaryLabelColor() Color {
	ptr := C.Color_TertiaryLabelColor()
	return Make(ptr)
}

func QuaternaryLabelColor() Color {
	ptr := C.Color_QuaternaryLabelColor()
	return Make(ptr)
}

func LinkColor() Color {
	ptr := C.Color_LinkColor()
	return Make(ptr)
}

func PlaceholderTextColor() Color {
	ptr := C.Color_PlaceholderTextColor()
	return Make(ptr)
}

func WindowFrameTextColor() Color {
	ptr := C.Color_WindowFrameTextColor()
	return Make(ptr)
}

func SelectedMenuItemTextColor() Color {
	ptr := C.Color_SelectedMenuItemTextColor()
	return Make(ptr)
}

func AlternateSelectedControlTextColor() Color {
	ptr := C.Color_AlternateSelectedControlTextColor()
	return Make(ptr)
}

func HeaderTextColor() Color {
	ptr := C.Color_HeaderTextColor()
	return Make(ptr)
}

func SeparatorColor() Color {
	ptr := C.Color_SeparatorColor()
	return Make(ptr)
}

func GridColor() Color {
	ptr := C.Color_GridColor()
	return Make(ptr)
}

func WindowBackgroundColor() Color {
	ptr := C.Color_WindowBackgroundColor()
	return Make(ptr)
}

func UnderPageBackgroundColor() Color {
	ptr := C.Color_UnderPageBackgroundColor()
	return Make(ptr)
}

func ControlBackgroundColor() Color {
	ptr := C.Color_ControlBackgroundColor()
	return Make(ptr)
}

func SelectedContentBackgroundColor() Color {
	ptr := C.Color_SelectedContentBackgroundColor()
	return Make(ptr)
}

func UnemphasizedSelectedContentBackgroundColor() Color {
	ptr := C.Color_UnemphasizedSelectedContentBackgroundColor()
	return Make(ptr)
}

func FindHighlightColor() Color {
	ptr := C.Color_FindHighlightColor()
	return Make(ptr)
}

func TextColor() Color {
	ptr := C.Color_TextColor()
	return Make(ptr)
}

func TextBackgroundColor() Color {
	ptr := C.Color_TextBackgroundColor()
	return Make(ptr)
}

func SelectedTextColor() Color {
	ptr := C.Color_SelectedTextColor()
	return Make(ptr)
}

func SelectedTextBackgroundColor() Color {
	ptr := C.Color_SelectedTextBackgroundColor()
	return Make(ptr)
}

func UnemphasizedSelectedTextBackgroundColor() Color {
	ptr := C.Color_UnemphasizedSelectedTextBackgroundColor()
	return Make(ptr)
}

func UnemphasizedSelectedTextColor() Color {
	ptr := C.Color_UnemphasizedSelectedTextColor()
	return Make(ptr)
}

func ControlColor() Color {
	ptr := C.Color_ControlColor()
	return Make(ptr)
}

func ControlTextColor() Color {
	ptr := C.Color_ControlTextColor()
	return Make(ptr)
}

func SelectedControlColor() Color {
	ptr := C.Color_SelectedControlColor()
	return Make(ptr)
}

func SelectedControlTextColor() Color {
	ptr := C.Color_SelectedControlTextColor()
	return Make(ptr)
}

func DisabledControlTextColor() Color {
	ptr := C.Color_DisabledControlTextColor()
	return Make(ptr)
}

func KeyboardFocusIndicatorColor() Color {
	ptr := C.Color_KeyboardFocusIndicatorColor()
	return Make(ptr)
}

func ScrubberTexturedBackgroundColor() Color {
	ptr := C.Color_ScrubberTexturedBackgroundColor()
	return Make(ptr)
}

func SystemRedColor() Color {
	ptr := C.Color_SystemRedColor()
	return Make(ptr)
}

func SystemGreenColor() Color {
	ptr := C.Color_SystemGreenColor()
	return Make(ptr)
}

func SystemBlueColor() Color {
	ptr := C.Color_SystemBlueColor()
	return Make(ptr)
}

func SystemOrangeColor() Color {
	ptr := C.Color_SystemOrangeColor()
	return Make(ptr)
}

func SystemYellowColor() Color {
	ptr := C.Color_SystemYellowColor()
	return Make(ptr)
}

func SystemBrownColor() Color {
	ptr := C.Color_SystemBrownColor()
	return Make(ptr)
}

func SystemPinkColor() Color {
	ptr := C.Color_SystemPinkColor()
	return Make(ptr)
}

func SystemPurpleColor() Color {
	ptr := C.Color_SystemPurpleColor()
	return Make(ptr)
}

func SystemGrayColor() Color {
	ptr := C.Color_SystemGrayColor()
	return Make(ptr)
}

func ControlAccentColor() Color {
	ptr := C.Color_ControlAccentColor()
	return Make(ptr)
}

func HighlightColor() Color {
	ptr := C.Color_HighlightColor()
	return Make(ptr)
}

func ShadowColor() Color {
	ptr := C.Color_ShadowColor()
	return Make(ptr)
}
