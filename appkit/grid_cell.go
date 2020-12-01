package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include "grid_cell.h"
import "C"

import (
	"github.com/hsiafan/cocoa/foundation"
	"unsafe"
)

// GridCell is an individual content area within a grid view, typically at the intersection of a row and a column
type GridCell interface {
	foundation.Object

	// Column return
	Column() GridColumn
	// Row return
	Row() GridRow
	// ContentView return
	ContentView() View
	// CustomPlacementConstraints return
	CustomPlacementConstraints() []LayoutConstraint
	// SetCustomPlacementConstraints set
	SetCustomPlacementConstraints(customPlacementConstraints []LayoutConstraint)
	// RowAlignment return
	RowAlignment() GridRowAlignment
	// SetRowAlignment set
	SetRowAlignment(rowAlignment GridRowAlignment)
	// XPlacement return
	XPlacement() GridCellPlacement
	// SetXPlacement set
	SetXPlacement(xPlacement GridCellPlacement)
	// YPlacement return
	YPlacement() GridCellPlacement
	// SetYPlacement set
	SetYPlacement(yPlacement GridCellPlacement)
}

var _ GridCell = (*NSGridCell)(nil)

type NSGridCell struct {
	foundation.NSObject
}

// MakeGridCell create a GridCell from native pointer
func MakeGridCell(ptr unsafe.Pointer) *NSGridCell {
	if ptr == nil {
		return nil
	}
	return &NSGridCell{
		NSObject: *foundation.MakeObject(ptr),
	}
}

func (g *NSGridCell) Column() GridColumn {
	return MakeGridColumn(C.GridCell_Column(g.Ptr()))
}

func (g *NSGridCell) Row() GridRow {
	return MakeGridRow(C.GridCell_Row(g.Ptr()))
}

func (g *NSGridCell) ContentView() View {
	return MakeView(C.GridCell_ContentView(g.Ptr()))
}

func (g *NSGridCell) CustomPlacementConstraints() []LayoutConstraint {
	var cArray C.Array = C.GridCell_CustomPlacementConstraints(g.Ptr())
	defer C.free(cArray.data)
	result := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(cArray.data))[:cArray.len:cArray.len]
	var goResult = make([]LayoutConstraint, len(result))
	for idx, r := range result {
		goResult[idx] = MakeLayoutConstraint(r)
	}
	return goResult
}

func (g *NSGridCell) SetCustomPlacementConstraints(customPlacementConstraints []LayoutConstraint) {
	cCustomPlacementConstraintsData := make([]unsafe.Pointer, len(customPlacementConstraints))
	for idx, v := range customPlacementConstraints {
		cCustomPlacementConstraintsData[idx] = v.Ptr()
	}
	cCustomPlacementConstraints := C.Array{data: unsafe.Pointer(&cCustomPlacementConstraintsData[0]), len: C.int(len(customPlacementConstraints))}
	C.GridCell_SetCustomPlacementConstraints(g.Ptr(), cCustomPlacementConstraints)
}

func (g *NSGridCell) RowAlignment() GridRowAlignment {
	return GridRowAlignment(C.GridCell_RowAlignment(g.Ptr()))
}

func (g *NSGridCell) SetRowAlignment(rowAlignment GridRowAlignment) {
	C.GridCell_SetRowAlignment(g.Ptr(), C.long(rowAlignment))
}

func (g *NSGridCell) XPlacement() GridCellPlacement {
	return GridCellPlacement(C.GridCell_XPlacement(g.Ptr()))
}

func (g *NSGridCell) SetXPlacement(xPlacement GridCellPlacement) {
	C.GridCell_SetXPlacement(g.Ptr(), C.long(xPlacement))
}

func (g *NSGridCell) YPlacement() GridCellPlacement {
	return GridCellPlacement(C.GridCell_YPlacement(g.Ptr()))
}

func (g *NSGridCell) SetYPlacement(yPlacement GridCellPlacement) {
	C.GridCell_SetYPlacement(g.Ptr(), C.long(yPlacement))
}
