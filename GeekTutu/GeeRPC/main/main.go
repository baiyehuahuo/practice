package main

import (
	"context"
	"geerpc"
	"geerpc/registry"
	"geerpc/xclient"
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

func call(registry string) {
	d := xclient.NewGeeRegistryDiscovery(registry, 0)
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

func broadcast(registry string) {
	d := xclient.NewGeeRegistryDiscovery(registry, 0)
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

func startRegistry(wg *sync.WaitGroup) {
	// 启动注册中心
	l, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Println("Registry listen port 9999 failed: ", err.Error())
		return
	}
	registry.HandleHTTP()
	wg.Done()
	if err = http.Serve(l, nil); err != nil {
		log.Println("Registry listen failed: ", err.Error())
	}
}

func startServer(registryAddr string, wg *sync.WaitGroup) {
	l, _ := net.Listen("tcp", ":0")
	server := geerpc.NewServer()
	err := server.Register(new(Foo))
	if err != nil {
		log.Printf("Register server %s failed: %s", l.Addr().String(), err.Error())
		return
	}
	// 注册服务器
	registry.Heartbeat(registryAddr, "tcp@"+l.Addr().String(), 0)
	wg.Done()
	server.Accept(l)
}

func main() {
	log.SetFlags(0)
	registryAddr := "http://localhost:9999/_geerpc_/registry"
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go startRegistry(wg)
	wg.Wait()
	time.Sleep(time.Second)

	wg.Add(2)
	go startServer(registryAddr, wg)
	go startServer(registryAddr, wg)
	wg.Wait()

	time.Sleep(time.Second)
	call(registryAddr)
	broadcast(registryAddr)
}
