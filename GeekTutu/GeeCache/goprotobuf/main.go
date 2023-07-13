package main

import (
	"github.com/golang/protobuf/proto"
	"log"
)

func main() {
	var err error
	test := &Student{
		Name:   "Geektutu",
		Male:   true,
		Scores: []int32{98, 85, 88},
	}
	var data []byte
	if data, err = proto.Marshal(test); err != nil {
		log.Fatal("marshaling error: ", err)
	}
	newTest := &Student{}
	if err = proto.Unmarshal(data, newTest); err != nil {
		log.Fatal("unmarshalling error: ", err)
	}
	if test.GetName() != newTest.GetName() {
		log.Fatalf("data mismatch %q != %q", test.GetName(), newTest.GetName())
	}
	log.Println("Test protobuf success.")
}
