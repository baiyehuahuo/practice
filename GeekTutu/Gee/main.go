package main

import (
	"fmt"
	"log"
	"main/gee"
	"net/http"
	"text/template"
	"time"
)

type Student struct {
	Name string
	Age  int8
}

func FormatAsData(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%d", year, month, day)
}

func main() {
	r := gee.New()
	// template 模版的规则 不是很理解
	r.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsData,
	})
	r.LoadHTMLGlob("templates/*")
	// 加载静态文件系统
	r.Static("/assets", "./static")
	stu1 := &Student{Name: "Geektutu", Age: 20}
	stu2 := &Student{Name: "FanWeiFeng", Age: 23}
	// Engine 里有个未命名的 RouterGroup 可以直接使用
	// 它的前缀为 "" ，即所有路由都会使用
	r.Use(gee.Logger())
	r.Use(gee.Recovery())
	r.GET("/", func(ctx *gee.Context) {
		ctx.HTML(http.StatusOK, "css.tmpl", nil)
	})
	r.GET("/students", func(ctx *gee.Context) {
		ctx.HTML(http.StatusOK, "arr.tmpl", gee.H{
			"title":  "gee",
			"stuArr": [2]*Student{stu1, stu2},
		})
	})
	r.GET("/date", func(ctx *gee.Context) {
		ctx.HTML(http.StatusOK, "custom_func.tmpl", gee.H{
			"title": "gee",
			"now":   time.Now(),
		})
	})
	r.GET("/panic", func(ctx *gee.Context) {
		names := []string{"geektutu"}
		ctx.String(http.StatusOK, names[100])
	})

	log.Fatal(r.Run(":9999"))
}

/*
http://localhost:9999/
http://localhost:9999/date
http://localhost:9999/students
http://localhost:9999/assets/file1.txt
curl http://localhost:9999/panic
*/
