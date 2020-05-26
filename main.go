package main

import (
	"fmt"

	simplepb "github.com/nitin1259/protobuf-learning/src/simple"
)
// example taken  - https://github.com/simplesteph/protobuf-example-go

func main() {
	fmt.Println("protobuf exaple in go lang")

	doSimple()
}

func doSimple(){

	sm:= simplepb.Simple{
		Id: 12345,
		Name: "sachindar tomar",
		IsSimple: true,
		SampleList: []int32{1,2,3,6},
	}

	fmt.Println(sm)
}