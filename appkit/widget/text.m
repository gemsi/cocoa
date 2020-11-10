#import <Appkit/NSText.h>
        #import "text.h"
        #include "_cgo_export.h"


void* Text_New(long id, NSRect frame) {
    NSText* text = [[[NSText alloc] initWithFrame:frame] autorelease];
    return (void*)text;
}

const char* Text_String(void* ptr) {
    NSText* text = (NSText*)ptr;
    return [text.string UTF8String];
}

void Text_SetString(void* ptr, const char* value) {
    NSText* text = (NSText*)ptr;
    [text setString:[NSString stringWithUTF8String:value]];
}

bool Text_Editable(void* ptr) {
    NSText* text = (NSText*)ptr;
    return text.editable;
}

void Text_SetEditable(void* ptr, bool value) {
    NSText* text = (NSText*)ptr;
    return [text setEditable:value];
}

bool Text_Selectable(void* ptr) {
    NSText* text = (NSText*)ptr;
    return text.selectable;
}

void Text_SetSelectable(void* ptr, bool value) {
    NSText* text = (NSText*)ptr;
    return [text setSelectable:value];
}

bool Text_FieldEditor(void* ptr) {
    NSText* text = (NSText*)ptr;
    return text.fieldEditor;
}

void Text_SetFieldEditor(void* ptr, bool value) {
    NSText* text = (NSText*)ptr;
    return [text setFieldEditor:value];
}

bool Text_RichText(void* ptr) {
    NSText* text = (NSText*)ptr;
    return text.richText;
}

void Text_SetRichText(void* ptr, bool value) {
    NSText* text = (NSText*)ptr;
    return [text setRichText:value];
}

NSSize Text_MinSize(void* ptr) {
    NSText* text = (NSText*)ptr;
    return text.minSize;
}

void Text_SetMinSize(void* ptr, NSSize value) {
    NSText* text = (NSText*)ptr;
    [text setMinSize:value];
}

NSSize Text_MaxSize(void* ptr) {
    NSText* text = (NSText*)ptr;
    return text.maxSize;
}

void Text_SetMaxSize(void* ptr, NSSize value) {
    NSText* text = (NSText*)ptr;
    [text setMaxSize:value];
}

bool Text_VerticallyResizable(void* ptr) {
    NSText* text = (NSText*)ptr;
    return text.verticallyResizable;
}

void Text_SetVerticallyResizable(void* ptr, bool value) {
    NSText* text = (NSText*)ptr;
    return [text setVerticallyResizable:value];
}

bool Text_HorizontallyResizable(void* ptr) {
    NSText* text = (NSText*)ptr;
    return text.horizontallyResizable;
}

void Text_SetHorizontallyResizable(void* ptr, bool value) {
    NSText* text = (NSText*)ptr;
    return [text setHorizontallyResizable:value];
}
