#import <Cocoa/Cocoa.h>

@interface ButtonHandler : NSObject

@property (assign) int goButtonID;
-(void) buttonClicked:(id) sender;

@end

void* Button_New();
void Button_SetTitle(void* btnPtr, const char* title);
void Button_SizeToFit(void* btnPtr);
void Button_SetBezelStyle(void* btnPtr, int style);
void Button_SetTarget(void* btnPtr, int buttonId);