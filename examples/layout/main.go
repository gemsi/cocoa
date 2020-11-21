package main

import (
	"fmt"
	"github.com/hsiafan/cocoa/appkit"
	"github.com/hsiafan/cocoa/foundation"
	"runtime"
)

// Arrange that main.main runs on main thread.
func init() {
	runtime.LockOSThread()
}

func main() {
	appkit.Init()
	w := appkit.NewPlainWindow(foundation.MakeRect(150, 150, 600, 400))
	w.SetTitle("Test Layout")

	label := appkit.NewLabel(foundation.MakeRect(20, 380, 100, 20))
	label.SetStringValue("label")
	tf := appkit.NewTextField(foundation.MakeRect(150, 380, 100, 20))
	tf.SetStringValue("text field")

	sv := appkit.NewStackViewWithViews([]appkit.View{label, tf})
	views := sv.Views()
	fmt.Println(views)
	w.ContentView().AddSubview(sv)

	w.MakeKeyAndOrderFront(nil)
	w.Center()

	appkit.Run()
}
