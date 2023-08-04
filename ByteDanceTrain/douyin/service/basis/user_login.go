package basis

import (
	"douyin/configs"
	"douyin/pb"
	"github.com/gin-gonic/gin"
)

func ServeUserLogin(c *gin.Context) *pb.DouyinUserLoginResponse {
	return &pb.DouyinUserLoginResponse{
		StatusCode: &configs.DefaultInt32,
		StatusMsg:  &configs.DefaultString,
		UserId:     &configs.DefaultInt64,
		Token:      &configs.DefaultString,
	}
}
