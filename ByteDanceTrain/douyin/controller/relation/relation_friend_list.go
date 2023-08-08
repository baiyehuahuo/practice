package relation

import (
	"douyin/configs"
	"douyin/pb"
	"github.com/gin-gonic/gin"
	"strconv"
)

// ServeRelationFriendList handle relation friend list request
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
	return &pb.DouyinRelationFriendListResponse{
		StatusCode: &configs.DefaultInt32,
		StatusMsg:  &configs.DefaultString,
		UserList:   []*pb.User{configs.DefaultUser},
	}, nil
}

func checkRelationFriendListParams(c *gin.Context, pUserID *int64, pToken *string) error {
	userID, token := c.Query("user_id"), c.Query("token")
	if userID == "" || token == "" {
		return configs.ParamEmptyError
	}
	id, err := strconv.Atoi(userID)
	if err != nil {
		return configs.ParamInputTypeError
	}
	*pUserID = int64(id)
	*pToken = token
	return nil
}
