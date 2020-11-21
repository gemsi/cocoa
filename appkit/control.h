#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>

bool Control_IsEnabled(void* ptr);
void Control_SetEnabled(void* ptr, bool enabled);
double Control_DoubleValue(void* ptr);
void Control_SetDoubleValue(void* ptr, double doubleValue);
float Control_FloatValue(void* ptr);
void Control_SetFloatValue(void* ptr, float floatValue);
long Control_IntValue(void* ptr);
void Control_SetIntValue(void* ptr, long intValue);
long Control_IntegerValue(void* ptr);
void Control_SetIntegerValue(void* ptr, long integerValue);
const char* Control_StringValue(void* ptr);
void Control_SetStringValue(void* ptr, const char* stringValue);
void* Control_Cell(void* ptr);
void Control_SetCell(void* ptr, void* cell);

void Control_SizeToFit(void* ptr);
