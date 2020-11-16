#import <stdlib.h>
#import <Foundation/NSGeometry.h>

void* Text_New(long id, NSRect frame);
const char* Text_String(void* ptr);
void Text_SetString(void* ptr, const char* string);
bool Text_IsEditable(void* ptr);
void Text_SetEditable(void* ptr, bool editable);
bool Text_IsSelectable(void* ptr);
void Text_SetSelectable(void* ptr, bool selectable);
bool Text_IsFieldEditor(void* ptr);
void Text_SetFieldEditor(void* ptr, bool fieldEditor);
bool Text_IsRichText(void* ptr);
void Text_SetRichText(void* ptr, bool richText);
NSSize Text_MinSize(void* ptr);
void Text_SetMinSize(void* ptr, NSSize minSize);
NSSize Text_MaxSize(void* ptr);
void Text_SetMaxSize(void* ptr, NSSize maxSize);
bool Text_IsVerticallyResizable(void* ptr);
void Text_SetVerticallyResizable(void* ptr, bool verticallyResizable);
bool Text_IsHorizontallyResizable(void* ptr);
void Text_SetHorizontallyResizable(void* ptr, bool horizontallyResizable);
