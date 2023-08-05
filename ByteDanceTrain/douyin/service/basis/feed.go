package basis

import (
	"douyin/configs"
	"douyin/pb"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"time"
)

func ServeFeed(c *gin.Context) (res *pb.DouyinFeedResponse, err error) {
	token := c.Query("token")
	var latestTime time.Time
	if latestTimeStr := c.Query("latest_time"); latestTimeStr != "" {
		t, err := strconv.Atoi(latestTimeStr)
		if err != nil {
			return nil, configs.LatestTimeParamError
		}
		latestTime = time.UnixMilli(int64(t))
	} else {
		latestTime = time.Now()
	}
	log.Printf("latestTime: %v, token: %v", latestTime, token)
	feedRes := pb.DouyinFeedResponse{
		StatusCode: &configs.DefaultInt32,
		StatusMsg:  &configs.DefaultString,
		VideoList:  []*pb.Video{configs.DefaultVideo},
		NextTime:   &configs.DefaultInt64,
	}

	return &feedRes, nil
}
