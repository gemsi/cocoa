package main

import (
	"github.com/hsiafan/cocoa"
	"github.com/hsiafan/cocoa/application"
	"github.com/hsiafan/cocoa/button"
	"github.com/hsiafan/cocoa/color"
	"github.com/hsiafan/cocoa/indicator"
	"github.com/hsiafan/cocoa/textfield"
	"github.com/hsiafan/cocoa/window"
	"time"
)

func main() {
	application.Init()
	w := window.New(150, 150, 600, 400)
	w.SetTitle("Test widgets")

	// text
	tf := textfield.New()
	tf.SetFrame(10, 100, 200, 25)
	w.AddView(tf)

	// button
	btn := button.New()
	btn.SetTitle("click me")
	btn.SetAction(func() {
		btn.SetTitle("clicked")
		tf.SetStringValue("button clicked")
		tf.SetTextColor(color.RedColor())
	})
	btn.SetFrame(10, 160, 80, 25)
	w.AddView(btn)

	quitBtn := button.New()
	quitBtn.SetFrame(10, 130, 80, 25)
	quitBtn.SetTitle("Quit")
	quitBtn.SetAction(func() {
		application.Terminate()
	})
	w.AddView(quitBtn)

	progressIndicator := indicator.New()
	progressIndicator.SetFrame(10, 70, 350, 25)
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
