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
	s.EmbedTempl("/", tmpldir, g)
	// s.EmbedExec(g)
}
