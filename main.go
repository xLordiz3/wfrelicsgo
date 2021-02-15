package main

import (
	"fyne.io/fyne/v2/app"
	"github.com/xLordiz3/wfrelicsgo/internal/ui"
)

func main() {
	a := app.New()
	w := a.NewWindow("WFRelics Go")
	w.SetContent(ui.Create(a, w))
	w.ShowAndRun()
}
