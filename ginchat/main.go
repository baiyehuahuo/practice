package main

import (
	"ginchat/router"
	"ginchat/utils"
	"log"
)

func main() {
	utils.InitConfig()
	utils.InitMySQL()

	r := router.Router()
	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}
