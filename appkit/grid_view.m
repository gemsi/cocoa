#import <AppKit/AppKit.h>
#import "grid_view.h"

long GridView_NumberOfRows(void* ptr) {
	NSGridView* gridView = (NSGridView*)ptr;
	return [gridView numberOfRows];
}

long GridView_NumberOfColumns(void* ptr) {
	NSGridView* gridView = (NSGridView*)ptr;
	return [gridView numberOfColumns];
}

double GridView_ColumnSpacing(void* ptr) {
	NSGridView* gridView = (NSGridView*)ptr;
	return [gridView columnSpacing];
}

void GridView_SetColumnSpacing(void* ptr, double columnSpacing) {
	NSGridView* gridView = (NSGridView*)ptr;
	[gridView setColumnSpacing:columnSpacing];
}

double GridView_RowSpacing(void* ptr) {
	NSGridView* gridView = (NSGridView*)ptr;
	return [gridView rowSpacing];
}

void GridView_SetRowSpacing(void* ptr, double rowSpacing) {
	NSGridView* gridView = (NSGridView*)ptr;
	[gridView setRowSpacing:rowSpacing];
}

long GridView_RowAlignment(void* ptr) {
	NSGridView* gridView = (NSGridView*)ptr;
	return [gridView rowAlignment];
}

void GridView_SetRowAlignment(void* ptr, long rowAlignment) {
	NSGridView* gridView = (NSGridView*)ptr;
	[gridView setRowAlignment:rowAlignment];
}

long GridView_XPlacement(void* ptr) {
	NSGridView* gridView = (NSGridView*)ptr;
	return [gridView xPlacement];
}

void GridView_SetXPlacement(void* ptr, long xPlacement) {
	NSGridView* gridView = (NSGridView*)ptr;
	[gridView setXPlacement:xPlacement];
}

long GridView_YPlacement(void* ptr) {
	NSGridView* gridView = (NSGridView*)ptr;
	return [gridView yPlacement];
}

void GridView_SetYPlacement(void* ptr, long yPlacement) {
	NSGridView* gridView = (NSGridView*)ptr;
	[gridView setYPlacement:yPlacement];
}

void* GridView_NewGridView(NSRect frameRect) {
	NSGridView* gridView = [NSGridView alloc];
	return [[gridView initWithFrame:frameRect] autorelease];
}

void* GridView_GridViewWithNumberOfColumns(long columnCount, long rowCount) {
	return [NSGridView gridViewWithNumberOfColumns:columnCount rows:rowCount];
}

long GridView_IndexOfColumn(void* ptr, void* column) {
	NSGridView* gridView = (NSGridView*)ptr;
	return [gridView indexOfColumn:(NSGridColumn*)column];
}

long GridView_IndexOfRow(void* ptr, void* row) {
	NSGridView* gridView = (NSGridView*)ptr;
	return [gridView indexOfRow:(NSGridRow*)row];
}

void* GridView_ColumnAtIndex(void* ptr, long index) {
	NSGridView* gridView = (NSGridView*)ptr;
	return [gridView columnAtIndex:index];
}

void* GridView_RowAtIndex(void* ptr, long index) {
	NSGridView* gridView = (NSGridView*)ptr;
	return [gridView rowAtIndex:index];
}

void* GridView_AddRowWithViews(void* ptr, Array views) {
	NSGridView* gridView = (NSGridView*)ptr;
    NSMutableArray* objcViews = [[NSMutableArray alloc] init];
    void** viewsData = (void**)views.data;
    for (int i = 0; i < views.len; i++) {
    	[objcViews addObject:(NSView*)viewsData[i]];
    }
	return [gridView addRowWithViews:objcViews];
}

void* GridView_InsertRowAtIndex(void* ptr, long index, Array views) {
	NSGridView* gridView = (NSGridView*)ptr;
    NSMutableArray* objcViews = [[NSMutableArray alloc] init];
    void** viewsData = (void**)views.data;
    for (int i = 0; i < views.len; i++) {
    	[objcViews addObject:(NSView*)viewsData[i]];
    }
	return [gridView insertRowAtIndex:index withViews:objcViews];
}

void GridView_RemoveRowAtIndex(void* ptr, long index) {
	NSGridView* gridView = (NSGridView*)ptr;
	[gridView removeRowAtIndex:index];
}

void GridView_MoveRowAtIndex(void* ptr, long fromIndex, long toIndex) {
	NSGridView* gridView = (NSGridView*)ptr;
	[gridView moveRowAtIndex:fromIndex toIndex:toIndex];
}

void* GridView_AddColumnWithViews(void* ptr, Array views) {
	NSGridView* gridView = (NSGridView*)ptr;
    NSMutableArray* objcViews = [[NSMutableArray alloc] init];
    void** viewsData = (void**)views.data;
    for (int i = 0; i < views.len; i++) {
    	[objcViews addObject:(NSView*)viewsData[i]];
    }
	return [gridView addColumnWithViews:objcViews];
}

void* GridView_InsertColumnAtIndex(void* ptr, long index, Array views) {
	NSGridView* gridView = (NSGridView*)ptr;
    NSMutableArray* objcViews = [[NSMutableArray alloc] init];
    void** viewsData = (void**)views.data;
    for (int i = 0; i < views.len; i++) {
    	[objcViews addObject:(NSView*)viewsData[i]];
    }
	return [gridView insertColumnAtIndex:index withViews:objcViews];
}

void GridView_RemoveColumnAtIndex(void* ptr, long index) {
	NSGridView* gridView = (NSGridView*)ptr;
	[gridView removeColumnAtIndex:index];
}

void GridView_MoveColumnAtIndex(void* ptr, long fromIndex, long toIndex) {
	NSGridView* gridView = (NSGridView*)ptr;
	[gridView moveColumnAtIndex:fromIndex toIndex:toIndex];
}

void* GridView_CellAt(void* ptr, long columnIndex, long rowIndex) {
	NSGridView* gridView = (NSGridView*)ptr;
	return [gridView cellAtColumnIndex:columnIndex rowIndex:rowIndex];
}

void* GridView_CellForView(void* ptr, void* view) {
	NSGridView* gridView = (NSGridView*)ptr;
	return [gridView cellForView:(NSView*)view];
}

void GridView_MergeCellsInHorizontalRange(void* ptr, NSRange hRange, NSRange vRange) {
	NSGridView* gridView = (NSGridView*)ptr;
	[gridView mergeCellsInHorizontalRange:hRange verticalRange:vRange];
}
