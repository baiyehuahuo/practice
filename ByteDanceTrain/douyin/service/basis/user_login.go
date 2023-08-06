package basis

import (
	"douyin/configs"
	"douyin/model"
	"douyin/pb"
	"douyin/repository"
	"github.com/gin-gonic/gin"
	"log"
)

// ServeUserLogin handle user login request
// 通过用户名和密码进行登录，登录成功后返回用户 id 和权限 token
// Method is POST
// username, password is required
func ServeUserLogin(c *gin.Context) (res *pb.DouyinUserLoginResponse, err error) {
	username, password := c.PostForm("username"), c.PostForm("password")
	if username == "" || password == "" {
		log.Printf("username: %v, password: %v", username, password)
		return nil, configs.ParamEmptyError
	}
	user := &model.User{
		Name: username,
	}
	repository.FindUser(user)
	if user.Password != password {
		errCode := int32(1)
		errMessage := "username or password is wrong"
		return &pb.DouyinUserLoginResponse{
			StatusCode: &errCode,
			StatusMsg:  &errMessage,
			UserId:     &configs.DefaultInt64,
			Token:      &configs.DefaultString,
		}, nil
	}

	return &pb.DouyinUserLoginResponse{
		StatusCode: &configs.DefaultInt32,
		StatusMsg:  &configs.DefaultString,
		UserId:     &configs.DefaultInt64,
		Token:      &configs.DefaultString,
	}, nil
}
