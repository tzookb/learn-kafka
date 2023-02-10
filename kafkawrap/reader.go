package kafkawrap

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type KafkaConsumer struct {
	kafka *kafka.Reader
}

func (k *KafkaConsumer) ReadMessage(ctx context.Context, object protoreflect.ProtoMessage) {
	msg, err := k.kafka.ReadMessage(ctx)
	if err != nil {
		fmt.Println("error here", err)
	} else {
		fmt.Println("consumed", msg.Offset, msg.Topic, string(msg.Value))
	}

	if err := proto.Unmarshal([]byte(msg.Value), object); err != nil {
		fmt.Println("Failed to parse address book:", err)
	}
}

func (k *KafkaConsumer) Close() {
	fmt.Println("closing kafka properly")
	k.kafka.Close()
}

func NewKafkaReader(kafkaURL, topic string) KafkaConsumer {
	kafka.TCP(kafkaURL)
	reader := kafka.NewReader(kafka.ReaderConfig{
		GroupID:   "9",
		Brokers:   []string{kafkaURL},
		Topic:     topic,
		Partition: 0,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
	})
	consumer := KafkaConsumer{
		kafka: reader,
	}
	return consumer
}
