var mqtt    = require('mqtt');
var ProtoBuf = require('protobufjs');

var builder = ProtoBuf.protoFromFile('../proto/message.proto');
var Message = builder.build('Message');
var msg = new Message();

var client  = mqtt.connect('mqtt://iot.eclipse.org');

client.on('connect', function () {
  client.subscribe('topsecret');
  
  setInterval(function(){
      var msg = new Message();
      msg.body = 'hola';
      msg.profile = new Message.Profile('Carlos', 'eroto', 28);
      var bb = msg.encodeNB();
      client.publish('topsecret', bb);
  }, 1000);
  
});
 
client.on('message', function (topic, message) {
  // message is Buffer 
  var dec = Message.decode(message);
  console.log(dec);
  //client.end();
});