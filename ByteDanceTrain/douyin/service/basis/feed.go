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
	latestTimeStr, token := c.Query("latest_time"), c.Query("token")
	var latestTime time.Time
	if latestTimeStr != "" {
		t, err := strconv.Atoi(latestTimeStr)
		if err != nil {
			log.Printf("latestTimeStr: %v, token: %v", latestTimeStr, token)
			return nil, configs.ParamInputTypeError
		}
		latestTime = time.UnixMilli(int64(t))
	} else {
		latestTime = time.Now()
	}
	log.Printf("latestTime: %v", latestTime)
	feedRes := pb.DouyinFeedResponse{
		StatusCode: &configs.DefaultInt32,
		StatusMsg:  &configs.DefaultString,
		VideoList:  []*pb.Video{configs.DefaultVideo},
		NextTime:   &configs.DefaultInt64,
	}

	return &feedRes, nil
}
