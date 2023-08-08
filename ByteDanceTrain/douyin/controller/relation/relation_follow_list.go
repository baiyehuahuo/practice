package relation

import (
	"douyin/constants"
	"douyin/pb"
	"github.com/gin-gonic/gin"
	"strconv"
)

// ServeRelationFollowList handle relation follow list request
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
	return &pb.DouyinRelationFollowListResponse{
		StatusCode: &constants.DefaultInt32,
		StatusMsg:  &constants.DefaultString,
		UserList:   []*pb.User{constants.DefaultUser},
	}, nil
}

func checkRelationFollowListParams(c *gin.Context, pUserID *int64, pToken *string) error {
	userID, token := c.Query("user_id"), c.Query("token")
	if userID == "" || token == "" {
		return constants.ParamEmptyError
	}
	id, err := strconv.Atoi(userID)
	if err != nil {
		return constants.ParamInputTypeError
	}
	*pUserID = int64(id)
	*pToken = token
	return nil
}
