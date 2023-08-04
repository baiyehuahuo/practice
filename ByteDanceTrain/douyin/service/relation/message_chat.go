package relation

import (
	"douyin/configs"
	"douyin/pb"
	"github.com/gin-gonic/gin"
)

func ServeMessageChat(c *gin.Context) *pb.DouyinMessageChatResponse {
	return &pb.DouyinMessageChatResponse{
		StatusCode:  &configs.DefaultInt32,
		StatusMsg:   &configs.DefaultString,
		MessageList: []*pb.Message{configs.DefaultMessage},
	}
}
