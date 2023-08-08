package relation

import (
	"douyin/configs"
	"douyin/pb"
	"github.com/gin-gonic/gin"
	"strconv"
)

// ServeRelationFollowerList handle relation follower list request
// Method is GET
// user_id, token is required
func ServeRelationFollowerList(c *gin.Context) (res *pb.DouyinRelationFollowerListResponse, err error) {
	var (
		userID int64
		token  string
	)
	if err = checkRelationFollowerListParams(c, &userID, &token); err != nil {
		return nil, err
	}
	return &pb.DouyinRelationFollowerListResponse{
		StatusCode: &configs.DefaultInt32,
		StatusMsg:  &configs.DefaultString,
		UserList:   []*pb.User{configs.DefaultUser},
	}, nil
}

func checkRelationFollowerListParams(c *gin.Context, pUserID *int64, pToken *string) error {
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
