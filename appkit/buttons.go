package appkit

import "github.com/hsiafan/cocoa/foundation"

// BezelStyle used by the bezelStyle property.
type BezelStyle uint

const (
	BezelStyleRounded           BezelStyle = 1
	BezelStyleRegularSquare     BezelStyle = 2
	BezelStyleDisclosure        BezelStyle = 5
	BezelStyleShadowlessSquare  BezelStyle = 6
	BezelStyleCircular          BezelStyle = 7
	BezelStyleTexturedSquare    BezelStyle = 8
	BezelStyleHelpButton        BezelStyle = 9
	BezelStyleSmallSquare       BezelStyle = 10
	BezelStyleTexturedRounded   BezelStyle = 11
	BezelStyleRoundRect         BezelStyle = 12
	BezelStyleRecessed          BezelStyle = 13
	BezelStyleRoundedDisclosure BezelStyle = 14
	BezelStyleInline            BezelStyle = 15
)

// ButtonType is button types that can be specified using setButtonType.
type ButtonType uint

const (
	ButtonTypeMomentaryLight        ButtonType = 0
	ButtonTypePushOnPushOff         ButtonType = 1
	ButtonTypeToggle                ButtonType = 2
	ButtonTypeSwitch                ButtonType = 3
	ButtonTypeRadio                 ButtonType = 4
	ButtonTypeMomentaryChange       ButtonType = 5
	ButtonTypeOnOff                 ButtonType = 6
	ButtonTypeMomentaryPushIn       ButtonType = 7
	ButtonTypeAccelerator           ButtonType = 8
	ButtonTypeMultiLevelAccelerator ButtonType = 9
)

// NewPlainButton return a new common used Button
func NewPlainButton(frame foundation.Rect) Button {
	btn := NewButton(frame)
	btn.SetBezelStyle(BezelStyleRounded)
	return btn
}

// NewCheckBox return a new common used switch Button
func NewCheckBox(frame foundation.Rect) Button {
	btn := NewButton(frame)
	btn.SetButtonType(ButtonTypeSwitch)
	return btn
}

// NewRadioButton return a new common used radio Button
func NewRadioButton(frame foundation.Rect) Button {
	btn := NewButton(frame)
	btn.SetButtonType(ButtonTypeRadio)
	return btn
}
