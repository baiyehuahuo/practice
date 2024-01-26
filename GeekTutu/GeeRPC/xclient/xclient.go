package xclient

import (
	"context"
	"geerpc"
	"io"
	"log"
	"reflect"
	"sync"
)

type XClient struct {
	d       Discovery
	mode    SelectMode
	opt     *geerpc.Option
	mu      sync.Mutex
	clients map[string]*geerpc.Client
}

var _ io.Closer = (*XClient)(nil)

func NewXClient(d Discovery, mode SelectMode, opt *geerpc.Option) *XClient {
	return &XClient{
		d:       d,
		mode:    mode,
		opt:     opt,
		mu:      sync.Mutex{},
		clients: make(map[string]*geerpc.Client),
	}
}

func (xc *XClient) Close() error {
	xc.mu.Lock()
	defer xc.mu.Unlock()
	for key, client := range xc.clients {
		if err := client.Close(); err != nil {
			log.Printf("client %s close failed: %s", key, err.Error())
		}
		delete(xc.clients, key)
	}
	return nil
}

func (xc *XClient) dial(rpcAddr string) (*geerpc.Client, error) {
	xc.mu.Lock()
	defer xc.mu.Unlock()

	client, ok := xc.clients[rpcAddr]
	if ok && !client.IsAvailable() {
		_ = client.Close()
		delete(xc.clients, rpcAddr)
		client = nil
	}
	if client == nil {
		var err error
		client, err = geerpc.XDial(rpcAddr, xc.opt)
		if err != nil {
			return nil, err
		}
		xc.clients[rpcAddr] = client
	}
	return client, nil
}

func (xc *XClient) call(rpcAddr string, ctx context.Context, serviceMethod string, args, reply interface{}) error {
	client, err := xc.dial(rpcAddr)
	if err != nil {
		return err
	}
	return client.Call(ctx, serviceMethod, args, reply)
}

// Call invokes the named function, waits for it to complete, and returns its error status.
// xc will choose a proper server.
func (xc *XClient) Call(ctx context.Context, serviceMode string, args, reply interface{}) error {
	rpcAddr, err := xc.d.Get(xc.mode)
	if err != nil {
		return err
	}
	return xc.call(rpcAddr, ctx, serviceMode, args, reply)
}

// Broadcast 将请求广播到所有的服务实例，如果任意一个实例发生错误，则返回其中一个错误；如果调用成功，则返回其中一个的结果。
// 为了提升性能，请求是并发的。
// 并发情况下需要使用互斥锁保证 error 和 reply 能被正确赋值。
// 借助 context.WithCancel 确保有错误发生时，快速失败。
func (xc *XClient) Broadcast(ctx context.Context, serviceMode string, args, reply interface{}) error {
	servers, err := xc.d.GetAll()
	if err != nil {
		return err
	}
	var wg sync.WaitGroup
	var mu sync.Mutex
	var e error
	replyDone := reply == nil
	ctx, cancel := context.WithCancel(ctx)
	for _, rpcAddr := range servers {
		wg.Add(1)
		go func(rpcAddr string) {
			defer wg.Done()
			var clonedReply interface{}
			if reply != nil {
				clonedReply = reflect.New(reflect.ValueOf(reply).Elem().Type()).Interface()
			}
			err := xc.call(rpcAddr, ctx, serviceMode, args, clonedReply)
			mu.Lock()
			defer mu.Unlock()
			if err != nil && e == nil {
				e = err
				cancel()
			}
			if err == nil && !replyDone {
				reflect.ValueOf(reply).Elem().Set(reflect.ValueOf(clonedReply).Elem())
				replyDone = true
			}
		}(rpcAddr)
	}
	wg.Wait()
	return e
}
