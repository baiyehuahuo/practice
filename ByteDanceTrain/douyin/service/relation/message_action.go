package relation

import (
	"douyin/configs"
	"douyin/pb"
	"github.com/gin-gonic/gin"
)

func ServeMessageAction(c *gin.Context) *pb.DouyinMessageActionResponse {
	return &pb.DouyinMessageActionResponse{
		StatusCode: &configs.DefaultInt32,
		StatusMsg:  &configs.DefaultString,
	}
}
