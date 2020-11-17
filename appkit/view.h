#import <Foundation/NSGeometry.h>
#import <stdbool.h>
#import <stdlib.h>

NSRect View_Frame(void* ptr);
void View_SetFrame(void* ptr, NSRect frame);
unsigned long View_AutoresizingMask(void* ptr);
void View_SetAutoresizingMask(void* ptr, unsigned long autoresizingMask);
bool View_NeedsDisplay(void* ptr);
void View_SetNeedsDisplay(void* ptr, bool needsDisplay);

void View_AddSubview(void* ptr, void* subView);
