package main

import (
	"github.com/hsiafan/cocoa/appkit"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"runtime"
	"time"
)

// Arrange that main.main runs on main thread.
func init() {
	runtime.LockOSThread()
}

func initAndRun() {
	app := appkit.InitSharedApplication()
	w := appkit.NewPlainWindow(foundation.MakeRect(150, 150, 600, 400))
	w.SetTitle("Test widgets")

	// buttons
	cb := appkit.NewCheckBox(foundation.MakeRect(10, 250, 80, 25))
	cb.SetTitle("check box")
	w.ContentView().AddSubview(cb)

	rb := appkit.NewRadioButton(foundation.MakeRect(150, 250, 120, 25))
	rb.SetTitle("radio button")
	w.ContentView().AddSubview(rb)

	btn := appkit.NewPlainButton(foundation.MakeRect(10, 160, 120, 25))
	btn.SetTitle("change color")
	w.ContentView().AddSubview(btn)

	quitBtn := appkit.NewPlainButton(foundation.MakeRect(10, 130, 80, 25))
	quitBtn.SetTitle("Quit")
	quitBtn.SetAction(func(sender foundation.Object) {
		app.Terminate(nil)
	})
	w.ContentView().AddSubview(quitBtn)

	// text field
	tf := appkit.NewPlainTextField(foundation.MakeRect(10, 100, 150, 25))
	w.ContentView().AddSubview(tf)

	// label
	label := appkit.NewLabel(foundation.MakeRect(170, 100, 150, 25))
	w.ContentView().AddSubview(label)
	tf.ControlTextDidChange(func(foundation.Notification) {
		objc.MainQueueAsync(func() {
			label.SetStringValue(tf.StringValue())
		})
	})
	btn.SetAction(func(sender foundation.Object) {
		label.SetTextColor(appkit.ColorRed)
	})

	// password
	stf := appkit.NewPlainSecureTextField(foundation.MakeRect(340, 100, 150, 25))
	stf.ControlTextDidChange(func(notification foundation.Notification) {
		label.SetStringValue(stf.StringValue())
	})
	w.ContentView().AddSubview(stf)

	// progress indicator
	progressIndicator := appkit.NewProgressIndicator(foundation.MakeRect(10, 70, 350, 25))
	progressIndicator.SetMinValue(0)
	progressIndicator.SetMaxValue(1)
	progressIndicator.SetIndeterminate(false)
	progressIndicator.SetDisplayedWhenStopped(false)
	w.ContentView().AddSubview(progressIndicator)
	go func() {
		objc.MainQueueAsync(func() {
			progressIndicator.StartAnimation(progressIndicator)
		})
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			objc.MainQueueAsync(func() {
				progressIndicator.SetDoubleValue(0.1 * float64(i))
			})
		}
		objc.MainQueueAsync(func() {
			progressIndicator.StopAnimation(progressIndicator)
		})
	}()

	// text view & scroll view
	sv := appkit.ScrollableTextView()
	sv.SetFrame(foundation.MakeRect(10, 200, 200, 30))
	appkit.MakeTextView(sv.DocumentView().Ptr()).SetAllowsUndo(true)
	w.ContentView().AddSubview(sv)

	sv2 := appkit.ScrollableTextView()
	sv2.SetFrame(foundation.MakeRect(250, 200, 200, 30))
	appkit.MakeTextView(sv2.DocumentView().Ptr()).SetAllowsUndo(true)
	w.ContentView().AddSubview(sv2)

	w.WindowDidMove(func(notification foundation.Notification) {
		label.SetStringValue("moved!")
	})
	w.MakeKeyAndOrderFront(nil)
	w.Center()

	app.ApplicationWillFinishLaunching(func(notification foundation.Notification) {
		menu := appkit.NewMenu("main")
		app.SetMainMenu(menu)

		testMenuItem := appkit.NewMenuItem("", "", func(sender foundation.Object) {
		})
		testMenu := appkit.NewMenu("Name1")
		testMenu.AddItem(appkit.NewMenuItem("test", "c", func(sender foundation.Object) {
			label.SetStringValue("clicked")
		}))
		testMenuItem.SetSubmenu(testMenu)
		menu.AddItem(testMenuItem)

		testMenuItem2 := appkit.NewMenuItem("", "", func(sender foundation.Object) {
		})
		testMenu2 := appkit.NewMenu("Name2")
		ii := appkit.NewMenuItem("test2", "d", func(sender foundation.Object) {
			label.SetStringValue("clicked2")
		})
		testMenu2.AddItem(ii)
		testMenuItem2.SetSubmenu(testMenu2)
		menu.AddItem(testMenuItem2)
	})

	app.ApplicationDidFinishLaunching(func(notification foundation.Notification) {
		app.SetActivationPolicy(appkit.ApplicationActivationPolicyRegular)
		app.ActivateIgnoringOtherApps(true)
	})

	app.Run()
}

func main() {
	objc.WithAutoreleasePool(func() {
		initAndRun()
	})
}
