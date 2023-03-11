package main

// kafka-topics.sh --create --topic payments --bootstrap-server localhost:9092

import (
	"fmt"
	kafkawrap "kafkaplay/kafkawrap"
	"os"
	"strconv"
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

	start := time.Now()

	p := &boom_tutorialpb.Person{
		Email:         "another@gmail.com",
		OptionalEmail: anyToPointer("zzz"),
		Name:          anyToPointer("tzook"),
		Id:            anyToPointer(int32(32)),
	}

	fmt.Println("before write")
	i := 0
	for i < 11 {
		go func(key int) {
			kafkaWriter.Write(p, strconv.Itoa(key))
			wg.Done()
		}(i)
		wg.Add(1)
		i++
	}
	wg.Wait()
	fmt.Println("after final write", time.Since(start))
}
