package main

import (
	"flag"
	"time"

	"github.com/sensorstation/otto/device"
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

func main() {
	flag.Parse()

	// mock the GPIO device if we are not running on
	// a raspberry pi or similar gpio based computer
	if mock {
		device.Mock(true)
	}

	done := make(chan any)
	gardner := initGardener(stationName, done)
	gardner.Start()
	<-done
	gardner.Stop()
}
