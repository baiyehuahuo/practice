package interaction

import (
	"douyin/configs"
	"douyin/pb"
	"github.com/gin-gonic/gin"
	"strconv"
)

// ServeCommentList handle comment list request
// 查看视频的所有评论，按发布时间倒序
// Method is GET
// token, video_id is required
func ServeCommentList(c *gin.Context) (res *pb.DouyinCommentListResponse, err error) {
	var (
		token   string
		videoID int64
	)
	if err = checkCommentListParams(c, &token, &videoID); err != nil {
		return nil, err
	}

	return &pb.DouyinCommentListResponse{
		StatusCode:  &configs.DefaultInt32,
		StatusMsg:   &configs.DefaultString,
		CommentList: []*pb.Comment{configs.DefaultComment},
	}, nil
}

func checkCommentListParams(c *gin.Context, pToken *string, pVideoID *int64) error {
	token, videoID := c.Query("token"), c.Query("video_id")
	if token == "" || videoID == "" {
		return configs.ParamEmptyError
	}
	id, err := strconv.Atoi(videoID)
	if err != nil {
		return configs.ParamInputTypeError
	}
	*pToken = token
	*pVideoID = int64(id)
	return nil
}
