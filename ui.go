package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type UIElements struct {
	CountDownText            *CustomText
	CountDownMinute          *CustomText
	CountDownSecond          *CustomText
	StartStopButton          *widget.Button
	Start5MinuteBreakButton  *widget.Button
	Start20MinuteBreakButton *widget.Button
	ResetButton              *widget.Button
	QuitButton               *widget.Button
	SoundSliderLabel         *widget.Label
	SoundSlider              *widget.Slider
	NotificationSliderLabel  *widget.Label
	NotificationSlider       *widget.Slider
}

type CustomText struct {
	canvas.Text
}

var _ fyne.CanvasObject = (*CustomText)(nil)

func NewCustomText(text string, c color.Color) *CustomText {
	size := fyne.CurrentApp().Settings().Theme().Size("custom_text")
	nct := &CustomText{}
	nct.Text.Text = text
	nct.Text.TextSize = size
	nct.Text.Color = c

	return nct
}

func (t *CustomText) UpdateText(text string) {
	t.Text.Text = text
	t.Text.Refresh()
}

func (clock *Pomodoro) Show(stack *fyne.Container) fyne.CanvasObject {

	clock.UIElements.CountDownText = NewCustomText("25 Minutes", &color.RGBA{0, 109, 255, 255})
	clock.UIElements.CountDownText.TextStyle.Bold = true
	clock.UIElements.CountDownText.TextStyle.Monospace = true
	clock.UIElements.CountDownText.Alignment = fyne.TextAlignCenter

	content := clock.Render()

	clock.UIElements.StartStopButton = widget.NewButton("Start 🍅", func() {
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
	clock.UIElements.Start5MinuteBreakButton = widget.NewButton("Start 5 Minutes Break", func() {
		clock.Reset(clock.MainWindow, "Go 🍅: 5 Minutes pause running")
		clock.Countdown.Minute = 5
		clock.Countdown.Second = 0
		clock.UpdateStartStopButton("", true)
		clock.Stop = false
		go clock.Animate(content, clock.MainWindow)
	})
	clock.UIElements.Start20MinuteBreakButton = widget.NewButton("Start 20 Minutes Break", func() {
		clock.Reset(clock.MainWindow, "Go 🍅: 20 Minutes pause running")
		clock.Countdown.Minute = 20
		clock.Countdown.Second = 00
		clock.UpdateStartStopButton("", true)
		clock.Stop = false
		go clock.Animate(content, clock.MainWindow)
	})
	clock.UIElements.ResetButton = widget.NewButton("Reset ", func() {
		clock.Reset(clock.MainWindow, "Go 🍅")
	})
	clock.UIElements.QuitButton = widget.NewButton("Quit ", func() {
		clock.App.Quit()
	})

	clock.UIElements.SoundSliderLabel = widget.NewLabelWithStyle("Sound:", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	clock.UIElements.SoundSlider = widget.NewSlider(0, 1)
	clock.UIElements.SoundSlider.Bind(binding.BindPreferenceFloat("withSound", clock.App.Preferences()))
	clock.UIElements.NotificationSliderLabel = widget.NewLabelWithStyle("Notification:", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	clock.UIElements.NotificationSlider = widget.NewSlider(0, 1)
	clock.UIElements.NotificationSlider.Bind(binding.BindPreferenceFloat("withNotification", clock.App.Preferences()))

	content.Add(layout.NewSpacer())

	content.Add(clock.UIElements.StartStopButton)
	content.Add(clock.UIElements.Start5MinuteBreakButton)
	content.Add(clock.UIElements.Start20MinuteBreakButton)
	content.Add(clock.UIElements.ResetButton)
	content.Add(clock.UIElements.QuitButton)

	content.Add(layout.NewSpacer())
	content.Add(container.New(
		layout.NewGridLayout(2),
		clock.UIElements.SoundSliderLabel,
		clock.UIElements.NotificationSliderLabel,
		clock.UIElements.SoundSlider,
		clock.UIElements.NotificationSlider))

	clock.Reset(clock.MainWindow, "Go 🍅")

	return content

}

func (c *Pomodoro) UpdateStartStopButton(msg string, withPauseIcon bool) {
	if withPauseIcon {
		c.UIElements.StartStopButton.SetIcon(theme.MediaPauseIcon())
	} else {
		c.UIElements.StartStopButton.SetIcon(nil)
	}
	c.UIElements.StartStopButton.SetText(msg)
}

func (c *Pomodoro) Render() *fyne.Container {

	co := container.NewVBox(&c.UIElements.CountDownText.Text)

	return co
}

func (c *Pomodoro) Layout(_ []fyne.CanvasObject, size fyne.Size) {
	diameter := fyne.Min(size.Width, size.Height)
	size = fyne.NewSize(diameter, diameter)
}
