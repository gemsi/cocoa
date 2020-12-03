#import <Appkit/NSApplication.h>
#import "application.h"
#import "application_delegate.h"

@implementation GoNSApplicationDelegate

- (void)applicationWillFinishLaunching:(NSNotification*)notification {
	return Application_Delegate_ApplicationWillFinishLaunching([self goID], notification);
}

- (void)applicationDidFinishLaunching:(NSNotification*)notification {
	return Application_Delegate_ApplicationDidFinishLaunching([self goID], notification);
}

- (BOOL)applicationShouldTerminateAfterLastWindowClosed:(NSApplication*)theApplication {
	return Application_Delegate_ApplicationShouldTerminateAfterLastWindowClosed([self goID], theApplication);
}

@end
void Application_RegisterDelegate(void *ptr, long goID) {
	NSApplication* application = (NSApplication*)ptr;
	GoNSApplicationDelegate* delegate = [[GoNSApplicationDelegate alloc] init];
	[delegate setGoID:goID];
	[application setDelegate:delegate];
}

void* Application_CurrentEvent(void* ptr) {
	NSApplication* application = (NSApplication*)ptr;
	return [application currentEvent];
}

bool Application_IsRunning(void* ptr) {
	NSApplication* application = (NSApplication*)ptr;
	return [application isRunning];
}

bool Application_IsActive(void* ptr) {
	NSApplication* application = (NSApplication*)ptr;
	return [application isActive];
}

void* Application_MainMenu(void* ptr) {
	NSApplication* application = (NSApplication*)ptr;
	return [application mainMenu];
}

void Application_SetMainMenu(void* ptr, void* mainMenu) {
	NSApplication* application = (NSApplication*)ptr;
	[application setMainMenu:(NSMenu*)mainMenu];
}

void* Application_WindowsMenu(void* ptr) {
	NSApplication* application = (NSApplication*)ptr;
	return [application windowsMenu];
}

void Application_SetWindowsMenu(void* ptr, void* windowsMenu) {
	NSApplication* application = (NSApplication*)ptr;
	[application setWindowsMenu:(NSMenu*)windowsMenu];
}

void* Application_ServicesMenu(void* ptr) {
	NSApplication* application = (NSApplication*)ptr;
	return [application servicesMenu];
}

void Application_SetServicesMenu(void* ptr, void* servicesMenu) {
	NSApplication* application = (NSApplication*)ptr;
	[application setServicesMenu:(NSMenu*)servicesMenu];
}

void* Application_MainWindow(void* ptr) {
	NSApplication* application = (NSApplication*)ptr;
	return [application mainWindow];
}

Array Application_Windows(void* ptr) {
	NSApplication* application = (NSApplication*)ptr;
	NSArray* array = [application windows];
	int count = [array count];
	void** data = malloc(count * sizeof(void*));
	for (int i = 0; i < count; i++) {
		 data[i] = [array objectAtIndex:i];
	}
	Array result;
	result.data = data;
	result.len = count;
	return result;
}

bool Application_IsHidden(void* ptr) {
	NSApplication* application = (NSApplication*)ptr;
	return [application isHidden];
}

void Application_Run(void* ptr) {
	NSApplication* application = (NSApplication*)ptr;
	[application run];
}

void Application_FinishLaunching(void* ptr) {
	NSApplication* application = (NSApplication*)ptr;
	[application finishLaunching];
}

void Application_Stop(void* ptr, void* sender) {
	NSApplication* application = (NSApplication*)ptr;
	[application stop:(NSObject*)sender];
}

void Application_Terminate(void* ptr, void* sender) {
	NSApplication* application = (NSApplication*)ptr;
	[application terminate:(NSObject*)sender];
}

void Application_SendEvent(void* ptr, void* event) {
	NSApplication* application = (NSApplication*)ptr;
	[application sendEvent:(NSEvent*)event];
}

void Application_ActivateIgnoringOtherApps(void* ptr, bool flag) {
	NSApplication* application = (NSApplication*)ptr;
	[application activateIgnoringOtherApps:flag];
}

void Application_Deactivate(void* ptr) {
	NSApplication* application = (NSApplication*)ptr;
	[application deactivate];
}

void* Application_SharedApplication() {
	return [NSApplication sharedApplication];
}

void Application_SetActivationPolicy(void* ptr, long activationPolicy) {
	NSApplication* application = (NSApplication*)ptr;
	[application setActivationPolicy:activationPolicy];
}

void Application_WindowWithWindowNumber(void* ptr, long windowNum) {
	NSApplication* application = (NSApplication*)ptr;
	[application windowWithWindowNumber:windowNum];
}

void Application_MiniaturizeAll(void* ptr, void* sender) {
	NSApplication* application = (NSApplication*)ptr;
	[application miniaturizeAll:(NSObject*)sender];
}

void Application_Hide(void* ptr, void* sender) {
	NSApplication* application = (NSApplication*)ptr;
	[application hide:(NSObject*)sender];
}

void Application_Unhide(void* ptr, void* sender) {
	NSApplication* application = (NSApplication*)ptr;
	[application unhide:(NSObject*)sender];
}

void Application_UnhideWithoutActivation(void* ptr) {
	NSApplication* application = (NSApplication*)ptr;
	[application unhideWithoutActivation];
}

void Application_UpdateWindows(void* ptr) {
	NSApplication* application = (NSApplication*)ptr;
	[application updateWindows];
}

void Application_SetWindowsNeedUpdate(void* ptr, bool needUpdate) {
	NSApplication* application = (NSApplication*)ptr;
	[application setWindowsNeedUpdate:needUpdate];
}
