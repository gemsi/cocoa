#import <dispatch/dispatch.h>
#import "base.h"
#include "_cgo_export.h"
#import <objc/runtime.h>


void Dispatch_MainQueueAsync(long id) {
    dispatch_async(dispatch_get_main_queue(), ^{
            runTask(id);
    });
}

void Run_WithAutoreleasePool(long id) {
    @autoreleasepool {
        runTask(id);
    }
}
