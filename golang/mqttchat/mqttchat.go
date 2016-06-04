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
  MessageHandler func(string, []byte)
  Log *log.Logger
}

func NewClient( Config map[string]string, MessageHandler func(string, []byte) ) Client {
  logger := log.New( os.Stdout, "MqttChat.Client", log.Lshortfile)
  client := Client{
    Log: logger,
  }

  client.Config = Config
  client.MessageHandler = MessageHandler
  client.Prepare()

  if client.Config["Topic"] == "" {
    client.Log.Println("Running in publish mode.")
  } else {
    client.Log.Println( "Running in subscribe mode.")
    go client.Subscribe()
  }

  return client
}

/* Prepare MQTT client */

func (c *Client) Prepare() {
  log.Println( "Preparing...")
  c.MqttClient = client.New(&client.Options{
    ErrorHandler: func(err error) {
      // c.Log.Fatalf( "Client error: %v", err)
      panic(err)
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
            c.MessageHandler(string(topicName), message)
            // c.Log.Printf("Message @ %s\n", topicName )
            // c.Log.Println( message )
          },
        },
      },
    })

    if err != nil {
      c.Log.Fatalf( "Can't subscribe: %v", err )
    }
}

/* Initialize the message protobuf */

func (c *Client) SendMessage( Text string, Topic string ) {

  m := &experiment.Message{
    Body: proto.String( Text),
  }

  data, err := proto.Marshal(m)

  if err != nil {
    c.Log.Fatal( "Can't marshal: %v", err )
  }

  err = c.MqttClient.Publish(&client.PublishOptions{
    QoS: mqtt.QoS0,
    TopicName: []byte( Topic ),
    Message: data,
  })

  if err != nil {
    c.Log.Fatal("Can't send message: %v", err)
  }

}
