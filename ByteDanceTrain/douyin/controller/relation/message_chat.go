package relation

import (
	"douyin/common"
	"douyin/constants"
	"douyin/model/dyerror"
	"douyin/pb"
	"douyin/service/MessageService"
	"douyin/service/TokenService"
	"github.com/gin-gonic/gin"
)

// ServeMessageChat handle message chat request
// 点击上面任意用户，进入详细聊天页面。在该页面下会定时轮询消息查询接口，获取最新消息列表。
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
	userID, err := TokenService.GetUserIDFromToken(token)
	if err != nil {
		return nil, err
	}
	messages := MessageService.QueryMessagesByIDs(userID, toUserID)
	dpMessages := make([]*pb.Message, 0, len(messages))
	for i := range messages {
		dpMessages = append(dpMessages, common.ConvertToPBMessage(messages[i]))
	}
	return &pb.DouyinMessageChatResponse{
		StatusCode:  &constants.DefaultInt32,
		StatusMsg:   &constants.DefaultString,
		MessageList: dpMessages,
	}, nil
}

func checkMessageChatParams(c *gin.Context, pToken *string, pToUserID *int64) error {
	body := struct {
		Token    string `form:"token" json:"token" binding:"required"`
		ToUserID int64  `form:"to_user_id" json:"to_user_id" binding:"required"`
	}{}
	if err := c.ShouldBindQuery(&body); err != nil {
		return dyerror.HandleBindError(err)
	}

	*pToken = body.Token
	*pToUserID = body.ToUserID
	return nil
}
