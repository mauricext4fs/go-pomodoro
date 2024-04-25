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

type MyTheme struct{}

var _ fyne.Theme = (*MyTheme)(nil)

type Pomodoro struct {
	App                      fyne.App
	MainWindow               fyne.Window
	TimeLabel                *widget.Label
	StartStopButton          *widget.Button
	Start5MinuteBreakButton  *widget.Button
	Start20MinuteBreakButton *widget.Button
	ResetButton              *widget.Button
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
	p.MainWindow.Resize(fyne.Size{Width: 290, Height: 275})
	p.MainWindow.CenterOnScreen()
	p.MainWindow.SetMaster()

	tomatoeIcon, err := fyne.LoadResourceFromPath("icon.png")
	if err == nil {
		a.SetIcon(tomatoeIcon)
	}
	if desk, ok := a.(desktop.App); ok {
		log.Println("On Desktop!!")
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
	c.Objects = []fyne.CanvasObject{p.Show()}

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

func (c *Pomodoro) Reset(win fyne.Window, newTitle string) {
	// Stop any existing counter (if any)
	c.Stop = true
	time.Sleep(1 * time.Second)
	c.Countdown.Minute = 24
	c.Countdown.Second = 59
	c.TimeLabel.SetText("25 Minutes")

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
			c.CountdownDown(&c.Countdown)
			c.TimeLabel.SetText(fmt.Sprintf("%d Minutes and %d Seconds", c.Countdown.Minute, c.Countdown.Second))
		}
		if c.Countdown.Minute == 0 && c.Countdown.Second == 0 {
			n := fyne.NewNotification("🍅 completed!", "🍅 completed!")
			app.New().SendNotification(n)
			PlayNotificationSound()
			c.Reset(win, "Go 🍅")
		}
	}()
}

func (c *Pomodoro) Layout(_ []fyne.CanvasObject, size fyne.Size) {
	diameter := fyne.Min(size.Width, size.Height)
	size = fyne.NewSize(diameter, diameter)
}

func (c *Pomodoro) CountdownDown(cd *Countdown) {
	cd.Second--
	if cd.Minute >= 1 && cd.Second <= 0 {
		cd.Minute--
		cd.Second = 59
	} else if cd.Minute == 0 && cd.Second <= 0 {
		c.Stop = true
	}
}
