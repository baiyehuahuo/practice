package basis

import (
	"douyin/constants"
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
	var (
		userID int64
		token  string
	)
	if err = checkPublishListParams(c, &userID, &token); err != nil {
		return nil, err
	}
	//if err = common.CheckToken(token, userID); err != nil {
	//	return nil, err
	//}

	return &pb.DouyinPublishListResponse{
		StatusCode: &constants.DefaultInt32,
		StatusMsg:  &constants.DefaultString,
		VideoList:  []*pb.Video{constants.DefaultVideo},
	}, nil
}

func checkPublishListParams(c *gin.Context, pUserID *int64, pToken *string) error {
	userID, token := c.Query("user_id"), c.Query("token")
	if userID == "" || token == "" {
		log.Printf("userID: %v, token: %v", userID, token)
		return constants.ParamEmptyError
	}

	id, err := strconv.Atoi(userID)
	if err != nil {
		log.Printf("userID: %v, token: %v", userID, token)
		return constants.ParamInputTypeError
	}
	*pUserID = int64(id)
	*pToken = token
	return nil
}
