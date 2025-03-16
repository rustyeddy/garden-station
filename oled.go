package main

import (
	"log/slog"

	"github.com/sensorstation/otto/device/oled"
)

func (g *Gardener) InitOLED() {
	oled, err := oled.New("oled", 128, 64)
	if err != nil {
		slog.Error("Failed to initialize OLED", "error", err)
		return
	}
	g.AddDevice(oled)
}
