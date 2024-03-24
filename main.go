package main

import (
	"fmt"
	"log"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type clock struct {
	timeLabel   *widget.Label
	actionButon *widget.Button
	countdown   countdown
	stop        bool
}

type countdown struct {
	minute int64
	second int64
}

func main() {
	a := app.New()
	w := a.NewWindow("Go Pomodoro")
	c := container.NewStack()

	c.Objects = []fyne.CanvasObject{Show(w)}

	w.Resize(fyne.Size{Width: 1024, Height: 768})
	w.CenterOnScreen()
	w.SetContent(c)
	w.ShowAndRun()
}

func Show(win fyne.Window) fyne.CanvasObject {
	clock := &clock{}
	clock.timeLabel = widget.NewLabel("25 Minutes")
	clock.timeLabel.TextStyle.Bold = true
	clock.timeLabel.Importance = widget.HighImportance
	clock.stop = true

	content := clock.render()
	clock.actionButon = widget.NewButton("Start 🍎", func() {
		if clock.stop {
			log.Println("Starting 🍎")
			clock.actionButon.SetText("Stop 🍎")
			clock.stop = false
			go clock.animate(content)
		} else {
			log.Println("Stopping 🍎")
			clock.actionButon.SetText("Start 🍎")
			clock.stop = true
		}
	})
	content.Add(clock.actionButon)

	return content
}

func (c *clock) render() *fyne.Container {

	co := container.NewVBox(c.timeLabel)

	return co
}

func (c *clock) animate(co fyne.CanvasObject) {
	tick := time.NewTicker(time.Second)
	go func() {
		// Easier for testing
		//c.countdown.minute = 24
		c.countdown.minute = 1
		c.countdown.second = 60
		for !c.stop {
			c.Layout(nil, co.Size())
			<-tick.C
			c.countdownDown(&c.countdown)
			c.timeLabel.SetText(fmt.Sprintf("%d Minutes and %d Seconds", c.countdown.minute, c.countdown.second))
			fmt.Println(c.countdown.minute, " : ", c.countdown.second)
		}
		n := fyne.NewNotification("🍎 is over!", "🍎 is over")
		app.New().SendNotification(n)
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
		cd.second = 60
	} else if cd.minute == 0 && cd.second <= 0 {
		c.stop = true
	}
}
