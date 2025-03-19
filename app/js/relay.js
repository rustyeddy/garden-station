const clientId = 'mqttjs_' + Math.random().toString(16).substr(2, 8);
// const host = "ws://" + window.location.hostname + ":8080";
// const broker = os.getenv("MQTT_BROKER");
const host = "ws://10.11.1.11:8080";
console.log(host);

const options = {
    keepalive: 60,
    clientId: clientId,
    protocolId: 'MQTT',
    protocolVersion: 4,
    clean: true,
    reconnectPeriod: 1000,
    connectTimeout: 30 * 1000,
    will: {
        topic: 'WillMsg',
        payload: 'Connection Closed abnormally..!',
        qos: 0,
        retain: false
    },
}

console.log('Connecting mqtt client');

const client = mqtt.connect(host, options);
client.on('error', (err) => {
    console.log('Connection error: ', err);
    client.end();
})

client.on('reconnect', () => {
    console.log('Reconnecting...');
})

client.on('connect', () => {
    console.log(`Client connected: ${clientId}`);
    // Subscribe
    console.log("subscribing to ss/c/station/env");
    // client.subscribe('ss/c/station/relay', { qos: 0 });
    client.subscribe('ss/d/station/env', (err) => {
        if (err) {
            console.log("subscribe error: ", err)
            return;
        }
    });
    client.subscribe('ss/d/station/soil', (err) => {
        if (err) {
            console.log("subscribe error: ", err)
            return;
        }
    });
    client.subscribe('ss/c/station/pump', (err) => {
        if (err) {
            console.log("subscribe error: ", err)
            return;
        }
    });
})

client.on('message', (topic, message) => {
    const parts = topic.split('/');
    const lastPart = parts.pop();
    console.log(topic, " => ", message.toString());

    switch (lastPart) {
    case "env":
        var msg = JSON.parse(message);
        document.getElementById("temperature").innerHTML = msg.Temperature.toFixed(2);
        document.getElementById("pressure").innerHTML = msg.Pressure.toFixed(2);
        document.getElementById("humidity").innerHTML = msg.Humidity.toFixed(2);
        break;

    case "soil":
        document.getElementById("soil").innerHTML = message.toString()
        break;

    case "pump":
        document.getElementById("pump").innerHTML = message.toString()
        break;
    }

});

function On() {
    console.log("on")
    client.publish('ss/c/station/relay', "on", { qos: 0, retain: false })    
}

function Off() {
    console.log("off")
    client.publish('ss/c/station/relay', "off", { qos: 0, retain: false })    
}
