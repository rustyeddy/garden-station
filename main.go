package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/rustyeddy/devices"
	"github.com/rustyeddy/otto"
)

var (
	stationName = "gardener"
	mock        = false
	useLocal    = false
	mqttBroker  = "test.mosquitto.org"
)

func init() {
	flag.BoolVar(&mock, "mock", false, "mock gpio")
	flag.BoolVar(&useLocal, "local", false, "use local messaging (no MQTT)")
	flag.StringVar(&mqttBroker, "mqtt-broker", mqttBroker, "MQTT broker address (default: test.mosquitto.org)")
}

func main() {
	flag.Parse()

	log.Printf("starting %s (mock=%v, local=%v, mqtt=%q)", stationName, mock, useLocal, mqttBroker)

	// Enable mocking in devices if mock flag is set
	if mock {
		devices.SetMock(true)
	}

	gardener := &Gardener{}
	oTTo := &otto.OttO{
		Name:       stationName,
		Mock:       mock,
		UseLocal:   useLocal,
		MQTTBroker: mqttBroker,
	}
	gardener.OttO = oTTo
	gardener.Init()
	if err := gardener.Start(); err != nil {
		log.Printf("failed to start gardener: %v", err)
	}

	// Handle OS signals and call Stop() for graceful shutdown
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-signals
		log.Printf("received signal %s, stopping gardener", sig)
		gardener.Stop()
	}()

	<-gardener.Done()
	gardener.Stop()
	log.Println("gardener stopped")
}
