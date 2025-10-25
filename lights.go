package main

import (
	"github.com/rustyeddy/devices/relay"
	"github.com/rustyeddy/otto/messanger"
)

func (g *Gardener) InitLights() {
	lightsRelay := relay.New("lights", 8)
	lightsManaged := g.AddManagedDevice("lights", lightsRelay, messanger.GetTopics().Data("lights"))
	lightsManaged.Subscribe(messanger.GetTopics().Control("lights"), lightsRelay.Callback)
}
