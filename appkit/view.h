#import <Foundation/NSGeometry.h>
#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>

NSRect View_Frame(void* ptr);
void View_SetFrame(void* ptr, NSRect frame);
unsigned long View_AutoresizingMask(void* ptr);
void View_SetAutoresizingMask(void* ptr, unsigned long autoresizingMask);
bool View_NeedsDisplay(void* ptr);
void View_SetNeedsDisplay(void* ptr, bool needsDisplay);
bool View_TranslatesAutoresizingMaskIntoConstraints(void* ptr);
void View_SetTranslatesAutoresizingMaskIntoConstraints(void* ptr, bool translatesAutoresizingMaskIntoConstraints);
void* View_BottomAnchor(void* ptr);
void* View_CenterXAnchor(void* ptr);
void* View_CenterYAnchor(void* ptr);
void* View_FirstBaselineAnchor(void* ptr);
void* View_HeightAnchor(void* ptr);
void* View_LastBaselineAnchor(void* ptr);
void* View_LeadingAnchor(void* ptr);
void* View_LeftAnchor(void* ptr);
void* View_RightAnchor(void* ptr);
void* View_TopAnchor(void* ptr);
void* View_TrailingAnchor(void* ptr);
void* View_WidthAnchor(void* ptr);

void View_AddSubview(void* ptr, void* subView);
