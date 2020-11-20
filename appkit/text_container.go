package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include "text_container.h"
import "C"

import (
	"github.com/hsiafan/cocoa/foundation"
	"unsafe"
)

// TextContainer is a region where text layout occurs.
type TextContainer interface {
	foundation.Object

	// Size return the size of the text container’s bounding rectangle
	Size() foundation.Size
	// SetSize set the size of the text container’s bounding rectangle
	SetSize(size foundation.Size)
	// WidthTracksTextView return whether the text container adjusts the width of its bounding rectangle when its text view resizes
	WidthTracksTextView() bool
	// SetWidthTracksTextView set whether the text container adjusts the width of its bounding rectangle when its text view resizes
	SetWidthTracksTextView(widthTracksTextView bool)
	// HeightTracksTextView return whether the text container adjusts the height of its bounding rectangle when its text view resizes
	HeightTracksTextView() bool
	// SetHeightTracksTextView set whether the text container adjusts the height of its bounding rectangle when its text view resizes
	SetHeightTracksTextView(heightTracksTextView bool)
}

var _ TextContainer = (*NSTextContainer)(nil)

type NSTextContainer struct {
	foundation.NSObject
}

// MakeTextContainer create a TextContainer from native pointer
func MakeTextContainer(ptr unsafe.Pointer) *NSTextContainer {
	if ptr == nil {
		return nil
	}
	return &NSTextContainer{
		NSObject: *foundation.MakeObject(ptr),
	}
}

// NewTextContainer create new TextContainer
func NewTextContainer(size foundation.Size) TextContainer {
	id := resources.NextId()
	ptr := C.TextContainer_initWithSize(C.long(id), toNSSize(size))
	v := &NSTextContainer{
		NSObject: *foundation.MakeObject(ptr),
	}
	resources.Store(id, v)
	foundation.AddDeallocHook(v, func() {
		resources.Delete(id)
	})
	return v
}

func (t *NSTextContainer) Size() foundation.Size {
	return toSize(C.TextContainer_Size(t.Ptr()))
}

func (t *NSTextContainer) SetSize(size foundation.Size) {
	C.TextContainer_SetSize(t.Ptr(), toNSSize(size))
}

func (t *NSTextContainer) WidthTracksTextView() bool {
	return bool(C.TextContainer_WidthTracksTextView(t.Ptr()))
}

func (t *NSTextContainer) SetWidthTracksTextView(widthTracksTextView bool) {
	C.TextContainer_SetWidthTracksTextView(t.Ptr(), C.bool(widthTracksTextView))
}

func (t *NSTextContainer) HeightTracksTextView() bool {
	return bool(C.TextContainer_HeightTracksTextView(t.Ptr()))
}

func (t *NSTextContainer) SetHeightTracksTextView(heightTracksTextView bool) {
	C.TextContainer_SetHeightTracksTextView(t.Ptr(), C.bool(heightTracksTextView))
}
