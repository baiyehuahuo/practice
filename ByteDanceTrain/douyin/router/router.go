package router

import (
	"douyin/constants"
	"douyin/controller/basis"
	"douyin/controller/interaction"
	"douyin/controller/relation"
	"douyin/model/dyerror"
	"douyin/model/empty_response"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	_ "net/http/pprof"
)

// SetupRouter get a configured gin engine
func SetupRouter(r *gin.Engine) {
	//r.Static("static", "./assets")
	//r.Static("uploadfiles", "./uploadfiles")
	r.StaticFS("douyin/uploadfiles", http.Dir("./uploadfiles"))
	r.StaticFS("uploadfiles", http.Dir("./uploadfiles"))
	douyinGroup := r.Group(constants.ProjectGroup)
	{
		// base interfaces
		douyinGroup.GET(constants.RouteFeed, func(c *gin.Context) {
			res, err := basis.ServeFeed(c)
			if err != nil {
				res = empty_response.Feed()
				dyerr := err.(*dyerror.DouyinError)
				*res.StatusCode = dyerr.ErrCode
				*res.StatusMsg = dyerr.ErrMessage
				c.JSON(http.StatusOK, res)
				return
			}
			c.JSON(http.StatusOK, res)
		})
		douyinGroup.POST(constants.RouteUserRegister, func(c *gin.Context) {
			res, err := basis.ServeUserRegister(c)
			if err != nil {
				res = empty_response.UserRegister()
				dyerr := err.(*dyerror.DouyinError)
				*res.StatusCode = dyerr.ErrCode
				*res.StatusMsg = dyerr.ErrMessage
				c.JSON(http.StatusOK, res)
				return
			}
			c.JSON(http.StatusOK, res)
		})
		douyinGroup.POST(constants.RouteUserLogin, func(c *gin.Context) {
			res, err := basis.ServeUserLogin(c)
			if err != nil {
				res = empty_response.UserLogin()
				dyerr := err.(*dyerror.DouyinError)
				*res.StatusCode = dyerr.ErrCode
				*res.StatusMsg = dyerr.ErrMessage
				c.JSON(http.StatusOK, res)
				return
			}
			c.JSON(http.StatusOK, res)
		})
		douyinGroup.GET(constants.RouteUserInfo, func(c *gin.Context) {
			res, err := basis.ServeUserInfo(c)
			if err != nil {
				res = empty_response.User()
				dyerr := err.(*dyerror.DouyinError)
				*res.StatusCode = dyerr.ErrCode
				*res.StatusMsg = dyerr.ErrMessage
				c.JSON(http.StatusOK, res)
				return
			}
			c.JSON(http.StatusOK, res)
		})
		douyinGroup.POST(constants.RoutePublishAction, func(c *gin.Context) {
			res, err := basis.ServePublishAction(c)
			if err != nil {
				res = empty_response.PublishAction()
				dyerr, ok := err.(*dyerror.DouyinError)
				if !ok || dyerr == nil {
					panic(fmt.Sprintf("err is not dyerror: %T %s", err, err))
				}
				*res.StatusCode = dyerr.ErrCode
				*res.StatusMsg = dyerr.ErrMessage
				c.JSON(http.StatusOK, res)
				return
			}
			c.JSON(http.StatusOK, res)
		})
		douyinGroup.GET(constants.RoutePublishList, func(c *gin.Context) {
			res, err := basis.ServePublishList(c)
			if err != nil {
				res = empty_response.PublishList()
				dyerr := err.(*dyerror.DouyinError)
				*res.StatusCode = dyerr.ErrCode
				*res.StatusMsg = dyerr.ErrMessage
				c.JSON(http.StatusOK, res)
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
				res = empty_response.FavoriteAction()
				dyerr := err.(*dyerror.DouyinError)
				*res.StatusCode = dyerr.ErrCode
				*res.StatusMsg = dyerr.ErrMessage
				c.JSON(http.StatusOK, res)
				return
			}
			c.JSON(http.StatusOK, res)
		})
		douyinGroup.GET(constants.RouteFavoriteList, func(c *gin.Context) {
			res, err := interaction.ServeFavoriteList(c)
			if err != nil {
				res = empty_response.FavoriteList()
				dyerr := err.(*dyerror.DouyinError)
				*res.StatusCode = dyerr.ErrCode
				*res.StatusMsg = dyerr.ErrMessage
				c.JSON(http.StatusOK, res)
				return
			}
			c.JSON(http.StatusOK, res)
		})
		douyinGroup.POST(constants.RouteCommentAction, func(c *gin.Context) {
			res, err := interaction.ServeCommentAction(c)
			if err != nil {
				res = empty_response.CommentAction()
				dyerr := err.(*dyerror.DouyinError)
				*res.StatusCode = dyerr.ErrCode
				*res.StatusMsg = dyerr.ErrMessage
				c.JSON(http.StatusOK, res)
				return
			}
			c.JSON(http.StatusOK, res)
		})
		douyinGroup.GET(constants.RouteCommentList, func(c *gin.Context) {
			res, err := interaction.ServeCommentList(c)
			if err != nil {
				res = empty_response.CommentList()
				dyerr := err.(*dyerror.DouyinError)
				*res.StatusCode = dyerr.ErrCode
				*res.StatusMsg = dyerr.ErrMessage
				c.JSON(http.StatusOK, res)
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
				res = empty_response.RelationAction()
				dyerr := err.(*dyerror.DouyinError)
				*res.StatusCode = dyerr.ErrCode
				*res.StatusMsg = dyerr.ErrMessage
				c.JSON(http.StatusOK, res)
				return
			}
			c.JSON(http.StatusOK, res)
		})
		douyinGroup.GET(constants.RouteRelationFollowList, func(c *gin.Context) {
			res, err := relation.ServeRelationFollowList(c)
			if err != nil {
				res = empty_response.RelationFollowList()
				dyerr := err.(*dyerror.DouyinError)
				*res.StatusCode = dyerr.ErrCode
				*res.StatusMsg = dyerr.ErrMessage
				c.JSON(http.StatusOK, res)
				return
			}
			c.JSON(http.StatusOK, res)
		})
		douyinGroup.GET(constants.RouteRelationFollowerList, func(c *gin.Context) {
			res, err := relation.ServeRelationFollowerList(c)
			if err != nil {
				res = empty_response.RelationFollowerList()
				dyerr := err.(*dyerror.DouyinError)
				*res.StatusCode = dyerr.ErrCode
				*res.StatusMsg = dyerr.ErrMessage
				c.JSON(http.StatusOK, res)
				return
			}
			c.JSON(http.StatusOK, res)
		})
		douyinGroup.GET(constants.RouteRelationFriendList, func(c *gin.Context) {
			res, err := relation.ServeRelationFriendList(c)
			if err != nil {
				res = empty_response.RelationFriendList()
				dyerr := err.(*dyerror.DouyinError)
				*res.StatusCode = dyerr.ErrCode
				*res.StatusMsg = dyerr.ErrMessage
				c.JSON(http.StatusOK, res)
				return
			}
			c.JSON(http.StatusOK, res)
		})
		douyinGroup.GET(constants.RouteMessageChat, func(c *gin.Context) {
			res, err := relation.ServeMessageChat(c)
			if err != nil {
				res = empty_response.MessageChat()
				dyerr := err.(*dyerror.DouyinError)
				*res.StatusCode = dyerr.ErrCode
				*res.StatusMsg = dyerr.ErrMessage
				c.JSON(http.StatusOK, res)
				return
			}
			c.JSON(http.StatusOK, res)
		})
		douyinGroup.POST(constants.RouteMessageAction, func(c *gin.Context) {
			res, err := relation.ServeMessageAction(c)
			if err != nil {
				res = empty_response.MessageAction()
				dyerr := err.(*dyerror.DouyinError)
				*res.StatusCode = dyerr.ErrCode
				*res.StatusMsg = dyerr.ErrMessage
				c.JSON(http.StatusOK, res)
				return
			}
			c.JSON(http.StatusOK, res)
		})
	}

	{
		// base interfaces
		douyinGroup.GET(constants.RouteFeed+"/", func(c *gin.Context) {
			res, err := basis.ServeFeed(c)
			if err != nil {
				res = empty_response.Feed()
				dyerr := err.(*dyerror.DouyinError)
				*res.StatusCode = dyerr.ErrCode
				*res.StatusMsg = dyerr.ErrMessage
				c.JSON(http.StatusOK, res)
				return
			}
			c.JSON(http.StatusOK, res)
		})
		douyinGroup.POST(constants.RouteUserRegister+"/", func(c *gin.Context) {
			res, err := basis.ServeUserRegister(c)
			if err != nil {
				res = empty_response.UserRegister()
				dyerr := err.(*dyerror.DouyinError)
				*res.StatusCode = dyerr.ErrCode
				*res.StatusMsg = dyerr.ErrMessage
				c.JSON(http.StatusOK, res)
				return
			}
			c.JSON(http.StatusOK, res)
		})
		douyinGroup.POST(constants.RouteUserLogin+"/", func(c *gin.Context) {
			res, err := basis.ServeUserLogin(c)
			if err != nil {
				res = empty_response.UserLogin()
				dyerr := err.(*dyerror.DouyinError)
				*res.StatusCode = dyerr.ErrCode
				*res.StatusMsg = dyerr.ErrMessage
				c.JSON(http.StatusOK, res)
				return
			}
			c.JSON(http.StatusOK, res)
		})
		douyinGroup.GET(constants.RouteUserInfo+"/", func(c *gin.Context) {
			res, err := basis.ServeUserInfo(c)
			if err != nil {
				res = empty_response.User()
				dyerr := err.(*dyerror.DouyinError)
				*res.StatusCode = dyerr.ErrCode
				*res.StatusMsg = dyerr.ErrMessage
				c.JSON(http.StatusOK, res)
				return
			}
			c.JSON(http.StatusOK, res)
		})
		douyinGroup.POST(constants.RoutePublishAction+"/", func(c *gin.Context) {
			res, err := basis.ServePublishAction(c)
			if err != nil {
				res = empty_response.PublishAction()
				dyerr := err.(*dyerror.DouyinError)
				*res.StatusCode = dyerr.ErrCode
				*res.StatusMsg = dyerr.ErrMessage
				c.JSON(http.StatusOK, res)
				return
			}
			c.JSON(http.StatusOK, res)
		})
		douyinGroup.GET(constants.RoutePublishList+"/", func(c *gin.Context) {
			res, err := basis.ServePublishList(c)
			if err != nil {
				res = empty_response.PublishList()
				dyerr := err.(*dyerror.DouyinError)
				*res.StatusCode = dyerr.ErrCode
				*res.StatusMsg = dyerr.ErrMessage
				c.JSON(http.StatusOK, res)
				return
			}
			c.JSON(http.StatusOK, res)
		})
	}
	{
		// interaction interfaces
		douyinGroup.POST(constants.RouteFavoriteAction+"/", func(c *gin.Context) {
			res, err := interaction.ServeFavoriteAction(c)
			if err != nil {
				res = empty_response.FavoriteAction()
				dyerr := err.(*dyerror.DouyinError)
				*res.StatusCode = dyerr.ErrCode
				*res.StatusMsg = dyerr.ErrMessage
				c.JSON(http.StatusOK, res)
				return
			}
			c.JSON(http.StatusOK, res)
		})
		douyinGroup.GET(constants.RouteFavoriteList+"/", func(c *gin.Context) {
			res, err := interaction.ServeFavoriteList(c)
			if err != nil {
				res = empty_response.FavoriteList()
				dyerr := err.(*dyerror.DouyinError)
				*res.StatusCode = dyerr.ErrCode
				*res.StatusMsg = dyerr.ErrMessage
				c.JSON(http.StatusOK, res)
				return
			}
			c.JSON(http.StatusOK, res)
		})
		douyinGroup.POST(constants.RouteCommentAction+"/", func(c *gin.Context) {
			res, err := interaction.ServeCommentAction(c)
			if err != nil {
				res = empty_response.CommentAction()
				dyerr := err.(*dyerror.DouyinError)
				*res.StatusCode = dyerr.ErrCode
				*res.StatusMsg = dyerr.ErrMessage
				c.JSON(http.StatusOK, res)
				return
			}
			c.JSON(http.StatusOK, res)
		})
		douyinGroup.GET(constants.RouteCommentList+"/", func(c *gin.Context) {
			res, err := interaction.ServeCommentList(c)
			if err != nil {
				res = empty_response.CommentList()
				dyerr := err.(*dyerror.DouyinError)
				*res.StatusCode = dyerr.ErrCode
				*res.StatusMsg = dyerr.ErrMessage
				c.JSON(http.StatusOK, res)
				return
			}
			c.JSON(http.StatusOK, res)
		})
	}
	{
		// relation interfaces
		douyinGroup.POST(constants.RouteRelationAction+"/", func(c *gin.Context) {
			res, err := relation.ServeRelationAction(c)
			if err != nil {
				res = empty_response.RelationAction()
				dyerr := err.(*dyerror.DouyinError)
				*res.StatusCode = dyerr.ErrCode
				*res.StatusMsg = dyerr.ErrMessage
				c.JSON(http.StatusOK, res)
				return
			}
			c.JSON(http.StatusOK, res)
		})
		douyinGroup.GET(constants.RouteRelationFollowList+"/", func(c *gin.Context) {
			res, err := relation.ServeRelationFollowList(c)
			if err != nil {
				res = empty_response.RelationFollowList()
				dyerr := err.(*dyerror.DouyinError)
				*res.StatusCode = dyerr.ErrCode
				*res.StatusMsg = dyerr.ErrMessage
				c.JSON(http.StatusOK, res)
				return
			}
			c.JSON(http.StatusOK, res)
		})
		douyinGroup.GET(constants.RouteRelationFollowerList+"/", func(c *gin.Context) {
			res, err := relation.ServeRelationFollowerList(c)
			if err != nil {
				res = empty_response.RelationFollowerList()
				dyerr := err.(*dyerror.DouyinError)
				*res.StatusCode = dyerr.ErrCode
				*res.StatusMsg = dyerr.ErrMessage
				c.JSON(http.StatusOK, res)
				return
			}
			c.JSON(http.StatusOK, res)
		})
		douyinGroup.GET(constants.RouteRelationFriendList+"/", func(c *gin.Context) {
			res, err := relation.ServeRelationFriendList(c)
			if err != nil {
				res = empty_response.RelationFriendList()
				dyerr := err.(*dyerror.DouyinError)
				*res.StatusCode = dyerr.ErrCode
				*res.StatusMsg = dyerr.ErrMessage
				c.JSON(http.StatusOK, res)
				return
			}
			c.JSON(http.StatusOK, res)
		})
		douyinGroup.GET(constants.RouteMessageChat+"/", func(c *gin.Context) {
			res, err := relation.ServeMessageChat(c)
			if err != nil {
				res = empty_response.MessageChat()
				dyerr := err.(*dyerror.DouyinError)
				*res.StatusCode = dyerr.ErrCode
				*res.StatusMsg = dyerr.ErrMessage
				c.JSON(http.StatusOK, res)
				return
			}
			c.JSON(http.StatusOK, res)
		})
		douyinGroup.POST(constants.RouteMessageAction+"/", func(c *gin.Context) {
			res, err := relation.ServeMessageAction(c)
			if err != nil {
				res = empty_response.MessageAction()
				dyerr := err.(*dyerror.DouyinError)
				*res.StatusCode = dyerr.ErrCode
				*res.StatusMsg = dyerr.ErrMessage
				c.JSON(http.StatusOK, res)
				return
			}
			c.JSON(http.StatusOK, res)
		})
	}
}

//func handleError(c *gin.Context, err *dyerror.DouyinError) {
//	switch err {
//	case dyerror.ParamEmptyError, dyerror.ParamInputTypeError, dyerror.ParamUnknownActionTypeError, dyerror.ParamInputLengthExceededError:
//		c.String(http.StatusBadRequest, err.Error())
//	default:
//		c.String(http.StatusInternalServerError, err.Error())
//	}
//}
