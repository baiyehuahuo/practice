package UserService

import (
	"douyin/model/entity"
	"douyin/service/DBService"
)

// CreateUser create a new record in the mysql database
func CreateUser(user *entity.User) error {
	err := DBService.GetDB().Create(user).Error
	//if err != nil {
	//	log.Println(err)
	//}
	return err
}

// QueryUserByID query user information from mysql by id
func QueryUserByID(userID int64) (user *entity.User) {
	user = &entity.User{}
	DBService.GetDB().First(user, userID)
	return user
}

// QueryUserByName query user information from mysql by name
func QueryUserByName(username string) (user *entity.User) {
	user = &entity.User{}
	DBService.GetDB().Where("name = ?", username).First(user)
	return user
}
