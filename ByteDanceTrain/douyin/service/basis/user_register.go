package basis

import (
	"douyin/configs"
	"douyin/pb"
	"github.com/gin-gonic/gin"
	"log"
)

// ServeUserRegister handle user register request
// 新用户注册时提供用户名，密码即可，用户名需要保证唯一。创建成功后返回用户 id 和权限token
// Method is POST
// username, password is required
func ServeUserRegister(c *gin.Context) (res *pb.DouyinUserRegisterResponse, err error) {
	username, password := c.PostForm("username"), c.PostForm("password")
	if username == "" || password == "" {
		log.Printf("username: %v, password: %v", username, password)
		return nil, configs.ParamEmptyError
	}
	if len(username) > 32 || len(password) > 32 {
		log.Printf("username: %v, password: %v", username, password)
		return nil, configs.ParamInputLengthExceededError
	}

	return &pb.DouyinUserRegisterResponse{
		StatusCode: &configs.DefaultInt32,
		StatusMsg:  &configs.DefaultString,
		UserId:     &configs.DefaultInt64,
		Token:      &configs.DefaultString,
	}, nil
}
