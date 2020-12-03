#import <dispatch/dispatch.h>
#import <Foundation/NSObject.h>
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

@interface Parasite : NSObject
@property(nonatomic, assign) long hookId;
@end

@implementation Parasite
- (void)dealloc {
    runTask(self.hookId);
    [super dealloc];
}
@end

static void *kDeallocHookAssociation = &kDeallocHookAssociation;

void Dealloc_AddHook(void* ptr, long hookId) {
    Parasite *parasite = [Parasite alloc];
    parasite.hookId = hookId;
    objc_setAssociatedObject((NSObject*)ptr, &kDeallocHookAssociation, parasite, OBJC_ASSOCIATION_RETAIN_NONATOMIC);
}
