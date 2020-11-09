#import <stdlib.h>
#import <stdbool.h>

void* TextField_New(long id);
void* SecureTextField_New(long id);
void TextField_SetBezeled(void* ptr, bool bezeled);
void TextField_SetEditable(void* ptr, bool editable);
void TextField_SetSelectable(void* ptr, bool selectable);
void TextField_SetDrawsBackground(void* ptr, bool draws);
void* TextField_TextColor(void* ptr);
void TextField_SetTextColor(void* ptr, void* colorPtr);
void* TextField_BackgroundColor(void* ptr);
void TextField_SetBackgroundColor(void* ptr, void* colorPtr);