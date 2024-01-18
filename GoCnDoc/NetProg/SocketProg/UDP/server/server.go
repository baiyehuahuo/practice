package main

import (
	"fmt"
	"net"
)

func main() {
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 20000,
	})
	if err != nil {
		fmt.Println("listen failed err: ", err)
		return
	}
	defer listen.Close()
	var data [1024]byte
	for {
		n, addr, err := listen.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println("read UDP failed, err: ", err)
			continue
		}
		fmt.Printf("data:%v addr: %v count: %v\n", string(data[:n]), addr.String(), n)
		_, err = listen.WriteToUDP(data[:n], addr)
		if err != nil {
			fmt.Println("write to UDP failed, err: ", err)
			continue
		}
	}
}
