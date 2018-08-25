#import <Cocoa/Cocoa.h>

void* Window_New(int x, int y, int width, int height);
void Window_SetDelegate(void *wndPtr, int id);
void Window_SetTitle(void *wndPtr, const char* title);
void Window_MakeKeyAndOrderFront(void *wndPtr);
void Window_Center(void *wndPtr);
void Window_AddView(void *wndPtr, void* viewPtr);
void Window_Update(void *wndPtr);
NSRect Window_Frame(void *ptr);