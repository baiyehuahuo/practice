package relation

import (
	"douyin/common"
	"douyin/constants"
	"douyin/model/dyerror"
	"douyin/pb"
	"douyin/service/RelationService"
	"douyin/service/TokenService"
	"douyin/service/UserService"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

// ServeRelationFriendList handle relation friend list request
// 注册登录后，点击消息页面，会立即请求该接口，获取可聊天朋友列表，并且会带着和该用户的最新的一条消息
// 互相关注就是朋友
// Method is GET
// user_id, token is required
func ServeRelationFriendList(c *gin.Context) (res *pb.DouyinRelationFriendListResponse, dyerr *dyerror.DouyinError) {
	var (
		userID int64
		token  string
	)
	if dyerr = checkRelationFriendListParams(c, &userID, &token); dyerr != nil {
		return nil, dyerr
	}
	if dyerr = TokenService.CheckToken(token, userID); dyerr != nil {
		return nil, dyerr
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

func checkRelationFriendListParams(c *gin.Context, pUserID *int64, pToken *string) *dyerror.DouyinError {
	body := struct {
		common.TokenAuthFields
	}{}
	if err := c.ShouldBindQuery(&body); err != nil {
		fmt.Println(err)
	}
	//fmt.Printf("%+v\n", body)
	userID, token := c.Query("user_id"), c.Query("token")
	if userID == "" || token == "" {
		return dyerror.ParamEmptyError
	}
	id, err := strconv.Atoi(userID)
	if err != nil {
		return dyerror.ParamInputTypeError
	}
	*pUserID = int64(id)
	*pToken = token
	return nil
}
