package interaction

import (
	"douyin/configs"
	"douyin/pb"
	"github.com/gin-gonic/gin"
)

func ServeFavoriteAction(c *gin.Context) *pb.DouyinFavoriteActionResponse {
	return &pb.DouyinFavoriteActionResponse{
		StatusCode: &configs.DefaultInt32,
		StatusMsg:  &configs.DefaultString,
	}
}
