package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func someGet(ctx *gin.Context) {
	ctx.Writer.WriteString("this is someGet")
}

func somePost(ctx *gin.Context) {
	ctx.Writer.WriteString("this is somePost")
}

func getKey(ctx *gin.Context) {
	s := ctx.Query("name")
	ss := ctx.DefaultQuery("age", "xxoo")
	ctx.String(http.StatusOK, "Get Param Key name = %s\t%s", s, ss)
}

func postVal(ctx *gin.Context) {
	s := ctx.PostForm("name")
	ss := ctx.DefaultPostForm("age", "xxoo")
	ctx.String(http.StatusOK, "Get Param Key name = %s\t%s", s, ss)
}

func main() {
	r := gin.Default()                     // 拿到一个 Engine 引擎 在New的基础上加入 Logger 与 Recovery 中间件
	r.GET("ping", func(ctx *gin.Context) { // 获取 Get 连接的请求
		ctx.JSON(http.StatusOK, gin.H{"msg": "pong"})
	})
	r.GET("someGet", someGet)
	r.GET("getKey", getKey)
	r.POST("somePost", somePost)
	r.POST("postVal", postVal)
	fmt.Println("gin ... ")
	if err := r.Run(); err != nil { // 开启服务 默认监听127.0.0.1:8080
		log.Fatal(err)
	}
}
