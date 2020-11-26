package appkit

import (
	"github.com/hsiafan/cocoa/foundation"
)

// NewVerticallyScrollableTextView create and return a ScrollView,
// which embed a TextView, and can scroll Vertically.
func NewVerticallyScrollableTextView(frame foundation.Rect) ScrollView {

	sv := NewScrollView(frame)
	svSize := sv.ContentSize()
	tv := NewTextView(foundation.MakeRect(0, 0, svSize.Width, svSize.Height))
	tv.SetMinSize(foundation.Size{0.0, svSize.Height})
	tv.SetMaxSize(foundation.Size{foundation.SIZE_MAX, foundation.SIZE_MAX})
	tv.SetVerticallyResizable(true)
	tv.SetHorizontallyResizable(false)
	tv.SetAutoresizingMask(ViewWidthSizable)
	tv.TextContainer().SetSize(foundation.Size{svSize.Width, foundation.SIZE_MAX})
	tv.TextContainer().SetWidthTracksTextView(true)
	sv.SetHasHorizontalScroller(false)
	sv.SetHasVerticalScroller(true)
	sv.SetBorderType(NoBorder)
	sv.SetDocumentView(tv)
	sv.SetAutoresizingMask(ViewWidthSizable | ViewHeightSizable)
	return sv
}

// NewScrollableTextView create and return a ScrollView,
// which embed a TextView, and can scroll both Vertically and Horizontally.
func NewScrollableTextView(frame foundation.Rect) ScrollView {
	sv := NewScrollView(frame)
	svSize := sv.ContentSize()
	tv := NewTextView(foundation.MakeRect(0, 0, svSize.Width, svSize.Height))
	tv.SetMinSize(foundation.Size{0.0, svSize.Height})
	tv.SetMaxSize(foundation.Size{foundation.SIZE_MAX, foundation.SIZE_MAX})
	tv.SetVerticallyResizable(true)
	tv.SetHorizontallyResizable(true)
	tv.SetAutoresizingMask(ViewWidthSizable | ViewHeightSizable)
	tv.TextContainer().SetSize(foundation.Size{foundation.SIZE_MAX, foundation.SIZE_MAX})
	tv.TextContainer().SetWidthTracksTextView(false)
	sv.SetHasHorizontalScroller(true)
	sv.SetHasVerticalScroller(true)
	sv.SetBorderType(NoBorder)
	sv.SetDocumentView(tv)
	sv.SetAutoresizingMask(ViewWidthSizable | ViewHeightSizable)
	return sv
}
