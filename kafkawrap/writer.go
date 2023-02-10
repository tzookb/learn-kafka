package kafkawrap

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
)

type KafkaWriter struct {
	kafka *kafka.Writer
}

func (k *KafkaWriter) Close() {
	k.kafka.Close()
}

func NewKafkaWriter(kafkaURL, topic string) KafkaWriter {
	writer := kafka.Writer{
		BatchSize:              2,
		BatchTimeout:           time.Second * 10,
		AllowAutoTopicCreation: true,
		// WriteBackoffMin: 0,
		// BatchTimeout: 0,
		// WriteTimeout: 0,
		// ReadTimeout: time.Second,
		// WriteBackoffMax: time.Second * 2,
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
	consumer := KafkaWriter{
		kafka: &writer,
	}
	return consumer
}

func (k *KafkaWriter) Write(protoMsg proto.Message) {
	out, err := proto.Marshal(protoMsg)
	if err != nil {
		fmt.Println("Failed to encode address book:", err)
	}

	// randKey := []byte(fmt.Sprint(uuid.New()))
	msg := kafka.Message{
		// Key:   randKey,
		Value: out,
	}
	writeErr := k.kafka.WriteMessages(context.Background(), msg)
	if writeErr != nil {
		fmt.Println("error here", writeErr)
	} else {
		fmt.Println("produced event")
	}
}
