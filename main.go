package main

import (
	"fmt"
	"io/ioutil"
	"log"

	jsonpb "github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	complexpb "github.com/nitin1259/protobuf-learning/src/complex"
	simplepb "github.com/nitin1259/protobuf-learning/src/simple"
)

// example taken  - https://github.com/simplesteph/protobuf-example-go

func main() {
	fmt.Println("protobuf exaple in go lang")

	sm := doSimple()

	readAndWriteDemo(sm)

	jsonDemo(sm)

	doComplex()
}

func doComplex() {
	complex := complexpb.ComplexMessage{
		OneDummy: &complexpb.DummyMessage{
			Id:   1,
			Name: "Nitin",
		},
		MultipleDummy: []*complexpb.DummyMessage{
			&complexpb.DummyMessage{
				Id:   2,
				Name: "Vipin",
			},
			&complexpb.DummyMessage{
				Id:   3,
				Name: "Sachin",
			},
		},
	}

	fmt.Println("Complex message: ", complex)
}

func readAndWriteDemo(sm proto.Message) {

	doWritetoFile("simple.bin", sm)
	sm2 := &simplepb.Simple{}
	doReadFromFile("simple.bin", sm2)
	fmt.Println("read from the file sm2: ", sm2)

}

func jsonDemo(sm proto.Message) {
	smAsString := toJSON(sm)
	fmt.Println(smAsString)

	sm2 := &simplepb.Simple{}
	fromJSON(smAsString, sm2)
	fmt.Println("Successfully created proto struct:", sm2)
}

func toJSON(sm proto.Message) string {
	marshaler := jsonpb.Marshaler{}
	out, err := marshaler.MarshalToString(sm)
	if err != nil {
		log.Fatal("Cant convert to json: ", err)
	}
	return out
}

func fromJSON(in string, sm proto.Message) {
	err := jsonpb.UnmarshalString(in, sm)
	if err != nil {
		log.Fatalln("Couldn't unmarshal the JSON into the pb struct", err)
	}

}

func doReadFromFile(fileanme string, sm proto.Message) {
	in, err := ioutil.ReadFile(fileanme)
	if err != nil {
		log.Fatal("Cant serialize to bytes: ", err)
	}

	if err := proto.Unmarshal(in, sm); err != nil {
		log.Fatal("Cant serialize to bytes: ", err)
	}

	fmt.Println("reading from the file: ", sm)
}

func doWritetoFile(filename string, sm proto.Message) {
	out, err := proto.Marshal(sm)

	if err != nil {
		log.Fatal("Cant serialize to bytes: ", err)
	}

	if err := ioutil.WriteFile(filename, out, 0644); err != nil {
		log.Fatal("Cant serialize to bytes: ", err)
	}

	fmt.Println("data Written to file")
}

func doSimple() *simplepb.Simple {

	sm := simplepb.Simple{
		Id:         12345,
		Name:       "sachindar tomar",
		IsSimple:   true,
		SampleList: []int32{1, 2, 3, 6},
	}

	fmt.Println(sm)

	return &sm
}
