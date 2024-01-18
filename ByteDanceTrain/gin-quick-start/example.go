package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	return r
}

func main() {
	r := setupRouter()
	if err := r.Run("127.0.0.1:8080"); err != nil {
		panic(err)
	}
}
