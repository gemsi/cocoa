#import <Appkit/NSWindow.h>
#import "_cgo_export.h"

@interface GoNSWindowDelegate : NSObject <NSWindowDelegate>
@property (assign) long goID;
@end

@implementation GoNSWindowDelegate

- (void)windowDidResize:(NSNotification*)notification {
	return Window_Delegate_WindowDidResize([self goID], notification);
}

- (void)windowDidMove:(NSNotification*)notification {
	return Window_Delegate_WindowDidMove([self goID], notification);
}

- (void)windowDidMiniaturize:(NSNotification*)notification {
	return Window_Delegate_WindowDidMiniaturize([self goID], notification);
}

- (void)windowDidDeminiaturize:(NSNotification*)notification {
	return Window_Delegate_WindowDidDeminiaturize([self goID], notification);
}

@end
