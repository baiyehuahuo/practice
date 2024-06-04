package models

import (
	"ginchat/utils"
	"gorm.io/gorm"
)

type Contact struct {
	gorm.Model
	OwnerID  uint // who's contact
	TargetID uint // target user id
	Type     int
	Desc     string `gorm:"type:varchar(255);"`
}

func (c *Contact) TableName() string {
	return "contact"
}

func AutoMigrateContact() error {
	return utils.GetDB().AutoMigrate(&Contact{})
}
