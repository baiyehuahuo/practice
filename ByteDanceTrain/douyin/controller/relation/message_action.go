package relation

import (
	"douyin/constants"
	"douyin/model/dyerror"
	"douyin/model/entity"
	"douyin/pb"
	"douyin/service/MessageService"
	"douyin/service/TokenService"
	"github.com/gin-gonic/gin"
	"time"
)

// ServeMessageAction handle message action request
// 点击发送，会通过消息发送接口发送该消息。
// Method is POST
// token, to_user_id, action_type, content is required
func ServeMessageAction(c *gin.Context) (res *pb.DouyinMessageActionResponse, dyerr *dyerror.DouyinError) {
	var (
		token, content string
		toUserID       int64
		actionType     int
	)
	if dyerr = checkMessageActionParams(c, &token, &toUserID, &actionType, &content); dyerr != nil {
		return nil, dyerr
	}
	userID, dyerr := TokenService.GetUserIDFromToken(token)
	if dyerr != nil {
		return nil, dyerr
	}
	switch actionType {
	case 1:
		msg := entity.Message{
			ToUserID:   toUserID,
			FromUserID: userID,
			Content:    content,
			CreateTime: time.Now().Format("01-02"),
		}
		if dyerr = MessageService.CreateMessageEvent(&msg); dyerr != nil {
			return nil, dyerr
		}
	}
	return &pb.DouyinMessageActionResponse{
		StatusCode: &constants.DefaultInt32,
		StatusMsg:  &constants.DefaultString,
	}, nil
}

func checkMessageActionParams(c *gin.Context, pToken *string, pToUserID *int64, pActionType *int, pContent *string) *dyerror.DouyinError {
	body := struct {
		Token      string `form:"token" json:"token" binding:"required"`
		ToUserID   int64  `form:"to_user_id" json:"to_user_id" binding:"required"`
		ActionType int    `form:"action_type" json:"action_type" binding:"required,oneof=1"`
		Content    string `form:"content" json:"content" binding:"required"`
	}{}
	if err := c.ShouldBind(&body); err != nil {
		return dyerror.HandleBindError(err)
	}

	*pToken = body.Token
	*pToUserID = body.ToUserID
	*pActionType = body.ActionType
	*pContent = body.Content
	return nil
}
