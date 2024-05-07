package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	"github.com/gopxl/beep"
	"github.com/gopxl/beep/speaker"
	"github.com/gopxl/beep/wav"
)

type Pomodoro struct {
	App                      fyne.App
	MainWindow               fyne.Window
	CountDownText            *CustomText
	StartStopButton          *widget.Button
	Start5MinuteBreakButton  *widget.Button
	Start20MinuteBreakButton *widget.Button
	ResetButton              *widget.Button
	QuitButton               *widget.Button
	SoundSliderLabel         *widget.Label
	SoundSlider              *widget.Slider
	NotificationSliderLabel  *widget.Label
	NotificationSlider       *widget.Slider
	Countdown                Countdown
	Stop                     bool
}

type Countdown struct {
	Minute int64
	Second int64
}

func main() {
	var p Pomodoro
	a := app.NewWithID("ch.mauricext4fs.gopomodoro")
	p.App = a
	a.Settings().SetTheme(&MyTheme{})

	// Window
	p.MainWindow = a.NewWindow("Go 🍅")
	p.MainWindow.Resize(fyne.Size{Width: 290, Height: 350})
	p.MainWindow.CenterOnScreen()
	p.MainWindow.SetMaster()

	tomatoeIcon := resourceIconPng
	a.SetIcon(tomatoeIcon)

	if desk, ok := a.(desktop.App); ok {
		p.MainWindow.SetCloseIntercept(func() {
			p.MainWindow.Hide()
		})
		m := fyne.NewMenu("Go Pomodoro",
			fyne.NewMenuItem("Show", func() {
				p.MainWindow.Show()
			}))
		desk.SetSystemTrayMenu(m)
		tomatoeSystrayIcon, err := fyne.LoadResourceFromPath("icon_systray.png")
		if err == nil {
			desk.SetSystemTrayIcon(tomatoeSystrayIcon)
		}
	}
	c := container.NewStack()
	//c.Objects = []fyne.CanvasObject{Show(p.MainWindow)}
	//c.Objects = []fyne.CanvasObject{p.Show(c)}
	c.Add(p.Show(c))

	p.MainWindow.SetContent(c)
	p.MainWindow.ShowAndRun()
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

func (c *Pomodoro) Reset(win fyne.Window, newTitle string) {
	// Stop any existing counter (if any)
	c.Stop = true
	time.Sleep(1 * time.Second)
	c.Countdown.Minute = 24
	c.Countdown.Second = 59
	c.CountDownText.UpdateText("25 Minutes")

	c.UpdateStartStopButton("Start 🍅", false)
	if win != nil && newTitle != "" {
		fyne.Window.SetTitle(win, newTitle)
	}
}

func (c *Pomodoro) Animate(co fyne.CanvasObject, win fyne.Window) {
	tick := time.NewTicker(time.Second)
	go func() {
		for !c.Stop {
			c.Layout(nil, co.Size())
			<-tick.C
			c.CountdownDown()
			c.CountDownText.UpdateText(fmt.Sprintf("%d Minutes and %d Seconds", c.Countdown.Minute, c.Countdown.Second))
		}
		if c.Countdown.Minute == 0 && c.Countdown.Second == 0 {
			if c.App.Preferences().FloatWithFallback("withNotification", 1) == 1 {
				n := fyne.NewNotification("🍅 completed!", "🍅 completed!")
				app.New().SendNotification(n)
			}

			if c.App.Preferences().FloatWithFallback("withSound", 1) == 1 {
				PlayNotificationSound()
			}
			c.Reset(win, "Go 🍅")
		}
	}()
}

func (c *Pomodoro) CountdownDown() {
	c.Countdown.Second--
	if c.Countdown.Minute >= 1 && c.Countdown.Second <= 0 {
		c.Countdown.Minute--
		c.Countdown.Second = 59
	} else if c.Countdown.Minute == 0 && c.Countdown.Second <= 0 {
		c.Stop = true
	}
}
