package main

import(
  "log"
  "time"

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
    "Topic": Topic,
  })

  for {
    // log.Println( "Publishing a test message" )
    // ChatClient.SendMessage("")
    time.Sleep( 3 * time.Second )
  }
}
