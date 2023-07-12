package singleflight

import "sync"

type call struct {
	wg  sync.WaitGroup
	val interface{}
	err error
}

type Group struct {
	mu sync.Mutex
	m  map[string]*call
}

// Do 针对相同的 key，无论 Do 被调用多少次，函数 fn 都只会被调用一次
func (g *Group) Do(key string, fn func() (interface{}, error)) (interface{}, error) {
	g.mu.Lock()
	if g.m == nil {
		g.m = make(map[string]*call)
	}
	if c, ok := g.m[key]; ok {
		g.mu.Unlock()
		c.wg.Wait() // 等相同请求结束 再返回
		return c.val, c.err
	}
	c := new(call)
	c.wg.Add(1)  // 同一时间 只有一个 key 在处理 请求前加锁
	g.m[key] = c // 添加到哈希表中，表示 key 正在处理
	g.mu.Unlock()

	c.val, c.err = fn() // 发起请求
	c.wg.Done()

	g.mu.Lock()
	delete(g.m, key)
	g.mu.Unlock()

	return c.val, c.err
}
