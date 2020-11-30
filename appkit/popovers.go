package appkit

// the appearance and disappearance behavior of a popover.
type PopoverBehavior int

const (
	//  Your application assumes responsibility for closing the popover. AppKit will still close the popover in a limited number of circumstances. For instance, AppKit will attempt to close the popover when the window of its positioningView is closed.  The exact interactions in which AppKit will close the popover are not guaranteed.  You may consider implementing -cancel: to close the popover when the escape key is pressed.
	PopoverBehaviorApplicationDefined = 0

	//  AppKit will close the popover when the user interacts with a user interface element outside the popover.  Note that interacting with menus or panels that become key only when needed will not cause a transient popover to close.  The exact interactions that will cause transient popovers to close are not specified.
	PopoverBehaviorTransient = 1

	//  AppKit will close the popover when the user interacts with user interface elements in the window containing the popover's positioning view.  Semi-transient popovers cannot be shown relative to views in other popovers, nor can they be shown relative to views in child windows.  The exact interactions that cause semi-transient popovers to close are not specified.
	PopoverBehaviorSemitransient = 2
)
