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

// QueryUserByID query user information from mysql by id
//func QueryUserByID(user *entity.User) {
//	DBService.GetDB().First(user, user.ID)
//}

// QueryUserByName query user information from mysql by name
//func QueryUserByName(user *entity.User) {
//	DBService.GetDB().Where("name = ?", user.Name).First(user)
//}
