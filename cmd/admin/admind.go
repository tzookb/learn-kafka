package main

// kafka-topics.sh --create --topic payments --bootstrap-server localhost:9092

import (
	"context"
	"fmt"
	"io/ioutil"
	"kafkaplay/api"
	"kafkaplay/api/zoom"
	kafkacli "kafkaplay/cli/kafka"
	boom_tutorialpb "kafkaplay/out/boom.tutorialpb"
	"log"

	"google.golang.org/protobuf/proto"

	// "kafkaplay/out/boom_tutorialpb"
	"os"
	"time"

	kafka "github.com/segmentio/kafka-go"
	"github.com/urfave/cli/v2" // imports as package "cli"
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

func WriteProtobufToBinaryFile(message proto.Message, filename string) error {
	data, err := proto.Marshal(message)
	if err != nil {
		return fmt.Errorf("cannot marshal proto message to binary: %w", err)
	}

	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("cannot write binary data to file: %w", err)
	}

	return nil
}

func anyToPointer[k interface{}](s k) *k {
	return &s
}

func main() {

	kafkaCommand := kafkacli.GetCommand()
	app := &cli.App{
		Commands: []*cli.Command{
			&kafkaCommand,
			// {
			// 	Name:    "add",
			// 	Aliases: []string{"a"},
			// 	Usage:   "add a task to the list",
			// 	Action: func(cCtx *cli.Context) error {
			// 		fmt.Println("added task: ", cCtx.Args().First())
			// 		return nil
			// 	},
			// },
			// {
			// 	Name:    "complete",
			// 	Aliases: []string{"c"},
			// 	Usage:   "complete a task on the list",
			// 	Action: func(cCtx *cli.Context) error {
			// 		fmt.Println("completed task: ", cCtx.Args().First())
			// 		return nil
			// 	},
			// },
			// {
			// 	Name:    "template",
			// 	Aliases: []string{"t"},
			// 	Usage:   "options for task templates",
			// 	Subcommands: []*cli.Command{
			// 		{
			// 			Name:  "add",
			// 			Usage: "add a new template",
			// 			Action: func(cCtx *cli.Context) error {
			// 				fmt.Println("new task template: ", cCtx.Args().First())
			// 				return nil
			// 			},
			// 		},
			// 		{
			// 			Name:  "remove",
			// 			Usage: "remove an existing template",
			// 			Action: func(cCtx *cli.Context) error {
			// 				fmt.Println("removed task template: ", cCtx.Args().First())
			// 				return nil
			// 			},
			// 		},
			// 	},
			// },
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

	fileName := "thefile"

	if false {
		p := &boom_tutorialpb.Person{
			Email:         "tzookb@gmail.com",
			OptionalEmail: anyToPointer("zzz"),
			Name:          anyToPointer("tzook"),
			Id:            anyToPointer(int32(32)),
		}

		api.Print()
		zoom.Print()
		fmt.Println(p)
		fmt.Println(p.String())

		out, err := proto.Marshal(p)
		if err != nil {
			fmt.Println("Failed to encode address book:", err)
		}
		if err := ioutil.WriteFile(fileName, out, 0644); err != nil {
			fmt.Println("Failed to write address book:", err)
		}
	}
	if false {
		// Read the existing address book.
		in, err := ioutil.ReadFile(fileName)
		if err != nil {
			fmt.Println("Error reading file:", err)
		}
		book := &boom_tutorialpb.Person{}
		if err := proto.Unmarshal(in, book); err != nil {
			fmt.Println("Failed to parse address book:", err)
		}
		fmt.Println(book)
	}

	if false {
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
}
