#import <Cocoa/Cocoa.h>

void* ProgressIndicator_New();

void ProgressIndicator_StartAnimation(void* ptr);

void ProgressIndicator_StopAnimation(void* ptr);

void ProgressIndicator_SetLimits(void* ptr, double minValue, double maxValue);
double ProgressIndicator_MaxValue(void* ptr);
double ProgressIndicator_MinValue(void* ptr);

void ProgressIndicator_SetValue(void* ptr, double value);

void ProgressIndicator_IncrementBy(void* ptr, double value);

double ProgressIndicator_Value(void* ptr);

void ProgressIndicator_SetIndeterminate(void* ptr, int value);

int ProgressIndicator_IsIndeterminate(void* ptr);

void ProgressIndicator_SetDisplayedWhenStopped(void* ptr, int value);

void ProgressIndicator_SetHidden(void* ptr, int hidden);

void ProgressIndicator_Remove(void* ptr);
