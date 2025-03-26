package main

import (
	"github.com/sensorstation/otto/device/relay"
	"github.com/sensorstation/otto/messanger"
)

func (g *Gardener) InitLights() {
	lights := relay.New("lights", 8)
	lights.Topic = messanger.GetTopics().Data("lights")
	lights.Subscribe(messanger.GetTopics().Control("lights"), lights.Callback)
	g.AddDevice(lights)
}
