package appkit

type ControlStateValue int

const (
	ControlStateValueMixed ControlStateValue = -1
	ControlStateValueOff   ControlStateValue = 0
	ControlStateValueOn    ControlStateValue = 1
)

// ControlSize constants for specifying a cell’s size.
type ControlSize uint

const (
	ControlSizeRegular ControlSize = 0
	ControlSizeSmall   ControlSize = 1
	ControlSizeMini    ControlSize = 2
	ControlSizeLarge   ControlSize = 3
)
