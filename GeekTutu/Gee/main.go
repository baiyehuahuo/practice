package main

import (
	"log"
	"main/gee"
	"net/http"
	"time"
)

func onlyForV2() gee.HandlerFunc {
	return func(ctx *gee.Context) {
		t := time.Now()
		ctx.Fail(500, "Internal Server Error")
		log.Printf("[%d] %s in %v for group v2", ctx.StatusCode, ctx.Req.RequestURI, time.Since(t))
	}
}

func main() {
	r := gee.New()
	r.GET("/index", func(ctx *gee.Context) {
		ctx.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})
	// Engine 里有个未命名的 RouterGroup 可以直接使用
	// 它的前缀为 "" ，即所有路由都会使用
	r.Use(gee.Logger())
	v1 := r.Group("/v1")
	{
		v1.GET("/", func(ctx *gee.Context) {
			ctx.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
		})
		v1.GET("/hello", func(ctx *gee.Context) {
			ctx.String(http.StatusOK, "hello %s, you're at %s\n", ctx.Query("name"), ctx.Path)
		})
	}

	v2 := r.Group("/v2")
	v2.Use(onlyForV2())
	{
		v2.GET("/hello/:name", func(ctx *gee.Context) {
			ctx.String(http.StatusOK, "hello %s, you're at %s\n", ctx.Param("name"), ctx.Path)
		})
		v2.POST("/login", func(ctx *gee.Context) {
			ctx.JSON(http.StatusOK, gee.H{
				"username": ctx.PostForm("username"),
				"password": ctx.PostForm("password"),
			})
		})
	}

	log.Fatal(r.Run(":9999"))
}

/*
curl -i http://localhost:9999/index
curl http://localhost:9999/v1/hello?name=geektutu
curl http://localhost:9999/v2/hello/geektutu
*/
