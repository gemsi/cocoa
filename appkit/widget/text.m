#import <Appkit/NSText.h>
#import "text.h"


const char* Text_String(void* ptr) {
    NSText* text = (NSText*)ptr;
    return [[text string] UTF8String];
}

void Text_SetEditable(void* ptr, bool value) {
    NSText* text = (NSText*)ptr;
    return [text setEditable:value];
}

void Text_SetSelectable(void* ptr, bool value) {
    NSText* text = (NSText*)ptr;
    return [text setSelectable:value];
}

void Text_SetFieldEditor(void* ptr, bool value) {
    NSText* text = (NSText*)ptr;
    return [text setFieldEditor:value];
}

void Text_SetRichText(void* ptr, bool value) {
    NSText* text = (NSText*)ptr;
    return [text setRichText:value];
}