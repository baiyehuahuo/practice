package main

import (
	"douyin/controller/views/base"
	"github.com/gin-gonic/gin"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Static("/static", "./public")
	douyinGroup := r.Group("/douyin")
	{
		// base interfaces
		douyinGroup.GET("/feed", func(c *gin.Context) {
			feedRes := base.DouyinFeedResponse{
				StatusCode: new(int32),
				StatusMsg:  new(string),
				VideoList:  nil,
				NextTime:   new(int64),
			}
			*feedRes.StatusCode = 0
			*feedRes.StatusMsg = "string"
			feedRes.VideoList = append(feedRes.VideoList, &base.Video{
				Id:            nil,
				Author:        nil,
				PlayUrl:       nil,
				CoverUrl:      nil,
				FavoriteCount: nil,
				CommentCount:  nil,
				IsFavorite:    nil,
				Title:         nil,
			})
			*feedRes.NextTime = time.Now().Unix()
			c.JSON(http.StatusOK, &feedRes)
		})
		douyinGroup.POST("/user/register", nil)
		douyinGroup.POST("/user/login", nil)
		douyinGroup.GET("/user", nil)
		douyinGroup.POST("/publish/action", nil)
		douyinGroup.GET("/publish/list", nil)
	}
	{
		// interaction interfaces
		douyinGroup.POST("/favorite/action", nil)
		douyinGroup.GET("/favorite/list", nil)
		douyinGroup.POST("/comment/action", nil)
		douyinGroup.GET("/comment/list", nil)
	}
	{
		// relation interfaces
		douyinGroup.POST("/relation/action", nil)
		douyinGroup.GET("/relation/follow/list", nil)
		douyinGroup.GET("/relation/follower/list", nil)
		douyinGroup.GET("/relation/friend/list", nil)
		douyinGroup.GET("/message/chat", nil)
		douyinGroup.POST("/message/action", nil)
	}
	return r
}
