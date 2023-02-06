package kafkawrap

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// kafka-topics.sh --create --topic payments --bootstrap-server localhost:9092

// "kafkaplay/out/boom_tutorialpb"

// imports as package "cli"

func Read(reader *kafka.Reader, object protoreflect.ProtoMessage) {

	msg, err := reader.ReadMessage(context.Background())
	if err != nil {
		fmt.Println("error here", err)
	} else {
		fmt.Println("consumed", msg.Offset, msg.Topic, string(msg.Value))
	}

	if err := proto.Unmarshal([]byte(msg.Value), object); err != nil {
		fmt.Println("Failed to parse address book:", err)
	}
}
