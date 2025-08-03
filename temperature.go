package main

import (
	"fmt"
	"time"

	"github.com/rustyeddy/otto/device/bme280"
	"github.com/rustyeddy/otto/messanger"
)

func InitBME280(done chan any) (*bme280.BME280, error) {
	bme := bme280.New("env", "/dev/i2c-1", 0x76)
	err := bme.Init()
	if err != nil {
		fmt.Printf("Failed to init BME280 %+v\n", err)
		return nil, err
	}

	bme.Topic = messanger.GetTopics().Data("env")
	go bme.TimerLoop(5*time.Second, done, bme.ReadPub)
	return bme, nil
}
