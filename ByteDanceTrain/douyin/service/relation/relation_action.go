package relation

import (
	"douyin/configs"
	"douyin/pb"
	"github.com/gin-gonic/gin"
	"strconv"
)

// ServeRelationAction handle relation action request
// Method is POST
// token, to_user_id, action_type is required
func ServeRelationAction(c *gin.Context) (res *pb.DouyinRelationActionResponse, err error) {
	token, toUserID, actionType := c.PostForm("token"), c.PostForm("to_user_id"), c.PostForm("action_type")
	if token == "" || toUserID == "" || actionType == "" {
		return nil, configs.ParamEmptyError
	}
	if _, err = strconv.Atoi(toUserID); err != nil {
		return nil, configs.ParamInputTypeError
	}
	if action, _ := strconv.Atoi(actionType); action != 1 && action != 2 {
		return nil, configs.ParamUnknownActionTypeError
	}
	return &pb.DouyinRelationActionResponse{
		StatusCode: &configs.DefaultInt32,
		StatusMsg:  &configs.DefaultString,
	}, nil
}
