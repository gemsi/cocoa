#import <Appkit/NSButton.h>
#import "button.h"
#include "_cgo_export.h"

@interface ButtonHandler : NSObject

@property (assign) int goID;
-(void) onAction:(id) sender;

@end

@implementation ButtonHandler

-(void) onAction:(id) sender {
    onButtonAction([self goID], sender);
}

@end


void* Button_New(long id) {
    NSButton* button = [[[NSButton alloc] init] autorelease];
    ButtonHandler* handler = [[ButtonHandler alloc] init];
    [handler setGoID:id];
    [handler autorelease];
    [button setTarget:handler];
    [button setAction:@selector(onAction:)];
    return (void*)button;
}


void Button_SetTitle(void* ptr, const char* title) {
    NSButton* button = (NSButton*)ptr;
    [button setTitle:[NSString stringWithUTF8String:title]];
}

void Button_SizeToFit(void* ptr) {
    NSButton* button = (NSButton*)ptr;
    [button sizeToFit];
}

void Button_SetBezelStyle(void* ptr, unsigned long style) {
    NSButton* button = (NSButton*)ptr;
    [button setBezelStyle:style];
}