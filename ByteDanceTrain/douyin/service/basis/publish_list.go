package basis

import (
	"douyin/configs"
	"douyin/pb"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

// ServePublishList handle publish list request
// 用户的视频发布列表，直接列出用户所有投稿过的视频
// Method is GET
// user_id, token is required
func ServePublishList(c *gin.Context) (res *pb.DouyinPublishListResponse, err error) {
	userID, token := c.Query("user_id"), c.Query("token")
	if userID == "" || token == "" {
		log.Printf("userID: %v, token: %v", userID, token)
		return nil, configs.ParamEmptyError
	}
	if _, err = strconv.Atoi(userID); err != nil {
		log.Printf("userID: %v, token: %v", userID, token)
		return nil, configs.ParamInputTypeError
	}
	return &pb.DouyinPublishListResponse{
		StatusCode: &configs.DefaultInt32,
		StatusMsg:  &configs.DefaultString,
		VideoList:  []*pb.Video{configs.DefaultVideo},
	}, nil
}
