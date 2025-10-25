# Garden Station 🌱

Garden Station is an automated watering system IoT project built with the Otto framework. It demonstrates clean device/messaging architecture with real-world sensors and actuators.

## Features

### 📊 **Sensors**
- **BME280**: Temperature, humidity, and pressure monitoring
- **VH400**: Soil moisture measurement with configurable wet/dry thresholds
- **GPIO Buttons**: Manual pump control override

### 🔧 **Actuators**  
- **Water Pump**: Automated watering based on soil moisture levels
- **LED Indicators**: Visual status indicators (white, blue, green, yellow, red)
- **OLED Display**: Real-time sensor data and system status

### 🌐 **Connectivity**
- **Web Interface**: Full-featured UI at http://localhost:8011
- **MQTT Integration**: Sensor data publishing and remote control
- **RESTful API**: JSON endpoints for integration

### 🧪 **Development Features**
- **Mock Mode**: Complete hardware simulation for testing
- **Local Messaging**: No external dependencies for development
- **Flexible MQTT**: Public broker support or custom broker configuration

## Quick Start

### Development Mode (No Hardware Required)
```bash
# Local messaging, fully mocked
./garden-station -mock -local

# With public MQTT broker
./garden-station -mock

# Web interface: http://localhost:8011
```

### Hardware Deployment
```bash
# Connect real sensors and run
./garden-station

# With custom MQTT broker  
./garden-station -mqtt-broker your.broker.com
```

## Command Line Options
- `-mock`: Enable hardware mocking for development/testing
- `-local`: Use local messaging (no MQTT broker required)
- `-mqtt-broker string`: Custom MQTT broker (default: test.mosquitto.org)

## How It Works

### Automated Watering Logic
1. **Monitor**: VH400 sensor continuously measures soil moisture
2. **Evaluate**: Compare readings against wet/dry thresholds (configurable)
3. **Act**: Start pump when soil is dry, stop when adequately watered
4. **Display**: Show status on OLED and web interface
5. **Report**: Publish sensor data via MQTT for monitoring

### Web Interface Features
- Real-time soil moisture display with pump status
- Environmental data (temperature, humidity, pressure)
- Manual pump control buttons
- Station information and network details

## Architecture

Built on the Otto framework demonstrating:
- **Clean Separation**: Hardware devices vs messaging infrastructure  
- **ManagedDevice Pattern**: Wraps simple devices with MQTT capabilities
- **Type-Safe Management**: Devices registered and retrieved by name/type
- **Graceful Fallbacks**: Local messaging when MQTT unavailable
- **Easy Testing**: Complete mock mode for development


