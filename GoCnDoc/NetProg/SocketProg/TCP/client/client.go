package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("dial err:", err)
		return
	}
	defer conn.Close()
	inputReader := bufio.NewReader(os.Stdin)
	var buf [512]byte
	for {
		input, _ := inputReader.ReadString('\n')
		inputInfo := strings.Trim(input, "\r\n")
		if strings.ToUpper(inputInfo) == "Q" {
			break
		}
		if _, err = conn.Write([]byte(inputInfo)); err != nil {
			fmt.Println("Write err: ", err)
			return
		}
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("recv failed err: ", err)
			break
		}
		fmt.Println(string(buf[:n]))
	}
}
