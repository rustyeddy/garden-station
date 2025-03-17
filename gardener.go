package main

import (
	"github.com/sensorstation/otto"
	"github.com/sensorstation/otto/messanger"
)

type Gardener struct {
	*otto.OttO
}

func (g *Gardener) Init() {

	// XXX make all of these functions panic on error. That will force
	// the decision on how to handle errors
	g.InitPump()
	g.InitLights()
	g.InitButtons()
	g.InitBME280()
	g.InitOLED()
	g.InitApp()

	soil := g.InitSoil(g.Done())
	g.Subscribe(soil.Topic, g.MsgHandler)

}

func (g *Gardener) SetOttO(o *otto.OttO) {
	g.OttO = o
}

func (g *Gardener) MsgHandler(msg *messanger.Msg) {
	moisture := msg.Float64()

	soil := g.GetSoil()
	pump := g.GetPump()

	if soil.IsDry(moisture) && pump.IsOff() {

		pump.Start()
	} else if soil.IsWet(moisture) && pump.IsOn() {
		pump.Stop()
	}
}
