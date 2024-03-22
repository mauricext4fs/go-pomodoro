package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type clock struct {
	canvas      fyne.CanvasObject
	MinuteLabel *widget.Label
	SecondLabel *widget.Label
	countdown   countdown
	stop        bool
}

type countdown struct {
	minute      int64
	second      int64
}

func main() {
	a := app.New()
	w := a.NewWindow("Go Pomodoro")
	c := container.NewStack();

	c.Add(widget.NewLabel("Remaining Time: "))

	w.Resize(fyne.Size{Width: 1024, Height: 768})
	w.CenterOnScreen()
	w.SetContent(c)
	w.ShowAndRun()
}

