package VideoService

import (
	"douyin/model/dyerror"
	"douyin/model/entity"
	"douyin/service/DBService"
	"time"
)

// CreateVideo create a new record in table videos
func CreateVideo(video *entity.Video) *dyerror.DouyinError {
	if err := DBService.GetDB().Create(video).Error; err != nil {
		return dyerror.DBCreateVideoError
	}
	return nil
}

// QueryWorkCountByAuthorID search video count published by author
func QueryWorkCountByAuthorID(authorID int64) (workCount int64) {
	DBService.GetDB().Model(&entity.Video{}).Where("author_id = ?", authorID).Count(&workCount)
	return workCount
}

// QueryVideoByVideoID search video where id = videoID
func QueryVideoByVideoID(videoID int64) (video *entity.Video) {
	DBService.GetDB().Where("id = ?", videoID).Find(&video)
	return video
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
