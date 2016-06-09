var mqtt    = require('mqtt');
var ProtoBuf = require('protobufjs');
var brotli = require('brotli');

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
      console.log(bb);
      var bbComp = new Buffer( brotli.compress(bb, false) );
      client.publish('topsecret', bbComp);
  }, 1000);
  
});
 
client.on('message', function (topic, message) {
  console.log("original length:", message.length)
  bbDecomp = brotli.decompress(message);
  console.log("decompressed length:", bbDecomp.length)
  var dec = Message.decode(bbDecomp);
  console.log(dec);
});
