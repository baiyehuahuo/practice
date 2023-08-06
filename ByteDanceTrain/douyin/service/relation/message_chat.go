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
	token, toUserID := c.Query("token"), c.Query("to_user_id")
	if token == "" || toUserID == "" {
		return nil, configs.ParamEmptyError
	}
	if _, err = strconv.Atoi(toUserID); err != nil {
		return nil, configs.ParamInputTypeError
	}
	return &pb.DouyinMessageChatResponse{
		StatusCode:  &configs.DefaultInt32,
		StatusMsg:   &configs.DefaultString,
		MessageList: []*pb.Message{configs.DefaultMessage},
	}, nil
}
