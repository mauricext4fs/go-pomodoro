package main

import (
	"fmt"
	"image/color"
	"log"
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/gopxl/beep"
	"github.com/gopxl/beep/speaker"
	"github.com/gopxl/beep/wav"
)

type myTheme struct{}

var _ fyne.Theme = (*myTheme)(nil)

type Pomodoro struct {
	timeLabel                *widget.Label
	startstopButton          *widget.Button
	start5MinuteBreakButton  *widget.Button
	start20MinuteBreakButton *widget.Button
	resetButton              *widget.Button
	countdown                Countdown
	stop                     bool
}

type Countdown struct {
	minute int64
	second int64
}

func main() {
	a := app.New()
	a.Settings().SetTheme(&myTheme{})
	w := a.NewWindow("Go 🍅")

	tomatoeIcon, err := fyne.LoadResourceFromPath("icon.png")
	if err == nil {
		a.SetIcon(tomatoeIcon)
	}
	if desk, ok := a.(desktop.App); ok {
		log.Println("On Desktop!!")
		w.SetCloseIntercept(func() {
			w.Hide()
		})
		m := fyne.NewMenu("Go Pomodoro",
			fyne.NewMenuItem("Show", func() {
				w.Show()
			}))
		desk.SetSystemTrayMenu(m)
		tomatoeSystrayIcon, err := fyne.LoadResourceFromPath("icon_systray.png")
		if err == nil {
			desk.SetSystemTrayIcon(tomatoeSystrayIcon)
		}
	}
	c := container.NewStack()
	c.Objects = []fyne.CanvasObject{Show(w)}

	w.Resize(fyne.Size{Width: 400, Height: 300})
	w.CenterOnScreen()
	w.SetContent(c)
	w.ShowAndRun()
}

func PlayNotificationSound() {
	f, err := os.Open("notification.wav")
	if err != nil {
		log.Fatal("Error: ", err)
	}

	streamer, format, err := wav.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done
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
	var clock Pomodoro
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
		clock.reset(win, "Go 🍅: 5 Minutes pause running")
		clock.countdown.minute = 5
		clock.countdown.second = 0
		clock.updateStartstopButton("", true)
		clock.stop = false
		go clock.animate(content, win)
	})
	clock.start20MinuteBreakButton = widget.NewButton("Start 20 Minutes Break", func() {
		clock.reset(win, "Go 🍅: 20 Minutes pause running")
		clock.countdown.minute = 20
		clock.countdown.second = 00
		clock.updateStartstopButton("", true)
		clock.stop = false
		go clock.animate(content, win)
	})
	clock.resetButton = widget.NewButton("Reset ", func() {
		clock.reset(win, "Go 🍅")
	})
	content.Add(clock.startstopButton)
	content.Add(clock.start5MinuteBreakButton)
	content.Add(clock.start20MinuteBreakButton)
	content.Add(clock.resetButton)

	clock.reset(win, "Go 🍅")

	return content
}

func (c *Pomodoro) updateStartstopButton(msg string, withPauseIcon bool) {
	if withPauseIcon {
		c.startstopButton.SetIcon(theme.MediaPauseIcon())
	} else {
		c.startstopButton.SetIcon(nil)
	}
	c.startstopButton.SetText(msg)
}

func (c *Pomodoro) render() *fyne.Container {

	co := container.NewVBox(c.timeLabel)

	return co
}

func (c *Pomodoro) reset(win fyne.Window, newTitle string) {
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

func (c *Pomodoro) animate(co fyne.CanvasObject, win fyne.Window) {
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
			PlayNotificationSound()
			c.reset(win, "Go 🍅")
		}
	}()
}

func (c *Pomodoro) Layout(_ []fyne.CanvasObject, size fyne.Size) {
	diameter := fyne.Min(size.Width, size.Height)
	size = fyne.NewSize(diameter, diameter)
}

func (c *Pomodoro) countdownDown(cd *Countdown) {
	cd.second--
	if cd.minute >= 1 && cd.second <= 0 {
		cd.minute--
		cd.second = 59
	} else if cd.minute == 0 && cd.second <= 0 {
		c.stop = true
	}
}
