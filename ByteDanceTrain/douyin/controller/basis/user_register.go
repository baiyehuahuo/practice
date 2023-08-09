package basis

import (
	"douyin/constants"
	"douyin/model/dyerror"
	"douyin/model/entity"
	"douyin/pb"
	"douyin/service/TokenService"
	"douyin/service/UserService"
	"github.com/gin-gonic/gin"
	"log"
)

// ServeUserRegister handle user register request
// 新用户注册时提供用户名，密码即可，用户名需要保证唯一。创建成功后返回用户 id 和权限token
// Method is POST
// username, password is required
func ServeUserRegister(c *gin.Context) (res *pb.DouyinUserRegisterResponse, err *dyerror.DouyinError) {
	var (
		username, password string
	)
	if err = checkUserRegisterParams(c, &username, &password); err != nil {
		return nil, err
	}
	user := &entity.User{
		Name:     username,
		Password: password,
	}
	if err := UserService.CreateUser(user); err != nil {
		return nil, dyerror.DBCreateUserError
	}
	UserService.QueryUserByName(user)
	token := TokenService.GenerateToken()
	TokenService.SetToken(token, user.ID)
	return &pb.DouyinUserRegisterResponse{
		StatusCode: &constants.DefaultInt32,
		StatusMsg:  &constants.DefaultString,
		UserId:     &user.ID,
		Token:      &token,
	}, nil
}

func checkUserRegisterParams(c *gin.Context, pUsername, pPassword *string) *dyerror.DouyinError {
	username, password := c.PostForm("username"), c.PostForm("password")
	if username == "" || password == "" {
		log.Printf("username: %v, password: %v", username, password)
		return dyerror.ParamEmptyError
	}
	if len(username) > 32 || len(password) > 32 {
		log.Printf("username: %v, password: %v", username, password)
		return dyerror.ParamInputLengthExceededError
	}
	*pUsername = username
	*pPassword = password
	return nil
}
