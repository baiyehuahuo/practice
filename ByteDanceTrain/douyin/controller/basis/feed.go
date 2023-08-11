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
	"strings"
	"time"
)

// ServeFeed handle feed request
// 不限制登录状态，返回按投稿时间倒序的视频列表，视频数由服务端控制，单次最多30个
// Method is GET
func ServeFeed(c *gin.Context) (res *pb.DouyinFeedResponse, dyerr *dyerror.DouyinError) {
	var (
		latestTime time.Time
		token      string
	)
	if dyerr = checkFeedParams(c, &latestTime, &token); dyerr != nil {
		return nil, dyerr
	}
	if _, dyerr = TokenService.GetUserIDFromToken(token); token != "" && dyerr != nil { // 如果有 token 要验证
		return nil, dyerr
	}
	videos := VideoService.QueryVideosByTimestamp(latestTime)
	pbVideoList := make([]*pb.Video, 0, len(videos))
	for i := range videos {
		//log.Printf("video title: %s, timestamp: %d", videos[i].Title, videos[i].PublishTime.Unix())
		pbAuthor := UserService.QueryUserByID(videos[i].AuthorID).GetPBUser()
		pbVideoList = append(pbVideoList, videos[i].GetPBVideo(pbAuthor))
	}
	nextTime := videos[len(videos)-1].PublishTime.Unix()
	feedRes := pb.DouyinFeedResponse{
		StatusCode: &constants.DefaultInt32,
		StatusMsg:  &constants.DefaultString,
		VideoList:  pbVideoList,
		NextTime:   &nextTime,
	}

	return &feedRes, nil
}

func checkFeedParams(c *gin.Context, pLatestTime *time.Time, pToken *string) *dyerror.DouyinError {
	latestTime, token := c.Query("latest_time"), strings.Trim(c.Query("token"), " ")
	if latestTime != "" {
		t, err := strconv.Atoi(latestTime)
		if err != nil {
			return dyerror.ParamInputTypeError
		}
		*pLatestTime = time.Unix(int64(t), 0)
	} else {
		*pLatestTime = time.Now()
	}
	*pToken = token
	return nil
}
