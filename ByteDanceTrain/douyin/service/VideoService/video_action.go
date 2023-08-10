package VideoService

import (
	"douyin/model/entity"
	"douyin/service/DBService"
	"log"
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
