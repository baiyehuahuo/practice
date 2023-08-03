package basis

import (
	"douyin/pb"
	"github.com/gin-gonic/gin"
	"time"
)

func GetFeed(c *gin.Context) *pb.DouyinFeedResponse {
	feedRes := pb.DouyinFeedResponse{
		StatusCode: new(int32),
		StatusMsg:  new(string),
		VideoList:  nil,
		NextTime:   new(int64),
	}
	*feedRes.StatusCode = 0
	*feedRes.StatusMsg = "string"
	feedRes.VideoList = append(feedRes.VideoList, &pb.Video{
		Id:            nil,
		Author:        nil,
		PlayUrl:       nil,
		CoverUrl:      nil,
		FavoriteCount: nil,
		CommentCount:  nil,
		IsFavorite:    nil,
		Title:         nil,
	})
	*feedRes.NextTime = time.Now().Unix()
	return &feedRes
}
