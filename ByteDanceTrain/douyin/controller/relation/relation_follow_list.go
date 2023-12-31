package relation

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

// ServeRelationFollowList handle relation follow list request
// 登录用户关注的所有用户列表
// 在个人页点击 粉丝/关注，能够打开该页面，会立即调用接口拉取关注用户和粉丝用户列表
// Method is GET
// user_id, token is required
func ServeRelationFollowList(c *gin.Context) (res *pb.DouyinRelationFollowListResponse, err error) {
	var (
		userID int64
		token  string
	)
	if err = checkRelationFollowListParams(c, &userID, &token); err != nil {
		return nil, err
	}
	if err = TokenService.CheckToken(token, userID); err != nil {
		return nil, err
	}
	relations := RelationService.QueryRelationEventByUserID(userID)
	pbUsers := make([]*pb.User, 0, len(relations))
	for i := range relations {
		user := common.ConvertToPBUser(UserService.QueryUserByID(relations[i].ToUserID))
		*user.IsFollow = true // RelationService.QueryFollowByIDs(userID, *user.Id) 本来就是关注目标
		pbUsers = append(pbUsers, user)
	}
	return &pb.DouyinRelationFollowListResponse{
		StatusCode: &constants.DefaultInt32,
		StatusMsg:  &constants.DefaultString,
		UserList:   pbUsers,
	}, nil
}

func checkRelationFollowListParams(c *gin.Context, pUserID *int64, pToken *string) error {
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
