#import <Foundation/NSGeometry.h>
#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>

bool TextView_AllowsUndo(void* ptr);
void TextView_SetAllowsUndo(void* ptr, bool allowsUndo);
void* TextView_TextContainer(void* ptr);
void TextView_SetTextContainer(void* ptr, void* textContainer);

void* TextView_NewTextView(NSRect frame);
void* TextView_ScrollableTextView();
