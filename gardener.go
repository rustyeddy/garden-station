package main

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/rustyeddy/devices"
	"github.com/rustyeddy/devices/button"
	"github.com/rustyeddy/devices/env"
	"github.com/rustyeddy/devices/oled"
	"github.com/rustyeddy/devices/relay"
	"github.com/rustyeddy/devices/vh400"
	"github.com/rustyeddy/otto/messanger"
	"github.com/rustyeddy/otto/server"
	"github.com/rustyeddy/otto/station"
)

type Gardener struct {
	messanger.Messanger
	*station.DeviceManager
	*station.StationManager
	*server.Server

	Done chan any
}

func (g *Gardener) GetMessanger() messanger.Messanger {
	if g.Messanger == nil {
		g.Messanger = messanger.NewMessanger(config.Msg.Broker)
	}
	return g.Messanger
}

func (g *Gardener) GetDeviceManager() *station.DeviceManager {
	if g.DeviceManager == nil {
		g.DeviceManager = station.NewDeviceManager()
	}
	return g.DeviceManager
}

var (
	pinmap = map[string]int{
		"on":   17,
		"off":  27,
		"soil": 22,
		"pump": 5,
		"env":  6,
	}
)

func (g *Gardener) Init() {
	g.Messanger = g.GetMessanger()
	g.DeviceManager = g.GetDeviceManager()
	g.StationManager = station.NewStationManager()
	g.Server = server.GetServer()
	g.Done = make(chan any)

	g.initButtons()
	g.InitSoil()
	g.initPump()
	g.initEnv()
	g.initDisplay()
}

func (g *Gardener) initButtons() {
	on, err := button.New("on", pinmap["on"])
	if err != nil {
		panic(err)
	}
	g.DeviceManager.Add(on)

	on.RegisterEventHandler(func(evt *devices.DeviceEvent) {
		switch evt.Type {
		case devices.DeviceEventRisingEdge:
			slog.Info("button pressed", "button", "on", "action", "pump_on")
			g.Messanger.Pub("pump", []byte("on"))
			g.Messanger.Pub("blue", []byte("on"))
			g.Messanger.Pub("display", []byte("Pump ON"))
		}
	})

	off, err := button.New("off", pinmap["off"])
	if err != nil {
		panic(err)
	}
	g.DeviceManager.Add(off)
	off.RegisterEventHandler(func(evt *devices.DeviceEvent) {
		switch evt.Type {
		case devices.DeviceEventRisingEdge:
			slog.Info("button pressed", "button", "off", "action", "pump_off")
			g.Messanger.Pub("pump", []byte("off"))
			g.Messanger.Pub("blue", []byte("off"))
			g.Messanger.Pub("display", []byte("Pump OFF"))
		}
	})
}

func (g *Gardener) InitSoil() {
	soil, err := vh400.New("soil", pinmap["soil"])
	if err != nil {
		panic(err)
	}
	g.DeviceManager.Add(soil)
	cb := func(t time.Time) {
		value, err := soil.Get()
		if err != nil {
			slog.Error("soil sensor read failed", "error", err)
			return
		}
		slog.Info("soil moisture reading", "value", value)
		g.Messanger.Pub("soil", []byte(fmt.Sprintf("%5.2f", value)))
	}
	soil.StartTicker(10*time.Second, &cb)
}

func (g *Gardener) initPump() {
	pump, err := relay.New("pump", pinmap["pump"])
	if err != nil {
		panic(err)
	}
	g.Messanger.Subscribe("pump", pump.HandleMsg)
}

func (g *Gardener) initEnv() {

	env, err := env.New("env", "/dev/i2c-1", 0x76)
	if err != nil {
		panic(err)
	}
	g.DeviceManager.Add(env)
	ticker := func(t time.Time) {
		resp, err := env.Get()
		if err != nil {
			slog.Error("env sensor read failed", "error", err)
			return
		}
		slog.Info("env sensor reading",
			"temperature", resp.Temperature,
			"humidity", resp.Humidity,
			"pressure", resp.Pressure)

		g.Messanger.Pub("env", []byte(fmt.Sprintf("Temp: %.2f, Hum: %.2f, Pres: %.2f", resp.Temperature, resp.Humidity, resp.Pressure)))
	}
	env.StartTicker(10*time.Second, &ticker)
}

func (g *Gardener) initDisplay() {
	display, err := oled.New("lcd", 0x27, 1)
	if err != nil {
		panic(err)
	}
	display.Clear()

	// Register devices
	g.DeviceManager.Add(display)
}

func (g *Gardener) Start() {
	// Implement start logic if needed
}

func (g *Gardener) Stop() {
	// Implement stop logic if needed
	g.Done <- true
}
