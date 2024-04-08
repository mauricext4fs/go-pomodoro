package main

import (
	"fmt"
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type myTheme struct{}

var _ fyne.Theme = (*myTheme)(nil)

type clock struct {
	timeLabel                *widget.Label
	startstopButton          *widget.Button
	start5MinuteBreakButton  *widget.Button
	start20MinuteBreakButton *widget.Button
	resetButton              *widget.Button
	countdown                countdown
	stop                     bool
}

type countdown struct {
	minute int64
	second int64
}

func main() {
	a := app.New()
	a.Settings().SetTheme(&myTheme{})
	w := a.NewWindow("Go 🍅")
	c := container.NewStack()

	c.Objects = []fyne.CanvasObject{Show(w)}

	w.Resize(fyne.Size{Width: 400, Height: 300})
	w.CenterOnScreen()
	w.SetContent(c)
	w.ShowAndRun()
}

func (m myTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	if name == theme.ColorNameBackground {
		if variant == theme.VariantLight {
			return color.White
		}
		return color.Black
	}
	return theme.DefaultTheme().Color(name, variant)
}

func (m myTheme) Icon(name fyne.ThemeIconName) fyne.Resource {

	return theme.DefaultTheme().Icon(name)
}

func (m myTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (m myTheme) Size(name fyne.ThemeSizeName) float32 {
	//return 22
	return theme.DefaultTheme().Size(name)
}

func Show(win fyne.Window) fyne.CanvasObject {
	clock := &clock{}
	clock.timeLabel = widget.NewLabelWithStyle("25 Minutes", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	clock.timeLabel.Importance = widget.HighImportance

	content := clock.render()
	clock.startstopButton = widget.NewButton("Start 🍅", func() {
		if clock.stop {
			fyne.Window.SetTitle(win, "Go 🍅: Pomodoro running")
			clock.updateStartstopButton("", true)
			clock.stop = false
			go clock.animate(content, win)
		} else {
			fyne.Window.SetTitle(win, "Go 🍅: Paused")
			clock.updateStartstopButton("Continue", false)
			clock.stop = true
		}
	})
	clock.start5MinuteBreakButton = widget.NewButton("Start 5 Minutes Break", func() {
		clock.reset(nil, "")
		clock.countdown.minute = 5
		clock.countdown.second = 0
		clock.updateStartstopButton("", true)
		clock.stop = false
		fyne.Window.SetTitle(win, "Go 🍅: 5 Minutes pause running")
		go clock.animate(content, win)
	})
	clock.start20MinuteBreakButton = widget.NewButton("Start 20 Minutes Break", func() {
		clock.reset(nil, "")
		clock.countdown.minute = 20
		clock.countdown.second = 00
		clock.updateStartstopButton("", true)
		clock.stop = false
		fyne.Window.SetTitle(win, "Go 🍅: 20 Minutes pause running")
		go clock.animate(content, win)
	})
	clock.resetButton = widget.NewButton("Reset ", func() {
		fyne.Window.SetTitle(win, "Go 🍅")
		clock.reset(nil, "")
	})
	content.Add(clock.startstopButton)
	content.Add(clock.start5MinuteBreakButton)
	content.Add(clock.start20MinuteBreakButton)
	content.Add(clock.resetButton)

	clock.reset(nil, "")

	return content
}

func (c *clock) updateStartstopButton(msg string, withPauseIcon bool) {
	if withPauseIcon {
		c.startstopButton.SetIcon(theme.MediaPauseIcon())
	} else {
		c.startstopButton.SetIcon(nil)
	}
	c.startstopButton.SetText(msg)
}

func (c *clock) render() *fyne.Container {

	co := container.NewVBox(c.timeLabel)

	return co
}

func (c *clock) reset(win fyne.Window, newTitle string) {
	// Stop any existing counter (if any)
	c.stop = true
	time.Sleep(1 * time.Second)
	c.countdown.minute = 24
	c.countdown.second = 59
	c.timeLabel.SetText("25 Minutes")

	c.updateStartstopButton("Start 🍅", false)
	if win != nil && newTitle != "" {
		fyne.Window.SetTitle(win, newTitle)
	}
}

func (c *clock) animate(co fyne.CanvasObject, win fyne.Window) {
	tick := time.NewTicker(time.Second)
	go func() {
		for !c.stop {
			c.Layout(nil, co.Size())
			<-tick.C
			c.countdownDown(&c.countdown)
			c.timeLabel.SetText(fmt.Sprintf("%d Minutes and %d Seconds", c.countdown.minute, c.countdown.second))
		}
		if c.countdown.minute == 0 && c.countdown.second == 0 {
			n := fyne.NewNotification("🍅 completed!", "🍅 completed!")
			app.New().SendNotification(n)
			c.reset(win, "Go 🍅")
		}
	}()
}

func (c *clock) Layout(_ []fyne.CanvasObject, size fyne.Size) {
	diameter := fyne.Min(size.Width, size.Height)
	size = fyne.NewSize(diameter, diameter)
}

func (c *clock) countdownDown(cd *countdown) {
	cd.second--
	if cd.minute >= 1 && cd.second <= 0 {
		cd.minute--
		cd.second = 59
	} else if cd.minute == 0 && cd.second <= 0 {
		c.stop = true
	}
}
