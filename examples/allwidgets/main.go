package main

import (
	"github.com/hsiafan/cocoa/appkit"
	"github.com/hsiafan/cocoa/dispatch"
	"github.com/hsiafan/cocoa/foundation"
	"runtime"
	"time"
)

// Arrange that main.main runs on main thread.
func init() {
	runtime.LockOSThread()
}

func main() {
	appkit.Init()
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
		appkit.Terminate()
	})
	w.ContentView().AddSubview(quitBtn)

	// text field
	tf := appkit.NewPlainTextField(foundation.MakeRect(10, 100, 150, 25))
	w.ContentView().AddSubview(tf)

	// label
	label := appkit.NewLabel(foundation.MakeRect(170, 100, 150, 25))
	w.ContentView().AddSubview(label)
	tf.ControlTextDidChange(func(foundation.Notification) {
		dispatch.MainQueueAsync(func() {
			label.SetStringValue(tf.StringValue())
		})
	})
	btn.SetAction(func(sender foundation.Object) {
		label.SetTextColor(appkit.ColorRed)
	})

	// password
	stf := appkit.NewPlainSecureTextField(foundation.MakeRect(340, 100, 150, 25))
	w.ContentView().AddSubview(stf)

	// progress indicator
	progressIndicator := appkit.NewProgressIndicator(foundation.MakeRect(10, 70, 350, 25))
	progressIndicator.SetMinValue(0)
	progressIndicator.SetMaxValue(1)
	progressIndicator.SetIndeterminate(false)
	progressIndicator.SetDisplayedWhenStopped(false)
	w.ContentView().AddSubview(progressIndicator)
	go func() {
		dispatch.MainQueueAsync(func() {
			progressIndicator.StartAnimation(progressIndicator)
		})
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			dispatch.MainQueueAsync(func() {
				progressIndicator.SetDoubleValue(0.1 * float64(i))
			})
		}
		dispatch.MainQueueAsync(func() {
			progressIndicator.StopAnimation(progressIndicator)
		})
	}()

	// text view & scroll view
	sv := appkit.NewVerticallyScrollableTextView(foundation.MakeRect(10, 200, 200, 30))
	w.ContentView().AddSubview(sv)
	/*
		Listing 4  Setting up a horizontal scroll bar

		[[theTextView enclosingScrollView] setHasHorizontalScroller:YES];
		[theTextView setHorizontallyResizable:YES];
		[theTextView setAutoresizingMask:(NSViewWidthSizable | NSViewHeightSizable)];
		[[theTextView textContainer] setContainerSize:NSMakeSize(FLT_MAX, FLT_MAX)];
		[[theTextView textContainer] setWidthTracksTextView:NO];
	*/

	w.WindowDidMove(func(notification foundation.Notification) {
		label.SetStringValue("moved!")
	})
	w.MakeKeyAndOrderFront(w)
	w.Center()

	appkit.Run()
}
