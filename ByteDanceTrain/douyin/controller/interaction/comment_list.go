package interaction

import (
	"douyin/constants"
	"douyin/model/dyerror"
	"douyin/pb"
	"github.com/gin-gonic/gin"
	"strconv"
)

// ServeCommentList handle comment list request
// 查看视频的所有评论，按发布时间倒序
// Method is GET
// token, video_id is required
func ServeCommentList(c *gin.Context) (res *pb.DouyinCommentListResponse, err *dyerror.DouyinError) {
	var (
		token   string
		videoID int64
	)
	if err = checkCommentListParams(c, &token, &videoID); err != nil {
		return nil, err
	}

	return &pb.DouyinCommentListResponse{
		StatusCode:  &constants.DefaultInt32,
		StatusMsg:   &constants.DefaultString,
		CommentList: []*pb.Comment{constants.DefaultComment},
	}, nil
}

func checkCommentListParams(c *gin.Context, pToken *string, pVideoID *int64) *dyerror.DouyinError {
	token, videoID := c.Query("token"), c.Query("video_id")
	if token == "" || videoID == "" {
		return dyerror.ParamEmptyError
	}
	id, err := strconv.Atoi(videoID)
	if err != nil {
		return dyerror.ParamInputTypeError
	}
	*pToken = token
	*pVideoID = int64(id)
	return nil
}
