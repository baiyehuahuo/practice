package basis

import (
	"douyin/common"
	"douyin/constants"
	"douyin/model/dyerror"
	"douyin/pb"
	"douyin/service/RelationService"
	"douyin/service/TokenService"
	"douyin/service/UserService"
	"github.com/gin-gonic/gin"
	"strconv"
)

// ServeUserInfo handle user request
// 获取用户的 id、昵称，如果实现社交部分的功能，还会返回关注数和粉丝数
// 在注册成功后会调用/douyin/user/接口拉取当前登录用户的全部信息，并存储到本地。
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
	user := common.ConvertToPBUser(UserService.QueryUserByID(userID)) // fwf 假设可能存在查询不到的情况？
	*user.IsFollow = RelationService.QueryFollowByIDs(userID, *user.Id)
	return &pb.DouyinUserResponse{
		StatusCode: &constants.DefaultInt32,
		StatusMsg:  &constants.DefaultString,
		User:       user,
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
