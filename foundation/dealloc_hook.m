#import <Foundation/NSObject.h>
#import "dealloc_hook.h"
#include "_cgo_export.h"
#import <objc/runtime.h>


@interface Parasite : NSObject
@property(nonatomic, assign) long hookId;
@end

@implementation Parasite
- (void)dealloc {
    runDeallocHook(self.hookId);
    [super dealloc];
}
@end

static void *kDeallocHookAssociation = &kDeallocHookAssociation;

void Dealloc_AddHook(void* ptr, long hookId) {
    Parasite *parasite = [Parasite alloc];
    parasite.hookId = hookId;
    objc_setAssociatedObject((NSObject*)ptr, &kDeallocHookAssociation, parasite, OBJC_ASSOCIATION_RETAIN_NONATOMIC);
}
