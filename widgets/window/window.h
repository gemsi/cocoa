#import <Cocoa/Cocoa.h>

void* Window_New(int x, int y, int width, int height, int goID);
void Window_SetTitle(void *ptr, const char* title);
void Window_MakeKeyAndOrderFront(void *ptr);
void Window_Center(void *ptr);
void Window_AddView(void *ptr, void* viewPtr);
void Window_Update(void *ptr);
NSRect Window_Frame(void *ptr);