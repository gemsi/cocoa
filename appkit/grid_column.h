#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>

long GridColumn_NumberOfCells(void* ptr);
bool GridColumn_IsHidden(void* ptr);
void GridColumn_SetHidden(void* ptr, bool hidden);
double GridColumn_LeadingPadding(void* ptr);
void GridColumn_SetLeadingPadding(void* ptr, double leadingPadding);
double GridColumn_TrailingPadding(void* ptr);
void GridColumn_SetTrailingPadding(void* ptr, double trailingPadding);
double GridColumn_Width(void* ptr);
void GridColumn_SetWidth(void* ptr, double width);
long GridColumn_XPlacement(void* ptr);
void GridColumn_SetXPlacement(void* ptr, long xPlacement);
void* GridColumn_GridView(void* ptr);

void* GridColumn_CellAtIndex(void* ptr, long index);
void GridColumn_MergeCellsInRange(void* ptr, NSRange r);
