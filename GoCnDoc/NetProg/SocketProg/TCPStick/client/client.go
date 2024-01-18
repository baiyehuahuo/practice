package main

import (
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
	msg := []byte("Hello, Hello. How are you?")
	for i := 0; i < 20; i++ {
		if _, err = conn.Write(msg); err != nil {
			fmt.Println("send message failed, err: ", err)
		}
	}
}
