package main

import (
	"flag"
	"time"

	"github.com/sensorstation/otto/devices"
	"github.com/sensorstation/otto/devices/relay"
	"github.com/sensorstation/otto/devices/vh400"
	"github.com/sensorstation/otto/messanger"
)

var (
	stationName = "gardener"
	mock        = true
)

func init() {
	flag.BoolVar(&mock, "mock", false, "mock gpio")
}

func main() {
	flag.Parse()

	// mock the GPIO device if we are not running on
	// a raspberry pi or similar gpio based computer
	if mock {
		devices.GetMockGPIO()
	}

	// make the done channel to properly terminate devices and loops
	done := make(chan bool)

	// set up the VH400 soil moisture sensor and go into a timer loop
	soil := vh400.New("soil", 4)
	soil.AddPub(messanger.TopicData("soil"))
	soil.Period = 1 * time.Second
	go soil.TimerLoop(done, soil.ReadPub)

	// setup the pump and subscribe to the pump value
	pump := relay.New("pump", 5)
	pump.AddPub(messanger.TopicData("pump"))
	pump.Subscribe(messanger.TopicControl("pump"), pump.Callback)

	// create the controller that will subscribe to the soil moisture
	// data then make the decision to either ingore the data or send
	// a message to the pump to either turn on or off
	ctl := &controller{
		soil: soil,
		pump: pump,
		done: done,
	}
	ctl.Subscribe(soil.Pub, ctl.MsgHandler)

	ms := messanger.GetMsgSaver()
	ms.Saving = true
	for {
		time.Sleep(1 * time.Minute)
		ms.Dump()
	}
	ctl.Wait()
}

type controller struct {
	soil *vh400.VH400
	pump *relay.Relay
	done chan bool
}

func (c *controller) Wait() {
	<-c.done
}

func (c *controller) Subscribe(topic string, cb func(msg *messanger.Msg)) {
	mqtt := messanger.GetMQTT()
	mqtt.Subscribe(topic, cb)
}

func (c *controller) MsgHandler(msg *messanger.Msg) {
	topic := c.pump.Subs[0]
	val := msg.Float64()
	pval, err := c.pump.Get()

	mqtt := messanger.GetMQTT()
	if err != nil {
		mqtt.Publish(topic, "off")
		return
	}

	if val < 60.0 && pval == 0 {
		mqtt.Publish(topic, "on")
	} else if val > 60 && pval == 1 {
		mqtt.Publish(topic, "off")
	}
}
