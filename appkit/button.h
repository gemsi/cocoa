#import <Foundation/NSGeometry.h>
#import <stdlib.h>
#import <utils.h>

void Button_RegisterDelegate(void *ptr, long goID);
const char* Button_Title(void* ptr);
void Button_SetTitle(void* ptr, const char* title);
unsigned long Button_BezelStyle(void* ptr);
void Button_SetBezelStyle(void* ptr, unsigned long bezelStyle);
long Button_State(void* ptr);
void Button_SetState(void* ptr, long state);

void* Button_InitWithFrame(NSRect frame);
void Button_SetButtonType(void* ptr, unsigned long buttonType);
