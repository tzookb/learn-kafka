package main

// kafka-topics.sh --create --topic payments --bootstrap-server localhost:9092

import (
	"fmt"
	kafkawrap "kafkaplay/kafkawrap"
	"os"

	boom_tutorialpb "kafkaplay/out/boom.tutorialpb"

	kafka "github.com/segmentio/kafka-go"
)

func newKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

func anyToPointer[k interface{}](s k) *k {
	return &s
}

func main() {
	kafkaURL := os.Getenv("kafkaURL")
	topic := os.Getenv("topic")

	fmt.Println("producing topic: ", topic, " on", kafkaURL)

	writer := newKafkaWriter(kafkaURL, topic)
	defer writer.Close()
	fmt.Println("start producing ... !!")

	p := &boom_tutorialpb.Person{
		Email:         "tzookb@gmail.com",
		OptionalEmail: anyToPointer("zzz"),
		Name:          anyToPointer("tzook"),
		Id:            anyToPointer(int32(32)),
	}

	kafkawrap.Write(writer, p)
}
