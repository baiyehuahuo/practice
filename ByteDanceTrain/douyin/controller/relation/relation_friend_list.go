package relation

import (
	"douyin/constants"
	"douyin/model/dyerror"
	"douyin/pb"
	"github.com/gin-gonic/gin"
	"strconv"
)

// ServeRelationFriendList handle relation friend list request
// Method is GET
// user_id, token is required
func ServeRelationFriendList(c *gin.Context) (res *pb.DouyinRelationFriendListResponse, err *dyerror.DouyinError) {
	var (
		userID int64
		token  string
	)
	if err = checkRelationFriendListParams(c, &userID, &token); err != nil {
		return nil, err
	}
	return &pb.DouyinRelationFriendListResponse{
		StatusCode: &constants.DefaultInt32,
		StatusMsg:  &constants.DefaultString,
		UserList:   []*pb.User{constants.DefaultUser},
	}, nil
}

func checkRelationFriendListParams(c *gin.Context, pUserID *int64, pToken *string) *dyerror.DouyinError {
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
