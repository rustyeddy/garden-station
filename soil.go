package main

import (
	"fmt"
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

	// Open the device for reading
	if err := soil.VH400.Open(); err != nil {
		fmt.Printf("Failed to open VH400 soil sensor: %v\n", err)
	}

	soilManaged := g.AddManagedDevice("soil", soil, messanger.GetTopics().Data("soil"))

	// Start timer loop for periodic soil moisture readings
	soilManaged.StartTimerLoop(1*time.Second, done)

	return soil
}

func (s *Soil) IsDry(vwc float64) bool {
	return vwc <= s.DryThreshold
}

func (s *Soil) IsWet(vwc float64) bool {
	return vwc >= s.WetThreshold
}
