package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/golang/protobuf/proto"
	simplepb "github.com/nitin1259/protobuf-learning/src/simple"
)

// example taken  - https://github.com/simplesteph/protobuf-example-go

func main() {
	fmt.Println("protobuf exaple in go lang")

	sm := doSimple()

	readAndWriteDemo(sm)
}

func readAndWriteDemo(sm proto.Message) {

	doWritetoFile("simple.bin", sm)
	sm2 := &simplepb.Simple{}
	doReadFromFile("simple.bin", sm2)
	fmt.Println("read from the file sm2: ", sm2)

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
