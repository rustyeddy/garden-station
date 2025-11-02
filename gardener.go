package main

import (
	"fmt"
	"time"

	"github.com/rustyeddy/otto"
	"github.com/rustyeddy/otto/messanger"
)

type Gardener struct {
	*otto.OttO
}

func (g *Gardener) Init() {

	g.OttO.Init()

	display, err := InitDisplay()
	if err != nil {
		fmt.Printf("Failed to initialize display")
	} else {
		g.AddManagedDevice("display", display, messanger.GetTopics().Control("display"))

		// Set up custom message handler for display
		displayHandler := messanger.MsgHandler(func(msg *messanger.Msg) error {
			display.MsgHandler(msg)
			return nil
		})
		g.Subscribe(messanger.GetTopics().Control("display"), displayHandler)
	}
	g.InitPump()
	g.InitLights()
	g.InitButtons()
	g.InitLEDs()
	g.InitApp()

	bme, err := InitBME280(g.Done())
	if err != nil {
		fmt.Printf("Failed to initialize bme280")
	} else {
		bmeManaged := g.AddManagedDevice("env", bme, messanger.GetTopics().Data("env"))

		// Start timer loop for periodic sensor readings
		bmeManaged.StartTimerLoop(5*time.Second, g.Done())
	}

	g.InitSoil(g.Done())

	msgHandler := messanger.MsgHandler(func(msg *messanger.Msg) error {
		g.MsgHandler(msg)
		return nil
	})
	g.Subscribe(messanger.GetTopics().Data("soil"), msgHandler)
}

func (g *Gardener) Start() {
	g.OttO.Start()
}

func (g *Gardener) Stop() {
	g.OttO.Stop()
}

func (g *Gardener) GetSoil() *Soil {
	md := g.GetManagedDevice("soil")
	if md == nil {
		return nil
	}
	soil := md.Device.(*Soil)
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
