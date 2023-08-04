package basis

import (
	"douyin/configs"
	"douyin/pb"
	"github.com/gin-gonic/gin"
)

func ServeUserInfo(c *gin.Context) *pb.DouyinUserResponse {
	return &pb.DouyinUserResponse{
		StatusCode: &configs.DefaultInt32,
		StatusMsg:  &configs.DefaultString,
		User:       configs.DefaultUser,
	}
}
