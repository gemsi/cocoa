package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include "appearance.h"
import "C"

import (
	"github.com/hsiafan/cocoa/foundation"
	"unsafe"
)

// Appearance is an object that manages standard appearance attributes for UI elements in an app
type Appearance interface {
	foundation.Object

	// Name return the name of the appearance
	Name() AppearanceName
}

var _ Appearance = (*NSAppearance)(nil)

type NSAppearance struct {
	foundation.NSObject
}

// MakeAppearance create a Appearance from native pointer
func MakeAppearance(ptr unsafe.Pointer) *NSAppearance {
	if ptr == nil {
		return nil
	}
	return &NSAppearance{
		NSObject: *foundation.MakeObject(ptr),
	}
}

func (a *NSAppearance) Name() AppearanceName {
	return AppearanceName(C.GoString(C.Appearance_Name(a.Ptr())))
}

func AppearanceNamed(name AppearanceName) Appearance {
	cName := C.CString(string(name))
	defer C.free(unsafe.Pointer(cName))
	return MakeAppearance(C.Appearance_AppearanceNamed(cName))
}

type AppearanceName string

var AppearanceNameAqua AppearanceName
var AppearanceNameDarkAqua AppearanceName
var AppearanceNameVibrantLight AppearanceName
var AppearanceNameVibrantDark AppearanceName
var AppearanceNameAccessibilityHighContrastAqua AppearanceName
var AppearanceNameAccessibilityHighContrastDarkAqua AppearanceName
var AppearanceNameAccessibilityHighContrastVibrantLight AppearanceName
var AppearanceNameAccessibilityHighContrastVibrantDark AppearanceName
