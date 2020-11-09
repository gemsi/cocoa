package scrollview

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include "scroll_view.h"
import "C"

import (
	"github.com/hsiafan/cocoa/appkit/view"
	"github.com/hsiafan/cocoa/foundation/geometry"
	"github.com/hsiafan/cocoa/foundation/object"
	"github.com/hsiafan/cocoa/internal"
	"unsafe"
)

// A view that displays a portion of a document view and provides scroll bars that allow the user to move the document view within the scroll view.
type ScrollView interface {
	view.View

	// HasVerticalScroller return whether the scroll view has a vertical scroller
	HasVerticalScroller() bool
	// SetHasVerticalScroller set whether the scroll view has a vertical scroller
	SetHasVerticalScroller(value bool)
	// HasHorizontalScroller return whether the scroll view has a horizontal scroller
	HasHorizontalScroller() bool
	// SetHasHorizontalScroller set whether the scroll view has a horizontal scroller
	SetHasHorizontalScroller(value bool)
	// DocumentView return the view the scroll view scrolls within its content view
	DocumentView() view.View
	// SetDocumentView set the view the scroll view scrolls within its content view
	SetDocumentView(value view.View)
	// BorderType return the appearance of the scroll view’s border
	BorderType() view.BorderType
	// SetBorderType set the appearance of the scroll view’s border
	SetBorderType(value view.BorderType)
}

var _ ScrollView = (*NSScrollView)(nil)

type NSScrollView struct {
	view.NSView
}

func (s *NSScrollView) HasVerticalScroller() bool {
	return bool(C.ScrollView_HasVerticalScroller(s.Ptr()))
}

func (s *NSScrollView) SetHasVerticalScroller(value bool) {
	C.ScrollView_SetHasVerticalScroller(s.Ptr(), C.bool(value))
}

func (s *NSScrollView) HasHorizontalScroller() bool {
	return bool(C.ScrollView_HasHorizontalScroller(s.Ptr()))
}

func (s *NSScrollView) SetHasHorizontalScroller(value bool) {
	C.ScrollView_SetHasHorizontalScroller(s.Ptr(), C.bool(value))
}

func (s *NSScrollView) DocumentView() view.View {
	return view.Make(C.ScrollView_DocumentView(s.Ptr()))
}

func (s *NSScrollView) SetDocumentView(value view.View) {
	C.ScrollView_SetDocumentView(s.Ptr(), value.Ptr())
}

func (s *NSScrollView) BorderType() view.BorderType {
	return view.BorderType(C.ScrollView_BorderType(s.Ptr()))
}

func (s *NSScrollView) SetBorderType(value view.BorderType) {
	C.ScrollView_SetBorderType(s.Ptr(), C.ulong(value))
}

var resources = internal.NewResourceRegistry()

// New create new ScrollView
func New(frame geometry.Rect) ScrollView {
	id := resources.NextId()
	ptr := C.ScrollView_New(C.long(id), toNSRect(frame))

	v := &NSScrollView{
		NSView: *view.Make(ptr),
	}

	resources.Store(id, v)

	object.AddDeallocHook(v, func() {
		resources.Delete(id)
	})

	return v
}

func toNSRect(rect geometry.Rect) C.NSRect {
	return *(*C.NSRect)(unsafe.Pointer(&rect))
}
func toRect(nsRect C.NSRect) geometry.Rect {
	return *(*geometry.Rect)(unsafe.Pointer(&nsRect))
}
