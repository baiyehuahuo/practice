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
	"time"
)

// ServeFeed handle feed request
// 不限制登录状态，返回按投稿时间倒序的视频列表，视频数由服务端控制，单次最多30个
// 打开App会立即调用接口/douyin/feed/获取视频列表信息
// Method is GET
func ServeFeed(c *gin.Context) (res *pb.DouyinFeedResponse, dyerr *dyerror.DouyinError) {
	var (
		latestTime time.Time
		token      string
	)
	if dyerr = checkFeedParams(c, &latestTime, &token); dyerr != nil {
		return nil, dyerr
	}
	userID, dyerr := TokenService.GetUserIDFromToken(token)
	if token != "" && dyerr != nil { // 如果有 token 要验证
		return nil, dyerr
	}
	videos := VideoService.QueryVideosByTimestamp(latestTime)
	pbVideoList := make([]*pb.Video, 0, len(videos))
	for i := range videos {
		//log.Printf("video title: %s, timestamp: %d", videos[i].Title, videos[i].PublishTime.Unix())
		pbAuthor := common.ConvertToPBUser(UserService.QueryUserByID(videos[i].AuthorID))
		*pbAuthor.IsFollow = RelationService.QueryFollowByIDs(userID, *pbAuthor.Id)
		pbVideo := common.ConvertToPBVideo(videos[i], pbAuthor)
		if userID != 0 {
			*pbVideo.IsFavorite = FavoriteService.QueryFavoriteByIDs(userID, *pbVideo.Id)
		} else {
			*pbVideo.IsFavorite = false
		}
		pbVideoList = append(pbVideoList, pbVideo)
	}
	var nextTime int64
	if len(videos) > 0 {
		nextTime = videos[len(videos)-1].PublishTime.Unix()
	}

	return &pb.DouyinFeedResponse{
		StatusCode: &constants.DefaultInt32,
		StatusMsg:  &constants.DefaultString,
		VideoList:  pbVideoList,
		NextTime:   &nextTime,
	}, nil
}

func checkFeedParams(c *gin.Context, pLatestTime *time.Time, pToken *string) *dyerror.DouyinError {
	body := struct {
		Token      string `form:"token" json:"token"`
		LatestTime int64  `form:"latest_time" json:"latest_time"`
	}{}

	if err := c.ShouldBindQuery(&body); err != nil {
		return dyerror.HandleBindError(err)
	}
	latestTime := body.LatestTime
	if latestTime != 0 {
		*pLatestTime = time.Unix(latestTime, 0)
	} else {
		*pLatestTime = time.Now()
	}
	*pToken = body.Token
	return nil
}
