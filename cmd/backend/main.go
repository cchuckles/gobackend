package main

import (
	"log"

	"github.com/goodchuckle/gobackend"
)

func main() {
	log.Println("initializing backend...")

	app, err := gobackend.NewBackend()
	if err != nil {
		log.Fatalf("failed to initialize backend: %v", err)
	}

	if err := app.Start(); err != nil {
		app.Logger.Fatalf("server error: %v", err)
	}
}
