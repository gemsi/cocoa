#import <stdlib.h>
#import <Foundation/NSGeometry.h>

void* Text_New(long id, NSRect frame);

const char* Text_String(void* ptr); 

void Text_SetString(void* ptr, const char* value); 

bool Text_Editable(void* ptr); 

void Text_SetEditable(void* ptr, bool value); 

bool Text_Selectable(void* ptr); 

void Text_SetSelectable(void* ptr, bool value); 

bool Text_FieldEditor(void* ptr); 

void Text_SetFieldEditor(void* ptr, bool value); 

bool Text_RichText(void* ptr); 

void Text_SetRichText(void* ptr, bool value); 

NSSize Text_MinSize(void* ptr); 

void Text_SetMinSize(void* ptr, NSSize value); 

NSSize Text_MaxSize(void* ptr); 

void Text_SetMaxSize(void* ptr, NSSize value); 

bool Text_VerticallyResizable(void* ptr); 

void Text_SetVerticallyResizable(void* ptr, bool value); 

bool Text_HorizontallyResizable(void* ptr); 

void Text_SetHorizontallyResizable(void* ptr, bool value); 
