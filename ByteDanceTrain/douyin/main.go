package main

import (
	"douyin/router"
	"log"
)

func main() {
	r := router.SetupRouter()
	if err := r.Run("127.0.0.1:20000"); err != nil {
		log.Fatal(err)
	}
}
