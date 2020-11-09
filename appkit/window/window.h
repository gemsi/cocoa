#import <stdlib.h>
#import <stdbool.h>
#import <Foundation/NSGeometry.h>

void* Window_New(long goID, NSRect rect);
void Window_SetTitle(void *ptr, const char* title);
void Window_MakeKeyAndOrderFront(void *ptr);
void Window_Center(void *ptr);
void Window_AddView(void *ptr, void* viewPtr);
void Window_Update(void *ptr);
NSRect Window_Frame(void *ptr);
void Window_SetFrame(void *ptr, NSRect rect, bool display);