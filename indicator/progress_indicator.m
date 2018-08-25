#import "progress_indicator.h"
#include "_cgo_export.h"

void* ProgressIndicator_New() {
    id nsProgressIndicator = [[NSProgressIndicator alloc] init];
    [nsProgressIndicator setUsesThreadedAnimation:YES];
    [nsProgressIndicator autorelease];
    return (void*)nsProgressIndicator;
}

void ProgressIndicator_StartAnimation(void* ptr)
{
    NSProgressIndicator* indicator = (NSProgressIndicator*)ptr;
    [indicator startAnimation:nil];
    [indicator setNeedsDisplay:YES];
}

void ProgressIndicator_StopAnimation(void* ptr)
{
    NSProgressIndicator* indicator = (NSProgressIndicator*)ptr;
    [indicator stopAnimation:nil];
    [indicator setNeedsDisplay:YES];
}

void ProgressIndicator_SetLimits(void* ptr, double minValue, double maxValue)
{
    NSProgressIndicator* indicator = (NSProgressIndicator*)ptr;
    [indicator setMinValue:minValue];
    [indicator setMaxValue:maxValue];
}

double ProgressIndicator_MaxValue(void* ptr)
{
    NSProgressIndicator* indicator = (NSProgressIndicator*)ptr;
    return [indicator maxValue];
}

double ProgressIndicator_MinValue(void* ptr)
{
    NSProgressIndicator* indicator = (NSProgressIndicator*)ptr;
    return [indicator minValue];
}

void ProgressIndicator_SetValue(void* ptr, double value)
{
    NSProgressIndicator* indicator = (NSProgressIndicator*)ptr;
    [indicator setDoubleValue:value];
    [indicator setNeedsDisplay:YES];
}

void ProgressIndicator_IncrementBy(void* ptr, double value)
{
    NSProgressIndicator* indicator = (NSProgressIndicator*)ptr;
    [indicator incrementBy:value];
    [indicator setNeedsDisplay:YES];
}

double ProgressIndicator_Value(void* ptr)
{
    NSProgressIndicator* indicator = (NSProgressIndicator*)ptr;
    return [indicator doubleValue];
}

void ProgressIndicator_SetIndeterminate(void* ptr, int value)
{
    NSProgressIndicator* indicator = (NSProgressIndicator*)ptr;
    [indicator setIndeterminate: value > 0?YES:NO];
}

int ProgressIndicator_IsIndeterminate(void* ptr)
{
    NSProgressIndicator* indicator = (NSProgressIndicator*)ptr;
    if ([indicator isIndeterminate]) {
        return 1;
    }
    return 0;
}

void ProgressIndicator_SetDisplayedWhenStopped(void* ptr, int value)
{
    NSProgressIndicator* indicator = (NSProgressIndicator*)ptr;
    [indicator setDisplayedWhenStopped: value > 0?YES:NO];
}


void ProgressIndicator_SetHidden(void* ptr, int hidden)
{
    NSProgressIndicator* indicator = (NSProgressIndicator*)ptr;
    indicator.hidden = hidden > 0?YES:NO;
}

void ProgressIndicator_Remove(void* ptr)
{
    NSProgressIndicator* indicator = (NSProgressIndicator*)ptr;
    [indicator removeFromSuperview];
}
