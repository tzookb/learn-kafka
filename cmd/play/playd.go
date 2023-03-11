package main

import (
	"fmt"
	boom_tutorialpb "kafkaplay/out/boom.tutorialpb"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/structpb"
)

func main() {
	jsonString := `{"field1":"string","field2":55,"field3":false}`

	// create a new FreeStyle message
	freeStyle := &boom_tutorialpb.FreeStyle{}
	freeStyle.Data = &structpb.Struct{}
	fmt.Println("dasdasdasd", freeStyle.Data)

	// unmarshal the JSON string to a Protobuf object
	if err := protojson.Unmarshal([]byte(jsonString), freeStyle.Data); err != nil {
		fmt.Println("Error unmarshaling JSON: ", err)
		return
	}

	fmt.Println("freeStylefreeStyle", freeStyle)

	newJsonString, err := protojson.Marshal(freeStyle.Data)
	if err != nil {
		fmt.Println("Error marshaling to JSON: ", err)
		return
	}

	fmt.Println("JSON string: ", string(newJsonString))
}
