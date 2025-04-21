package main

import (
	"encoding/json"
	"fmt"

	"github.com/sensorstation/otto"
	"github.com/sensorstation/otto/device/bme280"
	"github.com/sensorstation/otto/messanger"
)

type Gardener struct {
	*otto.OttO
}

func (g *Gardener) Init() {

	g.OttO.Init()

	// XXX make all of these functions panic on error. That will force
	// the decision on how to handle errors
	g.InitDisplay()
	g.InitPump()
	g.InitLights()
	g.InitButtons()
	g.InitLEDs()
	g.InitApp()
	bme := g.InitBME280()

	soil := g.InitSoil(g.Done())
	g.Subscribe(soil.Topic, g.MsgHandler)
	g.Subscribe(bme.Topic, g.MsgHandler)
}

func (g *Gardener) Start() error {
	g.Display("title", "Gardener")
	return g.OttO.Start()
}

func (g *Gardener) Stop() {
	g.OttO.Stop()
}

func (g *Gardener) Display(topic string, val any) {
	d := g.GetDevice("display").(*Display)
	d.displayQ <- DisplayMsg{topic, val}
}

func (g *Gardener) MsgHandler(msg *messanger.Msg) {

	top := msg.Path[3]
	switch top {
	case "soil":
		g.handleSoil(msg)

	case "env":
		g.handleEnv(msg)
	}
}

func (g *Gardener) handleSoil(msg *messanger.Msg) {
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

	pstr := "off"
	if g.IsPumpOn() {
		pstr = "on"
	}

	g.Display("soil", fmt.Sprintf("%5.2f", moisture))
	g.Display("pump", pstr)

}

func (g *Gardener) handleEnv(msg *messanger.Msg) {
	// fmt.Printf("env msg: %+v\n", msg)
	var env bme280.Env
	err := json.Unmarshal(msg.Data, &env)
	if err != nil {
		fmt.Printf("Failed to unmarshal: %+v\n", err)
		return
	}
	g.Display("env", env)
}

func (g *Gardener) IsPumpOn() bool {
	pump := g.GetPump()
	return pump.IsOn()
}
