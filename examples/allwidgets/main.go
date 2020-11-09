package main

import (
	"github.com/hsiafan/cocoa/appkit/application"
	"github.com/hsiafan/cocoa/appkit/button"
	"github.com/hsiafan/cocoa/appkit/color"
	"github.com/hsiafan/cocoa/appkit/indicator"
	"github.com/hsiafan/cocoa/appkit/textfield"
	"github.com/hsiafan/cocoa/appkit/textview"
	"github.com/hsiafan/cocoa/appkit/window"
	"github.com/hsiafan/cocoa/core/dispatch"
	"github.com/hsiafan/cocoa/foundation/geometry"
	"github.com/hsiafan/cocoa/foundation/notification"
	"github.com/hsiafan/cocoa/foundation/object"
	"runtime"
	"time"
)

// Arrange that main.main runs on main thread.
func init() {
	runtime.LockOSThread()
}

func main() {
	application.Init()
	w := window.New(geometry.MkRect(150, 150, 600, 400))
	w.SetTitle("Test widgets")

	// button
	btn := button.New()
	btn.SetTitle("click me")
	btn.SetFrame(geometry.MkRect(10, 160, 80, 25))
	w.AddView(btn)

	quitBtn := button.New()
	quitBtn.SetFrame(geometry.MkRect(10, 130, 80, 25))
	quitBtn.SetTitle("Quit")
	quitBtn.SetAction(func(sender object.Object) {
		application.Terminate()
	})
	w.AddView(quitBtn)

	// text field
	tf := textfield.New()
	tf.SetFrame(geometry.MkRect(10, 100, 150, 25))
	w.AddView(tf)

	// label
	label := textfield.NewLabel()
	label.SetFrame(geometry.MkRect(170, 100, 150, 25))
	w.AddView(label)
	tf.TextDidChange(func(notification.Notification) {
		dispatch.MainQueueAsync(func() {
			label.SetStringValue(tf.StringValue())
		})
	})
	btn.SetAction(func(sender object.Object) {
		label.SetTextColor(color.RedColor())
	})

	// password
	stf := textfield.NewSecure()
	stf.SetFrame(geometry.MkRect(340, 100, 150, 25))
	w.AddView(stf)

	// progress indicator
	progressIndicator := indicator.New()
	progressIndicator.SetFrame(geometry.MkRect(10, 70, 350, 25))
	progressIndicator.SetLimits(0, 1)
	progressIndicator.SetIndeterminate(false)
	w.AddView(progressIndicator)
	go func() {
		dispatch.MainQueueAsync(func() {
			progressIndicator.StartAnimation()
		})
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			dispatch.MainQueueAsync(func() {
				progressIndicator.SetValue(0.1 * float64(i))
			})
		}
		dispatch.MainQueueAsync(func() {
			progressIndicator.StartAnimation()
		})
	}()

	// text view
	tv := textview.New(geometry.MkRect(10, 200, 200, 30))
	w.AddView(tv)

	w.MakeKeyAndOrderFront()
	w.Center()

	application.Run()
}
