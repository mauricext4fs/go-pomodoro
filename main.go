package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"time"
)

type clock struct {
	canvas    fyne.CanvasObject
	timeLabel *widget.Label
	countdown countdown
	stop      bool
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

	content := clock.render()
	go clock.animate(content)

	return content
}

func (c *clock) render() *fyne.Container {

	container := container.NewStack(c.timeLabel)

	return container
}

func (c *clock) animate(co fyne.CanvasObject) {
	tick := time.NewTicker(time.Second)
	go func() {
		c.countdown.minute = 24
		c.countdown.second = 59
		for !c.stop {
			c.Layout(nil, co.Size())
			<-tick.C
			c.countdownDown(&c.countdown)
			c.timeLabel.SetText(fmt.Sprintf("%d Minutes and %d Seconds", c.countdown.minute, c.countdown.second))
		}
	}()
}

func (c *clock) Layout(_ []fyne.CanvasObject, size fyne.Size) {
	diameter := fyne.Min(size.Width, size.Height)
	size = fyne.NewSize(diameter, diameter)
	fmt.Println(c.countdown.minute, " : ", c.countdown.second)
}

func (c *clock) countdownDown(cd *countdown) {
	if cd.minute > 0 && cd.second == 1 {
		cd.minute--
		cd.second = 60
	} else if cd.minute == 0 && cd.second <= 1 {
		c.stop = true
	}
	cd.second--
}
