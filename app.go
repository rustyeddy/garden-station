package main

import (
	"embed"

	"github.com/sensorstation/otto/server"
)

//go:embed app
var content embed.FS

func startApp(done chan any) {
	s := server.GetServer()

	// The following line is commented out because
	// var data any
	// s.EmbedTempl("/", data, content)
	s.Appdir("/", "app")
	go s.Start(done)
}
