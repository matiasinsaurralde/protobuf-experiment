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
  go client.Subscribe()
  return client
}

/* Prepare MQTT client */

func (c *Client) Prepare() {
  log.Println( "Preparing...")
  c.MqttClient = client.New(&client.Options{
    ErrorHandler: func(err error) {
      c.Log.Fatalf( "Client error: %v", err)
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

}

func (c *Client) Subscribe() {
  log.Println( "Subscribing...")
  err := c.MqttClient.Subscribe(&client.SubscribeOptions{
      SubReqs: []*client.SubReq{
        &client.SubReq{
          TopicFilter: []byte( c.Config["Topic"] ),
          QoS: mqtt.QoS0,
          Handler: func( topicName, message []byte) {
            c.Log.Printf("Message @%s: %s\n", topicName, string(message))
          },
        },
      },
    })

    if err != nil {
      c.Log.Fatalf( "Can't subscribe: %v", err )
    }
}

/* Initialize the message protobuf */

func (c *Client) SendMessage( text string ) {
  m := &experiment.Message{
    Body: proto.String( text),
  }

  err := c.MqttClient.Publish(&client.PublishOptions{
    QoS: mqtt.QoS0,
    TopicName: []byte(c.Config["Topic"]),
    Message: []byte("xd"),
  })

  if err != nil {
    c.Log.Fatal("Can't send message: %v", err)
  }

  c.Log.Println("Message", m)
}
