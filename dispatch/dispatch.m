#import <dispatch/dispatch.h>
#import "dispatch.h"
#include "_cgo_export.h"


void Dispatch_MainQueueAsync(long id) {
    dispatch_async(dispatch_get_main_queue(), ^{
            runTask(id);
    });
}
