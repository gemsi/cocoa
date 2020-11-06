#import "button.h"
#include "_cgo_export.h"

@implementation ButtonHandler
-(void) buttonClicked:(id) sender
{
    onButtonClicked([self goButtonID]);
}
@end

void* Button_New() {
    id nsButton = [[[NSButton alloc] init] autorelease];
    return (void*)nsButton;
}


void Button_SetTitle(void* btnPtr, const char* title) {
    NSButton* button = (NSButton*)btnPtr;
    [button setTitle:[NSString stringWithUTF8String:title]];
}

void Button_SizeToFit(void* btnPtr) {
    NSButton* button = (NSButton*)btnPtr;
    [button sizeToFit];
}

void Button_SetBezelStyle(void* btnPtr, int style) {
    NSButton* button = (NSButton*)btnPtr;
    [button setBezelStyle:style];
}

void Button_SetTarget(void* btnPtr, int buttonId) {
    NSButton* nsButton = (NSButton*)btnPtr;
    ButtonHandler* handler = [[ButtonHandler alloc] init];
    [handler setGoButtonID:buttonId];
    [handler autorelease];
    [nsButton setTarget:handler];
    [nsButton setAction:@selector(buttonClicked:)];

}