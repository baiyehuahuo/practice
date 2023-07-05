package main

import (
	"fmt"
	"log"
	"net/http"
)

type Engine struct {
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%v] = %v\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 Not Found: %v\n", req.URL.Path)
	}
}

func main() {
	log.Fatal(http.ListenAndServe(":9999", &Engine{}))
}
