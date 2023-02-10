package main

import (
	"fmt"
	boom_tutorialpb "kafkaplay/out/boom.tutorialpb"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

func main() {
	src := proto.Message{}
	any, _ := anypb.New(src)
	arr := []*anypb.Any{any}
	h := &boom_tutorialpb.House{
		Details: arr,
	}
	fmt.Println(h)

}
