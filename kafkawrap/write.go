package kafkawrap

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
)

// kafka-topics.sh --create --topic payments --bootstrap-server localhost:9092

// "kafkaplay/out/boom_tutorialpb"

// imports as package "cli"

func Write(writer *kafka.Writer, protoMsg proto.Message) {
	out, err := proto.Marshal(protoMsg)
	if err != nil {
		fmt.Println("Failed to encode address book:", err)
	}

	randKey := []byte(fmt.Sprint(uuid.New()))
	msg := kafka.Message{
		// Key:   randKey,
		Value: out,
	}
	writeErr := writer.WriteMessages(context.Background(), msg)
	if writeErr != nil {
		fmt.Println("error here", writeErr)
	} else {
		fmt.Println("produced", string(randKey), string(out))
	}
}
