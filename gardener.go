package main

import (
	"fmt"
	"time"

	"github.com/sensorstation/otto/device"
	"github.com/sensorstation/otto/device/bme280"
	"github.com/sensorstation/otto/device/button"
	"github.com/sensorstation/otto/device/oled"
	"github.com/sensorstation/otto/device/relay"
	"github.com/sensorstation/otto/device/vh400"
	"github.com/sensorstation/otto/messanger"
	"github.com/sensorstation/otto/station"
	"github.com/warthog618/go-gpiocdev"
)

type gardener struct {
	*station.Station
}

func initGardener(name string, done chan any) *gardener {

	gardner := &gardener{
		Station: station.NewStation(name),
	}
	gardner.initVH400(done)
	gardner.initPump()
	gardner.initLights()
	gardner.initButtons(done)
	gardner.initBME280(done)
	gardner.initOLED()
	gardner.initController()
	gardner.initMessanger()
	return gardner
}

func (g *gardener) initVH400(done chan any) {
	// set up the VH400 vh400 moisture sensor and go into a timer loop
	vh400 := vh400.New("vh400", 0)
	vh400.AddPub(messanger.TopicData("vh400"))
	go vh400.TimerLoop(1*time.Second, done, vh400.ReadPub)
	g.AddDevice(vh400)
}

func (g *gardener) initPump() {
	// setup the pump and subscribe to the pump value
	pump := relay.New("pump", 5)
	pump.AddPub(messanger.TopicData("pump"))
	pump.Subscribe(messanger.TopicControl("pump"), pump.Callback)
	g.AddDevice(pump)
}

func (g *gardener) initLights() {
	lights := relay.New("lights", 8)
	lights.AddPub(messanger.TopicData("lights"))
	lights.Subscribe(messanger.TopicControl("lights"), lights.Callback)
	g.AddDevice(lights)
}

func (g *gardener) initButtons(done chan any) {
	on := button.New("on", 23, gpiocdev.WithRisingEdge)
	on.AddPub(messanger.TopicControl("button"))
	go on.EventLoop(done, on.ReadPub)
	g.AddDevice(on)

	off := button.New("off", 27, gpiocdev.WithRisingEdge)
	off.AddPub(messanger.TopicControl("button"))
	go off.EventLoop(done, off.ReadPub)
	g.AddDevice(off)
}

func (g *gardener) initBME280(done chan any) {
	bme := bme280.New("env", "/dev/i2c-1", 0x76)
	err := bme.Init()
	if err != nil {
		fmt.Printf("Failed to init BME280 %+v\n", err)
		return
	}

	bme.AddPub(messanger.TopicData("env"))
	go bme.TimerLoop(5*time.Second, done, bme.ReadPub)
	g.AddDevice(bme)
}

func (g *gardener) initOLED() {
	oled, _ := oled.New("oled", 128, 64)
	g.AddDevice(oled)
}

func (g *gardener) initController() {

	vh400 := g.GetDevice("vh400").(*vh400.VH400)

	controller := device.NewDevice("controller")
	controller.AddPub(messanger.TopicData("gardener"))
	controller.Subscribe(vh400.GetPub(), g.MsgHandler)
	g.AddDevice(controller)
}

func (g *gardener) initMessanger() {
	ms := messanger.GetMsgSaver()
	ms.Saving = true
}

func (g *gardener) MsgHandler(msg *messanger.Msg) {

	mqtt := messanger.GetMQTT()

	pump := g.GetDevice("pump").(*relay.Relay)

	topic := pump.GetSubs()[0]
	val := msg.Float64()
	pval, err := pump.Get()
	if err != nil {
		mqtt.Publish(topic, "off")
		return
	}

	if val < 3.0 && pval == 0 {
		mqtt.Publish(topic, "on")
	} else if val > 4.0 && pval == 1 {
		mqtt.Publish(topic, "off")
	}

}
