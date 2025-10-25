package main

import (
	"time"

	"github.com/rustyeddy/devices/vh400"
	"github.com/rustyeddy/otto/messanger"
)

type Soil struct {
	*vh400.VH400
	WetThreshold float64
	DryThreshold float64
}

func (s *Soil) ID() string {
	return "soil"
}

func (g *Gardener) InitSoil(done chan any) *Soil {
	// set up the VH400 vh400 moisture sensor and go into a timer loop
	soil := &Soil{
		VH400:        vh400.New("soil", 0),
		DryThreshold: 1.5,
		WetThreshold: 2.5,
	}

	soilManaged := g.AddManagedDevice("soil", soil, messanger.GetTopics().Data("soil"))

	// Check if the device has TimerLoop method, if so use it
	if timerLooper, ok := interface{}(soil.VH400).(interface {
		TimerLoop(time.Duration, chan any, func())
	}); ok {
		go timerLooper.TimerLoop(1*time.Second, done, soilManaged.ReadPub)
	}

	return soil
}

func (s *Soil) IsDry(vwc float64) bool {
	return vwc <= s.DryThreshold
}

func (s *Soil) IsWet(vwc float64) bool {
	return vwc >= s.WetThreshold
}
