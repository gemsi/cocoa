#import <Cocoa/Cocoa.h>
#import "application.h"
#include "_cgo_export.h"

@interface AppDelegate : NSObject <NSApplicationDelegate>

@end

@implementation AppDelegate

- (void)dealloc {
    [super dealloc];
}

- (void)applicationDidFinishLaunching:(NSNotification *)aNotification {
    callOnApplicationDidFinishLaunchingHandler();
}

- (BOOL)applicationShouldTerminateAfterLastWindowClosed:(NSApplication *)theApplication {
    return YES;
}

@end

AppDelegate *cocoa_applicationDelegate = nil;

// InitSharedApplication calls NSApplication.sharedApplication() which initializes the 
// global application instance NSApp.
void InitSharedApplication() {
    static bool hasBeenInitialized = false; // false first time function is called
    if (hasBeenInitialized)
        return;
    [NSApplication sharedApplication];
    cocoa_applicationDelegate = [[AppDelegate alloc] init];
    [NSApp setDelegate:cocoa_applicationDelegate];
    hasBeenInitialized = true;
}

void releaseSharedApplication() {
    if (cocoa_applicationDelegate != nil) {
        [cocoa_applicationDelegate release];
    }
}

void RunApplication() {
    @autoreleasepool {
        [NSApp setActivationPolicy:NSApplicationActivationPolicyRegular];
        [NSApp activateIgnoringOtherApps:YES];
        [NSApp run];
        releaseSharedApplication();
    }
}

void TerminateApplication() {
    [NSApp terminate:nil];
}