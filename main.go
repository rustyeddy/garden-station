package main

import (
	"flag"
	"time"

	"github.com/rustyeddy/devices"
	"github.com/rustyeddy/otto"
)

var (
	stationName = "gardener"
	mock        = true
	period      = 15 * time.Minute
	useLocal    = false
	mqttBroker  = ""
)

func init() {
	flag.BoolVar(&mock, "mock", false, "mock gpio")
	flag.BoolVar(&useLocal, "local", false, "use local messaging (no MQTT)")
	flag.StringVar(&mqttBroker, "mqtt-broker", "", "MQTT broker address (default: test.mosquitto.org)")
}

func main() {
	flag.Parse()

	// Enable mocking in devices if mock flag is set
	if mock {
		devices.SetMock(true)
	}

	gardener := &Gardener{}
	otto := &otto.OttO{
		Name:       "gardener",
		Mock:       mock,
		UseLocal:   useLocal,
		MQTTBroker: mqttBroker,
	}
	gardener.OttO = otto
	gardener.Init()
	gardener.Start()
	<-gardener.Done()
	gardener.Stop()
}
