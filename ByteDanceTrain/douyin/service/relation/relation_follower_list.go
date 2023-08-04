package relation

import (
	"douyin/configs"
	"douyin/pb"
	"github.com/gin-gonic/gin"
)

func ServeRelationFollowerList(c *gin.Context) *pb.DouyinRelationFollowerListResponse {
	return &pb.DouyinRelationFollowerListResponse{
		StatusCode: &configs.DefaultInt32,
		StatusMsg:  &configs.DefaultString,
		UserList:   []*pb.User{configs.DefaultUser},
	}
}
