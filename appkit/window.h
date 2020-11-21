#import <Foundation/NSGeometry.h>
#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>

void* Window_initWithContentRect(long goID, NSRect rect, unsigned long styleMask, unsigned long backing, bool Defer);

const char* Window_Title(void* ptr);
void Window_SetTitle(void* ptr, const char* title);
void* Window_ContentView(void* ptr);
void Window_SetContentView(void* ptr, void* contentView);
unsigned long Window_StyleMask(void* ptr);
void Window_SetStyleMask(void* ptr, unsigned long styleMask);

void Window_Center(void* ptr);
void Window_MakeKeyAndOrderFront(void* ptr, void* sender);
bool Window_MakeFirstResponder(void* ptr, void* responder);
