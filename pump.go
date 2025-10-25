package main

import (
	"log/slog"

	"github.com/rustyeddy/devices/relay"
	"github.com/rustyeddy/otto/messanger"
)

type Pump struct {
	*relay.Relay
}

func (p *Pump) ID() string {
	return "pump"
}

func (g *Gardener) InitPump() {
	// setup the pump and subscribe to the pump value
	pumpRelay := relay.New("pump", 23)
	pump := &Pump{Relay: pumpRelay}

	pumpManaged := g.AddManagedDevice("pump", pump, messanger.GetTopics().Control("pump"))
	pumpManaged.Subscribe(messanger.GetTopics().Control("pump"), pump.Callback)
}

func (g *Gardener) GetPump() *Pump {
	md := g.GetManagedDevice("pump")
	if md == nil {
		return nil
	}
	pump := md.Device.(*Pump)
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
	slog.Info("Starting pump")
	p.Set(1) // Turn on the relay
}

func (p *Pump) Stop() {
	slog.Info("Stopping pump")
	p.Set(0) // Turn off the relay
}
