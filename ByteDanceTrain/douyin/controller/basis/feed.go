package basis

import (
	"douyin/configs"
	"douyin/pb"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"time"
)

// ServeFeed handle feed request
// 不限制登录状态，返回按投稿时间倒序的视频列表，视频数由服务端控制，单次最多30个
// Method is GET
func ServeFeed(c *gin.Context) (res *pb.DouyinFeedResponse, err error) {
	var (
		latestTime time.Time
		token      string
	)
	if err = checkFeedParams(c, &latestTime, &token); err != nil {
		return nil, err
	}
	feedRes := pb.DouyinFeedResponse{
		StatusCode: &configs.DefaultInt32,
		StatusMsg:  &configs.DefaultString,
		VideoList:  []*pb.Video{configs.DefaultVideo},
		NextTime:   &configs.DefaultInt64,
	}

	return &feedRes, nil
}

func checkFeedParams(c *gin.Context, pLatestTime *time.Time, pToken *string) error {
	latestTime, token := c.Query("latest_time"), c.Query("token")
	if latestTime != "" {
		t, err := strconv.Atoi(latestTime)
		if err != nil {
			log.Printf("latestTimeStr: %v, token: %v", latestTime, token)
			return configs.ParamInputTypeError
		}
		*pLatestTime = time.UnixMilli(int64(t))
	} else {
		*pLatestTime = time.Now()
	}
	*pToken = token
	return nil
}
