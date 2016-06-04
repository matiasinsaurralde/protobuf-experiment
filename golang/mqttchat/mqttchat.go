package MqttChat

import(
  "log"
  "os"

  "github.com/golang/protobuf/proto"
  "github.com/yosssi/gmq/mqtt"
  "github.com/yosssi/gmq/mqtt/client"

  "github.com/matiasinsaurralde/protobuf-experiment/proto"
)

type Client struct {
  MqttClient *client.Client
  Log *log.Logger
}

func NewClient( Params map[string]string ) Client {
  logger := log.New( os.Stdout, "MqttChat.Client", log.Lshortfile)
  client := Client{
    Log: logger,
  }
  client.Prepare()
  return client
}

/* Prepare MQTT client */

func (c *Client) Prepare() {
  c.MqttClient = client.New(&client.Options{
    ErrorHandler: func(err error) {
      c.Log.Println( "ERROR")
    },
  })

  defer c.MqttClient.Terminate()

  err := c.MqttClient.Connect(&client.ConnectOptions{
    Network: "tcp",
    Address: "127.0.0.1",
    ClientID: []byte("mqttchat"),
  })

  if err != nil {
    c.Log.Fatalf( "Can't connect: %v", err )
  }

  if mqtt.QoS0 == 1 {

  }
}

/* Initialize the message protobuf */

func (c *Client) SendMessage( text string ) {
  m := &experiment.Message{
    Body: proto.String( text),
  }
  c.Log.Println("Message", m)
}
