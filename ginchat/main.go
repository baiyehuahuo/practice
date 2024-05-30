package main

import (
	"ginchat/router"
	"log"
)

func main() {
	r := router.Router()
	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}
