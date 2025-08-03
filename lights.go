package main

import (
	"github.com/rustyeddy/otto/device/relay"
	"github.com/rustyeddy/otto/messanger"
)

func (g *Gardener) InitLights() {
	lights := relay.New("lights", 8)
	lights.Topic = messanger.GetTopics().Data("lights")
	lights.Subscribe(messanger.GetTopics().Control("lights"), lights.Callback)
	g.AddDevice(lights)
}
