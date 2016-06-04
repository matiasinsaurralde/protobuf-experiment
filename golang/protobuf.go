package main

import(
  "log"
  "github.com/golang/protobuf/proto"
  "github.com/matiasinsaurralde/protobuf-experiment/proto"
)

func main() {
  log.Println("main()")
  m := &experiment.Message{
    Body: proto.String("Hello world"),
  }

  log.Println("m", m)
}
