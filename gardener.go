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
	g.InitPump()
	g.InitLights()
	g.InitButtons()
	g.InitBME280()
	g.InitOLED()
	g.InitLEDs()
	g.InitApp()

	soil := g.InitSoil(g.Done())
	g.Subscribe(soil.Topic, g.MsgHandler)

}

func (g *Gardener) Start() error {
	return g.OttO.Start()
}

func (g *Gardener) Stop() {
	g.OttO.Stop()
}

func (g *Gardener) MsgHandler(msg *messanger.Msg) {

	fmt.Printf("MSG: %#v\n", msg)

	moisture := msg.Float64()

	soil := g.GetSoil()
	pump := g.GetPump()
	blue := g.GetLED("blue")

	if soil.IsDry(moisture) && pump.IsOff() {
		blue.PubData("on") // turn the blue LED on
		pump.Start()
	} else if soil.IsWet(moisture) && pump.IsOn() {
		pump.Stop()
		blue.PubData("off") // turn the blue LED off
	}
}
