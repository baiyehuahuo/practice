package main

import (
	"context"
	"geerpc"
	"geerpc/xclient"
	"log"
	"net"
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

func (f Foo) Sleep(args Args, reply *int) error {
	time.Sleep(time.Second * time.Duration(args.Num1))
	*reply = args.Num1 + args.Num2
	return nil
}

func foo(xc *xclient.XClient, ctx context.Context, typ, serviceMethod string, args *Args) {
	var reply int
	var err error
	switch typ {
	case "call":
		err = xc.Call(ctx, serviceMethod, args, &reply)
	case "broadcast":
		err = xc.Broadcast(ctx, serviceMethod, args, &reply)
	}
	if err != nil {
		log.Printf("%s %s error: %v", typ, serviceMethod, err.Error())
	} else {
		log.Printf("%s %s success: %d + %d = %d", typ, serviceMethod, args.Num1, args.Num2, reply)
	}
}

func call(addr1, addr2 string) {
	d := xclient.NewMultiServersDiscovery([]string{"tcp@" + addr1, "tcp@" + addr2})
	xc := xclient.NewXClient(d, xclient.RandomSelect, nil)
	defer func() {
		if err := xc.Close(); err != nil {
			log.Printf("close xclient failed: %s", err.Error())
		}
	}()
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			foo(xc, context.Background(), "call", "Foo.Sum", &Args{
				Num1: i,
				Num2: i * i,
			})
		}(i)
	}
	wg.Wait()
}

func broadcast(addr1, addr2 string) {
	d := xclient.NewMultiServersDiscovery([]string{"tcp@" + addr1, "tcp@" + addr2})
	xc := xclient.NewXClient(d, xclient.RandomSelect, nil)
	defer func() {
		if err := xc.Close(); err != nil {
			log.Printf("close xclient failed: %s", err.Error())
		}
	}()
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			foo(xc, context.Background(), "broadcast", "Foo.Sum", &Args{
				Num1: i,
				Num2: i * i,
			})
			ctx, _ := context.WithTimeout(context.Background(), time.Second*2)
			foo(xc, ctx, "broadcast", "Foo.Sleep", &Args{
				Num1: i,
				Num2: i * i,
			})
		}(i)
	}
	wg.Wait()
}

func startServer(addr chan string) {
	l, _ := net.Listen("tcp", ":0")
	server := geerpc.NewServer()
	err := server.Register(new(Foo))
	if err != nil {
		log.Printf("Register server %s failed: %s", l.Addr().String(), err.Error())
		return
	}
	addr <- l.Addr().String()
	server.Accept(l)
}

func main() {
	log.SetFlags(0)
	ch1 := make(chan string)
	ch2 := make(chan string)
	go startServer(ch1)
	go startServer(ch2)

	addr1 := <-ch1
	addr2 := <-ch2
	time.Sleep(time.Second)
	call(addr1, addr2)
	broadcast(addr1, addr2)
}
