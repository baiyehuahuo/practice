package relation

import (
	"douyin/constants"
	"douyin/model/dyerror"
	"douyin/pb"
	"github.com/gin-gonic/gin"
	"strconv"
)

// ServeMessageChat handle message chat request
// Method is GET
// token, to_user_id is required
func ServeMessageChat(c *gin.Context) (res *pb.DouyinMessageChatResponse, err *dyerror.DouyinError) {
	var (
		token    string
		toUserID int64
	)
	if err = checkMessageChatParams(c, &token, &toUserID); err != nil {
		return nil, err
	}
	return &pb.DouyinMessageChatResponse{
		StatusCode:  &constants.DefaultInt32,
		StatusMsg:   &constants.DefaultString,
		MessageList: []*pb.Message{constants.DefaultMessage},
	}, nil
}

func checkMessageChatParams(c *gin.Context, pToken *string, pToUserID *int64) *dyerror.DouyinError {
	token, toUserID := c.Query("token"), c.Query("to_user_id")
	if token == "" || toUserID == "" {
		return dyerror.ParamEmptyError
	}
	id, err := strconv.Atoi(toUserID)
	if err != nil {
		return dyerror.ParamInputTypeError
	}
	*pToken = token
	*pToUserID = int64(id)
	return nil
}
