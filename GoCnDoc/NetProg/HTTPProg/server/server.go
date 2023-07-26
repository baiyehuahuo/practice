package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/go", myHandler)
	http.ListenAndServe("127.0.0.1:8000", nil)
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RemoteAddr, "connection success")
	fmt.Println("method: ", r.Method)
	fmt.Println("url: ", r.URL.Path)
	fmt.Println("header: ", r.Header)
	fmt.Println("body: ", r.Body)
	w.Write([]byte("Hello world, here is server."))
}
