package events

import (
	"time"

	"github.com/wailsapp/wails/v3/pkg/application"
)

func Time(app *application.App) {
	if app == nil {
		panic("app is nil")
	}

	go func() {
		for {
			app.EmitEvent("time", time.Now().Format(time.RFC1123))
			time.Sleep(time.Second)
		}
	}()
}
