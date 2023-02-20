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
	pos   int
}

func (k *KafkaWriter) Close() {
	k.kafka.Close()
}

func NewKafkaWriter(kafkaURL, topic string) KafkaWriter {
	balancer := kafka.CRC32Balancer{}
	writer := kafka.Writer{
		BatchSize:              1,
		BatchTimeout:           time.Second * 10,
		AllowAutoTopicCreation: true,
		// WriteBackoffMin: 0,
		// BatchTimeout: 0,
		// WriteTimeout: 0,
		// ReadTimeout: time.Second,
		// WriteBackoffMax: time.Second * 2,
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topic,
		Balancer: balancer,
	}
	consumer := KafkaWriter{
		kafka: &writer,
	}
	return consumer
}

func (k *KafkaWriter) Write(protoMsg proto.Message, theKey string) {
	out, err := proto.Marshal(protoMsg)
	if err != nil {
		fmt.Println("Failed to encode address book:", err)
	}
	fmt.Println("the key", theKey)

	if theKey == "" {
		theKey = "basekey"
	}
	keyBytes := []byte(theKey)

	msg := kafka.Message{
		Key:   keyBytes,
		Value: out,
	}
	writeErr := k.kafka.WriteMessages(context.Background(), msg)
	if writeErr != nil {
		fmt.Println("error here", writeErr)
	} else {
		// fmt.Println("produced event")
	}
}
