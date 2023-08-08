package main

import (
	"douyin/constants"
	"douyin/router"
	_ "douyin/service/DBService"
	"log"
	"os"
)

func main() {
	r := router.SetupRouter()
	if err := r.Run("127.0.0.1:20000"); err != nil {
		log.Fatal(err)
	}
}

func init() {
	if _, err := os.Stat(constants.UploadFileDir); err != nil {
		if err = os.Mkdir(constants.UploadFileDir, 0777); err != nil {
			panic(err)
		}
	}
}
