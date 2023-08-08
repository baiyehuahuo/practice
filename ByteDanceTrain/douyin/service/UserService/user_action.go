package UserService

import (
	"douyin/model/entity"
	"douyin/service/DBService"
	"log"
)

// CreateUser create a new record in the mysql database
func CreateUser(user *entity.User) error {
	err := DBService.GetDB().Create(user).Error
	if err != nil {
		log.Println(err)
	}
	return err
}

// QueryUser query user information from mysql by id or name
func QueryUser(user *entity.User) {
	DBService.GetDB().First(user)
}
