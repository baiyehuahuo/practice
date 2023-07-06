package main

import (
	"log"
	"main/gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.GET("/index", func(ctx *gee.Context) {
		ctx.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})
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
