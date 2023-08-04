package basis

import (
	"douyin/configs"
	"douyin/pb"
	"github.com/gin-gonic/gin"
)

func ServePublishAction(c *gin.Context) *pb.DouyinPublishActionResponse {
	return &pb.DouyinPublishActionResponse{
		StatusCode: &configs.DefaultInt32,
		StatusMsg:  &configs.DefaultString,
	}
}
