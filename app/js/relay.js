const clientId = 'mqttjs_' + Math.random().toString(16).substr(2, 8);
const host = "ws://" + window.location.hostname + ":8080";
console.log(window.location.hostname);

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
    client.subscribe('ss/c/station/relay', { qos: 0 });
    client.subscribe('ss/d/station/env', { qos: 0 });
    client.subscribe('ss/d/station/soil', { qos: 0 });
    client.subscribe('ss/c/station/pump', { qos: 0 });
})

client.on("message", (topic, message) => {
    // message is Buffer
    switch (topic) {
    case "ss/d/station/env":
        var j = JSON.parse(message);
        var temp = j.Temperature.toFixed(2);
        var hum  = j.Humidity.toFixed(2);
        var press= j.Pressure.toFixed(2);

        document.getElementById("temperature").innerHTML = temp;
        document.getElementById("humidity").innerHTML = hum;
        document.getElementById("pressure").innerHTML = press;
        break;
        
    case "ss/d/station/soil":
        var soil = message.toString();
        document.getElementById("soil").innerHTML = soil;
        break;

    case "ss/c/station/pump":
        var pump = message.toString();
        document.getElementById("pump").innerHTML = pump;
        break;
    }

    console.log(topic + " => " + message.toString());
});

// Unsubscribe
/* client.unsubscribe('tt', () => {
 *   console.log('Unsubscribed');
 * })
 */

function On() {
    console.log("on")
    client.publish('ss/c/station/relay', "on", { qos: 0, retain: false })    
}

function Off() {
    console.log("off")
    client.publish('ss/c/station/relay', "off", { qos: 0, retain: false })    
}
