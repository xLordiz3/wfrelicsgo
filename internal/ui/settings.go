package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type AppSettings struct {
	Theme string
}

type settings struct {
	fileLocPath *widget.Entry

	appSettings *AppSettings
	window      *fyne.Window
	app         *fyne.App
}

func newSettings(a fyne.App, w fyne.Window, as *AppSettings) *settings {
	return &settings{app: a, window: w, appSettings: as}
}

func (s *settings) buildUI() *container.Scroll {
	s.fileLocPath = &widget.Entry{}

	filepContainer := container.NewGridWithColumns(2,
		newSettingLabel("Files Directory"), s.fileLocPath,
	)
	filepGroup := widget.NewCard("Path Settings", "", filepContainer)

	return container.NewScroll(container.NewVBox(
		filepGroup,
	))
}

func (s *settings) tabItem() *container.TabItem {
	return container.NewTabItem("Settings", s.buildUI())
}

func newSettingLabel(text string) *widget.Label {
	return &widget.Label{Text: text}
}
