package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func someGet(ctx *gin.Context) {
	_, _ = ctx.Writer.WriteString("this is someGet")
}

func somePost(ctx *gin.Context) {
	_, _ = ctx.Writer.WriteString("this is somePost")
}

func getKey(ctx *gin.Context) {
	s := ctx.Query("name")
	ss := ctx.DefaultQuery("age", "xxoo")
	ctx.String(http.StatusOK, "Get Key name = %s\t%s", s, ss)
}

func postVal(ctx *gin.Context) {
	name := ctx.PostForm("name")
	age := ctx.DefaultPostForm("age", "xxoo")
	//ctx.String(http.StatusOK, "Post Val name = %s\t%s", s, ss)
	ctx.HTML(http.StatusOK, "welcome.html", gin.H{"name": name, "age": age})
}

func getParam(ctx *gin.Context) {
	s := ctx.Param("username")
	ctx.String(http.StatusOK, "Get Param Key name = %s", s)
}

func getParams(ctx *gin.Context) {
	name := ctx.Param("username")
	age := ctx.Param("age")
	ctx.String(http.StatusOK, "Get Param Key name = %s, age = %s", name, age)
}

func search(ctx *gin.Context) {
	page := ctx.DefaultQuery("page", "default string")
	key := ctx.PostForm("key")
	age := ctx.PostForm("age")
	hobby := ctx.PostFormArray("hobby")
	ctx.String(http.StatusOK, "this is search() page=%s\tkey=%s\tage=%s\thobby=%v\tcount of hobby=%d", page, key, age, hobby, len(hobby))
}

func main() {
	r := gin.Default()                     // 拿到一个 Engine 引擎 在New的基础上加入 Logger 与 Recovery 中间件
	r.GET("ping", func(ctx *gin.Context) { // 获取 Get 连接的请求
		ctx.JSON(http.StatusOK, gin.H{"msg": "pong"})
	})
	r.GET("/someGet", someGet)
	r.GET("/getKey", getKey)
	r.GET("/getParam/:username", getParam)
	r.GET("/getParam/:username/:age", getParams)
	r.POST("/somePost", somePost)
	r.POST("/postVal", postVal)
	r.POST("/search", search)
	r.LoadHTMLGlob("templates/*")
	fmt.Println("gin ... ")
	if err := r.Run(); err != nil { // 开启服务 默认监听127.0.0.1:8080
		log.Fatal(err)
	}
}
