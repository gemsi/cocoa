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

var Black = Make(C.Color_BlackColor())

var DarkGray = Make(C.Color_DarkGrayColor())

var LightGray = Make(C.Color_LightGrayColor())

var White = Make(C.Color_WhiteColor())

var Gray = Make(C.Color_GrayColor())

var Red = Make(C.Color_RedColor())

var Green = Make(C.Color_GreenColor())

var Blue = Make(C.Color_BlueColor())

var Cyan = Make(C.Color_CyanColor())

var Yellow = Make(C.Color_YellowColor())

var Magenta = Make(C.Color_MagentaColor())

var Orange = Make(C.Color_OrangeColor())

var Purple = Make(C.Color_PurpleColor())

var Brown = Make(C.Color_BrownColor())

var Clear = Make(C.Color_ClearColor())

var Label = Make(C.Color_LabelColor())

var SecondaryLabel = Make(C.Color_SecondaryLabelColor())

var TertiaryLabel = Make(C.Color_TertiaryLabelColor())

var QuaternaryLabel = Make(C.Color_QuaternaryLabelColor())

var Link = Make(C.Color_LinkColor())

var PlaceholderText = Make(C.Color_PlaceholderTextColor())

var WindowFrameText = Make(C.Color_WindowFrameTextColor())

var SelectedMenuItemText = Make(C.Color_SelectedMenuItemTextColor())

var AlternateSelectedControlText = Make(C.Color_AlternateSelectedControlTextColor())

var HeaderText = Make(C.Color_HeaderTextColor())

var Separator = Make(C.Color_SeparatorColor())

var Grid = Make(C.Color_GridColor())

var WindowBackground = Make(C.Color_WindowBackgroundColor())

var UnderPageBackground = Make(C.Color_UnderPageBackgroundColor())

var ControlBackground = Make(C.Color_ControlBackgroundColor())

var SelectedContentBackground = Make(C.Color_SelectedContentBackgroundColor())

var UnemphasizedSelectedContentBackground = Make(C.Color_UnemphasizedSelectedContentBackgroundColor())

var FindHighlight = Make(C.Color_FindHighlightColor())

var Text = Make(C.Color_TextColor())

var TextBackground = Make(C.Color_TextBackgroundColor())

var SelectedText = Make(C.Color_SelectedTextColor())

var SelectedTextBackground = Make(C.Color_SelectedTextBackgroundColor())

var UnemphasizedSelectedTextBackground = Make(C.Color_UnemphasizedSelectedTextBackgroundColor())

var UnemphasizedSelectedText = Make(C.Color_UnemphasizedSelectedTextColor())

var Control = Make(C.Color_ControlColor())

var ControlText = Make(C.Color_ControlTextColor())

var SelectedControl = Make(C.Color_SelectedControlColor())

var SelectedControlText = Make(C.Color_SelectedControlTextColor())

var DisabledControlText = Make(C.Color_DisabledControlTextColor())

var KeyboardFocusIndicator = Make(C.Color_KeyboardFocusIndicatorColor())

var ScrubberTexturedBackground = Make(C.Color_ScrubberTexturedBackgroundColor())

var SystemRed = Make(C.Color_SystemRedColor())

var SystemGreen = Make(C.Color_SystemGreenColor())

var SystemBlue = Make(C.Color_SystemBlueColor())

var SystemOrange = Make(C.Color_SystemOrangeColor())

var SystemYellow = Make(C.Color_SystemYellowColor())

var SystemBrown = Make(C.Color_SystemBrownColor())

var SystemPink = Make(C.Color_SystemPinkColor())

var SystemPurple = Make(C.Color_SystemPurpleColor())

var SystemGray = Make(C.Color_SystemGrayColor())

var ControlAccent = Make(C.Color_ControlAccentColor())

var Highlight = Make(C.Color_HighlightColor())

var Shadow = Make(C.Color_ShadowColor())
