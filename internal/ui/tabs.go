package ui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/v2/container"
)

func Create(app fyne.App, window fyne.Window) *container.AppTabs {
	appSettings := &AppSettings{}

	return container.NewAppTabs(
		newRelicSelect(app, window, appSettings).tabItem(),
		newSettings(app, window, appSettings).tabItem(),
	)
}
