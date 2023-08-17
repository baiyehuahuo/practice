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

// ServeRelationFollowerList handle relation follower list request
// 在个人页点击 粉丝/关注，能够打开该页面，会立即调用接口拉取关注用户和粉丝用户列表
// Method is GET
// user_id, token is required
func ServeRelationFollowerList(c *gin.Context) (res *pb.DouyinRelationFollowerListResponse, dyerr *dyerror.DouyinError) {
	var (
		userID int64
		token  string
	)
	if dyerr = checkRelationFollowerListParams(c, &userID, &token); dyerr != nil {
		return nil, dyerr
	}
	if dyerr = TokenService.CheckToken(token, userID); dyerr != nil {
		return nil, dyerr
	}
	relations := RelationService.QueryRelationEventByToUserID(userID)
	pbUsers := make([]*pb.User, 0, len(relations))
	for i := range relations {
		user := common.ConvertToPBUser(UserService.QueryUserByID(relations[i].UserID))
		*user.IsFollow = RelationService.QueryFollowByIDs(userID, *user.Id)
		pbUsers = append(pbUsers, user)
	}
	return &pb.DouyinRelationFollowerListResponse{
		StatusCode: &constants.DefaultInt32,
		StatusMsg:  &constants.DefaultString,
		UserList:   pbUsers,
	}, nil
}

func checkRelationFollowerListParams(c *gin.Context, pUserID *int64, pToken *string) *dyerror.DouyinError {
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
