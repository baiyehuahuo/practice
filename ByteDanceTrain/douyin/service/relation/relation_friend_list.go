package relation

import (
	"douyin/configs"
	"douyin/pb"
	"github.com/gin-gonic/gin"
)

func ServeRelationFriendList(c *gin.Context) *pb.DouyinRelationFriendListResponse {
	return &pb.DouyinRelationFriendListResponse{
		StatusCode: &configs.DefaultInt32,
		StatusMsg:  &configs.DefaultString,
		UserList:   []*pb.User{configs.DefaultUser},
	}
}
