package basis

import (
	"douyin/constants"
	"douyin/model/dyerror"
	"douyin/pb"
	"douyin/service/TokenService"
	"douyin/service/UserService"
	"github.com/gin-gonic/gin"
)

// ServeUserLogin handle user login request
// 通过用户名和密码进行登录，登录成功后返回用户 id 和权限 token
// 注册账号和登录账号的页面，该页面可以切换登录和注册两种模式，分别验证两个接口
// Method is POST
// username, password is required
func ServeUserLogin(c *gin.Context) (res *pb.DouyinUserLoginResponse, err error) {
	var (
		username, password string
	)
	if err = checkUserLoginParams(c, &username, &password); err != nil {
		return nil, err
	}
	user := UserService.QueryUserByName(username)
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

func checkUserLoginParams(c *gin.Context, pUsername, pPassword *string) error {
	body := struct {
		Username string `form:"username" json:"username" binding:"required"`
		Password string `form:"password" json:"password" binding:"required"`
	}{}
	if err := c.ShouldBind(&body); err != nil {
		return dyerror.HandleBindError(err)
	}

	*pUsername = body.Username
	*pPassword = body.Password
	return nil
}
