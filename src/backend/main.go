package main

import (
	_ "embed"
	"log"

	"github.com/kodflow/vpn/src/backend/ui"
)

func main() {
	if err := ui.App().Run(); err != nil {
		log.Fatal(err)
	}
}
