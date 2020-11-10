#import <stdlib.h>
#import <Foundation/NSGeometry.h>

void* TextContainer_New(long id, NSSize size);

NSSize TextContainer_Size(void* ptr); 

void TextContainer_SetSize(void* ptr, NSSize value); 

bool TextContainer_WidthTracksTextView(void* ptr); 

void TextContainer_SetWidthTracksTextView(void* ptr, bool value); 

bool TextContainer_HeightTracksTextView(void* ptr); 

void TextContainer_SetHeightTracksTextView(void* ptr, bool value); 
