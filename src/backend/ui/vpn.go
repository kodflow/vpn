package ui

import (
	"github.com/kodflow/vpn/src/backend/components"
	"github.com/kodflow/vpn/src/frontend"
	"github.com/wailsapp/wails/v3/pkg/application"
)

var (
	app *application.App

	ENTRYPOINT_WINDOW = "primary"
)

func VPN() *application.App {
	if app == nil {
		app = application.New(application.Options{
			Name:        "vpn",
			Description: "A demo of using raw HTML & CSS",
			Services: []application.Service{
				application.NewService(&components.Hello{}),
			},
			Assets: application.AssetOptions{
				Handler: application.AssetFileServerFS(frontend.Assets),
			},
			Mac: application.MacOptions{
				ApplicationShouldTerminateAfterLastWindowClosed: true,
			},
		})

		Window(app)
	}

	return app
}

func Window(app *application.App) *application.WebviewWindow {
	if app == nil {
		panic("app is nil")
	}

	return app.NewWebviewWindowWithOptions(application.WebviewWindowOptions{
		Title: ENTRYPOINT_WINDOW,
		Mac: application.MacWindow{
			InvisibleTitleBarHeight: 50,
			Backdrop:                application.MacBackdropTranslucent,
			TitleBar:                application.MacTitleBarHiddenInset,
		},
		BackgroundColour: application.NewRGB(27, 38, 54),
		URL:              "/",
	})
}
