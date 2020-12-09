#import "selector.h"
#import <objc/runtime.h>


void* Selector_RegisterSelector(const char* name) {
    return (void*)sel_registerName(name);
}