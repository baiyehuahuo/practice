package VideoService

import (
	"douyin/model/entity"
	"douyin/service/DBService"
	"log"
	"time"
)

// CreateVideo create a new record in table videos
func CreateVideo(video *entity.Video) error {
	err := DBService.GetDB().Create(video).Error
	if err != nil {
		log.Println(err)
	}
	return err
}

// QueryVideosByAuthorID query publish videos where author_id == authorID
func QueryVideosByAuthorID(authorID int64) (videos []*entity.Video) {
	DBService.GetDB().Where("author_id = ?", authorID).Find(&videos)
	return videos
}

// QueryVideosByTimestamp query publish videos before timeStamp (limit 30)
func QueryVideosByTimestamp(t time.Time) (videos []*entity.Video) {
	DBService.GetDB().Where("publish_time <= ?", t.Format("2006-01-02 15:04:05")).Order("publish_time desc").Limit(30).Find(&videos)
	return videos
}
