package consistenthash

import (
	"hash/crc32"
	"sort"
	"strconv"
)

// Hash map bytes to uint32
type Hash func(data []byte) uint32

// Map contains all hashed keys
type Map struct {
	hash     Hash
	replicas int            // 虚拟节点倍数
	keys     []int          // 哈希环
	hashMap  map[int]string // 虚拟节点与真实节点的映射表
}

// New creates a Map instance
func New(replicas int, fn Hash) *Map {
	if fn == nil {
		fn = crc32.ChecksumIEEE
	}
	return &Map{
		hash:     fn,
		replicas: replicas,
		keys:     nil,
		hashMap:  make(map[int]string),
	}
}

// Add adds some keys to the hash
func (m *Map) Add(keys ...string) {
	for _, key := range keys {
		for i := 0; i < m.replicas; i++ { // 对每一个真实节点 key，对应创建 m.replicas 个虚拟副节点
			hashCode := int(m.hash([]byte(strconv.Itoa(i) + key)))
			m.keys = append(m.keys, hashCode) // 没考虑碰撞？
			m.hashMap[hashCode] = key         // 每个虚拟副节点上都保存了该值
		}
	}
	sort.Ints(m.keys)
}

// Get gets the closest item in the hash to the provided key
func (m *Map) Get(key string) string {
	if len(m.keys) == 0 {
		return ""
	}
	hashCode := int(m.hash([]byte(key)))
	idx := sort.Search(len(m.keys), func(i int) bool {
		return m.keys[i] >= hashCode
	})
	return m.hashMap[m.keys[idx%len(m.keys)]] // 当数值大于上限时，取模构成循环 取大于该值 且最接近的哈希值
}
