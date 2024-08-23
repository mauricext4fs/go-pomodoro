package main

//go:generate fyne bundle -o bundled.go icon.png
//go:generate fyne bundle -o bundled.go -append icon_systray.png
//go:generate fyne bundle -o bundled.go -append notification.wav

import (
	"database/sql"
	"fmt"
	"go-pomodoro/repository"
	"log"
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"

	_ "github.com/glebarez/go-sqlite"
)

type Pomodoro struct {
	App        fyne.App
	DB         repository.Repository
	MainWindow fyne.Window
	UIElements UIElements
	Countdown  Countdown
	Stop       bool
	ID         int64
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

	// DB
	sqlDB, err := p.connectSQL()
	if err != nil {
		log.Panic(err)
	}

	p.setupDB(sqlDB)

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
		desk.SetSystemTrayIcon(resourceIconsystrayPng)
	}
	c := container.NewStack()
	c.Add(p.Show(c))

	p.MainWindow.SetContent(c)
	p.MainWindow.ShowAndRun()
}

func (p *Pomodoro) Reset(win fyne.Window, newTitle string) {
	// Stop any existing counter (if any)
	p.Stop = true
	time.Sleep(1 * time.Second)
	p.Countdown.Minute = 24
	p.Countdown.Second = 59
	p.UIElements.CountDownText.UpdateText("25 Minutes")

	p.UpdateStartStopButton("Start 🍅", false)
	if win != nil && newTitle != "" {
		fyne.Window.SetTitle(win, newTitle)
	}
}

func (p *Pomodoro) UpdatePomodoroCount() {
	count, err := p.DB.CountCompletedPomodoro()
	if err != nil {
		log.Fatal("Error getting count of Pomodoro from sqlite DB: ", err)
	}
	p.UIElements.PomodoroCountLabel.SetText(fmt.Sprintf("Completed Pomodoro: %d", count))
}

func (p *Pomodoro) Animate(co fyne.CanvasObject, win fyne.Window) {
	tick := time.NewTicker(time.Second)
	go func() {
		for !p.Stop {
			p.Layout(nil, co.Size())
			<-tick.C
			p.CountdownDown()
			p.UIElements.CountDownText.UpdateText(fmt.Sprintf("%d Minutes and %d Seconds", p.Countdown.Minute, p.Countdown.Second))
		}
		if p.Countdown.Minute == 0 && p.Countdown.Second == 0 {
			err := p.DB.UpdateActivity(p.ID, repository.Activities{ID: p.ID, EndTimestamp: time.Now()})
			if err != nil {
				log.Fatal("Error updating activity to sqlite DB: ", err)
			}
			p.UpdatePomodoroCount()

			if p.App.Preferences().FloatWithFallback("withSound", 1) == 1 {
				PlayNotificationSound()
			}

			if p.App.Preferences().FloatWithFallback("withNotification", 0) == 1 {
				count, err := p.DB.CountCompletedPomodoro()
				if err != nil {
					n := fyne.NewNotification("🍅 finished!", "Another 🍅 completed. Congrats!")
					app.New().SendNotification(n)
				} else {
					nMsg := fmt.Sprintf("%d 🍅 completed! Congrats!", count)
					n := fyne.NewNotification("🍅 finished!", nMsg)
					app.New().SendNotification(n)
				}
			}

			p.Reset(win, "Go 🍅")
		}
	}()
}

func (p *Pomodoro) CountdownDown() {
	p.Countdown.Second--
	if p.Countdown.Minute >= 1 && p.Countdown.Second <= 0 {
		p.Countdown.Minute--
		p.Countdown.Second = 59
	} else if p.Countdown.Minute == 0 && p.Countdown.Second <= 0 {
		p.Stop = true
	}
}

func (p *Pomodoro) connectSQL() (*sql.DB, error) {
	path := ""

	if os.Getenv("DB_PATH") != "" {
		path = os.Getenv("DB_PATH")
	} else {
		path = p.App.Storage().RootURI().Path() + "/sql.db"
	}

	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (p *Pomodoro) setupDB(sqlDB *sql.DB) {
	p.DB = repository.NewSQLiteRepository(sqlDB)

	err := p.DB.Migrate()
	if err != nil {
		log.Panic(err)
	}
}
