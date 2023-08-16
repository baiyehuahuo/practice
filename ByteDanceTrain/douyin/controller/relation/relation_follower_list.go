package relation

import (
	"douyin/constants"
	"douyin/model/dyerror"
	"douyin/pb"
	"github.com/gin-gonic/gin"
	"strconv"
)

// ServeRelationFollowerList handle relation follower list request
// 在个人页点击 粉丝/关注，能够打开该页面，会立即调用接口拉取关注用户和粉丝用户列表
// Method is GET
// user_id, token is required
func ServeRelationFollowerList(c *gin.Context) (res *pb.DouyinRelationFollowerListResponse, err *dyerror.DouyinError) {
	var (
		userID int64
		token  string
	)
	if err = checkRelationFollowerListParams(c, &userID, &token); err != nil {
		return nil, err
	}
	return &pb.DouyinRelationFollowerListResponse{
		StatusCode: &constants.DefaultInt32,
		StatusMsg:  &constants.DefaultString,
		UserList:   []*pb.User{constants.DefaultUser},
	}, nil
}

func checkRelationFollowerListParams(c *gin.Context, pUserID *int64, pToken *string) *dyerror.DouyinError {
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
