#import <stdlib.h>
#import <Foundation/NSGeometry.h>

void* ScrollView_New(long id, NSRect frame);
bool ScrollView_HasVerticalScroller(void* ptr);
void ScrollView_SetHasVerticalScroller(void* ptr, bool hasVerticalScroller);
bool ScrollView_HasHorizontalScroller(void* ptr);
void ScrollView_SetHasHorizontalScroller(void* ptr, bool hasHorizontalScroller);
void* ScrollView_DocumentView(void* ptr);
void ScrollView_SetDocumentView(void* ptr, void* documentView);
unsigned long ScrollView_BorderType(void* ptr);
void ScrollView_SetBorderType(void* ptr, unsigned long borderType);
NSSize ScrollView_ContentSize(void* ptr);
