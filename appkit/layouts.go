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
	LayoutPriorityDefaultHigh                LayoutPriority = 750  // this is the priority level with which a button resists compressing its content.  Note that it is higher than LayoutPriorityWindowSizeStayPut.  Thus dragging to resize a window will not make buttons clip.  Rather the window frame is constrained.
	LayoutPriorityDragThatCanResizeWindow    LayoutPriority = 510  // This is the appropriate priority level for a drag that may end up resizing the window.  This needn't be a drag whose explicit purpose is to resize the window. The user might be dragging around window contents, and it might be desirable that the window get bigger to accommodate.
	LayoutPriorityWindowSizeStayPut          LayoutPriority = 500  // This is the priority level at which the window prefers to stay the same size.  It's generally not appropriate to make a constraint at exactly this priority. You want to be higher or lower.
	LayoutPriorityDragThatCannotResizeWindow LayoutPriority = 490  // This is the priority level at which a split view divider, say, is dragged.  It won't resize the window.
	LayoutPriorityDefaultLow                 LayoutPriority = 250  // this is the priority level at which a button hugs its contents horizontally.
	LayoutPriorityFittingSizeCompression     LayoutPriority = 50   // When you issue -[NSView fittingSize], the smallest size that is large enough for the view's contents is computed.  This is the priority level with which the view wants to be as small as possible in that computation.  It's quite low.  It is generally not appropriate to make a constraint at exactly this priority.  You want to be higher or lower.
)

type LayoutRelation int

const (
	LayoutRelationLessThanOrEqual    LayoutRelation = -1
	LayoutRelationEqual              LayoutRelation = 0
	LayoutRelationGreaterThanOrEqual LayoutRelation = 1
)

type LayoutFormatOptions uint

const (
	LayoutFormatAlignAllLeft          = 1 << LayoutAttributeLeft
	LayoutFormatAlignAllRight         = 1 << LayoutAttributeRight
	LayoutFormatAlignAllTop           = 1 << LayoutAttributeTop
	LayoutFormatAlignAllBottom        = 1 << LayoutAttributeBottom
	LayoutFormatAlignAllLeading       = 1 << LayoutAttributeLeading
	LayoutFormatAlignAllTrailing      = 1 << LayoutAttributeTrailing
	LayoutFormatAlignAllCenterX       = 1 << LayoutAttributeCenterX
	LayoutFormatAlignAllCenterY       = 1 << LayoutAttributeCenterY
	LayoutFormatAlignAllLastBaseline  = 1 << LayoutAttributeLastBaseline
	LayoutFormatAlignAllFirstBaseline = 1 << LayoutAttributeFirstBaseline
	LayoutFormatAlignAllBaseline      = LayoutFormatAlignAllLastBaseline

	LayoutFormatAlignmentMask = 0xFFFF

	/* choose only one of these three */
	LayoutFormatDirectionLeadingToTrailing = 0 << 16 // default
	LayoutFormatDirectionLeftToRight       = 1 << 16
	LayoutFormatDirectionRightToLeft       = 2 << 16

	LayoutFormatDirectionMask = 0x3 << 16

	// iphone begin
	/* choose only one spacing format */
	LayoutFormatSpacingEdgeToEdge = 0 << 19 // default

	/* Valid only for vertical layouts. Between views with text content the value
	   will be used to determine the distance from the last baseline of the view above
	   to the first baseline of the view below. For views without text content the top
	   or bottom edge will be used in lieu of the baseline position.
	   The default spacing "]-[" will be determined from the line heights of the fonts
	   involved in views with text content, when present.
	*/
	LayoutFormatSpacingBaselineToBaseline = 1 << 19

	LayoutFormatSpacingMask = 0x1 << 19
	// iphone end
)
