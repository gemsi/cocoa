#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>

void Application_RegisterDelegate(void *ptr, long goID);
void* Application_CurrentEvent(void* ptr);
bool Application_IsRunning(void* ptr);
bool Application_IsActive(void* ptr);
void* Application_MainMenu(void* ptr);
void Application_SetMainMenu(void* ptr, void* mainMenu);
void* Application_WindowsMenu(void* ptr);
void Application_SetWindowsMenu(void* ptr, void* windowsMenu);
void* Application_ServicesMenu(void* ptr);
void Application_SetServicesMenu(void* ptr, void* servicesMenu);
void* Application_MainWindow(void* ptr);
Array Application_Windows(void* ptr);
bool Application_IsHidden(void* ptr);

void Application_Run(void* ptr);
void Application_FinishLaunching(void* ptr);
void Application_Stop(void* ptr, void* sender);
void Application_Terminate(void* ptr, void* sender);
void Application_SendEvent(void* ptr, void* event);
void Application_ActivateIgnoringOtherApps(void* ptr, bool flag);
void Application_Deactivate(void* ptr);
void* Application_SharedApplication();
void Application_SetActivationPolicy(void* ptr, long activationPolicy);
void Application_WindowWithWindowNumber(void* ptr, long windowNum);
void Application_MiniaturizeAll(void* ptr, void* sender);
void Application_Hide(void* ptr, void* sender);
void Application_Unhide(void* ptr, void* sender);
void Application_UnhideWithoutActivation(void* ptr);
void Application_UpdateWindows(void* ptr);
void Application_SetWindowsNeedUpdate(void* ptr, bool needUpdate);
void Application_ArrangeInFront(void* ptr, void* sender);
void Application_PreventWindowOrdering(void* ptr);
