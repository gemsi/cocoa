#import <stdlib.h>
#import <utils.h>

const char* TabViewItem_Label(void* ptr);
void TabViewItem_SetLabel(void* ptr, const char* label);
const char* TabViewItem_ToolTip(void* ptr);
void TabViewItem_SetToolTip(void* ptr, const char* toolTip);
unsigned long TabViewItem_TabState(void* ptr);
void* TabViewItem_Identifier(void* ptr);
void TabViewItem_SetIdentifier(void* ptr, void* identifier);
void* TabViewItem_Color(void* ptr);
void TabViewItem_SetColor(void* ptr, void* color);
void* TabViewItem_View(void* ptr);
void TabViewItem_SetView(void* ptr, void* view);
void* TabViewItem_InitialFirstResponder(void* ptr);
void TabViewItem_SetInitialFirstResponder(void* ptr, void* initialFirstResponder);
void* TabViewItem_TabView(void* ptr);

void* TabViewItem_InitWithIdentifier(void* identifier);
