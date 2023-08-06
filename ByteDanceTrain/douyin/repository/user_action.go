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

// FindUser query user information from mysql by id or name
func FindUser(user *model.User) {
	db.First(user)
}
