package interaction

import (
	"douyin/constants"
	"douyin/model/dyerror"
	"douyin/pb"
	"github.com/gin-gonic/gin"
	"strconv"
)

// ServeFavoriteAction handle comment list request
// 登录用户对视频的点赞和取消点赞操作
// Method is POST
// token, video_id, action_type is required
func ServeFavoriteAction(c *gin.Context) (res *pb.DouyinFavoriteActionResponse, err *dyerror.DouyinError) {
	var (
		token      string
		videoID    int64
		actionType int
	)
	if err = checkFavoriteActionParams(c, &token, &videoID, &actionType); err != nil {
		return nil, err
	}
	return &pb.DouyinFavoriteActionResponse{
		StatusCode: &constants.DefaultInt32,
		StatusMsg:  &constants.DefaultString,
	}, nil
}

func checkFavoriteActionParams(c *gin.Context, pToken *string, pVideoID *int64, pActionType *int) *dyerror.DouyinError {
	token, videoID, actionType := c.PostForm("token"), c.PostForm("video_id"), c.PostForm("action_type")
	if token == "" || videoID == "" || actionType == "" {
		return dyerror.ParamEmptyError
	}
	id, err1 := strconv.Atoi(videoID)
	action, err2 := strconv.Atoi(actionType)
	if err1 != nil || err2 != nil {
		return dyerror.ParamInputTypeError
	}
	if action != 1 && action != 2 {
		return dyerror.ParamUnknownActionTypeError
	}
	*pToken = token
	*pVideoID = int64(id)
	*pActionType = action
	return nil
}
