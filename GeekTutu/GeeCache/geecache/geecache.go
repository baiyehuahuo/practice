// 负责与外部交互，控制缓存存储和获取主流程
package geecache

import (
	"fmt"
	"sync"
)

// A Getter loads data for a key
type Getter interface {
	Get(key string) ([]byte, error)
}

// A GetterFunc implements Getter with a function
type GetterFunc func(key string) ([]byte, error)

// Get implements Getter interface function
func (f GetterFunc) Get(key string) ([]byte, error) {
	return f(key)
}

// A Group is a cache namespace and associated data loaded spread over
type Group struct {
	name      string
	getter    Getter
	mainCache cache
}

var (
	mu     sync.RWMutex
	groups = make(map[string]*Group)
)

// NewGroup create a new instance of Group
func NewGroup(name string, cacheBytes int64, getter Getter) *Group {
	if getter == nil {
		panic("nil Getter")
	}
	mu.Lock()
	defer mu.Unlock()
	g := &Group{
		name:      name,
		getter:    getter,
		mainCache: cache{cacheBytes: cacheBytes},
	}
	groups[name] = g
	return g
}

// GetGroup returns the named group previously created with NewGroup or nil if there's no such group
func GetGroup(name string) *Group {
	mu.RLock()
	defer mu.RUnlock()
	g := groups[name]
	return g
}

// Get value for a key from cache
func (g *Group) Get(key string) (ByteView, error) {
	if key == "" {
		return ByteView{}, fmt.Errorf("key is required")
	}
	if v, ok := g.mainCache.get(key); ok { // 先查缓存 缓存查得到就返回
		return v, nil
	}
	return g.load(key) // 缓存查不到就查数据库
}

func (g *Group) load(key string) (ByteView, error) {
	return g.getLocally(key)
}

func (g *Group) getLocally(key string) (ByteView, error) {
	bytes, err := g.getter.Get(key) // 查数据库
	if err != nil {                 // 数据库查不到就返回错误信息
		return ByteView{}, err
	}
	value := ByteView{b: cloneBytes(bytes)}
	g.populateCache(key, value) // 数据库查得到就保存到缓存里
	return value, nil
}

func (g *Group) populateCache(key string, value ByteView) {
	g.mainCache.add(key, value)
}
