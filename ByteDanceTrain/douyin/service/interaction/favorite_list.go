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
	var (
		userID int64
		token  string
	)
	if err = checkFavoriteListParams(c, &userID, &token); err != nil {
		return nil, err
	}
	return &pb.DouyinFavoriteListResponse{
		StatusCode: &configs.DefaultInt32,
		StatusMsg:  &configs.DefaultString,
		VideoList:  []*pb.Video{configs.DefaultVideo},
	}, nil
}

func checkFavoriteListParams(c *gin.Context, pUserID *int64, pToken *string) error {
	userID, token := c.Query("user_id"), c.Query("token")
	if userID == "" || token == "" {
		return configs.ParamEmptyError
	}
	id, err := strconv.Atoi(userID)
	if err != nil {
		return configs.ParamInputTypeError
	}
	*pUserID = int64(id)
	*pToken = token
	return nil
}
