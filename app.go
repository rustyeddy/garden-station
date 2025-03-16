package main

import (
	"embed"
)

//go:embed app
var content embed.FS

func startApp(g *Gardener, done chan any) {
	// s := server.GetServer()

	// The following line is commented out because
	// s.EmbedTempl("/", content, g)
	// s.AppTempl("/", "app/index.html", g)
	// s.Appdir("/", "app")
	go g.Server.Start(done)
}
