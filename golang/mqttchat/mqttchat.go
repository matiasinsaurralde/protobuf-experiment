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
  Config map[string]string
  Log *log.Logger
}

func NewClient( Config map[string]string ) Client {
  logger := log.New( os.Stdout, "MqttChat.Client", log.Lshortfile)
  client := Client{
    Log: logger,
  }
  client.Config = Config
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
    Network: c.Config["Network"],
    Address: c.Config["Address"],
    ClientID: []byte( c.Config["ClientID"]),
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
