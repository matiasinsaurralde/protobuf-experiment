package main

import(
  "log"

  "github.com/golang/protobuf/proto"
  "github.com/yosssi/gmq/mqtt"
  "github.com/yosssi/gmq/mqtt/client"

  "github.com/matiasinsaurralde/protobuf-experiment/proto"
)

const MqttUrl string = "mqtt://test.mosquitto.org"

/* Initialize the message protobuf */

func SendMessage( text string ) {
  m := &experiment.Message{
    Body: proto.String( text),
  }
}

/* The main function */

func main() {
  log.Println("main()")
}
