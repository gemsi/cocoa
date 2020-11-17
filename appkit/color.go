package appkit

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
func MakeColor(ptr unsafe.Pointer) *NSColor {
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
	return MakeColor(ptr)
}

var ColorBlack = MakeColor(C.Color_BlackColor())

var ColorDarkGray = MakeColor(C.Color_DarkGrayColor())

var ColorLightGray = MakeColor(C.Color_LightGrayColor())

var ColorWhite = MakeColor(C.Color_WhiteColor())

var ColorGray = MakeColor(C.Color_GrayColor())

var ColorRed = MakeColor(C.Color_RedColor())

var ColorGreen = MakeColor(C.Color_GreenColor())

var ColorBlue = MakeColor(C.Color_BlueColor())

var ColorCyan = MakeColor(C.Color_CyanColor())

var ColorYellow = MakeColor(C.Color_YellowColor())

var ColorMagenta = MakeColor(C.Color_MagentaColor())

var ColorOrange = MakeColor(C.Color_OrangeColor())

var ColorPurple = MakeColor(C.Color_PurpleColor())

var ColorBrown = MakeColor(C.Color_BrownColor())

var ColorClear = MakeColor(C.Color_ClearColor())

var ColorLabel = MakeColor(C.Color_LabelColor())

var ColorSecondaryLabel = MakeColor(C.Color_SecondaryLabelColor())

var ColorTertiaryLabel = MakeColor(C.Color_TertiaryLabelColor())

var ColorQuaternaryLabel = MakeColor(C.Color_QuaternaryLabelColor())

var ColorLink = MakeColor(C.Color_LinkColor())

var ColorPlaceholderText = MakeColor(C.Color_PlaceholderTextColor())

var ColorWindowFrameText = MakeColor(C.Color_WindowFrameTextColor())

var ColorSelectedMenuItemText = MakeColor(C.Color_SelectedMenuItemTextColor())

var ColorAlternateSelectedControlText = MakeColor(C.Color_AlternateSelectedControlTextColor())

var ColorHeaderText = MakeColor(C.Color_HeaderTextColor())

var ColorSeparator = MakeColor(C.Color_SeparatorColor())

var ColorGrid = MakeColor(C.Color_GridColor())

var ColorWindowBackground = MakeColor(C.Color_WindowBackgroundColor())

var ColorUnderPageBackground = MakeColor(C.Color_UnderPageBackgroundColor())

var ColorControlBackground = MakeColor(C.Color_ControlBackgroundColor())

var ColorSelectedContentBackground = MakeColor(C.Color_SelectedContentBackgroundColor())

var ColorUnemphasizedSelectedContentBackground = MakeColor(C.Color_UnemphasizedSelectedContentBackgroundColor())

var ColorFindHighlight = MakeColor(C.Color_FindHighlightColor())

var ColorText = MakeColor(C.Color_TextColor())

var ColorTextBackground = MakeColor(C.Color_TextBackgroundColor())

var ColorSelectedText = MakeColor(C.Color_SelectedTextColor())

var ColorSelectedTextBackground = MakeColor(C.Color_SelectedTextBackgroundColor())

var ColorUnemphasizedSelectedTextBackground = MakeColor(C.Color_UnemphasizedSelectedTextBackgroundColor())

var ColorUnemphasizedSelectedText = MakeColor(C.Color_UnemphasizedSelectedTextColor())

var ColorControl = MakeColor(C.Color_ControlColor())

var ColorControlText = MakeColor(C.Color_ControlTextColor())

var ColorSelectedControl = MakeColor(C.Color_SelectedControlColor())

var ColorSelectedControlText = MakeColor(C.Color_SelectedControlTextColor())

var ColorDisabledControlText = MakeColor(C.Color_DisabledControlTextColor())

var ColorKeyboardFocusIndicator = MakeColor(C.Color_KeyboardFocusIndicatorColor())

var ColorScrubberTexturedBackground = MakeColor(C.Color_ScrubberTexturedBackgroundColor())

var ColorSystemRed = MakeColor(C.Color_SystemRedColor())

var ColorSystemGreen = MakeColor(C.Color_SystemGreenColor())

var ColorSystemBlue = MakeColor(C.Color_SystemBlueColor())

var ColorSystemOrange = MakeColor(C.Color_SystemOrangeColor())

var ColorSystemYellow = MakeColor(C.Color_SystemYellowColor())

var ColorSystemBrown = MakeColor(C.Color_SystemBrownColor())

var ColorSystemPink = MakeColor(C.Color_SystemPinkColor())

var ColorSystemPurple = MakeColor(C.Color_SystemPurpleColor())

var ColorSystemGray = MakeColor(C.Color_SystemGrayColor())

var ColorControlAccent = MakeColor(C.Color_ControlAccentColor())

var ColorHighlight = MakeColor(C.Color_HighlightColor())

var ColorShadow = MakeColor(C.Color_ShadowColor())
