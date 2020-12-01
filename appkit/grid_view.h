#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdlib.h>
#import <utils.h>

long GridView_NumberOfRows(void* ptr);
long GridView_NumberOfColumns(void* ptr);
double GridView_ColumnSpacing(void* ptr);
void GridView_SetColumnSpacing(void* ptr, double columnSpacing);
double GridView_RowSpacing(void* ptr);
void GridView_SetRowSpacing(void* ptr, double rowSpacing);
long GridView_RowAlignment(void* ptr);
void GridView_SetRowAlignment(void* ptr, long rowAlignment);
long GridView_XPlacement(void* ptr);
void GridView_SetXPlacement(void* ptr, long xPlacement);
long GridView_YPlacement(void* ptr);
void GridView_SetYPlacement(void* ptr, long yPlacement);

void* GridView_NewGridView(NSRect frameRect);
void* GridView_GridViewWithNumberOfColumns(long columnCount, long rowCount);
long GridView_IndexOfColumn(void* ptr, void* column);
long GridView_IndexOfRow(void* ptr, void* row);
void* GridView_ColumnAtIndex(void* ptr, long index);
void* GridView_RowAtIndex(void* ptr, long index);
void* GridView_AddRowWithViews(void* ptr, Array views);
void* GridView_InsertRowAtIndex(void* ptr, long index, Array views);
void GridView_RemoveRowAtIndex(void* ptr, long index);
void GridView_MoveRowAtIndex(void* ptr, long fromIndex, long toIndex);
void* GridView_AddColumnWithViews(void* ptr, Array views);
void* GridView_InsertColumnAtIndex(void* ptr, long index, Array views);
void GridView_RemoveColumnAtIndex(void* ptr, long index);
void GridView_MoveColumnAtIndex(void* ptr, long fromIndex, long toIndex);
void* GridView_CellAt(void* ptr, long columnIndex, long rowIndex);
void* GridView_CellForView(void* ptr, void* view);
void GridView_MergeCellsInHorizontalRange(void* ptr, NSRange hRange, NSRange vRange);
