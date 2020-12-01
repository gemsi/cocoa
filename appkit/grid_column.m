#import <Appkit/NSGridView.h>
#import "grid_column.h"

long GridColumn_NumberOfCells(void* ptr) {
	NSGridColumn* gridColumn = (NSGridColumn*)ptr;
	return [gridColumn numberOfCells];
}

bool GridColumn_IsHidden(void* ptr) {
	NSGridColumn* gridColumn = (NSGridColumn*)ptr;
	return [gridColumn isHidden];
}

void GridColumn_SetHidden(void* ptr, bool hidden) {
	NSGridColumn* gridColumn = (NSGridColumn*)ptr;
	[gridColumn setHidden:hidden];
}

double GridColumn_LeadingPadding(void* ptr) {
	NSGridColumn* gridColumn = (NSGridColumn*)ptr;
	return [gridColumn leadingPadding];
}

void GridColumn_SetLeadingPadding(void* ptr, double leadingPadding) {
	NSGridColumn* gridColumn = (NSGridColumn*)ptr;
	[gridColumn setLeadingPadding:leadingPadding];
}

double GridColumn_TrailingPadding(void* ptr) {
	NSGridColumn* gridColumn = (NSGridColumn*)ptr;
	return [gridColumn trailingPadding];
}

void GridColumn_SetTrailingPadding(void* ptr, double trailingPadding) {
	NSGridColumn* gridColumn = (NSGridColumn*)ptr;
	[gridColumn setTrailingPadding:trailingPadding];
}

double GridColumn_Width(void* ptr) {
	NSGridColumn* gridColumn = (NSGridColumn*)ptr;
	return [gridColumn width];
}

void GridColumn_SetWidth(void* ptr, double width) {
	NSGridColumn* gridColumn = (NSGridColumn*)ptr;
	[gridColumn setWidth:width];
}

long GridColumn_XPlacement(void* ptr) {
	NSGridColumn* gridColumn = (NSGridColumn*)ptr;
	return [gridColumn xPlacement];
}

void GridColumn_SetXPlacement(void* ptr, long xPlacement) {
	NSGridColumn* gridColumn = (NSGridColumn*)ptr;
	[gridColumn setXPlacement:xPlacement];
}

void* GridColumn_GridView(void* ptr) {
	NSGridColumn* gridColumn = (NSGridColumn*)ptr;
	return [gridColumn gridView];
}

void* GridColumn_CellAtIndex(void* ptr, long index) {
	NSGridColumn* gridColumn = (NSGridColumn*)ptr;
	return [gridColumn cellAtIndex:index];
}

void GridColumn_MergeCellsInRange(void* ptr, NSRange r) {
	NSGridColumn* gridColumn = (NSGridColumn*)ptr;
	[gridColumn mergeCellsInRange:r];
}
