#import <stdlib.h>
#import <Foundation/NSGeometry.h>

void* Control_New(long id, NSRect frame);
bool Control_IsEnabled(void* ptr);
void Control_SetEnabled(void* ptr, bool enabled);
double Control_DoubleValue(void* ptr);
void Control_SetDoubleValue(void* ptr, double doubleValue);
float Control_FloatValue(void* ptr);
void Control_SetFloatValue(void* ptr, float floatValue);
int Control_IntValue(void* ptr);
void Control_SetIntValue(void* ptr, int intValue);
long Control_IntegerValue(void* ptr);
void Control_SetIntegerValue(void* ptr, long integerValue);
const char* Control_StringValue(void* ptr);
void Control_SetStringValue(void* ptr, const char* stringValue);
