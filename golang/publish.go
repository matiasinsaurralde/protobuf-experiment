package main

import(
  "log"
  "time"

  "github.com/matiasinsaurralde/protobuf-experiment/golang/mqttchat"
)

const MqttNetwork string = "tcp"
const MqttAddress string = "iot.eclipse.org:1883"
const MqttClientID string = "testclient"

var ChatClient MqttChat.Client

/* The main function */

func main() {
  log.Println("main()")

  Topic := "topsecret"

  ChatClient = MqttChat.NewClient(map[string]string{
    "Network": MqttNetwork,
    "Address": MqttAddress,
    "ClientID": MqttClientID,
  }, nil)

  for {
    ChatClient.SendMessage( "hello", Topic)
    time.Sleep( 2 * time.Second )
  }
}
