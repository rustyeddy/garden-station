package main

import (
	"time"

	"github.com/rustyeddy/otto/device/vh400"
	"github.com/rustyeddy/otto/messanger"
)

type Soil struct {
	*vh400.VH400
	WetThreshold float64
	DryThreshold float64
}

func (g *Gardener) InitSoil(done chan any) *Soil {
	// set up the VH400 vh400 moisture sensor and go into a timer loop
	soil := &Soil{
		VH400:        vh400.New("soil", 0),
		DryThreshold: 1.5,
		WetThreshold: 2.5,
	}
	soil.Topic = messanger.GetTopics().Data("soil")
	go soil.TimerLoop(1*time.Second, done, soil.ReadPub)
	g.AddDevice(soil)
	return soil
}

func (s *Soil) IsDry(vwc float64) bool {
	return vwc <= s.DryThreshold
}

func (s *Soil) IsWet(vwc float64) bool {
	return vwc >= s.WetThreshold
}
