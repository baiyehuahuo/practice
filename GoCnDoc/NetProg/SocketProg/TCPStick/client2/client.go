package main

import (
	"TCPStick/proto"
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("dial server failed, err: ", err)
		return
	}
	defer conn.Close()
	msg := "Hello, Hello. How are you?"
	data, err := proto.Encode(msg)
	if err != nil {
		fmt.Println("encode msg failed, err: ", err)
		return
	}
	for i := 0; i < 20; i++ {
		if _, err = conn.Write(data); err != nil {
			fmt.Println("send message failed, err: ", err)
		}
	}
}
