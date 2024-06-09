package models

import (
	"errors"
	"ginchat/utils"
	"gorm.io/gorm"
)

type Community struct {
	gorm.Model
	Name    string `gorm:"type:varchar(255);not null"`
	OwnerID uint   `gorm:"not null"`
	Image   string `gorm:"type:varchar(255);not null"`
	Desc    string `gorm:"type:varchar(255);not null"`
}

func (Community) TableName() string {
	return "community"
}

func AutoMigrateCommunity() error {
	return utils.GetDB().AutoMigrate(&Community{})
}

func CreateCommunity(name string, ownerID uint) error {
	if len(name) == 0 {
		return errors.New("name is empty")
	}
	if ownerID == 0 {
		return errors.New("ownerID is empty")
	}
	return utils.GetDB().Create(&Community{Name: name, OwnerID: ownerID}).Error
}
