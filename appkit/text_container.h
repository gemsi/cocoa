#import <Foundation/NSGeometry.h>
#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>

NSSize TextContainer_Size(void* ptr);
void TextContainer_SetSize(void* ptr, NSSize size);
bool TextContainer_WidthTracksTextView(void* ptr);
void TextContainer_SetWidthTracksTextView(void* ptr, bool widthTracksTextView);
bool TextContainer_HeightTracksTextView(void* ptr);
void TextContainer_SetHeightTracksTextView(void* ptr, bool heightTracksTextView);

void* TextContainer_InitWithSize(NSSize size);
