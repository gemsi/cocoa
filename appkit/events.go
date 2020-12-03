package appkit

// EventModifierFlags is flags that represent key states in an event object
type EventModifierFlags uint

const (
	EventModifierFlagCapsLock   EventModifierFlags = 1 << 16 // Set if Caps Lock key is pressed.
	EventModifierFlagShift      EventModifierFlags = 1 << 17 // Set if Shift key is pressed.
	EventModifierFlagControl    EventModifierFlags = 1 << 18 // Set if Control key is pressed.
	EventModifierFlagOption     EventModifierFlags = 1 << 19 // Set if Option or Alternate key is pressed.
	EventModifierFlagCommand    EventModifierFlags = 1 << 20 // Set if Command key is pressed.
	EventModifierFlagNumericPad EventModifierFlags = 1 << 21 // Set if any key in the numeric keypad is pressed.
	EventModifierFlagHelp       EventModifierFlags = 1 << 22 // Set if the Help key is pressed.
	EventModifierFlagFunction   EventModifierFlags = 1 << 23 // Set if any function key is pressed.

	// Used to retrieve only the device-independent modifier flags, allowing applications to mask off the device-dependent modifier flags, including event coalescing information.
	EventModifierFlagDeviceIndependentFlagsMask EventModifierFlags = 0xffff0000
)
