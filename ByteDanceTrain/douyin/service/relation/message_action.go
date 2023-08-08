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
	var (
		token, content string
		toUserID       int64
		actionType     int
	)
	if err = checkMessageActionParams(c, &token, &toUserID, &actionType, &content); err != nil {
		return nil, err
	}

	return &pb.DouyinMessageActionResponse{
		StatusCode: &configs.DefaultInt32,
		StatusMsg:  &configs.DefaultString,
	}, nil
}

func checkMessageActionParams(c *gin.Context, pToken *string, pToUserID *int64, pActionType *int, pContent *string) error {
	token, toUserID, actionType, content := c.PostForm("token"), c.PostForm("to_user_id"), c.PostForm("action_type"), c.PostForm("content")
	if token == "" || toUserID == "" || actionType == "" || content == "" {
		return configs.ParamEmptyError
	}
	id, err1 := strconv.Atoi(toUserID)
	action, err2 := strconv.Atoi(actionType)
	if err1 != nil || err2 != nil {
		return configs.ParamInputTypeError
	}
	if action != 1 {
		return configs.ParamUnknownActionTypeError
	}
	*pToken = token
	*pToUserID = int64(id)
	*pActionType = action
	*pContent = content
	return nil
}
