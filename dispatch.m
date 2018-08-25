#import "dispatch.h"
#include "_cgo_export.h"


void Dispatch_RunOnUIThread(int id) {
    dispatch_async(dispatch_get_main_queue(), ^{
            runTask(id);
    });
}
