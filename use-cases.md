# Garden Station Use Cases

## As a gardner

### Automated Watering

> I want to automatically start watering my plants when the soil gets
too dry and turn the water off when the soil gets wet enough.

#### Requirements

1. A soil moisture sensor to collect the soil moisture level
3. Timer to periodically read the soil moisture level
2. Control a pump by turning it on and off
4. Controller to turn the pump on and off according to the moisture
   level 

### Configurable Moisture Levels

> I want to be able to set the level that I consider the soil too dry or
too wet.

#### Requirements

2. Configurable value to start watering
2. Configurable value to stop watering

### Configurable Frequency 

I want to be able to change the frequency the soil and other
environmental data gets checked

#### Requirements

1. Configurable collection timer
2. Configurable variable for the collection timer

### Saved Data

I want to be able to save the soil moisture levels every time the
soil is checked locally on the device and in the cloud

#### Requirements

3. Timeseries database limited storage on device locally
3. Timeseries database to store all historic data in the cloud 

### View Data

I want to be able to view graphs of moisture levels and see when the
pump was turned on and off

#### Requirements

1. User Interface to view moisture level data 
1. User Interface to view pump activation data
2. Web UI

#### Additional Environmental Data

I want to collect additional environmental data such as temperature,
humidity as well as others.

#### Requirements

1. Make data collection mechanism generic and configurable
2. Expand data collected according to the specific sensors
3. REST or GraphQL API

### Pump Control Manually

I want to manually control the pump, turn it on or off with a physical
button. 

#### Requirements

1. Add a physical button that can control the pumps
2. Button GPIO API 

### Pump Control from App

I want to manually control the pump from an application

#### Requirements

1. Pump control API

### Station Display

I want to be able to read soil moisture level and other environmental
data on the garden station directly

#### Requirements

1. A physical display on the garden station
2. A data display API

### Application Display

I want to be able to read soil moisture level and other environmental
data on the garden from an application

#### Requirements

1. Add a display via the application
2. Data display API

### Lighting

I want to be able to add lights to my garden-station the lights should
be able to be physically turned off and on

#### Requirements

1. Add a physical light (LED) to the garden station
2. Add a lighting GPIO API

### Automated Lighting

1. I want to be able to control the lighting from an application
2. I want the lighting to be programmable based on light levels

#### Requirements

1. Lumenecense sensor
2. Light application API

### Wireless Power

I want the garden station to work with out direct electricity or
network connection

1. Connect physical battery
2. Connect solar panel
3. Wireless data communication

### Fleet Management

I want to be able to control a whole bunch of watering stations

#### Requirements

1. Local fleet controller
2. MQTT communication protocol
3. Cloud based controller
4. Fleet management API

## As an administrator

I want to be able to update garden station software without
interupting the station functionality.

#### Requirements

1. Software update mechanism
2. Software update API


