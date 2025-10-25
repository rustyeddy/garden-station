package main

import (
	"fmt"

	"github.com/rustyeddy/devices/bme280"
)

func InitBME280(done chan any) (*bme280.BME280, error) {
	bme := bme280.New("env", "/dev/i2c-1", 0x76)
	err := bme.Open()
	if err != nil {
		fmt.Printf("Failed to open BME280 %+v\n", err)
		return nil, err
	}

	return bme, nil
}
