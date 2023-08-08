package relation

import (
	"douyin/configs"
	"douyin/pb"
	"github.com/gin-gonic/gin"
	"strconv"
)

// ServeMessageChat handle message chat request
// Method is GET
// token, to_user_id is required
func ServeMessageChat(c *gin.Context) (res *pb.DouyinMessageChatResponse, err error) {
	var (
		token    string
		toUserID int64
	)
	if err = checkMessageChatParams(c, &token, &toUserID); err != nil {
		return nil, err
	}
	return &pb.DouyinMessageChatResponse{
		StatusCode:  &configs.DefaultInt32,
		StatusMsg:   &configs.DefaultString,
		MessageList: []*pb.Message{configs.DefaultMessage},
	}, nil
}

func checkMessageChatParams(c *gin.Context, pToken *string, pToUserID *int64) error {
	token, toUserID := c.Query("token"), c.Query("to_user_id")
	if token == "" || toUserID == "" {
		return configs.ParamEmptyError
	}
	id, err := strconv.Atoi(toUserID)
	if err != nil {
		return configs.ParamInputTypeError
	}
	*pToken = token
	*pToUserID = int64(id)
	return nil
}
