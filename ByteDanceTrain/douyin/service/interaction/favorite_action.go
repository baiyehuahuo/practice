package interaction

import (
	"douyin/configs"
	"douyin/pb"
	"github.com/gin-gonic/gin"
	"strconv"
)

// ServeFavoriteAction handle comment list request
// 登录用户对视频的点赞和取消点赞操作
// Method is POST
// token, video_id, action_type is required
func ServeFavoriteAction(c *gin.Context) (res *pb.DouyinFavoriteActionResponse, err error) {
	token, videoID, actionType := c.PostForm("token"), c.PostForm("video_id"), c.PostForm("action_type")
	if token == "" || videoID == "" || actionType == "" {
		return nil, configs.ParamEmptyError
	}
	if _, err = strconv.Atoi(videoID); err != nil {
		return nil, configs.ParamInputTypeError
	}
	if action, _ := strconv.Atoi(actionType); action != 1 && action != 2 {
		return nil, configs.ParamUnknownActionTypeError
	}
	return &pb.DouyinFavoriteActionResponse{
		StatusCode: &configs.DefaultInt32,
		StatusMsg:  &configs.DefaultString,
	}, nil
}
