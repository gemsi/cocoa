package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include "application.h"
import "C"

var didFinishLaunchingFunc func()

// Init initializes the global application instance.
func Init() {
	C.InitSharedApplication()
}

// Run launches the main Cocoa runloop.
func Run() {
	C.RunApplication()
}

// OnDidFinishLaunching will be triggered after Application Launch is finished
func OnDidFinishLaunching(fn func()) {
	didFinishLaunchingFunc = fn
}

// Terminate  terminate the application
func Terminate() {
	C.TerminateApplication()
}

//export callOnApplicationDidFinishLaunchingHandler
func callOnApplicationDidFinishLaunchingHandler() {
	if didFinishLaunchingFunc != nil {
		didFinishLaunchingFunc()
	}
}
