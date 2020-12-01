package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include "grid_view.h"
import "C"

import (
	"github.com/hsiafan/cocoa/foundation"
	"unsafe"
)

// GridView is a container that aligns views in a flexible grid of rows and columns
type GridView interface {
	View

	// NumberOfRows return the number of rows in the grid view
	NumberOfRows() int
	// NumberOfColumns return the number of columns in the grid view
	NumberOfColumns() int
	// ColumnSpacing return the column spacing for the grid view
	ColumnSpacing() float64
	// SetColumnSpacing set the column spacing for the grid view
	SetColumnSpacing(columnSpacing float64)
	// RowSpacing return the row spacing for the grid view
	RowSpacing() float64
	// SetRowSpacing set the row spacing for the grid view
	SetRowSpacing(rowSpacing float64)
	// RowAlignment return the row alignment for the grid view
	RowAlignment() GridRowAlignment
	// SetRowAlignment set the row alignment for the grid view
	SetRowAlignment(rowAlignment GridRowAlignment)
	// XPlacement return The placement of the cell within the grid column
	XPlacement() GridCellPlacement
	// SetXPlacement set The placement of the cell within the grid column
	SetXPlacement(xPlacement GridCellPlacement)
	// YPlacement return The placement of the cell within the grid row
	YPlacement() GridCellPlacement
	// SetYPlacement set The placement of the cell within the grid row
	SetYPlacement(yPlacement GridCellPlacement)
	// IndexOfColumn returns the index of the specified grid column
	IndexOfColumn(column GridColumn) int
	// IndexOfRow returns the index of the specified grid row
	IndexOfRow(row GridRow) int
	// ColumnAtIndex returns the grid column object at the specified index
	ColumnAtIndex(index int) GridColumn
	// RowAtIndex returns the grid row object at the specified index
	RowAtIndex(index int) GridRow
	// AddRowWithViews adds an array of views to a new row
	AddRowWithViews(views []View) GridRow
	// InsertRowAtIndex inserts the array of view objects into the grid view at the index
	InsertRowAtIndex(index int, views []View) GridRow
	// RemoveRowAtIndex removes the row from the grid view at the index
	RemoveRowAtIndex(index int)
	// MoveRowAtIndex moves the specified row to the new row location
	MoveRowAtIndex(fromIndex int, toIndex int)
	// AddColumnWithViews adds an array of views to a new column
	AddColumnWithViews(views []View) GridColumn
	// InsertColumnAtIndex inserts the array of view objects into the grid view at the index
	InsertColumnAtIndex(index int, views []View) GridColumn
	// RemoveColumnAtIndex removes the column from the grid view at the index
	RemoveColumnAtIndex(index int)
	// MoveColumnAtIndex moves the specified column to the new column location
	MoveColumnAtIndex(fromIndex int, toIndex int)
	// CellAt returns the grid cell object at the specified column and row index
	CellAt(columnIndex int, rowIndex int) GridCell
	// CellForView returns the grid cell object that contains the given view or one of its ancestors
	CellForView(view View) GridCell
	// MergeCellsInHorizontalRange expands the cell at the top-leading corner of the horizontal and vertical range to cover the entire area
	MergeCellsInHorizontalRange(hRange foundation.Range, vRange foundation.Range)
}

var _ GridView = (*NSGridView)(nil)

type NSGridView struct {
	NSView
}

// MakeGridView create a GridView from native pointer
func MakeGridView(ptr unsafe.Pointer) *NSGridView {
	if ptr == nil {
		return nil
	}
	return &NSGridView{
		NSView: *MakeView(ptr),
	}
}

func (g *NSGridView) NumberOfRows() int {
	return int(C.GridView_NumberOfRows(g.Ptr()))
}

func (g *NSGridView) NumberOfColumns() int {
	return int(C.GridView_NumberOfColumns(g.Ptr()))
}

func (g *NSGridView) ColumnSpacing() float64 {
	return float64(C.GridView_ColumnSpacing(g.Ptr()))
}

func (g *NSGridView) SetColumnSpacing(columnSpacing float64) {
	C.GridView_SetColumnSpacing(g.Ptr(), C.double(columnSpacing))
}

func (g *NSGridView) RowSpacing() float64 {
	return float64(C.GridView_RowSpacing(g.Ptr()))
}

func (g *NSGridView) SetRowSpacing(rowSpacing float64) {
	C.GridView_SetRowSpacing(g.Ptr(), C.double(rowSpacing))
}

func (g *NSGridView) RowAlignment() GridRowAlignment {
	return GridRowAlignment(C.GridView_RowAlignment(g.Ptr()))
}

func (g *NSGridView) SetRowAlignment(rowAlignment GridRowAlignment) {
	C.GridView_SetRowAlignment(g.Ptr(), C.long(rowAlignment))
}

func (g *NSGridView) XPlacement() GridCellPlacement {
	return GridCellPlacement(C.GridView_XPlacement(g.Ptr()))
}

func (g *NSGridView) SetXPlacement(xPlacement GridCellPlacement) {
	C.GridView_SetXPlacement(g.Ptr(), C.long(xPlacement))
}

func (g *NSGridView) YPlacement() GridCellPlacement {
	return GridCellPlacement(C.GridView_YPlacement(g.Ptr()))
}

func (g *NSGridView) SetYPlacement(yPlacement GridCellPlacement) {
	C.GridView_SetYPlacement(g.Ptr(), C.long(yPlacement))
}

func NewGridView(frameRect foundation.Rect) GridView {
	return MakeGridView(C.GridView_NewGridView(toNSRect(frameRect)))
}

func GridViewWithNumberOfColumns(columnCount int, rowCount int) GridView {
	return MakeGridView(C.GridView_GridViewWithNumberOfColumns(C.long(columnCount), C.long(rowCount)))
}

func (g *NSGridView) IndexOfColumn(column GridColumn) int {
	return int(C.GridView_IndexOfColumn(g.Ptr(), toPointer(column)))
}

func (g *NSGridView) IndexOfRow(row GridRow) int {
	return int(C.GridView_IndexOfRow(g.Ptr(), toPointer(row)))
}

func (g *NSGridView) ColumnAtIndex(index int) GridColumn {
	return MakeGridColumn(C.GridView_ColumnAtIndex(g.Ptr(), C.long(index)))
}

func (g *NSGridView) RowAtIndex(index int) GridRow {
	return MakeGridRow(C.GridView_RowAtIndex(g.Ptr(), C.long(index)))
}

func (g *NSGridView) AddRowWithViews(views []View) GridRow {
	cViewsData := make([]unsafe.Pointer, len(views))
	for idx, v := range views {
		cViewsData[idx] = v.Ptr()
	}
	cViews := C.Array{data: unsafe.Pointer(&cViewsData[0]), len: C.int(len(views))}
	return MakeGridRow(C.GridView_AddRowWithViews(g.Ptr(), cViews))
}

func (g *NSGridView) InsertRowAtIndex(index int, views []View) GridRow {
	cViewsData := make([]unsafe.Pointer, len(views))
	for idx, v := range views {
		cViewsData[idx] = v.Ptr()
	}
	cViews := C.Array{data: unsafe.Pointer(&cViewsData[0]), len: C.int(len(views))}
	return MakeGridRow(C.GridView_InsertRowAtIndex(g.Ptr(), C.long(index), cViews))
}

func (g *NSGridView) RemoveRowAtIndex(index int) {
	C.GridView_RemoveRowAtIndex(g.Ptr(), C.long(index))
}

func (g *NSGridView) MoveRowAtIndex(fromIndex int, toIndex int) {
	C.GridView_MoveRowAtIndex(g.Ptr(), C.long(fromIndex), C.long(toIndex))
}

func (g *NSGridView) AddColumnWithViews(views []View) GridColumn {
	cViewsData := make([]unsafe.Pointer, len(views))
	for idx, v := range views {
		cViewsData[idx] = v.Ptr()
	}
	cViews := C.Array{data: unsafe.Pointer(&cViewsData[0]), len: C.int(len(views))}
	return MakeGridColumn(C.GridView_AddColumnWithViews(g.Ptr(), cViews))
}

func (g *NSGridView) InsertColumnAtIndex(index int, views []View) GridColumn {
	cViewsData := make([]unsafe.Pointer, len(views))
	for idx, v := range views {
		cViewsData[idx] = v.Ptr()
	}
	cViews := C.Array{data: unsafe.Pointer(&cViewsData[0]), len: C.int(len(views))}
	return MakeGridColumn(C.GridView_InsertColumnAtIndex(g.Ptr(), C.long(index), cViews))
}

func (g *NSGridView) RemoveColumnAtIndex(index int) {
	C.GridView_RemoveColumnAtIndex(g.Ptr(), C.long(index))
}

func (g *NSGridView) MoveColumnAtIndex(fromIndex int, toIndex int) {
	C.GridView_MoveColumnAtIndex(g.Ptr(), C.long(fromIndex), C.long(toIndex))
}

func (g *NSGridView) CellAt(columnIndex int, rowIndex int) GridCell {
	return MakeGridCell(C.GridView_CellAt(g.Ptr(), C.long(columnIndex), C.long(rowIndex)))
}

func (g *NSGridView) CellForView(view View) GridCell {
	return MakeGridCell(C.GridView_CellForView(g.Ptr(), toPointer(view)))
}

func (g *NSGridView) MergeCellsInHorizontalRange(hRange foundation.Range, vRange foundation.Range) {
	C.GridView_MergeCellsInHorizontalRange(g.Ptr(), toNSRange(hRange), toNSRange(vRange))
}
