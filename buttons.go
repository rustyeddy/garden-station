package main

import (
	"github.com/rustyeddy/devices/button"
	"github.com/rustyeddy/otto/messanger"
	"github.com/warthog618/go-gpiocdev"
)

func (g *Gardener) InitButtons() {
	onButton := button.New("on", 23, gpiocdev.WithRisingEdge)
	onManaged := g.AddManagedDevice("on", onButton, messanger.GetTopics().Control("button"))
	go onManaged.EventLoop(g.Done())

	offButton := button.New("off", 27, gpiocdev.WithRisingEdge)
	offManaged := g.AddManagedDevice("off", offButton, messanger.GetTopics().Control("button"))
	go offManaged.EventLoop(g.Done())
}
