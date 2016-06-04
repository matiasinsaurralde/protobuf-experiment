package main

import (
	// "github.com/golang/protobuf/proto"
  "gopkg.in/kothar/brotli-go.v0/enc"
  // "gopkg.in/kothar/brotli-go.v0/dec"
  "encoding/base64"
  "encoding/json"
	"log"
)

type Person struct {
  Name string
  Bio string
  Quote string
  Age int
}

func main() {

  p := Person{
    Name: "William Burroughs",
    Bio: "Beat it?",
    Quote: "Communication must become total and conscious before we can stop it",
    Age: 83,
  }

  log.Println( "initializing a Person{}", p )

  jsonP, err := json.Marshal( &p )

  if err != nil {
    panic(err)
  }

  log.Println( "marshalling a Person{}", string(jsonP) )

  log.Println( "marshaled (JSON) []byte length:", len(jsonP), ":/" )

  compressedJsonP, _ := enc.CompressBuffer( nil, jsonP, make([]byte, 0))

  log.Println( "compressedJsonP []byte length:", len(compressedJsonP))

  lenDiff := len(jsonP) - len(compressedJsonP)
  compressionPercentage := float64(lenDiff) / float64(len(jsonP)) * 100.0

  log.Printf( "compression: %.2f %%", compressionPercentage)

  b64P := base64.StdEncoding.EncodeToString( jsonP )

  log.Println( "jsonP -> base64:", b64P, "length:", len(b64P))

  b64CP := base64.StdEncoding.EncodeToString( compressedJsonP )

  log.Println( "compressedJsonP -> base64:", b64CP, "length:", len(b64CP))

}
