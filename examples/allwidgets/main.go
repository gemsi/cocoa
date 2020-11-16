package main

import (
	"github.com/hsiafan/cocoa/appkit/application"
	"github.com/hsiafan/cocoa/appkit/color"
	"github.com/hsiafan/cocoa/appkit/widget"
	"github.com/hsiafan/cocoa/core/dispatch"
	"github.com/hsiafan/cocoa/foundation"
	"runtime"
	"time"
)

// Arrange that main.main runs on main thread.
func init() {
	runtime.LockOSThread()
}

func main() {
	application.Init()
	w := widget.NewWindow(foundation.MkRect(150, 150, 600, 400))
	w.SetTitle("Test widgets")

	// button
	btn := widget.NewButton()
	btn.SetTitle("click me")
	btn.SetFrame(foundation.MkRect(10, 160, 80, 25))
	w.AddView(btn)

	quitBtn := widget.NewButton()
	quitBtn.SetFrame(foundation.MkRect(10, 130, 80, 25))
	quitBtn.SetTitle("Quit")
	quitBtn.SetAction(func(sender foundation.Object) {
		application.Terminate()
	})
	w.AddView(quitBtn)

	// text field
	tf := widget.NewTextField()
	tf.SetFrame(foundation.MkRect(10, 100, 150, 25))
	w.AddView(tf)

	// label
	label := widget.NewLabel()
	label.SetFrame(foundation.MkRect(170, 100, 150, 25))
	w.AddView(label)
	tf.TextDidChange(func(foundation.Notification) {
		dispatch.MainQueueAsync(func() {
			label.SetStringValue(tf.StringValue())
		})
	})
	btn.SetAction(func(sender foundation.Object) {
		label.SetTextColor(color.Red)
	})

	// password
	stf := widget.NewSecure()
	stf.SetFrame(foundation.MkRect(340, 100, 150, 25))
	w.AddView(stf)

	// progress indicator
	progressIndicator := widget.NewProgressIndicator()
	progressIndicator.SetFrame(foundation.MkRect(10, 70, 350, 25))
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

	// text view & scroll view
	sv := widget.NewVerticallyScrollableTextView(foundation.MkRect(10, 200, 200, 30))
	w.AddView(sv)
	/*
		Listing 4  Setting up a horizontal scroll bar

		[[theTextView enclosingScrollView] setHasHorizontalScroller:YES];
		[theTextView setHorizontallyResizable:YES];
		[theTextView setAutoresizingMask:(NSViewWidthSizable | NSViewHeightSizable)];
		[[theTextView textContainer] setContainerSize:NSMakeSize(FLT_MAX, FLT_MAX)];
		[[theTextView textContainer] setWidthTracksTextView:NO];
	*/

	w.MakeKeyAndOrderFront()
	w.Center()

	application.Run()
}
