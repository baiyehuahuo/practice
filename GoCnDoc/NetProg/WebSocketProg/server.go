package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	go h.run()
	router.HandleFunc("/ws", handleRequest)
	if err := http.ListenAndServe("127.0.0.1:8080", router); err != nil {
		log.Fatal(err)
	}
}
