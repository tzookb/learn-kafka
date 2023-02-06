package main

import (
	"context"
	"fmt"

	"os"
	"time"

	kafka "github.com/segmentio/kafka-go"
)

func newKafkaReader(kafkaURL, topic string) *kafka.Reader {
	kafka.TCP(kafkaURL)
	return kafka.NewReader(kafka.ReaderConfig{
		GroupID:   "1",
		Brokers:   []string{kafkaURL},
		Topic:     topic,
		Partition: 0,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
	})
}

func main() {
	kafkaURL := os.Getenv("kafkaURL")
	topic := os.Getenv("topic")

	fmt.Println("consuming topic: ", topic, " on", kafkaURL)

	reader := newKafkaReader(kafkaURL, topic)

	// reader.SetOffset(0)
	defer reader.Close()
	fmt.Println("start consuming ... !!")
	for i := 0; i < 1; i++ {
		// msg, err := reader.FetchMessage(context.Background())
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Println("error here", err)
		} else {
			fmt.Println("consumed", msg.Offset, msg.Topic, string(msg.Value))
		}
		time.Sleep(1 * time.Second)
	}
}
