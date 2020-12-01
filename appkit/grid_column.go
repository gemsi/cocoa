package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include "grid_column.h"
import "C"

import (
	"github.com/hsiafan/cocoa/foundation"
	"unsafe"
)

// GridColumn is a row within a grid view
type GridColumn interface {
	foundation.Object

	// NumberOfCells return
	NumberOfCells() int
	// IsHidden return
	IsHidden() bool
	// SetHidden set
	SetHidden(hidden bool)
	// LeadingPadding return
	LeadingPadding() float64
	// SetLeadingPadding set
	SetLeadingPadding(leadingPadding float64)
	// TrailingPadding return
	TrailingPadding() float64
	// SetTrailingPadding set
	SetTrailingPadding(trailingPadding float64)
	// Width return
	Width() float64
	// SetWidth set
	SetWidth(width float64)
	// XPlacement return
	XPlacement() GridCellPlacement
	// SetXPlacement set
	SetXPlacement(xPlacement GridCellPlacement)
	// GridView return
	GridView() GridView
	// CellAtIndex
	CellAtIndex(index int) GridCell
	// MergeCellsInRange
	MergeCellsInRange(r foundation.Range)
}

var _ GridColumn = (*NSGridColumn)(nil)

type NSGridColumn struct {
	foundation.NSObject
}

// MakeGridColumn create a GridColumn from native pointer
func MakeGridColumn(ptr unsafe.Pointer) *NSGridColumn {
	if ptr == nil {
		return nil
	}
	return &NSGridColumn{
		NSObject: *foundation.MakeObject(ptr),
	}
}

func (g *NSGridColumn) NumberOfCells() int {
	return int(C.GridColumn_NumberOfCells(g.Ptr()))
}

func (g *NSGridColumn) IsHidden() bool {
	return bool(C.GridColumn_IsHidden(g.Ptr()))
}

func (g *NSGridColumn) SetHidden(hidden bool) {
	C.GridColumn_SetHidden(g.Ptr(), C.bool(hidden))
}

func (g *NSGridColumn) LeadingPadding() float64 {
	return float64(C.GridColumn_LeadingPadding(g.Ptr()))
}

func (g *NSGridColumn) SetLeadingPadding(leadingPadding float64) {
	C.GridColumn_SetLeadingPadding(g.Ptr(), C.double(leadingPadding))
}

func (g *NSGridColumn) TrailingPadding() float64 {
	return float64(C.GridColumn_TrailingPadding(g.Ptr()))
}

func (g *NSGridColumn) SetTrailingPadding(trailingPadding float64) {
	C.GridColumn_SetTrailingPadding(g.Ptr(), C.double(trailingPadding))
}

func (g *NSGridColumn) Width() float64 {
	return float64(C.GridColumn_Width(g.Ptr()))
}

func (g *NSGridColumn) SetWidth(width float64) {
	C.GridColumn_SetWidth(g.Ptr(), C.double(width))
}

func (g *NSGridColumn) XPlacement() GridCellPlacement {
	return GridCellPlacement(C.GridColumn_XPlacement(g.Ptr()))
}

func (g *NSGridColumn) SetXPlacement(xPlacement GridCellPlacement) {
	C.GridColumn_SetXPlacement(g.Ptr(), C.long(xPlacement))
}

func (g *NSGridColumn) GridView() GridView {
	return MakeGridView(C.GridColumn_GridView(g.Ptr()))
}

func (g *NSGridColumn) CellAtIndex(index int) GridCell {
	return MakeGridCell(C.GridColumn_CellAtIndex(g.Ptr(), C.long(index)))
}

func (g *NSGridColumn) MergeCellsInRange(r foundation.Range) {
	C.GridColumn_MergeCellsInRange(g.Ptr(), toNSRange(r))
}
