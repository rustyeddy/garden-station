package main

import (
	"flag"
	"os"
	"time"

	"github.com/sensorstation/otto/devices"
	"github.com/sensorstation/otto/devices/button"
	"github.com/sensorstation/otto/devices/oled"
	"github.com/sensorstation/otto/devices/relay"
	"github.com/sensorstation/otto/devices/vh400"
	"github.com/sensorstation/otto/messanger"
	"github.com/sensorstation/otto/station"
	"github.com/warthog618/go-gpiocdev"
)

var (
	stationName  = "gardener"
	stationCount = 4
	mock         = true
	period       = 15 * time.Minute
)

func init() {
	flag.BoolVar(&mock, "mock", false, "mock gpio")
}

type gardener struct {
	*station.Station
}

func main() {
	flag.Parse()
	os.Exit(0)

	// mock the GPIO device if we are not running on
	// a raspberry pi or similar gpio based computer
	if mock {
		devices.GetMockGPIO()
	}

	done := make(chan any)
	gardner := initGardner(stationName, done)
	gardner.Start()
	<-done
	gardner.Stop()
}

func initGardner(name string, done chan any) *gardener {

	gardner := &gardener{
		Station: station.NewStation(name),
	}

	// set up the VH400 soil moisture sensor and go into a timer loop
	soil := vh400.New("soil", 0)
	soil.AddPub(messanger.TopicData("soil"))
	go soil.TimerLoop(1*time.Second, done, soil.ReadPub)
	gardner.AddDevice(soil)

	// setup the pump and subscribe to the pump value
	pump := relay.New("pump", 5)
	pump.AddPub(messanger.TopicData("pump"))
	pump.Subscribe(messanger.TopicControl("pump"), pump.Callback)
	gardner.AddDevice(pump)

	lights := relay.New("lights", 8)
	lights.AddPub(messanger.TopicData("lights"))
	lights.Subscribe(messanger.TopicControl("lights"), lights.Callback)
	gardner.AddDevice(lights)

	on := button.New("on", 23, gpiocdev.WithRisingEdge)
	on.AddPub(messanger.TopicControl("button"))
	go on.EventLoop(done, on.ReadPub)
	gardner.AddDevice(on)

	off := button.New("off", 27, gpiocdev.WithRisingEdge)
	off.AddPub(messanger.TopicControl("button"))
	go off.EventLoop(done, off.ReadPub)
	gardner.AddDevice(off)

	oled, _ := oled.New("oled", 128, 64)
	gardner.AddDevice(oled)

	controller := devices.NewDevice("controller")
	controller.AddPub(messanger.TopicData("gardner"))
	controller.Subscribe(soil.GetPub(), gardner.MsgHandler)
	gardner.AddDevice(controller)

	ms := messanger.GetMsgSaver()
	ms.Saving = true

	return gardner
}

func (s *gardener) MsgHandler(msg *messanger.Msg) {

	// topic := c.pump.Subs[0]
	// val := msg.Float64()
	// pval, err := c.pump.Get()

	// mqtt := messanger.GetMQTT()
	// if err != nil {
	// 	mqtt.Publish(topic, "off")
	// 	return
	// }

	// if val < 60.0 && pval == 0 {
	// 	mqtt.Publish(topic, "on")
	// } else if val > 60 && pval == 1 {
	// 	mqtt.Publish(topic, "off")
	// }

}
