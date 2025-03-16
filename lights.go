package main

import (
	"github.com/sensorstation/otto/device/relay"
	"github.com/sensorstation/otto/messanger"
)

func (g *Gardener) InitLights() {
	lights := relay.New("lights", 8)
	lights.Topic = messanger.TopicData("lights")
	lights.Subscribe(messanger.TopicControl("lights"), lights.Callback)
	g.AddDevice(lights)
}
