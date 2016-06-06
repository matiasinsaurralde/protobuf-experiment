package main

import (
	"encoding/base64"
	"encoding/json"

	"log"
	// "os"

	"gopkg.in/kothar/brotli-go.v0/enc"
	"github.com/golang/protobuf/proto"

	// "github.com/andrewstuart/yenc"

	"github.com/matiasinsaurralde/protobuf-experiment/proto/person"
)

type Person struct {
	Name  string
	Bio   string
	Quote string
	Age   int
}

func main() {

	p := Person{
		Name:  "William Burroughs",
		Bio:   "Beat it?",
		Quote: "Communication must become total and conscious before we can stop it",
		Age:   83,
	}

	log.Println("initializing a Person{}", p)

	jsonP, err := json.Marshal(&p)

	if err != nil {
		panic(err)
	}

	log.Println("marshalling a Person{}", string(jsonP))

	log.Println("marshaled (JSON) []byte length:", len(jsonP), ":/")

	compressedJsonP, _ := enc.CompressBuffer(nil, jsonP, make([]byte, 0))

	log.Println("compressedJsonP []byte length:", len(compressedJsonP))

	lenDiff := len(jsonP) - len(compressedJsonP)
	compressionPercentage := float64(lenDiff) / float64(len(jsonP)) * 100.0

	log.Printf("compression: %.2f %%", compressionPercentage)

	b64P := base64.StdEncoding.EncodeToString(jsonP)

	log.Println("jsonP -> base64:", b64P, "length:", len(b64P))

	b64CP := base64.StdEncoding.EncodeToString(compressedJsonP)

	log.Println("compressedJsonP -> base64:", b64CP, "length:", len(b64CP))

	protoP := &person.Person{
		Name: proto.String(p.Name),
		Bio: proto.String(p.Bio),
		Quote: proto.String(p.Quote),
		Age: proto.Int(p.Age),
	}

	protoBuf, err := proto.Marshal( protoP )

	log.Println( "Protobuf length:", len(protoBuf))

	if err != nil {
		panic(err)
	}

	compressedProtoBuf, _ := enc.CompressBuffer(nil, protoBuf, make([]byte, 0))

	log.Println("compressed protobuf length:", len(compressedProtoBuf))

	lenDiff = len(protoBuf) - len(compressedProtoBuf)
	compressionPercentage = float64(lenDiff) / float64(len(protoBuf)) * 100.0

	log.Printf("compression: %.2f %%", compressionPercentage)

	b64Protobuf := base64.StdEncoding.EncodeToString(compressedProtoBuf)

	log.Println("protobuf -> b64", b64Protobuf, "length:", len(b64Protobuf))

	/*
	yencWriter := yenc.NewWriter(os.Stdout)
	yencWriter.Write(compressedProtoBuf)
	yencWriter.Close()
	*/
}
