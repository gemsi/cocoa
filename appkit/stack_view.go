package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include "stack_view.h"
import "C"

import (
	"github.com/hsiafan/cocoa/foundation"
	"unsafe"
)

// StackView is a view that arranges an array of views horizontally or vertically and updates their placement and sizing when the window size changes
type StackView interface {
	View

	// Views return the array of views owned by the stack view
	Views() []View
}

var _ StackView = (*NSStackView)(nil)

type NSStackView struct {
	NSView
}

// MakeStackView create a StackView from native pointer
func MakeStackView(ptr unsafe.Pointer) *NSStackView {
	if ptr == nil {
		return nil
	}
	return &NSStackView{
		NSView: *MakeView(ptr),
	}
}

// NewStackViewWithViews create new StackView
func NewStackViewWithViews(views []View) StackView {
	cViews := make([]unsafe.Pointer, len(views))
	for idx, v := range views {
		cViews[idx] = unsafe.Pointer(v.Ptr())
	}
	id := resources.NextId()
	ptr := C.StackView_stackViewWithViews(C.long(id), &cViews[0], C.size_t(len(views)))
	v := &NSStackView{
		NSView: *MakeView(ptr),
	}
	resources.Store(id, v)
	foundation.AddDeallocHook(v, func() {
		resources.Delete(id)
	})
	return v
}

func (s *NSStackView) Views() []View {
	var cArray C.Array = C.StackView_Views(s.Ptr())
	defer C.free(cArray.data)
	result := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(cArray.data))[:cArray.len:cArray.len]
	var goResult = make([]View, len(result))
	for idx, r := range result {
		goResult[idx] = MakeView(r)
	}
	return goResult
}
