package main

import (
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/rustyeddy/devices/bme280"
	"github.com/rustyeddy/devices/oled"
	"github.com/rustyeddy/otto/messanger"
)

var (
	s Screen
)

type Display struct {
	*oled.OLED
}

func (d *Display) ID() string {
	return "display"
}

type Screen struct {
	Title    string
	Soil     string
	Temp     string
	Humidity string
	Pressure string
	Pump     string
}

func InitDisplay() (*Display, error) {
	o, err := oled.New("display", 128, 64)
	if err != nil {
		slog.Error("Failed to initialize OLED DISPLAY", "error", err)
		return nil, err
	}
	display := &Display{
		OLED: o,
	}
	return display, nil
}

func (d *Display) MsgHandler(msg *messanger.Msg) {
	defer d.Draw()
	d.Clear()

	topic := msg.Path[3]
	fmt.Printf("TOPIC: %s\n", topic)

	switch topic {
	case "title":
		s.Title = msg.String()

	case "soil":
		s.Soil = msg.String()

	case "pump":
		s.Pump = msg.String()

	case "env":
		var env bme280.Env
		err := json.Unmarshal(msg.Data, &env)
		if err != nil {
			fmt.Printf("Failed to unmarshal: %+v\n", err)
			return
		}

		s.Temp = env.Temperature
		s.Humidity = env.Humidity
		s.Pressure = env.Pressure
	}
	d.DrawScreen()
}

func (d *Display) DrawScreen() {
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
