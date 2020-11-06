#import <Cocoa/Cocoa.h>
#import "notification.h"
#include "_cgo_export.h"
#import <objc/runtime.h>


const char* Notification_Name(void* ptr) {
    NSNotification* notification = (NSNotification*)ptr;
    return [[notification name] UTF8String];
}
