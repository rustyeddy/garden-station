package main

import (
	"log/slog"

	"github.com/sensorstation/otto/device/bme280"
	"github.com/sensorstation/otto/device/oled"
)

var (
	s Screen
)

type Screen struct {
	Title    string
	Soil     string
	Temp     string
	Humidity string
	Pressure string
	Pump     string
}

type Display struct {
	*oled.OLED
	displayQ chan DisplayMsg
}

type DisplayMsg struct {
	Topic string
	Value any
}

func (g *Gardener) InitDisplay() {
	o, err := oled.New("display", 128, 64)
	if err != nil {
		slog.Error("Failed to initialize OLED DISPLAY", "error", err)
		return
	}

	display := &Display{
		OLED:     o,
		displayQ: make(chan DisplayMsg),
	}

	g.AddDevice(display)
	go func() {
		for {
			select {
			case msg := <-display.displayQ:
				display.DrawMsg(msg)
			}
		}
	}()
	display.displayQ <- DisplayMsg{
		Topic: "title",
		Value: "Initializing...",
	}
}

func (d *Display) DrawMsg(msg DisplayMsg) {
	defer d.Draw()
	d.Clear()

	switch msg.Topic {
	case "title":
		s.Title = msg.Value.(string)

	case "soil":
		s.Soil = msg.Value.(string)

	case "pump":
		s.Pump = msg.Value.(string)

	case "env":
		e := msg.Value.(bme280.Env)
		s.Temp = e.Temperature
		s.Humidity = e.Humidity
		s.Pressure = e.Pressure

	}

	y := 15
	d.DrawString(0, y, s.Title)
	d.DrawString(85, y, s.Pump)
	y += 12
	d.DrawString(0, y, "soil")
	d.DrawString(80, y, s.Soil)
	y += 12
	d.DrawString(0, y, "temp")
	d.DrawString(80, y, s.Temp)
	y += 12
	d.DrawString(0, y, "humidity")
	d.DrawString(80, y, s.Humidity+"%")
	y += 12
	d.DrawString(0, y, "pressure")
	d.DrawString(66, y, s.Pressure)
}
