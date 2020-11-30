#import <Foundation/NSGeometry.h>
#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>

long Popover_Behavior(void* ptr);
void Popover_SetBehavior(void* ptr, long behavior);
NSRect Popover_PositioningRect(void* ptr);
void Popover_SetPositioningRect(void* ptr, NSRect positioningRect);
bool Popover_Animates(void* ptr);
void Popover_SetAnimates(void* ptr, bool animates);
NSSize Popover_ContentSize(void* ptr);
void Popover_SetContentSize(void* ptr, NSSize contentSize);
bool Popover_IsShown(void* ptr);
bool Popover_IsDetached(void* ptr);
void* Popover_Appearance(void* ptr);
void Popover_SetAppearance(void* ptr, void* appearance);
void* Popover_EffectiveAppearance(void* ptr);

void* Popover_NewPopover();
void Popover_PerformClose(void* ptr, void* sender);
void Popover_Close(void* ptr);
void Popover_ShowRelativeTo(void* ptr, NSRect positioningRect, void* positioningView, long preferredEdge);
