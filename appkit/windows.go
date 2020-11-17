package appkit

import "github.com/hsiafan/cocoa/foundation"

// WindowStyleMask specify the style of a window, and can be combined using bitwise OR operator.
type WindowStyleMask uint

const (
	WindowStyleMaskBorderless     WindowStyleMask = 0
	WindowStyleMaskTitled         WindowStyleMask = 1 << 0
	WindowStyleMaskClosable       WindowStyleMask = 1 << 1
	WindowStyleMaskMiniaturizable WindowStyleMask = 1 << 2
	WindowStyleMaskResizable      WindowStyleMask = 1 << 3

	// Specifies a window whose titlebar and toolbar have a unified look - that is, a continuous background. Under the titlebar and toolbar a horizontal separator line will appear.
	WindowStyleMaskUnifiedTitleAndToolbaWindowStyleMaskr = 1 << 12

	// When set, the window will appear full screen. This mask is automatically toggled when toggleFullScreen: is called.
	WindowStyleMaskFullScreeWindowStyleMaskn = 1 << 14

	// If set, the contentView will consume the full size of the window; it can be combined with other window style masks, but is only respected for windows with a titlebar. Utilizing this mask opts-in to layer-backing. Utilize the contentLayoutRect or auto-layout contentLayoutGuide to layout views underneath the titlebar/toolbar area.
	WindowStyleMaskFullSizeContentVieWindowStyleMaskw = 1 << 15

	// The following are only applicable for NSPanel (or a subclass thereof)
	WindowStyleMaskUtilityWindow                     WindowStyleMask = 1 << 4
	WindowStyleMaskDocModalWindow                    WindowStyleMask = 1 << 6
	WindowStyleMaskNonactivatingPaneWindowStyleMaskl                 = 1 << 7  // Specifies that a panel that does not activate the owning application
	WindowStyleMaskHUDWindow                         WindowStyleMask = 1 << 13 // Specifies a heads up display panel
)

// BackingStoreType specify how the drawing done in a window is buffered by the window device.
type BackingStoreType uint

const (
	// Deprecated
	BackingStoreRetained BackingStoreType = 0
	// Deprecated
	BackingStoreNonretained BackingStoreType = 1
	BackingStoreBuffered    BackingStoreType = 2
)

// NewPlainWindow create a common window with close/minimize buttons
func NewPlainWindow(rect foundation.Rect) Window {
	return NewWindow(
		rect,
		WindowStyleMaskTitled|WindowStyleMaskClosable|WindowStyleMaskResizable|WindowStyleMaskMiniaturizable,
		BackingStoreBuffered,
		false,
	)
}
