#import <Foundation/NSGeometry.h>
#import <stdbool.h>
#import <stdlib.h>

void* TextField_initWithFrame(long goID, NSRect frame);

bool TextField_IsBezeled(void* ptr);
void TextField_SetBezeled(void* ptr, bool bezeled);
bool TextField_DrawsBackground(void* ptr);
void TextField_SetDrawsBackground(void* ptr, bool drawsBackground);
bool TextField_IsEditable(void* ptr);
void TextField_SetEditable(void* ptr, bool editable);
bool TextField_IsSelectable(void* ptr);
void TextField_SetSelectable(void* ptr, bool selectable);
void* TextField_TextColor(void* ptr);
void TextField_SetTextColor(void* ptr, void* textColor);
void* TextField_BackgroundColor(void* ptr);
void TextField_SetBackgroundColor(void* ptr, void* backgroundColor);

