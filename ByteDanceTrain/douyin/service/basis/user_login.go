package basis

import (
	"douyin/common"
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
	var (
		username, password string
	)
	if err = checkUserLoginParams(c, &username, &password); err != nil {
		return nil, err
	}
	user := &model.User{
		Name: username,
	}
	repository.QueryUser(user)
	if user.Password != password {
		return nil, configs.AuthUsernameOrPasswordFail
	}

	token := common.GenerateToken()
	common.SetToken(token, user.ID)

	return &pb.DouyinUserLoginResponse{
		StatusCode: &configs.DefaultInt32,
		StatusMsg:  &configs.DefaultString,
		UserId:     &user.ID,
		Token:      &token,
	}, nil
}

func checkUserLoginParams(c *gin.Context, pUsername, pPassword *string) error {
	username, password := c.PostForm("username"), c.PostForm("password")
	if username == "" || password == "" {
		log.Printf("username: %v, password: %v", username, password)
		return configs.ParamEmptyError
	}
	*pUsername = username
	*pPassword = password
	return nil
}
