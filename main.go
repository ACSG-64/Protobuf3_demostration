package main

import (
	"fmt"
	simple_msg "github.com/ACSG-64/Protobuf3_demostration/src/simple.msg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"log"
)

func main() {
	msg := doSimpleMessage()

	// Writing to disk
	_ = writeToFile("message.bin", msg)
	// Reading from disk
	newMsg1 := simple_msg.SimpleMessage{}
	_ = readFromFile("message.bin", &newMsg1)
	fmt.Println("Message read from disk:", newMsg1)

	// To JSON
	json, _ := toJSON(msg)
	fmt.Println("JSON format:", json)
	// Parse from JSON
	newMsg2 := simple_msg.SimpleMessage{}
	_ = fromJSON(json, &newMsg2)
	fmt.Println("Converted from JSON:", newMsg2)

}

func writeToFile(fileName string, pb proto.Message) error {
	bytes, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("ERROR: Cannot serialize to bytes", err)
		return err
	}

	err = ioutil.WriteFile(fileName, bytes, 0644)
	if err != nil {
		log.Fatalln("ERROR: Cannot save the file", err)
		return err
	}

	fmt.Println("Data has been written successfully!")
	return nil
}

func readFromFile(fileName string, pb proto.Message) error {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalln("ERROR: Cannot read the binary file", err)
		return err
	}

	err = proto.Unmarshal(bytes, pb)
	if err != nil {
		log.Fatalln("ERROR: Cannot serialize the binary in the an protocol buffer struct", err)
		return err
	}

	return nil
}

func toJSON(pb proto.Message) (string, error) {
	marshaler := protojson.MarshalOptions{}
	json, err := marshaler.Marshal(pb)
	if err != nil {
		log.Fatalln("ERROR: Cannot serialize to bytes", err)
		return "", err
	}

	return string(json), nil
}

func fromJSON(jsonMessage string, pb proto.Message) error {
	unmarshaler := protojson.UnmarshalOptions{}
	err := unmarshaler.Unmarshal([]byte(jsonMessage), pb)
	if err != nil {
		log.Fatalln("ERROR: Cannot read the JSON file", err)
		return err
	}

	return nil
}

func doSimpleMessage() *simple_msg.SimpleMessage {
	msg := simple_msg.SimpleMessage{
		Id:         1234,
		IsSimple:   true,
		Name:       "Basic message",
		SampleList: []int32{2, 4, 6, 8, 10},
	}

	fmt.Println("Message:", msg)
	fmt.Println("The ID of  the message is:", msg.GetId())

	return &msg
}
