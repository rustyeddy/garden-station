package main

import (
	"log/slog"

	"github.com/sensorstation/otto/device/relay"
	"github.com/sensorstation/otto/messanger"
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
		l.Topic = messanger.GetTopics().Data(n.name)
		l.Subscribe(messanger.GetTopics().Control(n.name), l.Callback)
		g.AddDevice(l)
	}
}

func (g *Gardener) GetLED(color string) *relay.Relay {
	d := g.GetDevice(color)
	if d == nil {
		slog.Error("LED not found", "color", color)
		return nil
	}
	led := d.(*relay.Relay)
	return led

}
