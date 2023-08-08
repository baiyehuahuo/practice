package relation

import (
	"douyin/constants"
	"douyin/pb"
	"github.com/gin-gonic/gin"
	"strconv"
)

// ServeRelationAction handle relation action request
// Method is POST
// token, to_user_id, action_type is required
func ServeRelationAction(c *gin.Context) (res *pb.DouyinRelationActionResponse, err error) {
	var (
		token      string
		toUserID   int64
		actionType int
	)
	if err = checkRelationActionParams(c, &token, &toUserID, &actionType); err != nil {
		return nil, err
	}
	return &pb.DouyinRelationActionResponse{
		StatusCode: &constants.DefaultInt32,
		StatusMsg:  &constants.DefaultString,
	}, nil
}

func checkRelationActionParams(c *gin.Context, pToken *string, pToUserID *int64, pActionType *int) error {
	token, toUserID, actionType := c.PostForm("token"), c.PostForm("to_user_id"), c.PostForm("action_type")
	if token == "" || toUserID == "" || actionType == "" {
		return constants.ParamEmptyError
	}
	id, err1 := strconv.Atoi(toUserID)
	action, err2 := strconv.Atoi(actionType)
	if err1 != nil || err2 != nil {
		return constants.ParamInputTypeError
	}
	if action != 1 && action != 2 {
		return constants.ParamUnknownActionTypeError
	}
	*pToken = token
	*pToUserID = int64(id)
	*pActionType = action
	return nil
}
