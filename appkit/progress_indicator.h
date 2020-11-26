#import <Foundation/NSGeometry.h>
#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>

bool ProgressIndicator_UsesThreadedAnimation(void* ptr);
void ProgressIndicator_SetUsesThreadedAnimation(void* ptr, bool usesThreadedAnimation);
double ProgressIndicator_DoubleValue(void* ptr);
void ProgressIndicator_SetDoubleValue(void* ptr, double doubleValue);
double ProgressIndicator_MinValue(void* ptr);
void ProgressIndicator_SetMinValue(void* ptr, double minValue);
double ProgressIndicator_MaxValue(void* ptr);
void ProgressIndicator_SetMaxValue(void* ptr, double maxValue);
bool ProgressIndicator_IsIndeterminate(void* ptr);
void ProgressIndicator_SetIndeterminate(void* ptr, bool indeterminate);
bool ProgressIndicator_IsBezeled(void* ptr);
void ProgressIndicator_SetBezeled(void* ptr, bool bezeled);
bool ProgressIndicator_IsDisplayedWhenStopped(void* ptr);
void ProgressIndicator_SetDisplayedWhenStopped(void* ptr, bool displayedWhenStopped);

void* ProgressIndicator_InitWithFrame(NSRect frame);
void ProgressIndicator_StartAnimation(void* ptr, void* sender);
void ProgressIndicator_StopAnimation(void* ptr, void* sender);
void ProgressIndicator_IncrementBy(void* ptr, double delta);
