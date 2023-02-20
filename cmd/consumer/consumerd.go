package main

import (
	"context"
	"fmt"
	"kafkaplay/kafkawrap"
	boom_tutorialpb "kafkaplay/out/boom.tutorialpb"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func main() {
	kafkaURL := os.Getenv("kafkaURL")
	topic := os.Getenv("topic")
	readAmountStr := os.Getenv("readAmount")

	readAmount, err := strconv.Atoi(readAmountStr)
	if err != nil {
		readAmount = 1
	}

	ctx := context.Background()

	kafkaReader := kafkawrap.NewKafkaReader(kafkaURL, topic)

	defer kafkaReader.Close()

	SetupCloseHandler(func() {
		kafkaReader.Close()
	})
	p := &boom_tutorialpb.Person{}
	p.Email = "xx"
	fmt.Println("before reading: ", readAmount)
	for readAmount > 0 {
		kafkaReader.ReadMessage(ctx, p)
		fmt.Println("read: ", p)
		readAmount -= 1
	}
	fmt.Println("after reading")
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
