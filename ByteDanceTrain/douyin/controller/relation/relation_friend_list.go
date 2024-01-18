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

// ServeRelationFriendList handle relation friend list request
// 注册登录后，点击消息页面，会立即请求该接口，获取可聊天朋友列表，并且会带着和该用户的最新的一条消息
// 互相关注就是朋友
// Method is GET
// user_id, token is required
func ServeRelationFriendList(c *gin.Context) (res *pb.DouyinRelationFriendListResponse, err error) {
	var (
		userID int64
		token  string
	)
	if err = checkRelationFriendListParams(c, &userID, &token); err != nil {
		return nil, err
	}
	if err = TokenService.CheckToken(token, userID); err != nil {
		return nil, err
	}
	var pbUsers []*pb.User
	// follow relation
	relation := RelationService.QueryRelationEventByUserID(userID)
	followSet := make(map[int64]struct{}, len(relation))
	for i := range relation {
		followSet[relation[i].ToUserID] = struct{}{}
	}
	// follower relation
	relation = RelationService.QueryRelationEventByToUserID(userID)
	for i := range relation {
		if _, ok := followSet[relation[i].UserID]; ok {
			user := common.ConvertToPBUser(UserService.QueryUserByID(relation[i].UserID))
			*user.IsFollow = true // RelationService.QueryFollowByIDs(userID, *user.Id) // 互相关注才是好友
			pbUsers = append(pbUsers, user)
		}
	}
	return &pb.DouyinRelationFriendListResponse{
		StatusCode: &constants.DefaultInt32,
		StatusMsg:  &constants.DefaultString,
		UserList:   pbUsers,
	}, nil
}

func checkRelationFriendListParams(c *gin.Context, pUserID *int64, pToken *string) error {
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
