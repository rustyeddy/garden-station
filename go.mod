module github.com/sensorstation/garden-station

go 1.23.3

require (
	github.com/sensorstation/otto v0.0.7
	github.com/warthog618/go-gpiocdev v0.9.1
)

replace github.com/sensorstation/otto v0.0.7 => ../otto

require (
	github.com/eclipse/paho.mqtt.golang v1.5.0 // indirect
	github.com/gorilla/websocket v1.5.3 // indirect
	github.com/maciej/bme280 v0.2.0 // indirect
	github.com/nfnt/resize v0.0.0-20180221191011-83c6a9932646 // indirect
	github.com/tarm/serial v0.0.0-20180830185346-98f6abe2eb07 // indirect
	golang.org/x/exp v0.0.0-20250103183323-7d7fa50e5329 // indirect
	golang.org/x/image v0.23.0 // indirect
	golang.org/x/net v0.27.0 // indirect
	golang.org/x/sync v0.10.0 // indirect
	golang.org/x/sys v0.24.0 // indirect
	periph.io/x/conn/v3 v3.7.1 // indirect
	periph.io/x/devices/v3 v3.7.2 // indirect
	periph.io/x/host/v3 v3.8.2 // indirect
)
