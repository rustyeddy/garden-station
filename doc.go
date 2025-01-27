/*
Documentation for the Garden Station.
The Garden Station consists of the following elements: one or
more sensor stations, a control hub, a user interface and a
cloud based hub.

The Gardener can scale to a single indoor plant to an entire
outdoor garden with numerous stations and a cloud backend.

It is possible and common to have the hub and station on the
same physical hardware. Or multple stations controled by a
single controller communicating with an external hub.

A station consists of the following optional devices:

1. Soil Moisture Sensor
2. Temperature, Humidity and Pressure sensor
3. Water relay
4. Light relay
5. OLED
6. On / Off buttons

## Operational Overview

 1. Periodicaly collect environmental data: temp, humidity and
    pressure samples
    a. Publish samples

 2. Peridoicly collect soil moisture
    a. publish soil moisture

 3. Subscribe to pump controls
    a. Respond to pump on and pump off commands

 4. Manual on and off buttons publish on/off messages

 4. Subscribe to light relay commands
    a. Respond to on off light commands

 5. Controller subscribes to env data
    a. when soil gets to dry send control message
    to turn on the pump
    b. when soil gets to we send control message to
    turn pump off

 6. Controller subscribes to light relay to turn lights
    on and off accordingly

 7. Controller saves all data subscribed and publish

 8. Controller sends all data to cloud storage

 9. UI can monitor all stations data

 10. UI can turn on or off lights and pumps

 11. UI can set the wet and dry levels for soil sensor
*/
package main
