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
	token, videoID := c.Query("token"), c.Query("video_id")
	if token == "" || videoID == "" {
		return nil, configs.ParamEmptyError
	}
	if _, err = strconv.Atoi(videoID); err != nil {
		return nil, configs.ParamInputTypeError
	}
	return &pb.DouyinCommentListResponse{
		StatusCode:  &configs.DefaultInt32,
		StatusMsg:   &configs.DefaultString,
		CommentList: []*pb.Comment{configs.DefaultComment},
	}, nil
}
