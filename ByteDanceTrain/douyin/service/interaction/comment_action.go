package interaction

import (
	"douyin/configs"
	"douyin/pb"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

// ServeCommentAction handle comment action request
// 登录用户对视频进行评论
// Method is POST
// token, video_id, action_type is required
// comment_text comment_id is optional
func ServeCommentAction(c *gin.Context) (res *pb.DouyinCommentActionResponse, err error) {
	token, videoID, actionType := c.PostForm("token"), c.PostForm("video_id"), c.PostForm("action_type")
	commentText, commentID := c.PostForm("comment_text"), c.PostForm("comment_id")
	if token == "" || videoID == "" || actionType == "" {
		log.Printf("token: %v, videoID: %v, actionType: %v", token, videoID, actionType)
		return nil, configs.ParamEmptyError
	}
	action, _ := strconv.Atoi(actionType)
	if action != 1 && action != 2 {
		return nil, configs.ParamUnknownActionTypeError
	}
	if action == 1 && commentText == "" || action == 2 && commentID == "" {
		return nil, configs.ParamEmptyError
	}
	if _, err = strconv.Atoi(videoID); err != nil {
		return nil, configs.ParamInputTypeError
	}

	return &pb.DouyinCommentActionResponse{
		StatusCode: &configs.DefaultInt32,
		StatusMsg:  &configs.DefaultString,
		Comment:    configs.DefaultComment,
	}, nil
}
