#import <Cocoa/Cocoa.h>

void* TextField_New(int id);
void TextField_SetBezeled(void* ptr, int bezeled);
void TextField_SetEditable(void* ptr, int editable);
void TextField_SetSelectable(void* ptr, int selectable);
void TextField_SetDrawsBackground(void* ptr, int draws);
void* TextField_TextColor(void* ptr);
void TextField_SetTextColor(void* ptr, void* colorPtr);
void* TextField_BackgroundColor(void* ptr);
void TextField_SetBackgroundColor(void* ptr, void* colorPtr);