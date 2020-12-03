#import <stdlib.h>

void Dispatch_MainQueueAsync(long id);
void Run_WithAutoreleasePool(long id);
void Dealloc_AddHook(void* ptr, long hookId);