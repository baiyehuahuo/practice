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

// ServeUserRegister handle user register request
// 新用户注册时提供用户名，密码即可，用户名需要保证唯一。创建成功后返回用户 id 和权限token
// 注册账号和登录账号的页面，该页面可以切换登录和注册两种模式，分别验证两个接口
// Method is POST
// username, password is required
func ServeUserRegister(c *gin.Context) (res *pb.DouyinUserRegisterResponse, err error) {
	var (
		username, password string
	)
	if err = checkUserRegisterParams(c, &username, &password); err != nil {
		return nil, err
	}
	if err = UserService.CreateUser(&entity.User{
		Name:     username,
		Password: password,
	}); err != nil {
		return nil, err
	}
	user := UserService.QueryUserByName(username)
	token := TokenService.GenerateToken()
	TokenService.SetToken(token, user.ID)
	return &pb.DouyinUserRegisterResponse{
		StatusCode: &constants.DefaultInt32,
		StatusMsg:  &constants.DefaultString,
		UserId:     &user.ID,
		Token:      &token,
	}, nil
}

func checkUserRegisterParams(c *gin.Context, pUsername, pPassword *string) error {
	body := struct {
		Username string `form:"username" json:"username" binding:"required,lte=32"` // todo limit length
		Password string `form:"password" json:"password" binding:"required,lte=32"`
	}{}
	if err := c.ShouldBind(&body); err != nil {
		return dyerror.HandleBindError(err)
	}
	*pUsername = body.Username
	*pPassword = body.Password
	return nil
}
