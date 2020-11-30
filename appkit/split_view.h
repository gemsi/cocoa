#import <Foundation/NSGeometry.h>
#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>

bool SplitView_ArrangesAllSubviews(void* ptr);
void SplitView_SetArrangesAllSubviews(void* ptr, bool arrangesAllSubviews);
bool SplitView_IsVertical(void* ptr);
void SplitView_SetVertical(void* ptr, bool vertical);
void* SplitView_DividerColor(void* ptr);
double SplitView_DividerThickness(void* ptr);
long SplitView_DividerStyle(void* ptr);
void SplitView_SetDividerStyle(void* ptr, long dividerStyle);

void* SplitView_NewSplitView(NSRect frame);
double SplitView_MinPossiblePositionOfDividerAtIndex(void* ptr, long dividerIndex);
double SplitView_MaxPossiblePositionOfDividerAtIndex(void* ptr, long dividerIndex);
void SplitView_SetPosition(void* ptr, double position, long dividerIndex);
