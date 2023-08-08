package DBService

import (
	"douyin/constants"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open(mysql.Open(constants.DatabasePath))
	if err != nil {
		panic(err)
	}
	if err = db.Error; err != nil {
		panic(err)
	}
}

func GetDB() *gorm.DB {
	return db
}
