package basis

import (
	"douyin/constants"
	"douyin/model/dyerror"
	"douyin/pb"
	"douyin/service/TokenService"
	"douyin/service/UserService"
	"douyin/service/VideoService"
	"github.com/gin-gonic/gin"
	"strconv"
)

// ServePublishList handle publish list request
// 用户的视频发布列表，直接列出用户所有投稿过的视频
// Method is GET
// user_id, token is required
func ServePublishList(c *gin.Context) (res *pb.DouyinPublishListResponse, err *dyerror.DouyinError) {
	var (
		userID int64
		token  string
	)
	if err = checkPublishListParams(c, &userID, &token); err != nil {
		return nil, err
	}
	if err = TokenService.CheckToken(token, userID); err != nil {
		return nil, err
	}

	author := UserService.QueryUserByID(userID).GetPBUser()
	videos := VideoService.QueryVideosByAuthorID(*author.Id)
	pbVideos := make([]*pb.Video, 0, len(videos))
	for i := range videos {
		pbVideos = append(pbVideos, videos[i].GetPBVideo(author))
	}
	return &pb.DouyinPublishListResponse{
		StatusCode: &constants.DefaultInt32,
		StatusMsg:  &constants.DefaultString,
		VideoList:  pbVideos,
	}, nil
}

func checkPublishListParams(c *gin.Context, pUserID *int64, pToken *string) *dyerror.DouyinError {
	userID, token := c.Query("user_id"), c.Query("token")
	if userID == "" || token == "" {
		//log.Printf("userID: %v, token: %v", userID, token)
		return dyerror.ParamEmptyError
	}

	id, err := strconv.Atoi(userID)
	if err != nil {
		//log.Printf("userID: %v, token: %v", userID, token)
		return dyerror.ParamInputTypeError
	}
	*pUserID = int64(id)
	*pToken = token
	return nil
}
