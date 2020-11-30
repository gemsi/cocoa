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
	tf := appkit.NewPlainButton(foundation.MakeRect(150, 380, 100, 20))
	tf.SetStringValue("button")
	tf2 := appkit.NewPlainButton(foundation.MakeRect(150, 380, 100, 20))
	tf2.SetStringValue("button2")

	sv := appkit.StackViewWithViews([]appkit.View{label, tf, tf2})
	sv.SetOrientation(appkit.UserInterfaceLayoutOrientationVertical)
	sv.SetDistribution(appkit.NSStackViewDistributionFillEqually)
	sv.SetAlignment(appkit.LayoutAttributeCenterX)
	sv.SetSpacing(10)
	//sv.LeadingAnchor().ConstraintEqualTo(label.LeadingAnchor()).Active()
	views := sv.Views()
	fmt.Println(views)
	w.ContentView().AddSubview(sv)

	w.MakeKeyAndOrderFront(nil)
	w.Center()

	appkit.Run()
}
