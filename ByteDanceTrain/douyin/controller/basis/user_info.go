package basis

import (
	"douyin/constants"
	"douyin/model/dyerror"
	"douyin/model/entity"
	"douyin/pb"
	"douyin/service/TokenService"
	"douyin/service/UserService"
	"github.com/gin-gonic/gin"
	"strconv"
)

// ServeUserInfo handle user request
// 获取用户的 id、昵称，如果实现社交部分的功能，还会返回关注数和粉丝数
// Method is GET
// user_id, token is required
func ServeUserInfo(c *gin.Context) (res *pb.DouyinUserResponse, dyerr *dyerror.DouyinError) {
	var (
		userID int64
		token  string
	)
	if dyerr = checkUserInfoParams(c, &userID, &token); dyerr != nil {
		return nil, dyerr
	}
	if dyerr = TokenService.CheckToken(token, userID); dyerr != nil {
		return nil, dyerr
	}
	user := &entity.User{
		ID: userID,
	}
	UserService.QueryUserByID(user)
	return &pb.DouyinUserResponse{
		StatusCode: &constants.DefaultInt32,
		StatusMsg:  &constants.DefaultString,
		User:       user.GetPBUser(),
	}, nil
}

func checkUserInfoParams(c *gin.Context, pUserID *int64, pToken *string) *dyerror.DouyinError {
	userID, token := c.Query("user_id"), c.Query("token")
	if userID == "" || token == "" {
		//log.Printf("userID: %v, token: %v", userID, token)
		return dyerror.ParamEmptyError
	}

	id, err := strconv.Atoi(userID)
	if err != nil {
		//log.Printf("userID: %v, token: %v", userID, token)
		return dyerror.ParamInputTypeError
	}
	*pUserID = int64(id)
	*pToken = token
	return nil
}
