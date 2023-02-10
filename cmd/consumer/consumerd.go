package main

import (
	"context"
	"fmt"
	"kafkaplay/kafkawrap"
	boom_tutorialpb "kafkaplay/out/boom.tutorialpb"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	kafkaURL := os.Getenv("kafkaURL")
	topic := os.Getenv("topic")
	ctx := context.Background()

	kafkaReader := kafkawrap.NewKafkaReader(kafkaURL, topic)

	defer kafkaReader.Close()

	SetupCloseHandler(func() {
		kafkaReader.Close()
	})
	p := &boom_tutorialpb.Person{}
	p.Email = "xx"
	fmt.Println("before", p)
	fmt.Printf("t1: %T\n", p)
	kafkaReader.ReadMessage(ctx, p)
	fmt.Println("after", p)
	fmt.Println("after specific", p.Email)

}

func SetupCloseHandler(f func()) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\r- Ctrl+C pressed in Terminal")
		f()
		os.Exit(0)
	}()
}
