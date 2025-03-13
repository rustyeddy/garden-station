package main

import (
	"embed"

	"github.com/sensorstation/otto/server"
)

//go:embed app
var content embed.FS

func startApp(g *gardener, done chan any) {
	s := server.GetServer()

	// The following line is commented out because
	// s.EmbedTempl("/", content, g)
	// s.AppTempl("/", "app/index.html", g)
	// s.Appdir("/", "app")
	go s.Start(done)
}
