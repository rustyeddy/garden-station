package main

import (
	"github.com/sensorstation/otto/device/button"
	"github.com/sensorstation/otto/messanger"
	"github.com/warthog618/go-gpiocdev"
)

func (g *Gardener) InitButtons() {
	on := button.New("on", 23, gpiocdev.WithRisingEdge)
	on.Topic = messanger.GetTopics().Control("button")
	go on.EventLoop(g.Done(), on.ReadPub)
	g.AddDevice(on)

	off := button.New("off", 27, gpiocdev.WithRisingEdge)
	off.Topic = messanger.GetTopics().Control("button")
	go off.EventLoop(g.Done(), off.ReadPub)
	g.AddDevice(off)
}
