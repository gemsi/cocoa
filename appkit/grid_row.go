package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include "grid_row.h"
import "C"

import (
	"github.com/hsiafan/cocoa/foundation"
	"unsafe"
)

// GridRow is a row within a grid view
type GridRow interface {
	foundation.Object

	// NumberOfCells return
	NumberOfCells() int
	// IsHidden return
	IsHidden() bool
	// SetHidden set
	SetHidden(hidden bool)
	// TopPadding return
	TopPadding() float64
	// SetTopPadding set
	SetTopPadding(topPadding float64)
	// BottomPadding return
	BottomPadding() float64
	// SetBottomPadding set
	SetBottomPadding(bottomPadding float64)
	// Height return
	Height() float64
	// SetHeight set
	SetHeight(height float64)
	// RowAlignment return
	RowAlignment() GridRowAlignment
	// SetRowAlignment set
	SetRowAlignment(rowAlignment GridRowAlignment)
	// YPlacement return
	YPlacement() GridCellPlacement
	// SetYPlacement set
	SetYPlacement(yPlacement GridCellPlacement)
	// GridView return
	GridView() GridView
	// CellAtIndex
	CellAtIndex(index int) GridCell
	// MergeCellsInRange
	MergeCellsInRange(r foundation.Range)
}

var _ GridRow = (*NSGridRow)(nil)

type NSGridRow struct {
	foundation.NSObject
}

// MakeGridRow create a GridRow from native pointer
func MakeGridRow(ptr unsafe.Pointer) *NSGridRow {
	if ptr == nil {
		return nil
	}
	return &NSGridRow{
		NSObject: *foundation.MakeObject(ptr),
	}
}

func (g *NSGridRow) NumberOfCells() int {
	return int(C.GridRow_NumberOfCells(g.Ptr()))
}

func (g *NSGridRow) IsHidden() bool {
	return bool(C.GridRow_IsHidden(g.Ptr()))
}

func (g *NSGridRow) SetHidden(hidden bool) {
	C.GridRow_SetHidden(g.Ptr(), C.bool(hidden))
}

func (g *NSGridRow) TopPadding() float64 {
	return float64(C.GridRow_TopPadding(g.Ptr()))
}

func (g *NSGridRow) SetTopPadding(topPadding float64) {
	C.GridRow_SetTopPadding(g.Ptr(), C.double(topPadding))
}

func (g *NSGridRow) BottomPadding() float64 {
	return float64(C.GridRow_BottomPadding(g.Ptr()))
}

func (g *NSGridRow) SetBottomPadding(bottomPadding float64) {
	C.GridRow_SetBottomPadding(g.Ptr(), C.double(bottomPadding))
}

func (g *NSGridRow) Height() float64 {
	return float64(C.GridRow_Height(g.Ptr()))
}

func (g *NSGridRow) SetHeight(height float64) {
	C.GridRow_SetHeight(g.Ptr(), C.double(height))
}

func (g *NSGridRow) RowAlignment() GridRowAlignment {
	return GridRowAlignment(C.GridRow_RowAlignment(g.Ptr()))
}

func (g *NSGridRow) SetRowAlignment(rowAlignment GridRowAlignment) {
	C.GridRow_SetRowAlignment(g.Ptr(), C.long(rowAlignment))
}

func (g *NSGridRow) YPlacement() GridCellPlacement {
	return GridCellPlacement(C.GridRow_YPlacement(g.Ptr()))
}

func (g *NSGridRow) SetYPlacement(yPlacement GridCellPlacement) {
	C.GridRow_SetYPlacement(g.Ptr(), C.long(yPlacement))
}

func (g *NSGridRow) GridView() GridView {
	return MakeGridView(C.GridRow_GridView(g.Ptr()))
}

func (g *NSGridRow) CellAtIndex(index int) GridCell {
	return MakeGridCell(C.GridRow_CellAtIndex(g.Ptr(), C.long(index)))
}

func (g *NSGridRow) MergeCellsInRange(r foundation.Range) {
	C.GridRow_MergeCellsInRange(g.Ptr(), toNSRange(r))
}
