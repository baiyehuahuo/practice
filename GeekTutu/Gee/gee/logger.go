package gee

import (
	"log"
	"time"
)

func Logger() HandlerFunc {
	return func(ctx *Context) {
		// start time
		t := time.Now()
		// Process request
		ctx.Next() // 先去执行下一个中间件 执行结束再回来
		log.Printf("[%d] %s in %v", ctx.StatusCode, ctx.Req.RequestURI, time.Since(t))
	}
}
