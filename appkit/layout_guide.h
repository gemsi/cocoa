#import <Foundation/NSGeometry.h>
#import <stdlib.h>
#import <utils.h>

NSRect LayoutGuide_Frame(void* ptr);
void* LayoutGuide_BottomAnchor(void* ptr);
void* LayoutGuide_CenterXAnchor(void* ptr);
void* LayoutGuide_CenterYAnchor(void* ptr);
void* LayoutGuide_HeightAnchor(void* ptr);
void* LayoutGuide_LeadingAnchor(void* ptr);
void* LayoutGuide_LeftAnchor(void* ptr);
void* LayoutGuide_RightAnchor(void* ptr);
void* LayoutGuide_TopAnchor(void* ptr);
void* LayoutGuide_TrailingAnchor(void* ptr);
void* LayoutGuide_WidthAnchor(void* ptr);
void* LayoutGuide_OwningView(void* ptr);

void* LayoutGuide_NewLayoutGuide();
