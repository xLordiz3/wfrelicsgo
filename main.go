package main

import (
	"./internal/ui"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	w := a.NewWindow("WFRelics Go")
	w.SetContent(ui.Create(a, w))
	w.ShowAndRun()
}
