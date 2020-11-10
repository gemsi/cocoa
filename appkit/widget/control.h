#import <stdlib.h>
#import <Foundation/NSGeometry.h>

void* Control_New(long id, NSRect frame);

bool Control_Enabled(void* ptr); 

void Control_SetEnabled(void* ptr, bool value); 

double Control_DoubleValue(void* ptr); 

void Control_SetDoubleValue(void* ptr, double value); 

float Control_FloatValue(void* ptr); 

void Control_SetFloatValue(void* ptr, float value); 

int Control_IntValue(void* ptr); 

void Control_SetIntValue(void* ptr, int value); 

long Control_IntegerValue(void* ptr); 

void Control_SetIntegerValue(void* ptr, long value); 

const char* Control_StringValue(void* ptr); 

void Control_SetStringValue(void* ptr, const char* value); 
