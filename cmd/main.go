package main

import (
	"gkafka/model"
	"log"
	"github.com/golang/protobuf/proto"
)

func main() {
	message := &model.Message{
		Id: proto.Int32(1),
		Name: proto.String("Foo Bar"),
		Email: proto.String("foo@bar.bac"),
	}
	data, err := proto.Marshal(message)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	log.Println(data)
	log.Println(string(data))
}
