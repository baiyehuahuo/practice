package main

import (
	"log"
	"main/gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.GET("/", func(ctx *gee.Context) {
		ctx.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	r.GET("/hello", func(ctx *gee.Context) {
		// expect /hello?name=geektutu
		ctx.String(http.StatusOK, "hello %s, you're at %s\n", ctx.Query("name"), ctx.Path)
	})

	r.GET("/hello/:name", func(ctx *gee.Context) {
		// expect /hello/geektutu
		ctx.String(http.StatusOK, "hello %s, you're at %s\n", ctx.Param("name"), ctx.Path)
	})

	r.GET("/assets/*filepath", func(ctx *gee.Context) {
		ctx.JSON(http.StatusOK, gee.H{"filepath": ctx.Param("filepath")})
	})

	log.Fatal(r.Run(":9999"))
}

/*
curl -i http://localhost:9999/
curl http://localhost:9999/hello/geektutu
curl http://localhost:9999/assets/css/geektutu.css
curl "http://localhost:9999/hello?name=geektutu"
curl "http://localhost:9999/xxx"
*/
