package main

import (
	"fmt"

	"github.com/sensorstation/otto"
	"github.com/sensorstation/otto/messanger"
)

type Gardener struct {
	*otto.OttO
}

func (g *Gardener) Init() {

	g.OttO.Init()

	// XXX make all of these functions panic on error. That will force
	// the decision on how to handle errors
	display, err := InitDisplay()
	if err != nil {
		fmt.Printf("Failed to initialize display")
	}
	g.AddDevice(display)

	g.InitPump()
	g.InitLights()
	g.InitButtons()
	g.InitLEDs()
	g.InitApp()

	bme, err := InitBME280(g.Done())
	if err != nil {
		fmt.Printf("Failed to initialize bme280")
	}
	g.AddDevice(bme)

	soil := g.InitSoil(g.Done())
	g.Subscribe(soil.Topic, g.MsgHandler)
}

func (g *Gardener) Start() error {
	return g.OttO.Start()
}

func (g *Gardener) Stop() {
	g.OttO.Stop()
}

func (g *Gardener) GetSoil() *Soil {
	soil := g.GetDevice("soil").(*Soil)
	return soil
}

func (g *Gardener) MsgHandler(msg *messanger.Msg) {

	top := msg.Path[3]
	switch top {
	case "soil":
		g.handleSoil(msg)
	}
}

func (g *Gardener) handleSoil(msg *messanger.Msg) {
	moisture := msg.Float64()
	soil := g.GetSoil()
	pump := g.GetPump()

	if soil.IsDry(moisture) && pump.IsOff() {
		pump.Start()
	} else if soil.IsWet(moisture) && pump.IsOn() {
		pump.Stop()
	}
}

func (g *Gardener) IsPumpOn() bool {
	pump := g.GetPump()
	return pump.IsOn()
}
