package relation

import (
	"douyin/configs"
	"douyin/pb"
	"github.com/gin-gonic/gin"
	"strconv"
)

// ServeRelationFollowList handle relation follow list request
// Method is GET
// user_id, token is required
func ServeRelationFollowList(c *gin.Context) (res *pb.DouyinRelationFollowListResponse, err error) {
	userID, token := c.Query("user_id"), c.Query("token")
	if userID == "" || token == "" {
		return nil, configs.ParamEmptyError
	}
	if _, err = strconv.Atoi(userID); err != nil {
		return nil, configs.ParamInputTypeError
	}
	return &pb.DouyinRelationFollowListResponse{
		StatusCode: &configs.DefaultInt32,
		StatusMsg:  &configs.DefaultString,
		UserList:   []*pb.User{configs.DefaultUser},
	}, nil
}
