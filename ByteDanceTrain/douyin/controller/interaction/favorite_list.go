package interaction

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
)

// ServeFavoriteList handle favorite action request
// 用户的所有点赞视频
// 打开个人页，会立即调用这两个接口，分别将内容显示在“作品”和“喜欢”两个栏目下
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
	if err = TokenService.CheckToken(token, userID); err != nil {
		return nil, err
	}
	favorites := FavoriteService.QueryFavoritesByUserID(userID)
	pbVideoList := make([]*pb.Video, 0, len(favorites))
	for i := range favorites {
		video := VideoService.QueryVideoByVideoID(favorites[i].VideoID)
		author := common.ConvertToPBUser(UserService.QueryUserByID(video.AuthorID)) // more precision
		*author.IsFollow = RelationService.QueryFollowByIDs(userID, *author.Id)
		pbVideo := common.ConvertToPBVideo(video, author)
		*pbVideo.IsFavorite = true // 本来就是查 favorite 的
		pbVideoList = append(pbVideoList, pbVideo)
	}
	return &pb.DouyinFavoriteListResponse{
		StatusCode: &constants.DefaultInt32,
		StatusMsg:  &constants.DefaultString,
		VideoList:  pbVideoList,
	}, nil
}

func checkFavoriteListParams(c *gin.Context, pUserID *int64, pToken *string) error {
	body := struct {
		UserID int64  `form:"user_id" json:"user_id" binding:"required"`
		Token  string `form:"token" json:"token" binding:"required"`
	}{}
	if err := c.ShouldBindQuery(&body); err != nil {
		return dyerror.HandleBindError(err)
	}

	*pUserID = body.UserID
	*pToken = body.Token
	return nil
}
