package router

import (
	"douyin/service/basis"
	"github.com/gin-gonic/gin"
	"net/http"
	_ "net/http/pprof"
)

// SetupRouter get a configured gin engine
func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Static("/static", "./assets")
	douyinGroup := r.Group("/douyin")
	{
		// baseproto interfaces
		douyinGroup.GET("/feed", func(c *gin.Context) {
			c.JSON(http.StatusOK, basis.GetFeed(c))
		})
		douyinGroup.POST("/user/register", nil)
		douyinGroup.POST("/user/login", nil)
		douyinGroup.GET("/user", nil)
		douyinGroup.POST("/publish/action", nil)
		douyinGroup.GET("/publish/list", nil)
	}
	{
		// interactionproto interfaces
		douyinGroup.POST("/favorite/action", nil)
		douyinGroup.GET("/favorite/list", nil)
		douyinGroup.POST("/comment/action", nil)
		douyinGroup.GET("/comment/list", nil)
	}
	{
		// relationproto interfaces
		douyinGroup.POST("/relationproto/action", nil)
		douyinGroup.GET("/relationproto/follow/list", nil)
		douyinGroup.GET("/relationproto/follower/list", nil)
		douyinGroup.GET("/relationproto/friend/list", nil)
		douyinGroup.GET("/message/chat", nil)
		douyinGroup.POST("/message/action", nil)
	}
	return r
}
