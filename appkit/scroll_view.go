package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include "scroll_view.h"
import "C"

import (
	"github.com/hsiafan/cocoa/foundation"
	"unsafe"
)

// ScrollView is a view that displays a portion of a document view and provides scroll bars that allow the user to move the document view within the scroll view
type ScrollView interface {
	View

	// HasVerticalScroller return whether the scroll view has a vertical scroller
	HasVerticalScroller() bool
	// SetHasVerticalScroller set whether the scroll view has a vertical scroller
	SetHasVerticalScroller(hasVerticalScroller bool)
	// HasHorizontalScroller return whether the scroll view has a horizontal scroller
	HasHorizontalScroller() bool
	// SetHasHorizontalScroller set whether the scroll view has a horizontal scroller
	SetHasHorizontalScroller(hasHorizontalScroller bool)
	// DocumentView return the view the scroll view scrolls within its content view
	DocumentView() View
	// SetDocumentView set the view the scroll view scrolls within its content view
	SetDocumentView(documentView View)
	// BorderType return the appearance of the scroll view’s border
	BorderType() BorderType
	// SetBorderType set the appearance of the scroll view’s border
	SetBorderType(borderType BorderType)
	// ContentSize return the size of the scroll view’s content view
	ContentSize() foundation.Size
}

var _ ScrollView = (*NSScrollView)(nil)

type NSScrollView struct {
	NSView
}

// MakeScrollView create a ScrollView from native pointer
func MakeScrollView(ptr unsafe.Pointer) *NSScrollView {
	if ptr == nil {
		return nil
	}
	return &NSScrollView{
		NSView: *MakeView(ptr),
	}
}

// NewScrollView create new ScrollView
func NewScrollView(frame foundation.Rect) ScrollView {
	id := resources.NextId()
	ptr := C.ScrollView_initWithFrame(C.long(id), toNSRect(frame))
	v := &NSScrollView{
		NSView: *MakeView(ptr),
	}
	resources.Store(id, v)
	foundation.AddDeallocHook(v, func() {
		resources.Delete(id)
	})
	return v
}

func (s *NSScrollView) HasVerticalScroller() bool {
	return bool(C.ScrollView_HasVerticalScroller(s.Ptr()))
}

func (s *NSScrollView) SetHasVerticalScroller(hasVerticalScroller bool) {
	C.ScrollView_SetHasVerticalScroller(s.Ptr(), C.bool(hasVerticalScroller))
}

func (s *NSScrollView) HasHorizontalScroller() bool {
	return bool(C.ScrollView_HasHorizontalScroller(s.Ptr()))
}

func (s *NSScrollView) SetHasHorizontalScroller(hasHorizontalScroller bool) {
	C.ScrollView_SetHasHorizontalScroller(s.Ptr(), C.bool(hasHorizontalScroller))
}

func (s *NSScrollView) DocumentView() View {
	return MakeView(C.ScrollView_DocumentView(s.Ptr()))
}

func (s *NSScrollView) SetDocumentView(documentView View) {
	C.ScrollView_SetDocumentView(s.Ptr(), toPointer(documentView))
}

func (s *NSScrollView) BorderType() BorderType {
	return BorderType(C.ScrollView_BorderType(s.Ptr()))
}

func (s *NSScrollView) SetBorderType(borderType BorderType) {
	C.ScrollView_SetBorderType(s.Ptr(), C.ulong(borderType))
}

func (s *NSScrollView) ContentSize() foundation.Size {
	return toSize(C.ScrollView_ContentSize(s.Ptr()))
}
