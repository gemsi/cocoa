#import <Foundation/NSObjCRuntime.h>
#import <Foundation/NSString.h>
#import "runtime.h"

void* Selector_SelectorFromString(const char* name) {
    return NSSelectorFromString([NSString stringWithUTF8String:name]);
}

const char* Selector_StringFromSelector(void *selector) {
    return [NSStringFromSelector((SEL)selector) UTF8String];
}