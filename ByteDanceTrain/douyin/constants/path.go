package constants

import (
	"fmt"
)

const (
	ProjectGroup = "/douyin"

	// basis
	RouteFeed          = "/feed"
	RouteUserRegister  = "/user/register"
	RouteUserLogin     = "/user/login"
	RouteUserInfo      = "/user"
	RoutePublishAction = "/publish/action"
	RoutePublishList   = "/publish/list"

	// interaction
	RouteFavoriteAction = "/favorite/action"
	RouteFavoriteList   = "/favorite/list"
	RouteCommentAction  = "/comment/action"
	RouteCommentList    = "/comment/list"

	// relation
	RouteRelationAction       = "/relation/action"
	RouteRelationFollowList   = "/relation/follow/list"
	RouteRelationFollowerList = "/relation/follower/list"
	RouteRelationFriendList   = "/relation/friend/list"
	RouteMessageChat          = "/message/chat"
	RouteMessageAction        = "/message/action"
)

const (
	UploadFileDir = "uploadfiles"

	Assets              = "assets"
	UserSQLFileName     = "user.sql"
	VideoSQLFileName    = "video.sql"
	FavoriteSQLFileName = "favorite_event.sql"
	CommentSQLFileName  = "comment_event.sql"
)

const (
	// Database
	userName = "root"
	password = "rootpwd"
	ip       = "127.0.0.1"
	port     = 3306
	dbName   = "douyin"
)

var (
	DatabasePath = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", userName, password, ip, port, dbName)
)
