package main

import (
	"bufio"
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	var buf [128]byte // 缓冲区
	for {
		reader := bufio.NewReader(conn) // 读取器
		n, err := reader.Read(buf[:])   // 收
		if err != nil {
			fmt.Println("read from client failed.")
			break
		}
		recv := string(buf[:n])
		fmt.Println("Get Message from client: ", recv)
		conn.Write([]byte(recv)) // 发
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed, err", err)
			continue
		}
		go process(conn)
	}
}
