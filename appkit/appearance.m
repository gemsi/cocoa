#import <Appkit/NSAppearance.h>
#import <Foundation/NSString.h>
#import "appearance.h"

const char* Appearance_Name(void* ptr) {
	NSAppearance* appearance = (NSAppearance*)ptr;
	return [[appearance name] UTF8String];
}

void* Appearance_AppearanceNamed(const char* name) {
	return [NSAppearance appearanceNamed:[NSString stringWithUTF8String:name]];
}
