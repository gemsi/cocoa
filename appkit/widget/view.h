#import <stdlib.h>
#import <Foundation/NSGeometry.h>

void* View_New(long id, NSRect frame);
NSRect View_Frame(void* ptr);
void View_SetFrame(void* ptr, NSRect frame);
unsigned long View_AutoresizingMask(void* ptr);
void View_SetAutoresizingMask(void* ptr, unsigned long autoresizingMask);
