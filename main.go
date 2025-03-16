package main

import (
	"flag"
	"time"
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

	gardener := Gardener{}
	if mock {
		gardener.Mock(true)
	}
	gardener.Init()
	gardener.Start()
	<-gardener.Done()
	gardener.Stop()
}
