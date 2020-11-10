#import <stdlib.h>

void* Button_New(long id);
void Button_SetTitle(void* ptr, const char* title);
void Button_SizeToFit(void* ptr);
void Button_SetBezelStyle(void* ptr, unsigned long style);