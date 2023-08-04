package basis

import (
	"douyin/configs"
	"douyin/pb"
	"github.com/gin-gonic/gin"
)

func ServeUserRegister(c *gin.Context) *pb.DouyinUserRegisterResponse {
	return &pb.DouyinUserRegisterResponse{
		StatusCode: &configs.DefaultInt32,
		StatusMsg:  &configs.DefaultString,
		UserId:     &configs.DefaultInt64,
		Token:      &configs.DefaultString,
	}
}
