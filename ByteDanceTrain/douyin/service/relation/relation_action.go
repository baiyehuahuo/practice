package relation

import (
	"douyin/configs"
	"douyin/pb"
	"github.com/gin-gonic/gin"
)

func ServeRelationAction(c *gin.Context) *pb.DouyinRelationActionResponse {
	return &pb.DouyinRelationActionResponse{
		StatusCode: &configs.DefaultInt32,
		StatusMsg:  &configs.DefaultString,
	}
}
