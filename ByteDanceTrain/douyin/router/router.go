package router

import (
	"douyin/configs"
	"douyin/constants"
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
	douyinGroup := r.Group(constants.ProjectGroup)
	{
		// base interfaces
		douyinGroup.GET(constants.RouteFeed, func(c *gin.Context) {
			res, err := basis.ServeFeed(c)
			if err != nil {
				handleError(c, err)
				return
			}
			c.JSON(http.StatusOK, res)
		})
		douyinGroup.POST(constants.RouteUserRegister, func(c *gin.Context) {
			res, err := basis.ServeUserRegister(c)
			if err != nil {
				handleError(c, err)
				return
			}
			c.JSON(http.StatusOK, res)
		})
		douyinGroup.POST(constants.RouteUserLogin, func(c *gin.Context) {
			res, err := basis.ServeUserLogin(c)
			if err != nil {
				handleError(c, err)
				return
			}
			c.JSON(http.StatusOK, res)
		})
		douyinGroup.GET(constants.RouteUserInfo, func(c *gin.Context) {
			res, err := basis.ServeUserInfo(c)
			if err != nil {
				handleError(c, err)
				return
			}
			c.JSON(http.StatusOK, res)
		})
		douyinGroup.POST(constants.RoutePublishAction, func(c *gin.Context) {
			res, err := basis.ServePublishAction(c)
			if err != nil {
				handleError(c, err)
				return
			}
			c.JSON(http.StatusOK, res)
		})
		douyinGroup.GET(constants.RoutePublishList, func(c *gin.Context) {
			res, err := basis.ServePublishList(c)
			if err != nil {
				handleError(c, err)
				return
			}
			c.JSON(http.StatusOK, res)
		})
	}
	{
		// interaction interfaces
		douyinGroup.POST(constants.RouteFavoriteAction, func(c *gin.Context) {
			res, err := interaction.ServeFavoriteAction(c)
			if err != nil {
				handleError(c, err)
				return
			}
			c.JSON(http.StatusOK, res)
		})
		douyinGroup.GET(constants.RouteFavoriteList, func(c *gin.Context) {
			res, err := interaction.ServeFavoriteList(c)
			if err != nil {
				handleError(c, err)
				return
			}
			c.JSON(http.StatusOK, res)
		})
		douyinGroup.POST(constants.RouteCommentAction, func(c *gin.Context) {
			res, err := interaction.ServeCommentAction(c)
			if err != nil {
				handleError(c, err)
				return
			}
			c.JSON(http.StatusOK, res)
		})
		douyinGroup.GET(constants.RouteCommentList, func(c *gin.Context) {
			res, err := interaction.ServeCommentList(c)
			if err != nil {
				handleError(c, err)
				return
			}
			c.JSON(http.StatusOK, res)
		})
	}
	{
		// relation interfaces
		douyinGroup.POST(constants.RouteRelationAction, func(c *gin.Context) {
			res, err := relation.ServeRelationAction(c)
			if err != nil {
				handleError(c, err)
				return
			}
			c.JSON(http.StatusOK, res)
		})
		douyinGroup.GET(constants.RouteRelationFollowList, func(c *gin.Context) {
			res, err := relation.ServeRelationFollowList(c)
			if err != nil {
				handleError(c, err)
				return
			}
			c.JSON(http.StatusOK, res)
		})
		douyinGroup.GET(constants.RouteRelationFollowerList, func(c *gin.Context) {
			res, err := relation.ServeRelationFollowerList(c)
			if err != nil {
				handleError(c, err)
				return
			}
			c.JSON(http.StatusOK, res)
		})
		douyinGroup.GET(constants.RouteRelationFriendList, func(c *gin.Context) {
			res, err := relation.ServeRelationFriendList(c)
			if err != nil {
				handleError(c, err)
				return
			}
			c.JSON(http.StatusOK, res)
		})
		douyinGroup.GET(constants.RouteMessageChat, func(c *gin.Context) {
			res, err := relation.ServeMessageChat(c)
			if err != nil {
				handleError(c, err)
				return
			}
			c.JSON(http.StatusOK, res)
		})
		douyinGroup.POST(constants.RouteMessageAction, func(c *gin.Context) {
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
