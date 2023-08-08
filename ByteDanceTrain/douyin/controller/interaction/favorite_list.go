package interaction

import (
	"douyin/constants"
	"douyin/model/dyerror"
	"douyin/pb"
	"github.com/gin-gonic/gin"
	"strconv"
)

// ServeFavoriteList handle favorite action request
// 用户的所有点赞视频
// Method is GET
// user_id, token is required
func ServeFavoriteList(c *gin.Context) (res *pb.DouyinFavoriteListResponse, err *dyerror.DouyinError) {
	var (
		userID int64
		token  string
	)
	if err = checkFavoriteListParams(c, &userID, &token); err != nil {
		return nil, err
	}
	return &pb.DouyinFavoriteListResponse{
		StatusCode: &constants.DefaultInt32,
		StatusMsg:  &constants.DefaultString,
		VideoList:  []*pb.Video{constants.DefaultVideo},
	}, nil
}

func checkFavoriteListParams(c *gin.Context, pUserID *int64, pToken *string) *dyerror.DouyinError {
	userID, token := c.Query("user_id"), c.Query("token")
	if userID == "" || token == "" {
		return dyerror.ParamEmptyError
	}
	id, err := strconv.Atoi(userID)
	if err != nil {
		return dyerror.ParamInputTypeError
	}
	*pUserID = int64(id)
	*pToken = token
	return nil
}
