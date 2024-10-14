package ui

import (
	"github.com/kodflow/vpn/src/backend/config"
	"github.com/kodflow/vpn/src/backend/ui/service"
	"github.com/kodflow/vpn/src/backend/ui/view"
	"github.com/kodflow/vpn/src/frontend"
	"github.com/wailsapp/wails/v3/pkg/application"
)

var (
	app *application.App
)

func App() *application.App {
	if app == nil {
		app = application.New(application.Options{
			Name:        config.APP_NAME,
			Description: config.APP_DESCRIPTION,
			Services: []application.Service{
				application.NewService(&service.Hello{}),
			},
			Assets: application.AssetOptions{
				Handler: application.AssetFileServerFS(frontend.Assets),
			},
			Mac: application.MacOptions{
				ApplicationShouldTerminateAfterLastWindowClosed: true,
			},
		})

		view.Entrypoint(app)
	}

	return app
}
