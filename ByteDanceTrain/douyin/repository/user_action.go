package repository

import (
	"douyin/model"
	"log"
)

// CreateUser create a new record in the mysql database
func CreateUser(user *model.User) error {
	err := db.Create(user).Error
	if err != nil {
		log.Println(err)
	}
	return err
}

// QueryUser query user information from mysql by id or name
func QueryUser(user *model.User) {
	db.First(user)
}
