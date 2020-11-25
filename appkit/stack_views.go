package appkit

import "math"

// StackViewGravity is the gravity areas available in a stack view.
type StackViewGravity int

const (
	StackViewGravityTop      StackViewGravity = 1 // The top-most gravity area, should only be used when orientation = NSUserInterfaceLayoutOrientationVertical StackViewGravityLeading = 1, // The leading gravity area (as described by userInterfaceLayoutDirection), should only be used when orientation = NSUserInterfaceLayoutOrientationHorizontal
	StackViewGravityCenter   StackViewGravity = 2 // The center gravity area, this is the center regardless of orientation
	StackViewGravityBottom   StackViewGravity = 3 // The bottom-most gravity area, should only be used when orientation = NSUserInterfaceLayoutOrientationVertical
	StackViewGravityTrailing StackViewGravity = 3 // The trailing gravity area (as described by userInterfaceLayoutDirection), should only be used when orientation = NSUserInterfaceLayoutOrientationHorizontal
)

const StackViewSpacingUseDefault = math.MaxFloat32

/* Distribution—the layout along the stacking axis.
All NSStackViewDistribution enum values fit first and last stacked views tightly to the container, except for NSStackViewDistributionGravityAreas.
*/
type StackViewDistribution int

const (
	// Default value. NSStackView will not have any special distribution behavior, relying on behavior described by gravity areas and set hugging priorities along the stacking axis.
	NSStackViewDistributionGravityAreas StackViewDistribution = -1

	// The effective hugging priority in the stacking axis is NSLayoutPriorityRequired, causing the stacked views to tightly fill the container along the stacking axis.
	NSStackViewDistributionFill StackViewDistribution = iota

	// Stacked views will have sizes maintained to be equal as much as possible along the stacking axis. The effective hugging priority in the stacking axis is NSLayoutPriorityRequired.
	NSStackViewDistributionFillEqually

	// Stacked views will have sizes maintained to be equal, proportionally to their intrinsicContentSizes, as much as possible. The effective hugging priority in the stacking axis is NSLayoutPriorityRequired.
	NSStackViewDistributionFillProportionally

	// The space separating stacked views along the stacking axis are maintained to be equal as much as possible while still maintaining the minimum spacing.
	NSStackViewDistributionEqualSpacing

	// Equal center-to-center spacing of the items is maintained as much as possible while still maintaining the minimum spacing between each view.
	NSStackViewDistributionEqualCentering
)

/*
 Visibility Priority describes the priority at which a view should be held (aka, not be detached).
 In the case that clippingResistancePriority is optional (< NSLayoutPriorityRequired) and there's not enough space to display all of StackView's subviews, views are able to be detached from the StackView.
 Views will be detached in order (from lowest to highest) of their visibility priority, and reattached in the reverse order (FILO).
 If multiple views share the lowest visibility priority, all those views will be dropped when one needs to be. Likewise, groups of views with equal visibility priorities will wait to be reattached until they can all be readded.

 A view with a higher visibility priority will never be detached while a lower priority view remains visible

 See also:
 - visibilityPriorityForView:
 - setVisibilityPriority:ForView:
 - clippingResistancePriority
 - detachedViews
*/
type StackViewVisibilityPriority float32

const (
	NSStackViewVisibilityPriorityMustHold              StackViewVisibilityPriority = 1000 //Maximum, default - the view will never be detached
	NSStackViewVisibilityPriorityDetachOnlyIfNecessary StackViewVisibilityPriority = 900
	NSStackViewVisibilityPriorityNotVisible            StackViewVisibilityPriority = 0 //Minimum - will force a view to be detached
)
