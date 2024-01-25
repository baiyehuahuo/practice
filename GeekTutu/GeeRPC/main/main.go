package main

import (
	"context"
	"geerpc"
	"log"
	"net"
	"net/http"
	"sync"
	"time"
)

type Foo int

type Args struct {
	Num1, Num2 int
}

func (f Foo) Sum(args Args, reply *int) error {
	*reply = args.Num1 + args.Num2
	return nil
}

func startServer(addr chan string) {
	l, _ := net.Listen("tcp", ":9999")
	geerpc.Register(new(Foo))
	geerpc.HandleHTTP()
	log.Println("start rpc server on ", l.Addr())
	addr <- l.Addr().String()
	//go geerpc.Accept(l)
	_ = http.Serve(l, nil)
}

func call(addrCh chan string) {
	client, err := geerpc.DialHTTP("tcp", <-addrCh)
	if err != nil {
		panic(err)
	}
	defer func() { _ = client.Close() }()

	time.Sleep(time.Second)
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			args := &Args{Num1: i, Num2: i * i}
			var reply int
			if err := client.Call(context.Background(), "Foo.Sum", args, &reply); err != nil {
				log.Println("call Foo.Sum error: ", err)
			}
			log.Printf("%d + %d = %d", args.Num1, args.Num2, reply)
		}(i)
	}
	wg.Wait()
}

func main() {
	ch := make(chan string)
	go call(ch)
	startServer(ch)
}
