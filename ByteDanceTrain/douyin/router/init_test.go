package router

import "github.com/gin-gonic/gin"

var (
	buf = make([]byte, 1e5)
	r   = gin.New()
)

func init() {
	SetupRouter(r)
}
