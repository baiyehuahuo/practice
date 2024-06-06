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

func SearchFriends(userID uint) []*UserBasic {
	var contacts []*Contact
	utils.GetDB().Where("owner_id = ? and type = 1", userID).Find(&contacts)
	friendIDs := make([]int, len(contacts))
	for i, con := range contacts {
		friendIDs[i] = int(con.TargetID)
	}
	return FindUserByIDs(friendIDs)
}

func AutoMigrateContact() error {
	return utils.GetDB().AutoMigrate(&Contact{})
}
