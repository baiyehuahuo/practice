package interaction

import (
	"douyin/constants"
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
	var (
		token, commentText string
		videoID, commentID int64
		actionType         int
	)
	if err = checkCommentActionParams(c, &token, &videoID, &actionType, &commentText, &commentID); err != nil {
		return nil, err
	}

	return &pb.DouyinCommentActionResponse{
		StatusCode: &constants.DefaultInt32,
		StatusMsg:  &constants.DefaultString,
		Comment:    constants.DefaultComment,
	}, nil
}

func checkCommentActionParams(c *gin.Context, pToken *string, pVideoID *int64, pAction *int, pCommentText *string, pCommentID *int64) error {
	token, videoID, actionType := c.PostForm("token"), c.PostForm("video_id"), c.PostForm("action_type")
	commentText, commentID := c.PostForm("comment_text"), c.PostForm("comment_id")
	if token == "" || videoID == "" || actionType == "" {
		log.Printf("token: %v, videoID: %v, actionType: %v", token, videoID, actionType)
		return constants.ParamEmptyError
	}
	action, _ := strconv.Atoi(actionType)
	if action != 1 && action != 2 {
		return constants.ParamUnknownActionTypeError
	}
	if action == 1 && commentText == "" || action == 2 && commentID == "" {
		return constants.ParamEmptyError
	}
	vid, err1 := strconv.Atoi(videoID)
	cid, err2 := strconv.Atoi(commentID)
	if err1 != nil || err2 != nil {
		return constants.ParamInputTypeError
	}

	*pToken = token
	*pVideoID = int64(vid)
	*pAction = action
	*pCommentText = commentText
	*pCommentID = int64(cid)
	return nil
}
