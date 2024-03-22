package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Go Pomodoro")


	w.Resize(fyne.Size{Width: 1024, Height: 768})
	w.SetContent(widget.NewLabel("Remaining Time: "))
	w.ShowAndRun()
}

