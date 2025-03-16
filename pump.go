package main

import (
	"github.com/sensorstation/otto/device/relay"
	"github.com/sensorstation/otto/messanger"
)

type Pump struct {
	*relay.Relay
}

func (g *Gardener) InitPump() {
	// setup the pubmp and subscribe to the pump value
	pump := &Pump{}
	pump.Relay = relay.New("pump", 5)
	pump.Topic = messanger.TopicControl("pump")
	pump.Subscribe(messanger.TopicControl("pump"), pump.Callback)
	g.AddDevice(pump)
}

func (g *Gardener) GetPump() *Pump {
	pump := g.GetDevice("pump").(*Pump)
	return pump
}

func (p *Pump) IsOn() bool {
	val, err := p.Get()
	on := (val == 1)
	if err != nil {
		return false
	}
	return on
}

func (p *Pump) IsOff() bool {
	off := !p.IsOn()
	return off
}

func (p *Pump) Start() {
	p.PubData("on")
}

func (p *Pump) Stop() {
	p.PubData("off")
}
