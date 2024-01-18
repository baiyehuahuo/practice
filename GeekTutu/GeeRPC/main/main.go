package main

import (
	"encoding/json"
	"fmt"
	GeeRPC "geerpc"
	"geerpc/codec"
	"log"
	"net"
	"time"
)

// pick a free port
func startServer(addr chan string) {
	lis, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		log.Fatal("network error: ", err)
	}
	log.Println("start rpc server on ", lis.Addr())
	addr <- lis.Addr().String()
	GeeRPC.Accept(lis)
}

func main() {
	addr := make(chan string)
	go startServer(addr) // 启动服务器

	// 客户端请求
	conn, err := net.Dial("tcp", <-addr)
	if err != nil {
		log.Fatal("client dial error: ", err)
	}
	defer conn.Close()

	time.Sleep(time.Second)
	// 发送选项
	json.NewEncoder(conn).Encode(GeeRPC.Option{MagicNumber: GeeRPC.MagicNumebr, CodecType: codec.JsonType})
	//json.NewEncoder(conn).Encode(GeeRPC.Option{MagicNumber: GeeRPC.MagicNumebr, CodecType: codec.GobType})
	cc := codec.NewJsonCodec(conn)
	//cc := codec.NewGobCodec(conn)
	for i := 0; i < 5; i++ {
		header := &codec.Header{
			ServiceMethod: "Foo.Sum",
			Seq:           uint64(i),
		}
		if err = cc.Write(header, fmt.Sprintf("geerpc req %d", header.Seq)); err != nil {
			log.Println("Write seq failed: ", err)
			break
		}
		if err = cc.ReadHeader(header); err != nil {
			log.Println("Read header failed: ", err)
			break
		}
		var reply string
		if err = cc.ReadBody(&reply); err != nil {
			log.Println("Read reply failed: ", err)
			break
		}
		log.Println("reply:", reply)
	}
}
