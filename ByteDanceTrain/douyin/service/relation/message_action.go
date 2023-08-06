package relation

import (
	"douyin/configs"
	"douyin/pb"
	"github.com/gin-gonic/gin"
	"strconv"
)

// ServeMessageAction handle message action request
// Method is POST
// token, to_user_id, action_type, content is required
func ServeMessageAction(c *gin.Context) (res *pb.DouyinMessageActionResponse, err error) {
	token, toUserID, actionType, content := c.PostForm("token"), c.PostForm("to_user_id"), c.PostForm("action_type"), c.PostForm("content")
	if token == "" || toUserID == "" || actionType == "" || content == "" {
		return nil, configs.ParamEmptyError
	}
	if _, err = strconv.Atoi(toUserID); err != nil {
		return nil, configs.ParamInputTypeError
	}
	if action, _ := strconv.Atoi(actionType); action != 1 {
		return nil, configs.ParamUnknownActionTypeError
	}
	return &pb.DouyinMessageActionResponse{
		StatusCode: &configs.DefaultInt32,
		StatusMsg:  &configs.DefaultString,
	}, nil
}
