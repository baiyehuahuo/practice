package relation

import (
	"douyin/configs"
	"douyin/pb"
	"github.com/gin-gonic/gin"
)

func ServeRelationFollowList(c *gin.Context) *pb.DouyinRelationFollowListResponse {
	return &pb.DouyinRelationFollowListResponse{
		StatusCode: &configs.DefaultInt32,
		StatusMsg:  &configs.DefaultString,
		UserList:   []*pb.User{configs.DefaultUser},
	}
}
