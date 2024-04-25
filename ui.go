package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type MyTheme struct{}

var _ fyne.Theme = (*MyTheme)(nil)

func (m MyTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	if name == theme.ColorNameBackground {
		if variant == theme.VariantLight {
			return color.White
		}
		return color.Black
	}
	return theme.DefaultTheme().Color(name, variant)
}

func (m MyTheme) Icon(name fyne.ThemeIconName) fyne.Resource {

	return theme.DefaultTheme().Icon(name)
}

func (m MyTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (m MyTheme) Size(name fyne.ThemeSizeName) float32 {
	//return 22
	return theme.DefaultTheme().Size(name)
}

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

	clock.SoundSlider = widget.NewSlider(0, 1)

	content.Add(clock.StartStopButton)
	content.Add(clock.Start5MinuteBreakButton)
	content.Add(clock.Start20MinuteBreakButton)
	content.Add(clock.ResetButton)
	content.Add(clock.SoundSlider)

	clock.Reset(clock.MainWindow, "Go 🍅")

	return content

}

func (c *Pomodoro) UpdateStartStopButton(msg string, withPauseIcon bool) {
	if withPauseIcon {
		c.StartStopButton.SetIcon(theme.MediaPauseIcon())
	} else {
		c.StartStopButton.SetIcon(nil)
	}
	c.StartStopButton.SetText(msg)
}

func (c *Pomodoro) Render() *fyne.Container {

	co := container.NewVBox(c.TimeLabel)

	return co
}

func (c *Pomodoro) Layout(_ []fyne.CanvasObject, size fyne.Size) {
	diameter := fyne.Min(size.Width, size.Height)
	size = fyne.NewSize(diameter, diameter)
}
