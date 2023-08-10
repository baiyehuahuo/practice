package basis

import (
	"douyin/constants"
	"douyin/model/dyerror"
	"douyin/model/entity"
	"douyin/pb"
	"douyin/service/TokenService"
	"douyin/service/UserService"
	"github.com/gin-gonic/gin"
)

// ServeUserLogin handle user login request
// 通过用户名和密码进行登录，登录成功后返回用户 id 和权限 token
// Method is POST
// username, password is required
func ServeUserLogin(c *gin.Context) (res *pb.DouyinUserLoginResponse, dyerr *dyerror.DouyinError) {
	var (
		username, password string
	)
	if dyerr = checkUserLoginParams(c, &username, &password); dyerr != nil {
		return nil, dyerr
	}
	user := &entity.User{
		Name: username,
	}
	UserService.QueryUserByName(user)
	if user.Password != password {
		return nil, dyerror.AuthUsernameOrPasswordFailError
	}

	token := TokenService.GenerateToken()
	TokenService.SetToken(token, user.ID)

	return &pb.DouyinUserLoginResponse{
		StatusCode: &constants.DefaultInt32,
		StatusMsg:  &constants.DefaultString,
		UserId:     &user.ID,
		Token:      &token,
	}, nil
}

func checkUserLoginParams(c *gin.Context, pUsername, pPassword *string) *dyerror.DouyinError {
	username, password := c.PostForm("username"), c.PostForm("password")
	if username == "" || password == "" {
		//log.Printf("username: %v, password: %v", username, password)
		return dyerror.ParamEmptyError
	}
	*pUsername = username
	*pPassword = password
	return nil
}
