#import <stdlib.h>
#import <Foundation/NSGeometry.h>

void* TextView_New(long id, NSRect frame);
void* TextView_TextContainer(void* ptr);
void TextView_SetTextContainer(void* ptr, void* textContainer);
