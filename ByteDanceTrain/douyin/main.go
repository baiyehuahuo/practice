package main

import (
	"douyin/constants"
	"douyin/router"
	_ "douyin/service/DBService"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	r := gin.Default()
	router.SetupRouter(r)
	if err := r.Run(":20000"); err != nil {
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
