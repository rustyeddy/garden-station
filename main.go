package main

import (
	"flag"
	"time"

	"github.com/rustyeddy/otto"
)

var (
	stationName = "gardener"
	mock        = true
	period      = 15 * time.Minute
)

func init() {
	flag.BoolVar(&mock, "mock", false, "mock gpio")
}

func main() {
	flag.Parse()

	gardener := &Gardener{}
	otto := &otto.OttO{
		Name: "gardener",
		Mock: mock,
	}
	gardener.OttO = otto
	gardener.Init()
	gardener.Start()
	<-gardener.Done()
	gardener.Stop()
}
