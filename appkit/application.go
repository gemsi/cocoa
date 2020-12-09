package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include "application.h"
import "C"

import (
	"github.com/hsiafan/cocoa/foundation"
	"unsafe"
)

// Application is An object that manages an app’s main event loop and resources used by all of that app’s objects.
type Application interface {
	Responder

	// CurrentEvent return the last event object that the app retrieved from the event queue
	CurrentEvent() Event
	// IsRunning return whether the main event loop is running
	IsRunning() bool
	// IsActive return whether this is the active app
	IsActive() bool
	// MainMenu return the app’s main menu bar
	MainMenu() Menu
	// SetMainMenu set the app’s main menu bar
	SetMainMenu(mainMenu Menu)
	// WindowsMenu return the Window menu of the app
	WindowsMenu() Menu
	// SetWindowsMenu set the Window menu of the app
	SetWindowsMenu(windowsMenu Menu)
	// ServicesMenu return the app’s Services menu
	ServicesMenu() Menu
	// SetServicesMenu set the app’s Services menu
	SetServicesMenu(servicesMenu Menu)
	// MainWindow return the app’s main window
	MainWindow() Window
	// Windows return an array of the app’s window objects
	Windows() []Window
	// IsHidden return whether the app is hidden
	IsHidden() bool
	// Run starts the main event loop
	Run()
	// FinishLaunching activates the app, opens any files specified by the NSOpen user default, and unhighlights the app’s icon
	FinishLaunching()
	// Stop stops the main event loop
	Stop(sender foundation.Object)
	// Terminate terminates the receiver
	Terminate(sender foundation.Object)
	// SendEvent dispatches an event to other objects
	SendEvent(event Event)
	// ActivateIgnoringOtherApps makes the receiver the active app
	ActivateIgnoringOtherApps(flag bool)
	// Deactivate deactivates the receiver
	Deactivate()
	// SetActivationPolicy attempts to modify the app’s activation policy
	SetActivationPolicy(activationPolicy ApplicationActivationPolicy)
	// WindowWithWindowNumber returns the window corresponding to the specified window number
	WindowWithWindowNumber(windowNum int)
	// MiniaturizeAll miniaturizes all the receiver’s windows
	MiniaturizeAll(sender foundation.Object)
	// Hide hides all the receiver’s windows, and the next app in line is activated
	Hide(sender foundation.Object)
	// Unhide restores hidden windows to the screen and makes the receiver active
	Unhide(sender foundation.Object)
	// UnhideWithoutActivation restores hidden windows without activating their owner (the receiver)
	UnhideWithoutActivation()
	// UpdateWindows sends an update message to each onscreen window
	UpdateWindows()
	SetWindowsNeedUpdate(needUpdate bool)
	ArrangeInFront(sender foundation.Object)
	PreventWindowOrdering()
	// ApplicationWillFinishLaunching sent by the default notification center immediately before the application object is initialized
	ApplicationWillFinishLaunching(callback func(notification foundation.Notification))
	_applicationWillFinishLaunching() func(notification foundation.Notification)
	// ApplicationDidFinishLaunching sent by the default notification center after the application has been launched and initialized but before it has received its first event
	ApplicationDidFinishLaunching(callback func(notification foundation.Notification))
	_applicationDidFinishLaunching() func(notification foundation.Notification)
	// ApplicationShouldTerminateAfterLastWindowClosed invoked when the user closes the last window the application has open.return NO if the application should not be terminated when its last window is closed; otherwise, YES to terminate the application
	ApplicationShouldTerminateAfterLastWindowClosed(callback func(theApplication Application) bool)
	_applicationShouldTerminateAfterLastWindowClosed() func(theApplication Application) bool
}

var _ Application = (*NSApplication)(nil)

type NSApplication struct {
	NSResponder
	applicationWillFinishLaunching                  func(notification foundation.Notification)
	applicationDidFinishLaunching                   func(notification foundation.Notification)
	applicationShouldTerminateAfterLastWindowClosed func(theApplication Application) bool
}

// MakeApplication create a Application from native pointer
func MakeApplication(ptr unsafe.Pointer) *NSApplication {
	if ptr == nil {
		return nil
	}
	return &NSApplication{
		NSResponder: *MakeResponder(ptr),
	}
}
func (a *NSApplication) setDelegate() {
	id := resources.NextId()
	C.Application_RegisterDelegate(a.Ptr(), C.long(id))
	resources.Store(id, a)
	foundation.AddDeallocHook(a, func() {
		resources.Delete(id)
	})
}

func (a *NSApplication) CurrentEvent() Event {
	return MakeEvent(C.Application_CurrentEvent(a.Ptr()))
}

func (a *NSApplication) IsRunning() bool {
	return bool(C.Application_IsRunning(a.Ptr()))
}

func (a *NSApplication) IsActive() bool {
	return bool(C.Application_IsActive(a.Ptr()))
}

func (a *NSApplication) MainMenu() Menu {
	return MakeMenu(C.Application_MainMenu(a.Ptr()))
}

func (a *NSApplication) SetMainMenu(mainMenu Menu) {
	C.Application_SetMainMenu(a.Ptr(), toPointer(mainMenu))
}

func (a *NSApplication) WindowsMenu() Menu {
	return MakeMenu(C.Application_WindowsMenu(a.Ptr()))
}

func (a *NSApplication) SetWindowsMenu(windowsMenu Menu) {
	C.Application_SetWindowsMenu(a.Ptr(), toPointer(windowsMenu))
}

func (a *NSApplication) ServicesMenu() Menu {
	return MakeMenu(C.Application_ServicesMenu(a.Ptr()))
}

func (a *NSApplication) SetServicesMenu(servicesMenu Menu) {
	C.Application_SetServicesMenu(a.Ptr(), toPointer(servicesMenu))
}

func (a *NSApplication) MainWindow() Window {
	return MakeWindow(C.Application_MainWindow(a.Ptr()))
}

func (a *NSApplication) Windows() []Window {
	var cArray C.Array = C.Application_Windows(a.Ptr())
	defer C.free(cArray.data)
	result := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(cArray.data))[:cArray.len:cArray.len]
	var goResult = make([]Window, len(result))
	for idx, r := range result {
		goResult[idx] = MakeWindow(r)
	}
	return goResult
}

func (a *NSApplication) IsHidden() bool {
	return bool(C.Application_IsHidden(a.Ptr()))
}

func (a *NSApplication) Run() {
	C.Application_Run(a.Ptr())
}

func (a *NSApplication) FinishLaunching() {
	C.Application_FinishLaunching(a.Ptr())
}

func (a *NSApplication) Stop(sender foundation.Object) {
	C.Application_Stop(a.Ptr(), toPointer(sender))
}

func (a *NSApplication) Terminate(sender foundation.Object) {
	C.Application_Terminate(a.Ptr(), toPointer(sender))
}

func (a *NSApplication) SendEvent(event Event) {
	C.Application_SendEvent(a.Ptr(), toPointer(event))
}

func (a *NSApplication) ActivateIgnoringOtherApps(flag bool) {
	C.Application_ActivateIgnoringOtherApps(a.Ptr(), C.bool(flag))
}

func (a *NSApplication) Deactivate() {
	C.Application_Deactivate(a.Ptr())
}

func SharedApplication() Application {
	return MakeApplication(C.Application_SharedApplication())
}

func (a *NSApplication) SetActivationPolicy(activationPolicy ApplicationActivationPolicy) {
	C.Application_SetActivationPolicy(a.Ptr(), C.long(activationPolicy))
}

func (a *NSApplication) WindowWithWindowNumber(windowNum int) {
	C.Application_WindowWithWindowNumber(a.Ptr(), C.long(windowNum))
}

func (a *NSApplication) MiniaturizeAll(sender foundation.Object) {
	C.Application_MiniaturizeAll(a.Ptr(), toPointer(sender))
}

func (a *NSApplication) Hide(sender foundation.Object) {
	C.Application_Hide(a.Ptr(), toPointer(sender))
}

func (a *NSApplication) Unhide(sender foundation.Object) {
	C.Application_Unhide(a.Ptr(), toPointer(sender))
}

func (a *NSApplication) UnhideWithoutActivation() {
	C.Application_UnhideWithoutActivation(a.Ptr())
}

func (a *NSApplication) UpdateWindows() {
	C.Application_UpdateWindows(a.Ptr())
}

func (a *NSApplication) SetWindowsNeedUpdate(needUpdate bool) {
	C.Application_SetWindowsNeedUpdate(a.Ptr(), C.bool(needUpdate))
}

func (a *NSApplication) ArrangeInFront(sender foundation.Object) {
	C.Application_ArrangeInFront(a.Ptr(), toPointer(sender))
}

func (a *NSApplication) PreventWindowOrdering() {
	C.Application_PreventWindowOrdering(a.Ptr())
}

func (a *NSApplication) ApplicationWillFinishLaunching(callback func(notification foundation.Notification)) {
	a.applicationWillFinishLaunching = callback
}

func (a *NSApplication) _applicationWillFinishLaunching() func(notification foundation.Notification) {
	return a.applicationWillFinishLaunching
}

func (a *NSApplication) ApplicationDidFinishLaunching(callback func(notification foundation.Notification)) {
	a.applicationDidFinishLaunching = callback
}

func (a *NSApplication) _applicationDidFinishLaunching() func(notification foundation.Notification) {
	return a.applicationDidFinishLaunching
}

func (a *NSApplication) ApplicationShouldTerminateAfterLastWindowClosed(callback func(theApplication Application) bool) {
	a.applicationShouldTerminateAfterLastWindowClosed = callback
}

func (a *NSApplication) _applicationShouldTerminateAfterLastWindowClosed() func(theApplication Application) bool {
	return a.applicationShouldTerminateAfterLastWindowClosed
}

//export Application_Delegate_ApplicationWillFinishLaunching
func Application_Delegate_ApplicationWillFinishLaunching(id int64, notification unsafe.Pointer) {
	application := resources.Get(id).(Application)
	callback := application._applicationWillFinishLaunching()
	if callback != nil {
		callback(foundation.MakeNotification(notification))
	}
}

//export Application_Delegate_ApplicationDidFinishLaunching
func Application_Delegate_ApplicationDidFinishLaunching(id int64, notification unsafe.Pointer) {
	application := resources.Get(id).(Application)
	callback := application._applicationDidFinishLaunching()
	if callback != nil {
		callback(foundation.MakeNotification(notification))
	}
}

//export Application_Delegate_ApplicationShouldTerminateAfterLastWindowClosed
func Application_Delegate_ApplicationShouldTerminateAfterLastWindowClosed(id int64, theApplication unsafe.Pointer) bool {
	application := resources.Get(id).(Application)
	callback := application._applicationShouldTerminateAfterLastWindowClosed()
	if callback != nil {
		return callback(MakeApplication(theApplication))
	}
	return true
}
