package relation

import (
	"douyin/common"
	"douyin/constants"
	"douyin/model/dyerror"
	"douyin/pb"
	"douyin/service/MessageService"
	"douyin/service/TokenService"
	"github.com/gin-gonic/gin"
	"strconv"
)

// ServeMessageChat handle message chat request
// 点击上面任意用户，进入详细聊天页面。在该页面下会定时轮询消息查询接口，获取最新消息列表。
// Method is GET
// token, to_user_id is required
func ServeMessageChat(c *gin.Context) (res *pb.DouyinMessageChatResponse, dyerr *dyerror.DouyinError) {
	var (
		token    string
		toUserID int64
	)
	if dyerr = checkMessageChatParams(c, &token, &toUserID); dyerr != nil {
		return nil, dyerr
	}
	userID, dyerr := TokenService.GetUserIDFromToken(token)
	if dyerr != nil {
		return nil, dyerr
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
