#import <Foundation/NSObject.h>
#import "object.h"
#include "_cgo_export.h"


void Object_Dealloc(void* ptr) {
    NSObject* obj = (NSObject*)ptr;
    [obj dealloc];
}
