package main

// kafka-topics.sh --create --topic payments --bootstrap-server localhost:9092

import (
	"fmt"
	kafkawrap "kafkaplay/kafkawrap"
	"os"
	"sync"
	"time"

	boom_tutorialpb "kafkaplay/out/boom.tutorialpb"
)

func anyToPointer[k interface{}](s k) *k {
	return &s
}

func main() {
	kafkaURL := os.Getenv("kafkaURL")
	topic := os.Getenv("topic")

	fmt.Println("producing topic: ", topic, " on", kafkaURL)
	kafkaWriter := kafkawrap.NewKafkaWriter(kafkaURL, topic)
	defer kafkaWriter.Close()

	var wg sync.WaitGroup
	wg.Add(2)

	start := time.Now()
	p := &boom_tutorialpb.Person{
		Email:         "newemail@gmail.com",
		OptionalEmail: anyToPointer("zzz"),
		Name:          anyToPointer("tzook"),
		Id:            anyToPointer(int32(32)),
	}

	fmt.Println("before write")
	go func() {
		kafkaWriter.Write(p)
		fmt.Println("after write", time.Since(start))
		wg.Done()
	}()

	time.Sleep(time.Second * 2)
	go func() {
		kafkaWriter.Write(p)
		fmt.Println("after write", time.Since(start))
		wg.Done()
	}()

	wg.Wait()
}
