package basis

import (
	"douyin/configs"
	"douyin/pb"
	"github.com/gin-gonic/gin"
)

func ServeFeed(c *gin.Context) *pb.DouyinFeedResponse {
	feedRes := pb.DouyinFeedResponse{
		StatusCode: &configs.DefaultInt32,
		StatusMsg:  &configs.DefaultString,
		VideoList:  []*pb.Video{configs.DefaultVideo},
		NextTime:   &configs.DefaultInt64,
	}

	return &feedRes
}
