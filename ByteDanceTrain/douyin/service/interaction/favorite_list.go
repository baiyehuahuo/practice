package interaction

import (
	"douyin/configs"
	"douyin/pb"
	"github.com/gin-gonic/gin"
)

func ServeFavoriteList(c *gin.Context) *pb.DouyinFavoriteListResponse {
	return &pb.DouyinFavoriteListResponse{
		StatusCode: &configs.DefaultInt32,
		StatusMsg:  &configs.DefaultString,
		VideoList:  []*pb.Video{configs.DefaultVideo},
	}
}
