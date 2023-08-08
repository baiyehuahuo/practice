package router

import (
	"douyin/configs"
	"douyin/controller/basis"
	"douyin/controller/interaction"
	"douyin/controller/relation"
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
		// base interfaces
		douyinGroup.GET("/feed", func(c *gin.Context) {
			res, err := basis.ServeFeed(c)
			if err != nil {
				handleError(c, err)
				return
			}
			c.JSON(http.StatusOK, res)
		})
		douyinGroup.POST("/user/register", func(c *gin.Context) {
			res, err := basis.ServeUserRegister(c)
			if err != nil {
				handleError(c, err)
				return
			}
			c.JSON(http.StatusOK, res)
		})
		douyinGroup.POST("/user/login", func(c *gin.Context) {
			res, err := basis.ServeUserLogin(c)
			if err != nil {
				handleError(c, err)
				return
			}
			c.JSON(http.StatusOK, res)
		})
		douyinGroup.GET("/user", func(c *gin.Context) {
			res, err := basis.ServeUserInfo(c)
			if err != nil {
				handleError(c, err)
				return
			}
			c.JSON(http.StatusOK, res)
		})
		douyinGroup.POST("/publish/action", func(c *gin.Context) {
			res, err := basis.ServePublishAction(c)
			if err != nil {
				handleError(c, err)
				return
			}
			c.JSON(http.StatusOK, res)
		})
		douyinGroup.GET("/publish/list", func(c *gin.Context) {
			res, err := basis.ServePublishList(c)
			if err != nil {
				handleError(c, err)
				return
			}
			c.JSON(http.StatusOK, res)
		})
	}
	{
		// interactionproto interfaces
		douyinGroup.POST("/favorite/action", func(c *gin.Context) {
			res, err := interaction.ServeFavoriteAction(c)
			if err != nil {
				handleError(c, err)
				return
			}
			c.JSON(http.StatusOK, res)
		})
		douyinGroup.GET("/favorite/list", func(c *gin.Context) {
			res, err := interaction.ServeFavoriteList(c)
			if err != nil {
				handleError(c, err)
				return
			}
			c.JSON(http.StatusOK, res)
		})
		douyinGroup.POST("/comment/action", func(c *gin.Context) {
			res, err := interaction.ServeCommentAction(c)
			if err != nil {
				handleError(c, err)
				return
			}
			c.JSON(http.StatusOK, res)
		})
		douyinGroup.GET("/comment/list", func(c *gin.Context) {
			res, err := interaction.ServeCommentList(c)
			if err != nil {
				handleError(c, err)
				return
			}
			c.JSON(http.StatusOK, res)
		})
	}
	{
		// relationproto interfaces
		douyinGroup.POST("/relation/action", func(c *gin.Context) {
			res, err := relation.ServeRelationAction(c)
			if err != nil {
				handleError(c, err)
				return
			}
			c.JSON(http.StatusOK, res)
		})
		douyinGroup.GET("/relation/follow/list", func(c *gin.Context) {
			res, err := relation.ServeRelationFollowList(c)
			if err != nil {
				handleError(c, err)
				return
			}
			c.JSON(http.StatusOK, res)
		})
		douyinGroup.GET("/relation/follower/list", func(c *gin.Context) {
			res, err := relation.ServeRelationFollowerList(c)
			if err != nil {
				handleError(c, err)
				return
			}
			c.JSON(http.StatusOK, res)
		})
		douyinGroup.GET("/relation/friend/list", func(c *gin.Context) {
			res, err := relation.ServeRelationFriendList(c)
			if err != nil {
				handleError(c, err)
				return
			}
			c.JSON(http.StatusOK, res)
		})
		douyinGroup.GET("/message/chat", func(c *gin.Context) {
			res, err := relation.ServeMessageChat(c)
			if err != nil {
				handleError(c, err)
				return
			}
			c.JSON(http.StatusOK, res)
		})
		douyinGroup.POST("/message/action", func(c *gin.Context) {
			res, err := relation.ServeMessageAction(c)
			if err != nil {
				handleError(c, err)
				return
			}
			c.JSON(http.StatusOK, res)
		})
	}
	return r
}

func handleError(c *gin.Context, err error) {
	switch err {
	case configs.ParamEmptyError, configs.ParamInputTypeError, configs.ParamUnknownActionTypeError, configs.ParamInputLengthExceededError:
		c.String(http.StatusBadRequest, err.Error())
	default:
		c.String(http.StatusInternalServerError, err.Error())
	}
}
