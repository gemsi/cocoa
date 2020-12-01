#import <stdlib.h>
#import <utils.h>

void* GridCell_Column(void* ptr);
void* GridCell_Row(void* ptr);
void* GridCell_ContentView(void* ptr);
Array GridCell_CustomPlacementConstraints(void* ptr);
void GridCell_SetCustomPlacementConstraints(void* ptr, Array customPlacementConstraints);
long GridCell_RowAlignment(void* ptr);
void GridCell_SetRowAlignment(void* ptr, long rowAlignment);
long GridCell_XPlacement(void* ptr);
void GridCell_SetXPlacement(void* ptr, long xPlacement);
long GridCell_YPlacement(void* ptr);
void GridCell_SetYPlacement(void* ptr, long yPlacement);

