package registry

import (
	"log"
	"net/http"
	"sort"
	"strings"
	"sync"
	"time"
)

type GeeRegistry struct {
	timeout time.Duration
	mu      sync.Mutex
	servers map[string]*ServerItem
}

type ServerItem struct {
	Addr  string
	start time.Time
}

const (
	defaultPath    = "/_geerpc_/registry"
	defaultTimeout = time.Minute * 5
	GeerpcHeader   = "X-Geerpc-Server"
)

func New(timeout time.Duration) *GeeRegistry {
	return &GeeRegistry{
		timeout: timeout,
		mu:      sync.Mutex{},
		servers: make(map[string]*ServerItem),
	}
}

var DefaultGeeRegister = New(defaultTimeout)

// 激活一下
func (r *GeeRegistry) putServer(addr string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	s := r.servers[addr]
	if s == nil {
		r.servers[addr] = &ServerItem{
			Addr:  addr,
			start: time.Now(),
		}
	} else {
		s.start = time.Now()
	}
}

// 获取在线服务器
func (r *GeeRegistry) aliveServers() []string {
	r.mu.Lock()
	defer r.mu.Unlock()
	var alive []string
	for addr, s := range r.servers {
		if r.timeout == 0 || s.start.Add(r.timeout).After(time.Now()) {
			alive = append(alive, addr)
		} else {
			delete(r.servers, addr)
		}
	}
	sort.Strings(alive)
	return alive
}

// 服务客户端
func (r *GeeRegistry) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		w.Header().Set(GeerpcHeader, strings.Join(r.aliveServers(), ","))
	case "POST":
		addr := req.Header.Get(GeerpcHeader)
		if addr == "" {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		r.putServer(addr)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (r *GeeRegistry) HandleHTTP(registryPath string) {
	// http 启动服务
	http.Handle(registryPath, r)
	log.Println("rpc registry path: ", registryPath)
}

func HandleHTTP() {
	// http 启动服务
	DefaultGeeRegister.HandleHTTP(defaultPath)
}

// Heartbeat 便于服务启动时定时向注册中心发送心跳，默认周期比注册中心设置的过期时间少 1 min。
// 同样可以在这里注册服务
func Heartbeat(registry, addr string, duration time.Duration) {
	if duration == 0 {
		duration = defaultTimeout - time.Minute
	}
	go func() {
		err := sendHeartbeat(registry, addr)
		t := time.NewTicker(duration)
		for err == nil {
			<-t.C
			err = sendHeartbeat(registry, addr)
		}
	}()
}

// 向注册中心发出心跳
func sendHeartbeat(registry, addr string) error {
	log.Println(addr, " send heart beat to registry ", registry)
	httpClient := &http.Client{}
	req, _ := http.NewRequest("POST", registry, nil)
	req.Header.Set(GeerpcHeader, addr)
	if _, err := httpClient.Do(req); err != nil {
		log.Println("rpc server: heat beat err: ", err)
		return err
	}
	return nil
}
