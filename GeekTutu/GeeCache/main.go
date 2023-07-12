package main

import (
	"flag"
	"fmt"
	"geecache"
	"log"
	"net/http"
)

var db = map[string]string{
	"Tom":  "630",
	"Jack": "589",
	"Sam":  " 567",
}

func createGroup() *geecache.Group {
	return geecache.NewGroup("scores", 2<<10, geecache.GetterFunc(
		func(key string) ([]byte, error) {
			log.Println("[SlowDB] search key", key)
			if v, ok := db[key]; ok {
				return []byte(v), nil
			}
			return nil, fmt.Errorf("%s not exist", key)
		}))
}

// 启动缓存服务器 有三个接口 但用户不感知
func startCacheServer(addr string, addrs []string, gee *geecache.Group) {
	peers := geecache.NewHTTPPool(addr)
	peers.Set(addrs...)
	gee.RegisterPeers(peers)
	log.Println("geecache is running at", addr)
	log.Fatal(http.ListenAndServe(addr[7:], peers))
}

// 与用户交互 用户感知
func startAPIServer(apiAddr string, gee *geecache.Group) {
	http.Handle("/api", http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			key := r.URL.Query().Get("key")
			view, err := gee.Get(key)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/octet-stream")
			w.Write(view.ByteSlice())
		}))
	log.Println("fontend server is running at", apiAddr)
	log.Fatal(http.ListenAndServe(apiAddr[7:], nil))
}

func main() {
	port := flag.Int("port", 8002, "Geecache server port")
	api := flag.Bool("api", false, "Start a api server?")
	flag.Parse()
	fmt.Println("Terminate input: ", *port, *api)

	apiAddr := "http://localhost:9999"
	addrMap := map[int]string{
		8001: "http://localhost:8001",
		8002: "http://localhost:8002",
		8003: "http://localhost:8003",
	}
	var addrs []string
	for _, v := range addrMap {
		addrs = append(addrs, v)
	}
	gee := createGroup()
	if *api {
		go startAPIServer(apiAddr, gee)
	}
	startCacheServer(addrMap[*port], addrs, gee)
}

/*

缓存雪崩：缓存在同一时刻全部失效，造成瞬时 DB 请求量大，造成雪崩。通常因为缓存服务器宕机、缓存的 key 设置了相同的过期时间等引起
缓存击穿：一个存在的 key，在缓存过期的瞬间，有大量请求，击穿缓存到 DB，造成 DB 瞬时请求量大
缓存穿透：查询一个不存在的数据，因为不存在，不会写在缓存中，不断请求 DB，如果瞬间流量过大，穿透到 DB，造成宕机

这三个要并行运行，开三个窗口：
go run . -port=8001
go run . -port=8002
go run . -port=8003 -api

curl "http://localhost:9999/api?key=Tom"
*/
