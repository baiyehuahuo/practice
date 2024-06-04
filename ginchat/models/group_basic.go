package models

import (
	"ginchat/utils"
	"gorm.io/gorm"
)

type GroupBasic struct {
	gorm.Model
	Name    string
	OwnerID uint
	Icon    string
	Type    int
	Desc    string
}

func (c *GroupBasic) TableName() string {
	return "group_basic"
}

func AutoMigrateGroupBasic() error {
	return utils.GetDB().AutoMigrate(&GroupBasic{})
}
