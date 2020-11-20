package appkit

// TabState describe the current display state of a tab
type TabState uint

const (
	SelectedTab   TabState = 0
	BackgroundTab TabState = 1
	PressedTab    TabState = 2
)

// the tab view’s type as used by the tabViewType property
type TabViewType uint

const (
	TopTabsBezelBorder    TabViewType = 0
	LeftTabsBezelBorder   TabViewType = 1
	BottomTabsBezelBorder TabViewType = 2
	RightTabsBezelBorder  TabViewType = 3
	NoTabsBezelBorder     TabViewType = 4
	NoTabsLineBorder      TabViewType = 5
	NoTabsNoBorder        TabViewType = 6
)

type TabPosition uint

const (
	TabPositionNone   TabPosition = 0
	TabPositionTop    TabPosition = 1
	TabPositionLeft   TabPosition = 2
	TabPositionBottom TabPosition = 3
	TabPositionRight  TabPosition = 4
)

type TabViewBorderType uint

const (
	TabViewBorderTypeNone  TabViewBorderType = 0
	TabViewBorderTypeLine  TabViewBorderType = 1
	TabViewBorderTypeBezel TabViewBorderType = 2
)
