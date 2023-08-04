package interaction

import (
	"douyin/configs"
	"douyin/pb"
	"github.com/gin-gonic/gin"
)

func ServeCommentAction(c *gin.Context) *pb.DouyinCommentActionResponse {
	return &pb.DouyinCommentActionResponse{
		StatusCode: &configs.DefaultInt32,
		StatusMsg:  &configs.DefaultString,
		Comment:    configs.DefaultComment,
	}
}
