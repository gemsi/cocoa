#import <Appkit/NSText.h>
        #import "text.h"
        #include "_cgo_export.h"


void* Text_New(long id, NSRect frame) {
    NSText* text = [[[NSText alloc] initWithFrame:frame] autorelease];
    return (void*)text;
}
const char* Text_String(void* ptr) {
	NSText* text = (NSText*)ptr;
	return [[text string] UTF8String];
}
void Text_SetString(void* ptr, const char* string) {
	NSText* text = (NSText*)ptr;
	[text setString:[NSString stringWithUTF8String:string]];
}
bool Text_IsEditable(void* ptr) {
	NSText* text = (NSText*)ptr;
	return [text isEditable];
}
void Text_SetEditable(void* ptr, bool editable) {
	NSText* text = (NSText*)ptr;
	[text setEditable:editable];
}
bool Text_IsSelectable(void* ptr) {
	NSText* text = (NSText*)ptr;
	return [text isSelectable];
}
void Text_SetSelectable(void* ptr, bool selectable) {
	NSText* text = (NSText*)ptr;
	[text setSelectable:selectable];
}
bool Text_IsFieldEditor(void* ptr) {
	NSText* text = (NSText*)ptr;
	return [text isFieldEditor];
}
void Text_SetFieldEditor(void* ptr, bool fieldEditor) {
	NSText* text = (NSText*)ptr;
	[text setFieldEditor:fieldEditor];
}
bool Text_IsRichText(void* ptr) {
	NSText* text = (NSText*)ptr;
	return [text isRichText];
}
void Text_SetRichText(void* ptr, bool richText) {
	NSText* text = (NSText*)ptr;
	[text setRichText:richText];
}
NSSize Text_MinSize(void* ptr) {
	NSText* text = (NSText*)ptr;
	return [text minSize];
}
void Text_SetMinSize(void* ptr, NSSize minSize) {
	NSText* text = (NSText*)ptr;
	[text setMinSize:minSize];
}
NSSize Text_MaxSize(void* ptr) {
	NSText* text = (NSText*)ptr;
	return [text maxSize];
}
void Text_SetMaxSize(void* ptr, NSSize maxSize) {
	NSText* text = (NSText*)ptr;
	[text setMaxSize:maxSize];
}
bool Text_IsVerticallyResizable(void* ptr) {
	NSText* text = (NSText*)ptr;
	return [text isVerticallyResizable];
}
void Text_SetVerticallyResizable(void* ptr, bool verticallyResizable) {
	NSText* text = (NSText*)ptr;
	[text setVerticallyResizable:verticallyResizable];
}
bool Text_IsHorizontallyResizable(void* ptr) {
	NSText* text = (NSText*)ptr;
	return [text isHorizontallyResizable];
}
void Text_SetHorizontallyResizable(void* ptr, bool horizontallyResizable) {
	NSText* text = (NSText*)ptr;
	[text setHorizontallyResizable:horizontallyResizable];
}
