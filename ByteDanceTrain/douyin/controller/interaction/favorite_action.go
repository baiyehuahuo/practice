package interaction

import (
	"douyin/constants"
	"douyin/model/dyerror"
	"douyin/model/entity"
	"douyin/pb"
	"douyin/service/FavoriteService"
	"douyin/service/TokenService"
	"douyin/service/VideoService"
	"github.com/gin-gonic/gin"
)

// ServeFavoriteAction handle comment list request
// 登录用户对视频的点赞和取消点赞操作
// 点击点赞按钮调用接口douyin/favorite/action/
// Method is POST
// token, video_id, action_type is required
func ServeFavoriteAction(c *gin.Context) (res *pb.DouyinFavoriteActionResponse, dyerr *dyerror.DouyinError) {
	var (
		token      string
		videoID    int64
		actionType int
	)
	if dyerr = checkFavoriteActionParams(c, &token, &videoID, &actionType); dyerr != nil {
		return nil, dyerr
	}
	userID, dyerr := TokenService.GetUserIDFromToken(token)
	if dyerr != nil {
		return nil, dyerr
	}
	video := VideoService.QueryVideoByVideoID(videoID)
	switch actionType {
	case 1:
		if dyerr = FavoriteService.CreateFavoriteEvent(&entity.Favorite{UserID: userID, VideoID: videoID, AuthorID: video.AuthorID}); dyerr != nil {
			return nil, dyerr
		}
	case 2:
		if dyerr = FavoriteService.DeleteFavoriteEvent(&entity.Favorite{UserID: userID, VideoID: videoID}); dyerr != nil {
			return nil, dyerr
		}
	}
	return &pb.DouyinFavoriteActionResponse{
		StatusCode: &constants.DefaultInt32,
		StatusMsg:  &constants.DefaultString,
	}, nil
}

func checkFavoriteActionParams(c *gin.Context, pToken *string, pVideoID *int64, pActionType *int) *dyerror.DouyinError {
	body := struct {
		Token      string `form:"token" json:"token" binding:"required"`
		VideoID    int64  `form:"video_id" json:"video_id" binding:"required"`
		ActionType int    `form:"action_type" json:"action_type" binding:"required,oneof=1 2"`
	}{}
	if err := c.ShouldBind(&body); err != nil {
		return dyerror.HandleBindError(err)
	}

	*pToken = body.Token
	*pVideoID = body.VideoID
	*pActionType = body.ActionType
	return nil
}
