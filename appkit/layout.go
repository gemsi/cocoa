package appkit

type UserInterfaceLayoutDirection int

const (
	UserInterfaceLayoutDirectionLeftToRight UserInterfaceLayoutDirection = 0
	UserInterfaceLayoutDirectionRightToLeft UserInterfaceLayoutDirection = 1
)

type UserInterfaceLayoutOrientation int

const (
	UserInterfaceLayoutOrientationHorizontal UserInterfaceLayoutOrientation = 0
	UserInterfaceLayoutOrientationVertical   UserInterfaceLayoutOrientation = 1
)

// The layout constraint orientation, either horizontal or vertical, that the constraint uses to enforce layout between objects.
type LayoutConstraintOrientation int

const (
	LayoutConstraintOrientationHorizontal LayoutConstraintOrientation = 0
	LayoutConstraintOrientationVertical   LayoutConstraintOrientation = 1
)

type LayoutAttribute int

const (
	LayoutAttributeNotAnAttribute LayoutAttribute = iota
	LayoutAttributeLeft
	LayoutAttributeRight
	LayoutAttributeTop
	LayoutAttributeBottom
	LayoutAttributeLeading
	LayoutAttributeTrailing
	LayoutAttributeWidth
	LayoutAttributeHeight
	LayoutAttributeCenterX
	LayoutAttributeCenterY
	LayoutAttributeLastBaseline
	LayoutAttributeFirstBaseline
	// for ios begin
	LayoutAttributeLeftMargin
	LayoutAttributeRightMargin
	LayoutAttributeTopMargin
	LayoutAttributeBottomMargin
	LayoutAttributeLeadingMargin
	LayoutAttributeTrailingMargin
	LayoutAttributeCenterXWithinMargins
	LayoutAttributeCenterYWithinMargins
	// ios end

	LayoutAttributeBaseline LayoutAttribute = LayoutAttributeLastBaseline
)

// LayoutPriority used to indicate the relative importance of constraints, allowing Auto Layout to make appropriate tradeoffs when satisfying the constraints of the system as a whole.
type LayoutPriority float32

const (
	LayoutPriorityRequired                   LayoutPriority = 1000 // a required constraint.  Do not exceed this.
	LayoutPriorityDefaultHigh                LayoutPriority = 750  // this is the priority level with which a button resists compressing its content.  Note that it is higher than NSLayoutPriorityWindowSizeStayPut.  Thus dragging to resize a window will not make buttons clip.  Rather the window frame is constrained.
	LayoutPriorityDragThatCanResizeWindow    LayoutPriority = 510  // This is the appropriate priority level for a drag that may end up resizing the window.  This needn't be a drag whose explicit purpose is to resize the window. The user might be dragging around window contents, and it might be desirable that the window get bigger to accommodate.
	LayoutPriorityWindowSizeStayPut          LayoutPriority = 500  // This is the priority level at which the window prefers to stay the same size.  It's generally not appropriate to make a constraint at exactly this priority. You want to be higher or lower.
	LayoutPriorityDragThatCannotResizeWindow LayoutPriority = 490  // This is the priority level at which a split view divider, say, is dragged.  It won't resize the window.
	LayoutPriorityDefaultLow                 LayoutPriority = 250  // this is the priority level at which a button hugs its contents horizontally.
	LayoutPriorityFittingSizeCompression     LayoutPriority = 50   // When you issue -[NSView fittingSize], the smallest size that is large enough for the view's contents is computed.  This is the priority level with which the view wants to be as small as possible in that computation.  It's quite low.  It is generally not appropriate to make a constraint at exactly this priority.  You want to be higher or lower.
)
