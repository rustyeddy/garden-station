package main

import (
	"embed"
)

//go:embed app
var tmpldir embed.FS

func (g *Gardener) InitApp() {
	s := g.Server
	s.EmbedTempl("/", tmpldir, g)
}
