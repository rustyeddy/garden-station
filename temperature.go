package main

import (
	"fmt"
	"time"

	"github.com/sensorstation/otto/device/bme280"
	"github.com/sensorstation/otto/messanger"
)

func (g *Gardener) InitBME280() {
	bme := bme280.New("env", "/dev/i2c-1", 0x76)
	err := bme.Init()
	if err != nil {
		fmt.Printf("Failed to init BME280 %+v\n", err)
		return
	}

	bme.Topic = messanger.TopicData("env")
	go bme.TimerLoop(5*time.Second, g.Done(), bme.ReadPub)
	g.AddDevice(bme)
}
