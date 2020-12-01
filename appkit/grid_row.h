#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>

long GridRow_NumberOfCells(void* ptr);
bool GridRow_IsHidden(void* ptr);
void GridRow_SetHidden(void* ptr, bool hidden);
double GridRow_TopPadding(void* ptr);
void GridRow_SetTopPadding(void* ptr, double topPadding);
double GridRow_BottomPadding(void* ptr);
void GridRow_SetBottomPadding(void* ptr, double bottomPadding);
double GridRow_Height(void* ptr);
void GridRow_SetHeight(void* ptr, double height);
long GridRow_RowAlignment(void* ptr);
void GridRow_SetRowAlignment(void* ptr, long rowAlignment);
long GridRow_YPlacement(void* ptr);
void GridRow_SetYPlacement(void* ptr, long yPlacement);
void* GridRow_GridView(void* ptr);

void* GridRow_CellAtIndex(void* ptr, long index);
void GridRow_MergeCellsInRange(void* ptr, NSRange r);
