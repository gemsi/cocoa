package appkit

import "github.com/hsiafan/cocoa/foundation"

// NewLabel create a text field, which looks like a Label
func NewLabel(frame foundation.Rect) TextField {
	tf := NewTextField(frame)
	tf.SetBezeled(false)
	tf.SetDrawsBackground(false)
	tf.SetEditable(false)
	tf.SetSelectable(false)
	return tf
}

// NewPlainTextField return a plain TextField
func NewPlainTextField(frame foundation.Rect) TextField {
	field := NewTextField(frame)
	cell := field.Cell()
	cell.SetScrollable(true)
	cell.SetWraps(false)
	return field
}

// NewPlainSecureTextField return a plain SecureTextField
func NewPlainSecureTextField(frame foundation.Rect) SecureTextField {
	field := NewSecureTextField(frame)
	cell := field.Cell()
	cell.SetScrollable(true)
	cell.SetWraps(false)
	return field
}
