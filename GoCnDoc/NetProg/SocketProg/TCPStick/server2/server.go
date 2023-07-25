package main

import (
	"TCPStick/proto"
	"bufio"
	"fmt"
	"io"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		msg, err := proto.Decode(reader)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read from client failed, err: ", err)
			break
		}
		fmt.Println("receive message from client: ", msg)
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("listen failed, err: ", err)
		return
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept client failed, err: ", err)
			continue
		}
		go process(conn)
	}
}
