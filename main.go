package main

import (
	"fmt"
	complex2 "github.com/ACSG-64/Protobuf3_demostration/src/protobuf/complex"
	"github.com/ACSG-64/Protobuf3_demostration/src/protobuf/enum"
	"github.com/ACSG-64/Protobuf3_demostration/src/protobuf/simple"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"log"
)

func main() {
	simpleMsg := doSimpleMessage()

	// Writing to disk
	_ = writeToFile("message.bin", simpleMsg)
	// Reading from disk
	newSimpleMsg1 := simple.SimpleMessage{}
	_ = readFromFile("message.bin", &newSimpleMsg1)
	fmt.Println("Message read from disk:", newSimpleMsg1)

	// To JSON
	json, _ := toJSON(simpleMsg)
	fmt.Println("JSON format:", json)
	// Parse from JSON
	newSimpleMsg2 := simple.SimpleMessage{}
	_ = fromJSON(json, &newSimpleMsg2)
	fmt.Println("Converted from JSON:", newSimpleMsg2)
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

func doSimpleMessage() *simple.SimpleMessage {
	msg := simple.SimpleMessage{
		Id:         1234,
		IsSimple:   true,
		Name:       "Basic message",
		SampleList: []int32{2, 4, 6, 8, 10},
	}

	fmt.Println("Message:", msg)
	fmt.Println("The ID of  the message is:", msg.GetId())

	return &msg
}

func doEnumMessage() *enum.EnumMessage {
	msg := enum.EnumMessage{
		Id:           5678,
		DayOfTheWeek: enum.DayOfTheWeek_WEDNESDAY, // Using enum field
	}

	fmt.Println("Message:", msg)
	fmt.Println("The ID of  the message is:", msg.GetId())

	return &msg
}

func doComplexMessage() *complex2.ComplexMessage {
	msg := complex2.ComplexMessage{
		OneDummy: &complex2.DummyMessage{
			Id:   1,
			Name: "A very first message!",
		},
		MultipleDummy: []*complex2.DummyMessage{
			&complex2.DummyMessage{
				Id:   2,
				Name: "A first sub message!",
			},
			&complex2.DummyMessage{
				Id:   3,
				Name: "A second sub message!",
			},
			&complex2.DummyMessage{
				Id:   4,
				Name: "A third sub message!",
			},
		},
	}

	fmt.Println("Message:", msg)
	fmt.Println("The ID of  the message is:", msg.GetOneDummy().GetId())

	return &msg
}
