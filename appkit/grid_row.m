#import <AppKit/AppKit.h>
#import "grid_row.h"

long GridRow_NumberOfCells(void* ptr) {
	NSGridRow* gridRow = (NSGridRow*)ptr;
	return [gridRow numberOfCells];
}

bool GridRow_IsHidden(void* ptr) {
	NSGridRow* gridRow = (NSGridRow*)ptr;
	return [gridRow isHidden];
}

void GridRow_SetHidden(void* ptr, bool hidden) {
	NSGridRow* gridRow = (NSGridRow*)ptr;
	[gridRow setHidden:hidden];
}

double GridRow_TopPadding(void* ptr) {
	NSGridRow* gridRow = (NSGridRow*)ptr;
	return [gridRow topPadding];
}

void GridRow_SetTopPadding(void* ptr, double topPadding) {
	NSGridRow* gridRow = (NSGridRow*)ptr;
	[gridRow setTopPadding:topPadding];
}

double GridRow_BottomPadding(void* ptr) {
	NSGridRow* gridRow = (NSGridRow*)ptr;
	return [gridRow bottomPadding];
}

void GridRow_SetBottomPadding(void* ptr, double bottomPadding) {
	NSGridRow* gridRow = (NSGridRow*)ptr;
	[gridRow setBottomPadding:bottomPadding];
}

double GridRow_Height(void* ptr) {
	NSGridRow* gridRow = (NSGridRow*)ptr;
	return [gridRow height];
}

void GridRow_SetHeight(void* ptr, double height) {
	NSGridRow* gridRow = (NSGridRow*)ptr;
	[gridRow setHeight:height];
}

long GridRow_RowAlignment(void* ptr) {
	NSGridRow* gridRow = (NSGridRow*)ptr;
	return [gridRow rowAlignment];
}

void GridRow_SetRowAlignment(void* ptr, long rowAlignment) {
	NSGridRow* gridRow = (NSGridRow*)ptr;
	[gridRow setRowAlignment:rowAlignment];
}

long GridRow_YPlacement(void* ptr) {
	NSGridRow* gridRow = (NSGridRow*)ptr;
	return [gridRow yPlacement];
}

void GridRow_SetYPlacement(void* ptr, long yPlacement) {
	NSGridRow* gridRow = (NSGridRow*)ptr;
	[gridRow setYPlacement:yPlacement];
}

void* GridRow_GridView(void* ptr) {
	NSGridRow* gridRow = (NSGridRow*)ptr;
	return [gridRow gridView];
}

void* GridRow_CellAtIndex(void* ptr, long index) {
	NSGridRow* gridRow = (NSGridRow*)ptr;
	return [gridRow cellAtIndex:index];
}

void GridRow_MergeCellsInRange(void* ptr, NSRange r) {
	NSGridRow* gridRow = (NSGridRow*)ptr;
	[gridRow mergeCellsInRange:r];
}
