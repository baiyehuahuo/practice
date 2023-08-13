package interaction

import (
	"douyin/common"
	"douyin/constants"
	"douyin/model/dyerror"
	"douyin/pb"
	"douyin/service/FavoriteService"
	"douyin/service/TokenService"
	"douyin/service/UserService"
	"douyin/service/VideoService"
	"github.com/gin-gonic/gin"
	"strconv"
)

// ServeFavoriteList handle favorite action request
// 用户的所有点赞视频
// Method is GET
// user_id, token is required
func ServeFavoriteList(c *gin.Context) (res *pb.DouyinFavoriteListResponse, dyerr *dyerror.DouyinError) {
	var (
		userID int64
		token  string
	)
	if dyerr = checkFavoriteListParams(c, &userID, &token); dyerr != nil {
		return nil, dyerr
	}
	if dyerr = TokenService.CheckToken(token, userID); dyerr != nil {
		return nil, dyerr
	}
	favorites := FavoriteService.QueryFavoritesByUserID(userID)
	pbVideoList := make([]*pb.Video, 0, len(favorites))
	for i := range favorites {
		video := VideoService.QueryVideoByVideoID(favorites[i].VideoID)
		author := common.ConvertToPBUser(UserService.QueryUserByID(video.AuthorID)) // more precision
		pbVideoList = append(pbVideoList, common.ConvertToPBVideo(video, author))
	}
	return &pb.DouyinFavoriteListResponse{
		StatusCode: &constants.DefaultInt32,
		StatusMsg:  &constants.DefaultString,
		VideoList:  pbVideoList,
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
