package interaction

import (
	"douyin/configs"
	"douyin/pb"
	"github.com/gin-gonic/gin"
	"strconv"
)

// ServeFavoriteList handle favorite action request
// 用户的所有点赞视频
// Method is GET
// user_id, token is required
func ServeFavoriteList(c *gin.Context) (res *pb.DouyinFavoriteListResponse, err error) {
	userID, token := c.Query("user_id"), c.Query("token")
	if userID == "" || token == "" {
		return nil, configs.ParamEmptyError
	}
	if _, err = strconv.Atoi(userID); err != nil {
		return nil, configs.ParamInputTypeError
	}
	return &pb.DouyinFavoriteListResponse{
		StatusCode: &configs.DefaultInt32,
		StatusMsg:  &configs.DefaultString,
		VideoList:  []*pb.Video{configs.DefaultVideo},
	}, nil
}
