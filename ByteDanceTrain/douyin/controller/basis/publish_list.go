package basis

import (
	"douyin/common"
	"douyin/constants"
	"douyin/model/dyerror"
	"douyin/pb"
	"douyin/service/FavoriteService"
	"douyin/service/RelationService"
	"douyin/service/TokenService"
	"douyin/service/UserService"
	"douyin/service/VideoService"
	"github.com/gin-gonic/gin"
	"strconv"
)

// ServePublishList handle publish list request
// 用户的视频发布列表，直接列出用户所有投稿过的视频
// 打开个人页，会立即调用这两个接口，分别将内容显示在“作品”和“喜欢”两个栏目下
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

	author := common.ConvertToPBUser(UserService.QueryUserByID(userID))
	*author.IsFollow = RelationService.QueryFollowByIDs(userID, *author.Id)
	videos := VideoService.QueryVideosByAuthorID(*author.Id)
	pbVideos := make([]*pb.Video, 0, len(videos))
	for i := range videos {
		pbVideo := common.ConvertToPBVideo(videos[i], author)
		*pbVideo.IsFavorite = FavoriteService.QueryFavoriteByIDs(userID, *pbVideo.Id)
		pbVideos = append(pbVideos, pbVideo)
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
