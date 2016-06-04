var mqtt    = require('mqtt');
var client  = mqtt.connect('mqtt://iot.eclipse.org');
 
client.on('connect', function () {
  client.subscribe('topsecret');
  client.publish('topsecret', 'hi from asuncion');
});
 
client.on('message', function (topic, message) {
  // message is Buffer 
  console.log(topic);
  console.log(message.toString());
  //client.end();
});