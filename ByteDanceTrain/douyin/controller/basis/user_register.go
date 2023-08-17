package basis

import (
	"douyin/constants"
	"douyin/model/dyerror"
	"douyin/model/entity"
	"douyin/pb"
	"douyin/service/TokenService"
	"douyin/service/UserService"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// ServeUserRegister handle user register request
// 新用户注册时提供用户名，密码即可，用户名需要保证唯一。创建成功后返回用户 id 和权限token
// 注册账号和登录账号的页面，该页面可以切换登录和注册两种模式，分别验证两个接口
// Method is POST
// username, password is required
func ServeUserRegister(c *gin.Context) (res *pb.DouyinUserRegisterResponse, dyerr *dyerror.DouyinError) {
	var (
		username, password string
	)
	if dyerr = checkUserRegisterParams(c, &username, &password); dyerr != nil {
		return nil, dyerr
	}
	user := &entity.User{
		Name:     username,
		Password: password,
	}
	if err := UserService.CreateUser(user); err != nil {
		return nil, dyerror.DBCreateUserError
	}
	user = UserService.QueryUserByName(user.Name)
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
	body := struct {
		Username string `form:"username" json:"username" binding:"required"` // todo limit length
		Password string `form:"password" json:"password" binding:"required"`
	}{}
	if err := c.ShouldBind(&body); err != nil {
		switch err.(type) {
		case validator.ValidationErrors:
			return dyerror.ParamEmptyError
		default:
			fmt.Printf("%T\n", err)
			dyerr := dyerror.UnknownError
			dyerr.ErrMessage = err.Error()
			return dyerr
		}
	}
	//fmt.Printf("%+v\n", body)
	username, password := body.Username, body.Password
	if len(username) > 32 || len(password) > 32 {
		//log.Printf("username: %v, password: %v", username, password)
		return dyerror.ParamInputLengthExceededError
	}
	*pUsername = username
	*pPassword = password
	return nil
}
