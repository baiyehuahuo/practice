package interaction

import (
	"douyin/configs"
	"douyin/pb"
	"github.com/gin-gonic/gin"
)

func ServeCommentList(c *gin.Context) *pb.DouyinCommentListResponse {
	return &pb.DouyinCommentListResponse{
		StatusCode:  &configs.DefaultInt32,
		StatusMsg:   &configs.DefaultString,
		CommentList: []*pb.Comment{configs.DefaultComment},
	}
}
