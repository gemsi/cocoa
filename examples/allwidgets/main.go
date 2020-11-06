package main

import (
	"github.com/hsiafan/cocoa"
	"github.com/hsiafan/cocoa/application"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/foundation/geometry"
	"github.com/hsiafan/cocoa/foundation/notification"
	"github.com/hsiafan/cocoa/graphics/color"
	"github.com/hsiafan/cocoa/widgets/button"
	"github.com/hsiafan/cocoa/widgets/indicator"
	"github.com/hsiafan/cocoa/widgets/textfield"
	"github.com/hsiafan/cocoa/widgets/window"
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
	quitBtn.SetAction(func(sender foundation.Object) {
		application.Terminate()
	})
	w.AddView(quitBtn)

	// text
	tf := textfield.New()
	tf.SetFrame(geometry.MkRect(10, 100, 150, 25))
	w.AddView(tf)

	// label
	label := textfield.NewLabel()
	label.SetFrame(geometry.MkRect(170, 100, 150, 25))
	w.AddView(label)
	tf.TextDidChange(func(notification.Notification) {
		cocoa.RunOnUIThread(func() {
			label.SetStringValue(tf.StringValue())
		})
	})
	btn.SetAction(func(sender foundation.Object) {
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
		cocoa.RunOnUIThread(func() {
			progressIndicator.StartAnimation()
		})
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			cocoa.RunOnUIThread(func() {
				progressIndicator.SetValue(0.1 * float64(i))
			})
		}
		cocoa.RunOnUIThread(func() {
			progressIndicator.StartAnimation()
		})
	}()

	w.MakeKeyAndOrderFront()
	w.Center()

	application.Run()
}
