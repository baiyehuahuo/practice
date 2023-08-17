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
	body := struct {
		UserID int64  `form:"user_id" json:"user_id" binding:"required"`
		Token  string `form:"token" json:"token" binding:"required"`
	}{}
	if err := c.ShouldBindQuery(&body); err != nil {
		return dyerror.HandleBindError(err)
	}

	*pUserID = body.UserID
	*pToken = body.Token
	return nil
}
