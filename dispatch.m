#import <Cocoa/Cocoa.h>
#import "dispatch.h"
#include "_cgo_export.h"


void Dispatch_RunOnUIThread(long id) {
    dispatch_async(dispatch_get_main_queue(), ^{
            runTask(id);
    });
}
