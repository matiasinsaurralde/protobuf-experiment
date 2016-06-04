package main

import(
  "log"

  "github.com/matiasinsaurralde/protobuf-experiment/golang/mqttchat"
)

const MqttNetwork string = "tcp"
const MqttAddress string = "test.mosquitto.org:1883"
const MqttClientID string = "testclient"

var ChatClient MqttChat.Client

/* The main function */

func main() {
  log.Println("main()")

  Topic := "TopSecret123"

  ChatClient = MqttChat.NewClient(map[string]string{
    "Network": MqttNetwork,
    "Address": MqttAddress,
    "ClientID": MqttClientID,
  })

  ChatClient.SendMessage( Topic, "hello")

  for {

  }
}
