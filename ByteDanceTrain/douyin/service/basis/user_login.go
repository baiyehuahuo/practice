package basis

import (
	"douyin/configs"
	"douyin/pb"
	"github.com/gin-gonic/gin"
	"log"
)

func ServeUserLogin(c *gin.Context) (res *pb.DouyinUserLoginResponse, err error) {
	username, password := c.PostForm("username"), c.PostForm("password")
	log.Printf("username: %v, password: %v", username, password)
	if username == "" || password == "" {
		return nil, configs.ParamEmptyError
	}
	return &pb.DouyinUserLoginResponse{
		StatusCode: &configs.DefaultInt32,
		StatusMsg:  &configs.DefaultString,
		UserId:     &configs.DefaultInt64,
		Token:      &configs.DefaultString,
	}, nil
}
