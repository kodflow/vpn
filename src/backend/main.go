package main

import (
	_ "embed"
	"log"
	"time"

	"github.com/kodflow/vpn/src/backend/ui"
)

func main() {
	app := ui.VPN()

	go func() {
		for {
			now := time.Now().Format(time.RFC1123)
			app.EmitEvent("time", now)
			time.Sleep(time.Second)
		}
	}()

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
