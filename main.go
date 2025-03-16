package main

import (
	"flag"
	"time"

	"github.com/sensorstation/otto"
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

	gardener := &Gardener{}
	otto := otto.OttO{
		Name:       "gardener",
		Controller: gardener,
	}

	if mock {
		otto.SetMock(true)
	}
	otto.Init()
	otto.Start()
	<-otto.Done()
	otto.Stop()
}
