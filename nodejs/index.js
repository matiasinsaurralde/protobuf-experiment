'use strict';

var ProtoBuf = require('protobufjs');

var builder = ProtoBuf.newBuilder({ convertFieldsToCamelCase: true });
ProtoBuf.loadProtoFile('../proto/message.proto', builder);
var root = builder.build();
var msg = new root.Message('hi');

console.log( msg );