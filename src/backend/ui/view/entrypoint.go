package view

import (
	"github.com/kodflow/vpn/src/backend/ui/events"
	"github.com/wailsapp/wails/v3/pkg/application"
)

func Entrypoint(app *application.App) *application.WebviewWindow {
	if app == nil {
		panic("app is nil")
	}

	events.Time(app)

	return app.NewWebviewWindowWithOptions(application.WebviewWindowOptions{
		Title: "entrypoint",
		URL:   "/",
		Mac: application.MacWindow{
			InvisibleTitleBarHeight: 50,
			Backdrop:                application.MacBackdropTranslucent,
			TitleBar:                application.MacTitleBarHiddenInset,
		},
		BackgroundColour: application.NewRGB(27, 38, 54),
	})
}
