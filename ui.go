package main

import (
	"go-pomodoro/repository"
	"image/color"
	"log"
	"time"

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
	StartStopButton          *widget.Button
	Start5MinuteBreakButton  *widget.Button
	Start20MinuteBreakButton *widget.Button
	ResetButton              *widget.Button
	QuitButton               *widget.Button
	PomodoroCountLabel       *widget.Label
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

func (p *Pomodoro) Show(stack *fyne.Container) fyne.CanvasObject {

	p.UIElements.CountDownText = NewCustomText("25 Minutes", &color.RGBA{0, 109, 255, 255})
	p.UIElements.CountDownText.TextStyle.Bold = true
	p.UIElements.CountDownText.TextStyle.Monospace = true
	p.UIElements.CountDownText.Alignment = fyne.TextAlignCenter

	content := p.Render()

	p.UIElements.StartStopButton = widget.NewButton("Start 🍅", func() {
		if p.Stop {
			result, err := p.DB.StartActivity(repository.Activities{ActivityType: 100, StartTimestamp: time.Now()})
			if err != nil {
				log.Fatal("Error adding activity to sqlite DB: ", err)
			}
			p.ID = result.ID
			fyne.Window.SetTitle(p.MainWindow, "Go 🍅: Pomodoro running")
			p.UpdateStartStopButton("", true)
			p.Stop = false
			go p.Animate(content, p.MainWindow)
		} else {
			fyne.Window.SetTitle(p.MainWindow, "Go 🍅: Paused")
			p.UpdateStartStopButton("Continue", false)
			p.Stop = true
		}
	})
	p.UIElements.Start5MinuteBreakButton = widget.NewButton("Start 5 Minutes Break", func() {
		result, err := p.DB.StartActivity(repository.Activities{ActivityType: 200, StartTimestamp: time.Now()})
		if err != nil {
			log.Fatal("Error adding activity to sqlite DB: ", err)
		}
		p.ID = result.ID

		p.Reset(p.MainWindow, "Go 🍅: 5 Minutes pause running")
		p.Countdown.Minute = 5
		p.Countdown.Second = 0
		p.UpdateStartStopButton("", true)
		p.Stop = false
		go p.Animate(content, p.MainWindow)
	})
	p.UIElements.Start20MinuteBreakButton = widget.NewButton("Start 20 Minutes Break", func() {
		result, err := p.DB.StartActivity(repository.Activities{ActivityType: 500, StartTimestamp: time.Now()})
		if err != nil {
			log.Fatal("Error adding activity to sqlite DB: ", err)
		}
		p.ID = result.ID
		p.Reset(p.MainWindow, "Go 🍅: 20 Minutes pause running")
		p.Countdown.Minute = 20
		p.Countdown.Second = 00
		p.UpdateStartStopButton("", true)
		p.Stop = false
		go p.Animate(content, p.MainWindow)
	})
	p.UIElements.ResetButton = widget.NewButton("Reset ", func() {
		p.Reset(p.MainWindow, "Go 🍅")
	})
	p.UIElements.QuitButton = widget.NewButton("Quit ", func() {
		p.App.Quit()
	})

	p.UIElements.PomodoroCountLabel = widget.NewLabelWithStyle("Completed Pomodoro: 0", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	p.UpdatePomodoroCount()

	p.UIElements.SoundSliderLabel = widget.NewLabelWithStyle("Sound:", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	p.UIElements.SoundSlider = widget.NewSlider(0, 1)
	p.UIElements.SoundSlider.Bind(binding.BindPreferenceFloat("withSound", p.App.Preferences()))
	p.UIElements.NotificationSliderLabel = widget.NewLabelWithStyle("Notification:", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	p.UIElements.NotificationSlider = widget.NewSlider(0, 1)
	p.UIElements.NotificationSlider.Bind(binding.BindPreferenceFloat("withNotification", p.App.Preferences()))

	content.Add(layout.NewSpacer())

	content.Add(p.UIElements.StartStopButton)
	content.Add(p.UIElements.Start5MinuteBreakButton)
	content.Add(p.UIElements.Start20MinuteBreakButton)
	content.Add(p.UIElements.ResetButton)
	content.Add(p.UIElements.QuitButton)

	content.Add(layout.NewSpacer())
	content.Add(p.UIElements.PomodoroCountLabel)
	content.Add(container.New(
		layout.NewGridLayout(2),
		p.UIElements.SoundSliderLabel,
		p.UIElements.NotificationSliderLabel,
		p.UIElements.SoundSlider,
		p.UIElements.NotificationSlider))

	p.Reset(p.MainWindow, "Go 🍅")

	return content

}

func (p *Pomodoro) UpdateStartStopButton(msg string, withPauseIcon bool) {
	if withPauseIcon {
		p.UIElements.StartStopButton.SetIcon(theme.MediaPauseIcon())
	} else {
		p.UIElements.StartStopButton.SetIcon(nil)
	}
	p.UIElements.StartStopButton.SetText(msg)
}

func (p *Pomodoro) Render() *fyne.Container {

	c := container.NewVBox(&p.UIElements.CountDownText.Text)

	return c
}

func (p *Pomodoro) Layout(_ []fyne.CanvasObject, size fyne.Size) {
	diameter := fyne.Min(size.Width, size.Height)
	size = fyne.NewSize(diameter, diameter)
}
