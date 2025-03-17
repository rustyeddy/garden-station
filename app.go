package main

import (
	"embed"
)

//go:embed app
var tmpldir embed.FS

//go:embed app/js/*.js
var jsdir embed.FS

//go:embed app/css/*.css
var cssdir embed.FS

func (g *Gardener) InitApp() {
	s := g.Server

	data := struct {
		Name string
		Age  int
	}{
		Name: "Burt",
		Age:  99,
	}
	s.EmbedTempl("/", tmpldir, data)

	// s.AppTempl("/", "app/index.html", g)
	// s.Appdir("/", "app")
	// go g.Server.Start(g.Done())
}
