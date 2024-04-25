package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func (clock *Pomodoro) Show() fyne.CanvasObject {
	clock.TimeLabel = widget.NewLabelWithStyle("25 Minutes", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	clock.TimeLabel.Importance = widget.HighImportance

	content := clock.Render()
	clock.StartStopButton = widget.NewButton("Start 🍅", func() {
		if clock.Stop {
			fyne.Window.SetTitle(clock.MainWindow, "Go 🍅: Pomodoro running")
			clock.UpdateStartStopButton("", true)
			clock.Stop = false
			go clock.Animate(content, clock.MainWindow)
		} else {
			fyne.Window.SetTitle(clock.MainWindow, "Go 🍅: Paused")
			clock.UpdateStartStopButton("Continue", false)
			clock.Stop = true
		}
	})
	clock.Start5MinuteBreakButton = widget.NewButton("Start 5 Minutes Break", func() {
		clock.Reset(clock.MainWindow, "Go 🍅: 5 Minutes pause running")
		clock.Countdown.Minute = 5
		clock.Countdown.Second = 0
		clock.UpdateStartStopButton("", true)
		clock.Stop = false
		go clock.Animate(content, clock.MainWindow)
	})
	clock.Start20MinuteBreakButton = widget.NewButton("Start 20 Minutes Break", func() {
		clock.Reset(clock.MainWindow, "Go 🍅: 20 Minutes pause running")
		clock.Countdown.Minute = 20
		clock.Countdown.Second = 00
		clock.UpdateStartStopButton("", true)
		clock.Stop = false
		go clock.Animate(content, clock.MainWindow)
	})
	clock.ResetButton = widget.NewButton("Reset ", func() {
		clock.Reset(clock.MainWindow, "Go 🍅")
	})
	content.Add(clock.StartStopButton)
	content.Add(clock.Start5MinuteBreakButton)
	content.Add(clock.Start20MinuteBreakButton)
	content.Add(clock.ResetButton)

	clock.Reset(clock.MainWindow, "Go 🍅")

	return content

}
