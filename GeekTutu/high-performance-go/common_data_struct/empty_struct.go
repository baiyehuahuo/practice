package main

import (
	"fmt"
	"unsafe"
)

type Set map[string]struct{}

func (s Set) Has(key string) bool {
	_, ok := s[key]
	return ok
}

func (s Set) Add(key string) {
	s[key] = struct{}{}
}

func SetTest() {
	fmt.Println(unsafe.Sizeof(struct{}{}))
	s := make(Set)
	s.Add("Tom")
	s.Add("Sam")
	fmt.Println(s.Has("Tom"))
	fmt.Println(s.Has("Jack"))
}

func (s Set) Delete(key string) {
	delete(s, key)
}

func worker(ch chan struct{}) {
	<-ch
	fmt.Println("do something")
	close(ch)
}

func ChannelTest() {
	ch := make(chan struct{})
	go worker(ch)
	ch <- struct{}{}
}

type Door struct {
}

func (d Door) Open() {
	fmt.Println("Open the door")
}

func (d Door) Close() {
	fmt.Println("Close the door")
}

func DoorTest() {
	d := new(Door)
	d.Open()
	d.Close()
}

//func main() {
//SetTest()
//ChannelTest()
//DoorTest()
//}
