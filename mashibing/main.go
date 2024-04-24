package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"mashibing/model"
	"net/http"
)

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

func index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", nil)
}

func toRegister(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "register.html", nil)
}

func register(ctx *gin.Context) {
	var user model.User
	err := ctx.ShouldBind(&user)
	if err != nil {
		log.Println(err)
		return
	}
	ctx.String(http.StatusOK, "get User %v", user)
}

func MyHandler(ctx *gin.Context) {
	fmt.Println("My handler ...")
}

func MyHandlerB() func(ctx *gin.Context) {
	counter := 0
	return func(ctx *gin.Context) {
		counter++
		pth := ctx.FullPath()
		method := ctx.Request.Method
		fmt.Printf("My handlerB ... called times: %d\tfull path: %s\tmethod: %s\n", counter, pth, method)
	}
}

func main() {
	r := gin.Default()                     // 拿到一个 Engine 引擎 在New的基础上加入 Logger 与 Recovery 中间件
	r.Use(MyHandler, MyHandlerB())         // 与加入 engine 的顺序有关
	r.GET("ping", func(ctx *gin.Context) { // 获取 Get 连接的请求
		ctx.JSON(http.StatusOK, gin.H{"msg": "pong"})
	})
	r.GET("/getParam/:username", getParam)
	r.GET("/getParam/:username/:age", getParams)
	r.POST("/search", search)

	r.GET("/toRegister", toRegister)
	r.POST("/register", register)

	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "assets")
	r.GET("/index", index)

	fmt.Println("gin ... ")
	if err := r.Run(); err != nil { // 开启服务 默认监听127.0.0.1:8080
		log.Fatal(err)
	}
}
