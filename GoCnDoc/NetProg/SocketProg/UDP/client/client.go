package main

import (
	"fmt"
	"net"
)

func main() {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 20000,
	})
	if err != nil {
		fmt.Println("Dial UDP server failed err: ", err)
		return
	}
	defer socket.Close()
	sendData := []byte("Hello server")
	if _, err = socket.Write(sendData); err != nil {
		fmt.Println("Send data failed err: ", err)
		return
	}
	data := make([]byte, 4096)
	n, remoteAddr, err := socket.ReadFromUDP(data)
	if err != nil {
		fmt.Println("receive data from server failed: ", err)
		return
	}
	fmt.Printf("recv: %v addr: %v count: %v\n", string(data[:n]), remoteAddr.String(), n)
}
