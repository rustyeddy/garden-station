package main

import (
	"log/slog"

	"github.com/rustyeddy/devices/relay"
	"github.com/rustyeddy/otto/messanger"
)

func (g *Gardener) InitLEDs() {
	leds := []struct {
		pin  int
		name string
	}{
		{12, "white"},
		{1, "blue"},
		{7, "green"},
		{8, "yellow"},
		{25, "red"},
	}

	for _, n := range leds {
		l := relay.New(n.name, n.pin)
		ledManaged := g.AddManagedDevice(n.name, l, messanger.GetTopics().Control(n.name))
		ledManaged.Subscribe(messanger.GetTopics().Control(n.name), l.Callback)
	}
}

func (g *Gardener) GetLED(color string) *relay.Relay {
	md := g.GetManagedDevice(color)
	if md == nil {
		slog.Error("LED not found", "color", color)
		return nil
	}
	led := md.Device.(*relay.Relay)
	return led
}
