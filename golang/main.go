package main

import(
  "log"

  "github.com/matiasinsaurralde/protobuf-experiment/golang/mqttchat"
)

const MqttUrl string = "mqtt://test.mosquitto.org"

var ChatClient MqttChat.Client

/* The main function */

func main() {
  log.Println("main()")

  ChatClient = MqttChat.NewClient(map[string]string{
    "Url": MqttUrl,
  })
}
