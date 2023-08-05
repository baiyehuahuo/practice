package router

import (
	"douyin/configs"
	"douyin/service/basis"
	"douyin/service/interaction"
	"douyin/service/relation"
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
			res, err := basis.ServeFeed(c)
			if err != nil {
				switch err {
				case configs.LatestTimeParamError:
					c.String(http.StatusBadRequest, err.Error())
				default:
					c.String(http.StatusInternalServerError, err.Error())
				}
				return
			}
			c.JSON(http.StatusOK, res)
		})
		douyinGroup.POST("/user/register", func(c *gin.Context) {
			res, err := basis.ServeUserRegister(c)
			if err != nil {
				switch err {
				case configs.ParamEmptyError:
					c.String(http.StatusBadRequest, err.Error())
				default:
					c.String(http.StatusInternalServerError, err.Error())
				}
				return
			}
			c.JSON(http.StatusOK, res)
		})
		douyinGroup.POST("/user/login", func(c *gin.Context) {
			res, err := basis.ServeUserLogin(c)
			if err != nil {
				switch err {
				case configs.ParamEmptyError:
					c.String(http.StatusBadRequest, err.Error())
				default:
					c.String(http.StatusInternalServerError, err.Error())
				}
				return
			}
			c.JSON(http.StatusOK, res)
		})
		douyinGroup.GET("/user", func(c *gin.Context) {
			c.JSON(http.StatusOK, basis.ServeUserInfo(c))
		})
		douyinGroup.POST("/publish/action", func(c *gin.Context) {
			c.JSON(http.StatusOK, basis.ServePublishAction(c))
		})
		douyinGroup.GET("/publish/list", func(c *gin.Context) {
			c.JSON(http.StatusOK, basis.ServePublishList(c))
		})
	}
	{
		// interactionproto interfaces
		douyinGroup.POST("/favorite/action", func(c *gin.Context) {
			c.JSON(http.StatusOK, interaction.ServeFavoriteAction(c))
		})
		douyinGroup.GET("/favorite/list", func(c *gin.Context) {
			c.JSON(http.StatusOK, interaction.ServeFavoriteList(c))
		})
		douyinGroup.POST("/comment/action", func(c *gin.Context) {
			c.JSON(http.StatusOK, interaction.ServeCommentAction(c))
		})
		douyinGroup.GET("/comment/list", func(c *gin.Context) {
			c.JSON(http.StatusOK, interaction.ServeCommentList(c))
		})
	}
	{
		// relationproto interfaces
		douyinGroup.POST("/relation/action", func(c *gin.Context) {
			c.JSON(http.StatusOK, relation.ServeRelationAction(c))
		})
		douyinGroup.GET("/relation/follow/list", func(c *gin.Context) {
			c.JSON(http.StatusOK, relation.ServeRelationFollowList(c))
		})
		douyinGroup.GET("/relation/follower/list", func(c *gin.Context) {
			c.JSON(http.StatusOK, relation.ServeRelationFollowerList(c))
		})
		douyinGroup.GET("/relation/friend/list", func(c *gin.Context) {
			c.JSON(http.StatusOK, relation.ServeRelationFriendList(c))
		})
		douyinGroup.GET("/message/chat", func(c *gin.Context) {
			c.JSON(http.StatusOK, relation.ServeMessageChat(c))
		})
		douyinGroup.POST("/message/action", func(c *gin.Context) {
			c.JSON(http.StatusOK, relation.ServeMessageAction(c))
		})
	}
	return r
}
